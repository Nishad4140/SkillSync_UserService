package usecase

import (
	"bytes"
	"context"
	"log"
	"os"
	"time"

	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
	"github.com/Nishad4140/SkillSync_UserService/internal/adapters"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type UserUsecase struct {
	userAdapter adapters.AdapterInterface
}

func NewUserUsecase(userAdapter adapters.AdapterInterface) *UserUsecase {
	return &UserUsecase{
		userAdapter: userAdapter,
	}
}

func (user *UserUsecase) UploadClientImage(req *pb.ImageRequest, profileId string) (string, error) {
	minioClient, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESSKEY"), os.Getenv("MINIO_SECRETKEY"), ""),
		Secure: false,
	})
	if err != nil {
		log.Println("error while initializing minio", err)
		return "", err
	}
	objectName := "/" + req.ObjectName
	contentType := `image/jpeg`

	n, err := minioClient.PutObject(context.Background(), os.Getenv("BUCKET_NAME"), objectName, bytes.NewReader(req.ImageData), int64(len(req.ImageData)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("error while uploading to minio", err)
		return "", err
	}
	log.Printf("Successfully uploaded %s of size %v to minio", objectName, n)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), os.Getenv("BUCKET_NAME"), objectName, time.Second*24*60*60, nil)
	if err != nil {
		log.Println("error while generating the url", err)
		return "", err
	}
	url, err := user.userAdapter.UploadClientProfileImage(presignedURL.String(), profileId)
	return url, err

}

func (user *UserUsecase) UploadFreelancerImage(req *pb.ImageRequest, profileId string) (string, error) {
	minioClient, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESSKEY"), os.Getenv("MINIO_SECRETKEY"), ""),
		Secure: false,
	})
	if err != nil {
		log.Println("error while initializing minio", err)
		return "", err
	}
	objectName := "images/" + req.ObjectName
	contentType := `image/jpeg`

	n, err := minioClient.PutObject(context.Background(), os.Getenv("BUCKET_NAME"), objectName, bytes.NewReader(req.ImageData), int64(len(req.ImageData)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("error while uploading to minio", err)
		return "", err
	}
	log.Printf("Successfully uploaded %s of size %v to minio", objectName, n)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), os.Getenv("BUCKET_NAME"), objectName, time.Second*24*60*60, nil)
	if err != nil {
		log.Println("error while generating the url", err)
		return "", err
	}
	url, err := user.userAdapter.UploadFreelancerProfileImage(presignedURL.String(), profileId)
	return url, err

}
