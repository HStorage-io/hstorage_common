package hstorage_common

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudflare/cloudflare-go"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
	"gorm.io/gorm"
)

type RequestMethod string
type FileType string

const (
	RequestMethodWeb  RequestMethod = "web"
	RequestMethodAPI  RequestMethod = "api"
	RequestMethodSFTP RequestMethod = "sftp"

	FileTypeVideo      FileType = "video"
	FileTypeImage      FileType = "image"
	FileTypeDocument   FileType = "document"
	FileTypePDF        FileType = "pdf"
	FileTypeWord       FileType = "word"
	FileTypeExcel      FileType = "excel"
	FileTypePowerPoint FileType = "powerpoint"
	FileTypeCSV        FileType = "csv"
	FileTypeArchive    FileType = "archive"
)

type Upload struct {
	Count                     uint           `gorm:"default:0" json:"count"`
	CreatedAt                 time.Time      `json:"created_at"`
	CreatedAtFormatted        string         `gorm:"-" json:"created_at_formatted"`
	DeleteDate                sql.NullTime   `gorm:"default:null" json:"delete_date"`
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
	DownloadLimitCount        uint           `gorm:"default:null" json:"download_limit_count"`
	DownloadNotificationCount uint           `gorm:"default:0" json:"notification_count"` // 通知回数、最大値を決めるために利用
	DownloadUrl               string         `gorm:"-" json:"download_url"`
	ExternalID                string         `gorm:"type:char(26);not null; index:idx_external_id,priority:1" json:"external_id"`
	FileName                  string         `gorm:"type:varchar(255); not null" json:"file_name"`
	FileSize                  uint64         `gorm:"not null; index:idx_user_id_state_file_size,priority:3; comment:byte" json:"file_size"`
	FileType                  FileType       `gorm:"-" json:"file_type"` // e.x video, photo
	Group                     []Group        `gorm:"-" json:"group"`     // List.vue の b-taginput でのみ利用する
	GroupID                   uint           `gorm:"not null; default:0; index:idx_group_id_state,priority:1" json:"group_id"`
	Hash                      string         `gorm:"-" json:"hash"`
	ID                        uint           `gorm:"primaryKey" json:"id"`
	IsBusiness                bool           `gorm:"default:0" json:"is_business"`
	IsEncrypt                 bool           `gorm:"default:0" json:"is_encrypt"`
	IsInfected                uint           `gorm:"type:tinyint(1);default:0;comment:0 未処理, 1 パス, 2 ウイルス" json:"is_infected"`
	IsPremium                 bool           `gorm:"default:0" json:"is_premium"`
	IsRedirect                bool           `gorm:"-" json:"is_redirect"`
	ListID                    uint           `gorm:"-" json:"list_id"` // List.vue の key-field で利用
	OriginalFileName          string         `gorm:"type:varchar(255); not null" json:"original_file_name"`
	Password                  string         `gorm:"type:varchar(255);default:null" json:"password"`
	ThumbURL                  string         `gorm:"-" json:"thumb_url"`
	UpdatedAt                 time.Time      `json:"updated_at"`
	UpdatedAtFormatted        string         `gorm:"-" json:"updated_at_formatted"`
	UploadedBy                RequestMethod  `gorm:"type:varchar(100);not null;default:web;size:10" json:"uploaded_by"`
	Url                       string         `gorm:"-" json:"url"`
	UserID                    string         `gorm:"type:varchar(255); not null; index:idx_user_id_state,priority:1; index:idx_user_id_state_file_size,priority:1" json:"user_id"`
}

type PreSignedReq struct {
	DeleteDate         *time.Time `json:"delete_date"`
	DownloadLimitCount uint       `json:"download_limit_count"`
	FileName           string     `json:"file_name" binding:"required"`
	GroupUID           string     `json:"group_uid"`
	IsEncrypt          *bool      `json:"is_encrypt" binding:"required"`
	IsGuest            bool       `json:"is_guest"`
	Password           string     `json:"password"`
}

type PreSignedResp struct {
	AWSKey     string `json:"aws_key" binding:"required"`
	AWSUrl     string `json:"aws_url" binding:"required"`
	Bucket     string `json:"bucket" binding:"required"`
	ExternalID string `json:"external_id" binding:"required"`
	FileName   string `json:"file_name" biding:"required"`
	Key        string `json:"key" binding:"required"`
	SseKey     string `json:"sseKey"`
	SseMD5     string `json:"sseMD5"`
}

type PreSignedRespV1 struct {
	//AWSKey       string `json:"aws_key" binding:"required"`
	//AWSUrl       string `json:"aws_url" binding:"required"`
	//Bucket       string `json:"bucket" binding:"required"`
	ExternalID string `json:"external_id" binding:"required"`
	FileName   string `json:"file_name" biding:"required"`
	//Key          string `json:"key" binding:"required"`
	SseKey       string `json:"sseKey"`
	SseMD5       string `json:"sseMD5"`
	PreSignedURL string `json:"presigned_url" binding:"required"`
	ShareURL     string `json:"share_url" binding:"required"`
	DirectURL    string `json:"direct_url" binding:"required"`
}

type UploadClient struct {
	S3Minio     *s3.Client
	Bucket      string
	FileBaseURL string
	CF          CF
}

type CF struct {
	API *cloudflare.API
	ID  string
}

