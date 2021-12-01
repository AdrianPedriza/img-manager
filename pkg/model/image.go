package model

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type Image struct {
	ID string
	Format string
	Data *image.Image
}

type ImageEncoderDecoder interface {
	Encode(io.Writer, image.Image) error
	Decode(io.Reader) (image.Image, error)
}

type PngEncoderDecoder struct {
}

func (PngEncoderDecoder) Encode(file io.Writer, img image.Image) error{
	return png.Encode(file, img)
}

func (PngEncoderDecoder) Decode(imgFile io.Reader) (image.Image, error){
	return png.Decode(imgFile)
}

type JpegEncoderDecoder struct {
}

func (JpegEncoderDecoder) Encode(file io.Writer, img image.Image) error{
	opt := jpeg.Options{
		Quality: 100,
	}
	return jpeg.Encode(file, img, &opt)
}

func (JpegEncoderDecoder) Decode(imgFile io.Reader) (image.Image, error){
	return jpeg.Decode(imgFile)
}

type GifEncoderDecoder struct {
}

func (GifEncoderDecoder) Encode(file io.Writer, img image.Image) error{
	opt := gif.Options {
		NumColors: 256,
	}
	return gif.Encode(file, img, &opt)
}

func (GifEncoderDecoder) Decode(imgFile io.Reader) (image.Image, error){
	return gif.Decode(imgFile)
}

func GetEncoderDecoder(format string) (ImageEncoderDecoder, error) {
	var encoderDecoder ImageEncoderDecoder
	switch format {
	case "png": {
		encoderDecoder = PngEncoderDecoder{}
	}
	case "jpeg": {
		encoderDecoder = JpegEncoderDecoder{}
	}
	case "gif": {
		encoderDecoder = GifEncoderDecoder{}
	}
	default:
		return nil, image.ErrFormat
	}
	return encoderDecoder, nil
}


