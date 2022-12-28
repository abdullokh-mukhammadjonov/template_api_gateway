package util

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	"github.com/minio/minio-go/v7"
	credentials "github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PullFromUrlUploadToMinio(ctx context.Context, url string) (string, error) {
	extension := "pdf"
	url_slices := strings.Split(url, "/")
	external_id := url_slices[len(url_slices)-1]
	if len(url_slices) > 1 {
		ext := strings.ToLower(url_slices[len(url_slices)-2])
		if ext == "docx" || ext == "xlsx" || ext == "doc" {
			extension = ext
		}
	}

	err := DownloadFile(external_id+"."+extension, url)
	if err != nil {
		return "", err
	}

	file_source, err := UploadFileToMinio(ctx, external_id+"."+extension)

	return file_source, err
}

func DownloadFile(filepath string, url string) (err error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// prevent saving broken pdf file as valid
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on parsing response body to byte")
		return err
	}

	if len(string(bodyBytes)) < 50 {
		err := errors.New("the pdf file is broken")
		fmt.Println(err)
		return err
	}

	// Writer the body to file
	if err := ioutil.WriteFile(filepath, bodyBytes, 0644); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//compressing all files and entity info in pdf and uploading to minio
func UploadFileToMinio(ctx context.Context, file_path string) (string, error) {
	cfg := config.Load()

	minioClient, err := minio.New(cfg.MinioDomain, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretAccesKey, ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println("1", err)
		return "", err
	}

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	file_name := primitive.NewObjectID().Hex() + "_" + file_path

	uploadInfo, err := minioClient.PutObject(ctx, cfg.FilesBucketName, file_name, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", err
	}
	fmt.Println("Succesfully uploaded:", uploadInfo)

	file.Close()
	err = os.Remove(file_path)
	if err != nil {
		fmt.Println("error 8", err)
		return "", err
	}
	return file_name, err
}
