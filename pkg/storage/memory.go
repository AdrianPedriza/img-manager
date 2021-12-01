package storage

import (
	"fmt"
	"github.com/AdrianPedriza/img-manager/pkg/model"
	"io"
	"os"
)

type memoryStorageDeviceHandler struct {
	images map[string]*model.Image
}

func (memoryHandler memoryStorageDeviceHandler) Insert(filename string, imgFile io.Reader) error {
	img := memoryHandler.images[filename]
	if img != nil {
		return os.ErrExist
	}

	fileId, fileFormat := GetFileInfoFromFilename(filename)

	encoderDecoder, err := model.GetEncoderDecoder(fileFormat)
	if encoderDecoder == nil {
		return err
	}
	imgDecoded, err := encoderDecoder.Decode(imgFile)
	if err != nil {
		return err
	}
	memoryHandler.images[filename] = &model.Image{
		ID:     fileId,
		Format: fileFormat,
		Data:   &imgDecoded,
	}
	return nil
}

func (memoryHandler memoryStorageDeviceHandler) GetImageById(id string) (*model.Image, error) {
	img := memoryHandler.images[id]
	if img == nil {
		return nil, fmt.Errorf("no available img with id=%s", id)
	}
	return img, nil
}