func NewUploadClient(bucket, minioAccessKeyID, minioSecretAccessKey, minioEndpoint, fileBaseURL, cfID, cfKey, cfEmail string) *UploadClient {
	s3MinioCfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	s3Minio := s3.NewFromConfig(s3MinioCfg, func(o *s3.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(minioAccessKeyID, minioSecretAccessKey, "")
		o.Region = "ap-northeast-1"
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s", minioEndpoint))
		o.HTTPClient = &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   3600 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        32,
				MaxIdleConnsPerHost: 16,
				IdleConnTimeout:     60 * time.Second,
			},
		}
		o.UsePathStyle = true
	})
	otelaws.AppendMiddlewares(&s3MinioCfg.APIOptions)

	cf, err := cloudflare.New(cfKey, cfEmail)
	if err != nil {
		log.Fatalf("unable to create cloudflare client, %v", err)
	}

	return &UploadClient{
		S3Minio:     s3Minio,
		Bucket:      bucket,
		FileBaseURL: fileBaseURL,
		CF: CF{
			API: cf,
			ID:  cfID,
		},
	}
}

func (c *UploadClient) DeleteFiles(ctx context.Context, db *gorm.DB, uploads []Upload) error {
	for i := range uploads {
		err := c.DeleteFile(ctx, db, &uploads[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *UploadClient) DeleteFile(ctx context.Context, db *gorm.DB, upload *Upload) (err error) {
	opts := &s3.DeleteObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(c.GetS3Key(upload)),
	}

	_, err = c.S3Minio.DeleteObject(ctx, opts)
	if err != nil {
		return err
	}

	err = c.deleteCacheFromCF(ctx, upload)
	if err != nil {
		return err
	}

	err = c.deleteFromDB(ctx, db, upload)
	if err != nil {
		return err
	}

	return nil
}

// Cloudflare のキャッシュを削除
func (c *UploadClient) deleteCacheFromCF(ctx context.Context, upload *Upload) error {
	fileURL := GetFileURL(c.FileBaseURL, upload, true)

	pcr := cloudflare.PurgeCacheRequest{
		Files: []string{fileURL},
	}

	_, err := c.CF.API.PurgeCache(ctx, c.CF.ID, pcr)

	return err
}

func (c *UploadClient) deleteFromDB(ctx context.Context, db *gorm.DB, upload *Upload) error {
	return db.WithContext(ctx).Delete(upload).Error
}

/*
UploadCountHistory is history of upload not last access history.
*/
type UploadCountHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UploadID  uint      `gorm:"not null; index:idx_upload_id_ip_address"`
	IPAddress string    `gorm:"not null; type:varchar(50); index:idx_upload_id_ip_address"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// GetFileType returns file type (eg: image, video) from file name
// return image, video, document
func GetFileType(fileName string) FileType {
	images := []string{"jpg", "jpeg", "png", "gif", "webp", "ico", "svg"}
	videos := []string{"mp4", "mov", "mpg", "mpeg", "avi", "mkv", "ts", "flv", "3gp", "wmv"}
	archives := []string{"zip", "tar.gz", "7z", "rar", "tar.bz2", "tar.xz", "tar"}
	words := []string{"doc", "docx"}
	excels := []string{"xls", "xlsx"}
	powerpoints := []string{"ppt", "pptx"}

	pos := strings.LastIndex(fileName, ".")
	fileExt := fileName[pos+1:]
	fileExt = strings.ToLower(fileExt)

	for _, image := range images {
		if fileExt == image {
			return FileTypeImage
		}
	}

	for _, video := range videos {
		if fileExt == video {
			return FileTypeVideo
		}
	}

	for _, word := range words {
		if fileExt == word {
			return FileTypeWord
		}
	}

	for _, excel := range excels {
		if fileExt == excel {
			return FileTypeExcel
		}
	}

	for _, powerpoint := range powerpoints {
		if fileExt == powerpoint {
			return FileTypePowerPoint
		}
	}

	for _, archive := range archives {
		if fileExt == archive {
			return FileTypeArchive
		}
	}

	switch FileType(fileExt) {
	case FileTypePDF:
		return FileTypePDF
	case FileTypeCSV:
		return FileTypeCSV
	default:
		return FileTypeDocument
	}
}

func (c *UploadClient) GetS3Key(upload *Upload) string {
	key := fmt.Sprintf("upload/%s/%s", upload.UserID, upload.FileName)

	if RequestMethod(upload.UploadedBy) == RequestMethodSFTP {
		key = fmt.Sprintf("upload/%s/%s", upload.UserID, upload.OriginalFileName)
	}

	return key
}

func GetFileURL(fileBaseURL string, upload *Upload, needPasswordQuery bool) string {
	fileName := upload.FileName
	fileURL := fmt.Sprintf("%s/%s", fileBaseURL, fileName)

	if upload.Password != "" && needPasswordQuery {
		fileURL = fmt.Sprintf("%s?password=%s", fileURL, upload.Password)
	}

	return fileURL
}

func GetDownloadURL(fileBaseURL string, upload *Upload) string {
	fileName := upload.FileName
	fileURL := fmt.Sprintf("%s/%s?download=true", fileBaseURL, fileName)

	if upload.Password != "" {
		fileURL = fmt.Sprintf("%s&password=%s", fileURL, upload.Password)
	}

	return fileURL
}

func GetThumbnailURL(fileBaseURL string, upload *Upload, password string, isLocal bool) string {
	var thumbURL string

	baseURL := "https://thumbnail.hstorage.io/unsafe/fit-in/200x0/"

	// needPasswordQuery=false: thumbnail の password はエスケープする必要があるため、ここではクエリは不要
	thumbURL = fmt.Sprintf("%s%s", baseURL, GetFileURL(fileBaseURL, upload, false))

	if password != "" {
		thumbURL = fmt.Sprintf("%s%s", thumbURL, url.QueryEscape(fmt.Sprintf("?password=%s", password)))
	}

	return thumbURL
}
