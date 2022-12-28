package v1

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_api_gateway"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_user_service"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/pkg/logger"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/pkg/security"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrAlreadyExists       = "ALREADY_EXISTS"
	ErrBadRequest          = "BAD_REQUEST"
	ErrNotFound            = "NOT_FOUND"
	ErrInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrServiceUnavailable  = "SERVICE_UNAVAILABLE"
	log                    = logger.New("DEBUG", "admin_api_gateway")
)

const (
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
)

type HandlerV1 struct {
	cfg        config.Config
	log        logger.Logger
	services   services.ServiceManager
	redisCache services.RedisManager
}

type HandlerV1Options struct {
	Cfg         config.Config
	Log         logger.Logger
	Services    services.ServiceManager
	RedisClient services.RedisManager
}

func New(options *HandlerV1Options) *HandlerV1 {
	return &HandlerV1{
		log:        options.Log,
		cfg:        options.Cfg,
		services:   options.Services,
		redisCache: options.RedisClient,
	}
}

// ProtoToStruct is ..
func ProtoToStruct(structure interface{}, protoMessage proto.Message) error {
	var jm jsonpb.Marshaler

	jm.EmitDefaults = true
	jm.OrigName = true

	message, err := jm.MarshalToString(protoMessage)

	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(message), &structure)

	return err
}

// ProtoToStructNumeric is ..
func ProtoToStructNumeric(structure interface{}, protoMessage proto.Message) error {
	byte, err := json.Marshal(protoMessage)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byte, structure)
	return err
}

// ProtoToStructNumeric is ..
func MarshalUnmarshal(structure interface{}, response interface{}) error {
	byte, err := json.Marshal(structure)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byte, &response)
	return err
}

func (h *HandlerV1) MakeProxy(c *gin.Context, proxyUrl, path string) (err error) {
	req := c.Request

	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		h.log.Error("error in parse addr: %v", logger.Error(err))
		c.String(http.StatusInternalServerError, "error")
		return
	}

	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	req.URL.Path = path
	transport := http.DefaultTransport

	resp, err := transport.RoundTrip(req)
	if HandleHTTPError(c, 500, "error in round trip:", err) {
		return
	}

	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	_, _ = bufio.NewReader(resp.Body).WriteTo(c.Writer)
	return
}
func (h *HandlerV1) UserInfo(c *gin.Context, required bool) (*var_user_service.LoginInfo, error) {
	var token = c.Request.Header.Get("Authorization")
	if token != "" {
		claims, err := security.ExtractClaims(token, h.cfg.LoginSecretAccessKey)
		if err != nil {
			HandleHTTPError(c, http.StatusUnauthorized, "please provide valid token", errors.New("unauthorized"))
			return &var_user_service.LoginInfo{}, err
		}
		if claims["user_type"].(string) == "staff" || claims["user_type"].(string) == "superadmin" {
			return &var_user_service.LoginInfo{
				ID:             claims["id"].(string),
				Login:          claims["login"].(string),
				UserType:       claims["user_type"].(string),
				RoleID:         claims["role_id"].(string),
				Soato:          claims["soato"].(string),
				OrganizationID: claims["organization_id"].(string),
			}, nil
		} else {
			return &var_user_service.LoginInfo{
				ID:       claims["id"].(string),
				Login:    claims["login"].(string),
				UserType: claims["user_type"].(string),
				FullName: claims["full_name"].(string),
			}, nil
		}
	} else if required {
		HandleHTTPError(c, http.StatusUnauthorized, "please provide token", errors.New("unauthorized"))
		return &var_user_service.LoginInfo{}, errors.New("not authorized")
	} else {
		return &var_user_service.LoginInfo{}, nil
	}
}

// This function is to parsing upcoming query params
// please, make sure you write proper message error
func ParseQueryParam(c *gin.Context, log logger.Logger, key string, defaultValue string) (int, error) {
	valueString := c.DefaultQuery(key, defaultValue)

	value, err := strconv.Atoi(valueString)
	if err != nil {
		message := "Error while parsing query"
		log.Error(message, logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": message,
			"error":   err.Error(),
		})
		c.Abort()
		return 0, err
	}
	return value, err
}

// This function is for parsing string query params
// with default value if given param does not exist
func StringQueryWithDefault(c *gin.Context, key string, defaultValue string) string {
	valueString := c.Query(key)
	if valueString == "" {
		return defaultValue
	}
	return valueString
}

