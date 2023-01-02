package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"io"
	"time"
)

type Reader interface {
	Read(r io.Reader) error
}

type Writer interface {
	Write(w io.Writer) error
}

type UploaderDownloader interface {
	UploaderClient
	DownloaderClient
}

type UploaderClient interface {
	Upload(fileName string, writer Writer) (string, error)
}

type DownloaderClient interface {
	Download(fileURL string, reader Reader) error
}

type GCSConfig struct {
	BucketID                string `json:"bucket_id"`
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"TokenURI"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

type downloaderClient struct {
	bucketHandle *storage.BucketHandle
}

type gcsUploadClient struct {
	bucketHandle *storage.BucketHandle
	bucketName   string
	option       *storage.SignedURLOptions
}

func NewDownloader(config GCSConfig) (DownloaderClient, error) {
	ctx := context.Background()

	jsonConfig := translateConfigToString(config)

	creds, err := google.CredentialsFromJSON(ctx, []byte(jsonConfig), "https://www.googleapis.com/auth/devstorage.read_only")
	if err != nil {
		return nil, fmt.Errorf("[NewDownloader] creating credential fail. %v", err)
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("[NewDownloader] Getting client fail. %v", err)
	}

	return &downloaderClient{
		bucketHandle: client.Bucket(config.BucketID),
	}, nil
}

func (ths *downloaderClient) Download(fileURL string, reader Reader) error {
	ctx := context.Background()

	bucketFileReader, err := ths.bucketHandle.Object(fileURL).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("[Download] fail read file Cause: %v", err)
	}

	defer func() {
		if bucketFileReader != nil {
			bucketFileReader.Close()
		}
	}()

	err = reader.Read(bucketFileReader)
	if err != nil {
		return fmt.Errorf("[Download] fail read file Cause: %v", err)
	}

	return nil
}

func (ths *downloaderClient) Upload(fileName string, writer Writer) (string, error) {
	return "", fmt.Errorf("this method is not supported")
}

func NewUploader(config GCSConfig) (UploaderClient, error) {
	ctx := context.Background()

	jsonConfig := translateConfigToString(config)

	creds, err := google.CredentialsFromJSON(ctx, []byte(jsonConfig), "https://www.googleapis.com/auth/devstorage.read_write")
	if err != nil {
		return nil, fmt.Errorf("[NewUploader] creating credential fail. %v", err)
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("[NewUploader] Getting client fail. %v", err)
	}

	conf, err := google.JWTConfigFromJSON([]byte(jsonConfig))
	if err != nil {
		return nil, fmt.Errorf("[NewUploader] Getting client fail. %v", err)
	}

	options := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(165 * time.Hour),
	}

	return &gcsUploadClient{
		bucketHandle: client.Bucket(config.BucketID),
		bucketName:   config.BucketID,
		option:       options,
	}, nil
}

func (ths *gcsUploadClient) Upload(fileName string, writer Writer) (string, error) {
	ctx := context.Background()

	bucketWriter := ths.bucketHandle.Object(fileName).NewWriter(ctx)

	err := writer.Write(bucketWriter)
	if err != nil {
		return "", fmt.Errorf("[Upload] Failed to Upload File. Cause: %v", err)
	}

	err = bucketWriter.Close()
	if err != nil {
		return "", fmt.Errorf("[Upload] Failed to Upload File. Cause: %v", err)
	}

	url, err := storage.SignedURL(ths.bucketName, fileName, ths.option)
	if err != nil {
		return "", fmt.Errorf("[Upload] Failed to Upload File. Cause: %v", err)
	}

	return url, nil
}

func (ths *gcsUploadClient) Download(fileURL string, reader Reader) error {
	return fmt.Errorf("this method is not supported")
}

func translateConfigToString(config GCSConfig) string {
	gcsConfigByte, _ := json.Marshal(config)

	return string(gcsConfigByte)
}
