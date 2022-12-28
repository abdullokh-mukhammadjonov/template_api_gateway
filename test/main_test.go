package v1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/api"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/pkg/logger"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type header struct {
	Key   string
	Value string
}

var (
	server *gin.Engine
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	logger := logger.New(cfg.LogLevel, "ek_api_gateway")
	grpcClients, _ := services.NewGrpcClients(cfg)
	redisConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
	})
	_, err := redisConn.Ping().Result()
	if err != nil {
		panic(err)
	}
	defer func() {
		redisConn.Close()
	}()

	redisClient := services.NewRedisClient(redisConn, grpcClients, time.Duration(cfg.RedisTimeoutInSeconds))
	// kafka := event.NewKafka(context.Background(), cfg, logger)
	server = api.New(&api.RouterOptions{
		Cfg:         cfg,
		Log:         logger,
		Services:    grpcClients,
		RedisClient: redisClient,
	})
	os.Exit(m.Run())
}

func PerformRequest(method, path string, req, res interface{}, headers ...header) (*httptest.ResponseRecorder, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	for _, h := range headers {
		request.Header.Add(h.Key, h.Value)
	}
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return response, nil
}
