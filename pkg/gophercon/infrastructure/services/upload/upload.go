package upload

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/application/common/helpers"
	"google.golang.org/api/option"
)

// ServiceUploadImpl represents upload service implementations
type ServiceUploadImpl struct {
	Client storage.Client
}

// NewServiceUpload returns new instance of upload service
func NewServiceUpload(ctx context.Context) *ServiceUploadImpl {
	credentials := helpers.MustGetEnvVar("GOOGLE_APPLICATION_CREDENTIALS")

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		fmt.Println(fmt.Errorf("error creating client: %w", err))
		panic(err)
	}

	defer client.Close()

	return &ServiceUploadImpl{
		Client: *client,
	}
}

// UploadMedia uploads media to GCS
func (u *ServiceUploadImpl) UploadMedia(ctx context.Context, name string, file io.Reader) (string, error) {
	bucketName := helpers.MustGetEnvVar("GOPHERCON_BUCKET_NAME")

	object := u.Client.Bucket(bucketName).Object(name)

	wc := object.NewWriter(ctx)
	// wc.ContentType = contentType
	wc.ChunkSize = 256 * 1024 // 256 KB chunk size

	if _, err := io.Copy(wc, file); err != nil {
		wc.Close()
		return "", err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	if err := wc.Close(); err != nil {
		return "", err
	}

	_, err := object.Attrs(timeoutCtx)
	if err != nil {
		return "", err
	}

	signedURL, err := u.Client.Bucket(bucketName).SignedURL(name, &storage.SignedURLOptions{
		Method:  http.MethodGet,
		Scheme:  storage.SigningSchemeV2,
		Expires: time.Now().Add(24 * time.Hour * 365 * 100),
	})
	if err != nil {
		return "", err
	}

	return signedURL, nil
}
