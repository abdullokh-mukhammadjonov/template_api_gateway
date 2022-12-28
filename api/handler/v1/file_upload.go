package v1

import (
	"context"
	"errors"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_api_gateway"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Router /v1/file-upload [post]
// @Summary Create file
// @Description API for creating file
// @Tags file_upload
// @Param file formData file true "file"
// @Param use_p query string false "use_p"
// @Param entity_id query string false "entity_id"
// @Param property_id query string false "property_id"
// @Param file_name_id query string false "file_name_id"
// @Param region body var_api_gateway.EntityFilesSwag  true "region"
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Success 200 {object} var_api_gateway.FilesResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) FileUpload(c *gin.Context) {
	var (
		fileURL     string
		entityFiles var_api_gateway.EntityFilesSwag
		use_p       = c.Query("use_p")
		fileID      = primitive.NewObjectID().Hex()
	)

	if err := c.ShouldBind(&entityFiles); HandleHTTPError(c, http.StatusBadRequest, "DiscussionLogicService.Action.Create.BindingAction", err) {
		return
	}

	bucketName := h.cfg.BucketName
	if use_p == "1" {
		bucketName = h.cfg.JackBucketName
		fileID = "enchanted_" + fileID
	}

	minioClient, err := minio.New(h.cfg.MinioDomain, &minio.Options{
		Creds:  credentials.NewStaticV4(h.cfg.MinioAccessKeyID, h.cfg.MinioSecretAccesKey, ""),
		Secure: true,
	})
	if HandleHTTPError(c, http.StatusInternalServerError, "error while creating bucket in minio", err) {
		return
	}

	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if HandleHTTPError(c, http.StatusInternalServerError, "error while checking existance of bucket"+bucketName, err) {
		return
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: h.cfg.MinioDomain})
		if HandleHTTPError(c, http.StatusInternalServerError, "error while creating bucket in minio", err) {
			return
		}
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 200*1024*1024)
	file, err := c.FormFile("file")
	if HandleHTTPError(c, http.StatusInternalServerError, "error while getting file", err) {
		return
	}
	contentType := file.Header["Content-Type"][0]
	if len(contentType) <= 1 {
		contentType = "multipart/form-data"
	}

	fileURL = fileID + filepath.Ext(file.Filename)
	object, err := file.Open()
	if HandleHTTPError(c, http.StatusInternalServerError, "error while creating opening file", err) {
		return
	}
	defer object.Close()

	if file.Size > 10*1024*1024 && use_p != "1" {
		c.JSON(http.StatusBadRequest, template_variables.FailureResponse{Success: false, Error: "fayl hajmi katta", Message: "kichikroq hajmli fayl yuklang (max=10 mb)"})
		return
	}
	_, err = minioClient.PutObject(context.Background(), bucketName, fileURL, object, file.Size,
		minio.PutObjectOptions{ContentType: contentType})
	if HandleHTTPError(c, http.StatusInternalServerError, "error while putting object in minio", err) {
		return
	}

	c.JSON(http.StatusOK, var_api_gateway.FilesResponse{
		FilePath: fileURL,
		URL:      h.cfg.MinioDomain + "/" + bucketName + "/" + fileURL,
	})
}

// @Router /v1/image-upload [post]
// @Tags file_upload
// @Param image formData file true "image"
// @Param entity_id query string false "entity_id"
// @Param file_name_id query string false "file_name_id"
// @Param property_id query string false "property_id"
// @Accept multipart/form-data
// @Success 200 {object} var_api_gateway.FilesResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) ImageUpload(c *gin.Context) {

	var (
		fileNameId = c.Query("file_name_id")
	)

	minioClient, err := minio.New(h.cfg.MinioDomain, &minio.Options{
		Creds:  credentials.NewStaticV4(h.cfg.MinioAccessKeyID, h.cfg.MinioSecretAccesKey, ""),
		Secure: true,
	})
	if HandleHTTPError(c, http.StatusInternalServerError, "error while creating bucket in minio", err) {
		return
	}

	exists, _ := minioClient.BucketExists(context.Background(), h.cfg.BucketName)
	if !exists {
		err = minioClient.MakeBucket(context.Background(), h.cfg.BucketName, minio.MakeBucketOptions{Region: h.cfg.MinioDomain})
		if HandleHTTPError(c, http.StatusInternalServerError, "error while creating bucket in minio", err) {
			return
		}
	}
	file, err := c.FormFile("image")
	if HandleHTTPError(c, http.StatusInternalServerError, "error while getting file", err) {
		return
	}
	contentType := file.Header["Content-Type"][0]
	if len(contentType) <= 1 {
		contentType = "multipart/form-data"
	}
	fileName := file.Filename

	ext := filepath.Ext(fileName)
	if ext == ".jpeg" || ext == ".jpg" || ext == ".png" {

		file.Filename = primitive.NewObjectID().Hex() + filepath.Ext(fileName)

		object, err := file.Open()
		if HandleHTTPError(c, http.StatusInternalServerError, "error while creating bucket in minio", err) {
			return
		}
		defer object.Close()
		_, err = minioClient.PutObject(context.Background(), h.cfg.BucketName, file.Filename, object, file.Size,
			minio.PutObjectOptions{ContentType: contentType})
		if HandleHTTPError(c, http.StatusInternalServerError, "error while creating bucket in minio", err) {
			return
		}

		if fileNameId == "" {
			fileNameId = uuid.Nil.String()
		}

		c.JSON(http.StatusOK, var_api_gateway.FilesResponse{
			FilePath: file.Filename,
			URL:      h.cfg.MinioDomain + "/" + h.cfg.BucketName + "/" + file.Filename,
		})
	} else {
		err = errors.New("error filetype,valid  only .png, .jpg, .jpeg")

		if HandleHTTPError(c, http.StatusBadRequest, "error file uploading", err) {
			return
		}

	}
}