// This function is to handle all RPC server errors
// please, make sure you write proper message error
func HandleRPCError(c *gin.Context, message string, err error) (hasError bool) {
	st, ok := status.FromError(err)
	if st.Code() == codes.Canceled {
		log.Error(message+", canceled ", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"success":     false,
			"service":     st.Message(),
			"api-gateway": message,
			"error":       err.Error(),
		})
		return true
	} else if st.Code() == codes.AlreadyExists || st.Code() == codes.InvalidArgument {
		log.Error(message+", already exists", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"success":     false,
			"service":     st.Message(),
			"api-gateway": message,
			"error":       err.Error(),
		})
		return true
	} else if st.Code() == codes.NotFound {
		log.Error(message+", not found", logger.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"success":     false,
			"service":     st.Message(),
			"api-gateway": message,
			"error":       err.Error(),
		})
		return true
	} else if st.Code() == codes.Unavailable {
		log.Error(message+", service unavailable", logger.Error(err))
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success":     false,
			"service":     st.Message(),
			"api-gateway": message,
			"error":       err.Error(),
		})
		return true
	} else if !ok || st.Code() == codes.Internal || st.Code() == codes.Unknown || err != nil {
		log.Error(message+", internal server error", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":     false,
			"service":     st.Message(),
			"api-gateway": message,
			"error":       err.Error(),
		})
		return true
	}
	return false
}

// This function is to handle all HTTP client errors
// please, make sure you write proper message error
func HandleHTTPError(c *gin.Context, code int, message string, err error) bool {
	if err != nil && code == http.StatusBadRequest {
		log.Error(message+" --> Error: ", logger.Error(err))
		c.JSON(http.StatusBadRequest, template_variables.FailureResponse{
			Success: false,
			Message: message,
			Error:   err.Error(),
		})
		return true
	} else if err != nil && code == http.StatusInternalServerError {
		log.Error(message+", --> Error	", logger.Error(err))
		c.JSON(http.StatusServiceUnavailable, template_variables.FailureResponse{
			Success: false,
			Message: message,
			Error:   err.Error(),
		})
		return true
	} else if err != nil && code == http.StatusUnauthorized {
		log.Error(message+" --> Error: ", logger.Error(err))
		c.JSON(http.StatusUnauthorized, template_variables.FailureResponse{
			Success: false,
			Message: message,
			Error:   err.Error(),
		})
		return true
	} else if err != nil && code == http.StatusConflict {
		log.Error(message+" --> Error: ", logger.Error(err))
		c.JSON(http.StatusConflict, template_variables.FailureResponse{
			Success: false,
			Message: message,
			Error:   err.Error(),
		})
		return true
	}
	return false
}

// This function is to handle all HTTP client validation errors
// please, make sure you write proper message error
func HandleHTTPValidationError(c *gin.Context, exist bool, code int, err error) bool {
	if !exist && code == http.StatusNotFound {
		log.Error(" --> Error: ", logger.Error(err))
		c.JSON(http.StatusNotFound, template_variables.FailureResponse{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		})
		return true
	} else if exist && code == http.StatusConflict {
		log.Error(" --> Error: ", logger.Error(err))
		c.JSON(http.StatusConflict, template_variables.FailureResponse{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		})
		return true
	} else if !exist && code == http.StatusInternalServerError {
		log.Error(", --> Error	", logger.Error(err))
		c.JSON(http.StatusServiceUnavailable, template_variables.FailureResponse{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		})
		return true
	}
	return false
}

func ValidateTime(date string) error {
	var layout = "2006-01-02"

	_, err := time.Parse(layout, date)
	if err != nil {
		return var_api_gateway.ErrInvalidDate
	}

	return nil
}

func ParseBoolParam(c *gin.Context, log logger.Logger, key, defaultValue string) (bool, error) {
	valueString := c.DefaultQuery(key, defaultValue)

	value, err := strconv.ParseBool(valueString)
	if err != nil {
		message := "Error while parsing query"
		log.Error(message, logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": message,
			"error":   err.Error(),
		})
		c.Abort()
		return false, err
	}
	return value, err
}

func GetAccessRefreshTokensToken(cfg config.Config) (time.Duration, time.Duration) {
	return time.Duration(cfg.AccessTokenExpireDuration) * time.Hour, time.Duration(cfg.RefreshTokenExpireDuration) * time.Hour
}
