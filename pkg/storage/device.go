package storage

import (
	"fmt"
	"github.com/AdrianPedriza/img-manager/pkg/model"
	"io"
	"os"
)

type deviceStorageDeviceHandler struct{}

func (deviceHandler deviceStorageDeviceHandler) Insert(filename string, imgFile io.Reader) error{

	fileId, fileFormat := GetFileInfoFromFilename(filename)

	filePath := fmt.Sprintf("%s/%s.%s", getImgStoreDir(), fileId, fileFormat)

	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return os.ErrExist
	}

	encoderDecoder, err := model.GetEncoderDecoder(fileFormat)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	imgDecoded, err := encoderDecoder.Decode(imgFile)
	if err != nil {
		return err
	}

	err = encoderDecoder.Encode(f, imgDecoded)
	if err != nil {
		return err
	}
	return nil
}

func (deviceHandler deviceStorageDeviceHandler) GetImageById(filename string) (*model.Image, error){
	fileId, fileFormat := GetFileInfoFromFilename(filename)

	imgFile, err := os.Open(fmt.Sprintf("%s/%s", getImgStoreDir(), filename))
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	encoderDecoder, err := model.GetEncoderDecoder(fileFormat)
	if err != nil {
		return nil, err
	}
	imgDecoded, err := encoderDecoder.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return &model.Image{
		ID:     fileId,
		Format: fileFormat,
		Data:   &imgDecoded,
	}, nil
}
