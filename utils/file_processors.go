package utils

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/kayprogrammer/socialnet-v4/config"

)

var cloudName string
var apiKey string
var apiSecret string 
var cld *cloudinary.Cloudinary
var baseFolder = "socialnet-v4/"

func init() {
	cfg := config.GetConfig()
	// Load environment variables here
	cloudName = cfg.CloudinaryCloudName
	apiKey = cfg.CloudinaryAPIKey
	apiSecret = cfg.CloudinaryAPISecret

	// Initialize the Cloudinary client
	cld, _ = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
}

type SignatureFormat struct {
	PublicId  string `json:"public_id" example:"images/f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Signature string `json:"signature" example:"e1ba4683fbbf90b75ca22e9f8e545b18c6b24eae"`
	Timestamp int64  `json:"timestamp" example:"1678828200"`
}

var ImageExtensions = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/gif":  "gif",
	"image/bmp":  "bmp",
	"image/webp": "webp",
	"image/tiff": "tiff",
	"image/svg+xml": "svg",
}

func GenerateFileSignature(key string, folder string) SignatureFormat {
	key = fmt.Sprintf("%s%s/%s", baseFolder, folder, key)
	timestamp := time.Now().Unix()
	params := map[string]interface{}{
		"public_id": key,
		"timestamp": timestamp,
	}

	// Convert the params to url.Values
	values := url.Values{}
	for k, v := range params {
		values.Add(k, fmt.Sprintf("%v", v))
	}
	resp, err := api.SignParameters(values, apiSecret)
	if err != nil {
		log.Fatal("Error signing params: ", err)
	}
	signatureResp := SignatureFormat{PublicId: key, Signature: resp, Timestamp: timestamp}
	return signatureResp
}

func GenerateFileUrl(key string, folder string, contentType string) string {
	key = fmt.Sprintf("%s%s/%s.%s", baseFolder, folder, key, ImageExtensions[contentType])

	// Generate the Cloudinary URL for the existing resource
	urls, err := cld.Media(key)
	if err != nil {
		log.Println("Error generating Cloudinary URL:", err)
	}
	url, err := urls.String()
	if err != nil {
		log.Println("Error converting to string:", err)
	}
	return url
}

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}

func UploadImage(file *os.File, key string, folder string) {
	key = fmt.Sprintf("%s%s/%s", baseFolder, folder, key)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: key, Overwrite: BoolAddr(true), Faces: BoolAddr(true)})
}
