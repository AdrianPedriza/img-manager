package rest

import (
	"errors"
	"fmt"
	"github.com/AdrianPedriza/img-manager/pkg/model"
	"github.com/AdrianPedriza/img-manager/pkg/storage"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func UploadImageRequest(writer http.ResponseWriter, request *http.Request) {
	imgFile, header, err := request.FormFile("image")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer imgFile.Close()
	err = storage.GetStorageHandler().Insert(header.Filename, imgFile)
	if err != nil {
		log.Println(err)
		if errors.Is(err, image.ErrFormat){
			writer.WriteHeader(http.StatusBadRequest)
		}else if errors.Is(err, os.ErrExist) {
			writer.WriteHeader(http.StatusConflict)
		}else{
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	writer.WriteHeader(http.StatusCreated)
	log.Println(fmt.Sprintf("Uploaded %s image.", header.Filename))
}

func GetImageByIDRequest(writer http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/images/")

	img, err := storage.GetStorageHandler().GetImageById(id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, image.ErrFormat){
			writer.WriteHeader(http.StatusBadRequest)
		}else if errors.Is(err, os.ErrNotExist){
			writer.WriteHeader(http.StatusNotFound)
		}else{
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	tmpFile, err := ioutil.TempFile("", fmt.Sprintf("%s.%s", img.ID, img.Format))
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFile.Name())

	encoderDecoder, err := model.GetEncoderDecoder(img.Format)
	if err!=nil{
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encoderDecoder.Encode(tmpFile, *img.Data)
	if err!=nil{
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnedFilename := fmt.Sprintf("%s.%s", img.ID, img.Format)
	writer.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(returnedFilename))
	writer.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(writer, request, tmpFile.Name())
	writer.WriteHeader(http.StatusOK)
	log.Println(fmt.Sprintf("Returned %s image.", returnedFilename))
}
