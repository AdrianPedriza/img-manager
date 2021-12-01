package storage

import (
	"fmt"
	"github.com/AdrianPedriza/img-manager/pkg/model"
	"io"
	"log"
	"os"
	"strings"
)

var	storageHandler StorageHandler

type StorageHandler interface {
	Insert(fileId string, reader io.Reader) error
	GetImageById(id string) (*model.Image, error)
}

func GetStorageHandler() StorageHandler {
	if storageHandler == nil {
		createStorageHandlerFromEnv()
	}
	return storageHandler
}

func getImgStoreDir() string{
	imgStoreDirEnvValue := os.Getenv("IMG_STORAGE_DIR")
	if len(imgStoreDirEnvValue) == 0{
		imgStoreDir, err := os.MkdirTemp("", "images")
		if err != nil{
			log.Fatal(err)
		}
		err = os.Setenv("IMG_STORAGE_DIR", imgStoreDir)
		log.Println(fmt.Sprintf("IMG_STORAGE_DIR value set to '%s'", imgStoreDir))
		if err != nil{
			log.Fatal(err)
		}
		return imgStoreDir
	}
	return imgStoreDirEnvValue
}

func createStorageHandlerFromEnv() {
	switch os.Getenv("IMG_STORAGE_DEVICE") {
	case "memory":
		storageHandler = memoryStorageDeviceHandler{
			images: map[string]*model.Image{},
		}
	case "device":
		storageHandler = deviceStorageDeviceHandler{}
	default:
		storageHandler = deviceStorageDeviceHandler{}
		log.Println("IMG_STORAGE_DEVICE environment variable not set properly. Accepted values [memory, device]. Using 'device' as default value.")
	}
}

func GetFileInfoFromFilename(filename string) (string, string) {
	filenameInfo := strings.Split(filename, ".")
	fileFormat := filenameInfo[len(filenameInfo)-1]
	fileId := strings.Join(filenameInfo[:len(filenameInfo)-1], ".")
	return fileId, fileFormat
}
