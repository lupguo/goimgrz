package goimgrz

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path"
)

// gir image data resource type maybe local image, http url, or stdin base64_encode
type ResourceType int

const (
	ResTypeLocal ResourceType = iota // local image
	ResTypeHttp                      // http url image
)

// Image support using specified parameters resize image to dst
type Image interface {
	ResizeTo(dst string, w, h, interp uint, qty int) (string, error)
}

// LocImage is an local image file whose image from local filesystem
type LocImage struct {
	filename string
}

// HttpImage is an image from http(s) url
type HttpImage struct {
	url string
}

func (img *LocImage) ResizeTo(dst string, w, h, interp uint, qty int) (save string, err error) {
	f, err := os.Open(img.filename)
	if err != nil {
		return "", NewError(ErrOpenLocalImage, "open local image", err.Error())
	}
	defer f.Close()

	// resize local image
	return ResizeImage(f, dst, path.Base(img.filename), w, h, interp, qty)
}

func (img *HttpImage) ResizeTo(dst string, w, h, interp uint, qty int) (save string, err error) {
	resp, err := http.Get(img.url)
	if err != nil {
		return "", NewError(ErrOpenHttpImage, "http get image", err.Error())
	}
	defer resp.Body.Close()

	// status check
	if resp.StatusCode != 200 {
		return "", NewError(ErrOpenHttpImage, "http error", fmt.Sprintf("Request Url(%s), StatusCode(%d), Status(%s)",
			img.url, resp.StatusCode, resp.Status))
	}

	// resize url image
	filename := path.Base(resp.Request.URL.Path)
	return ResizeImage(resp.Body, dst, filename, w, h, interp, qty)
}

// ResizeImage resize the image with the specified parameters, return the successful file or error
// support resize an image to dst with specified width and height, interpolation functions and quality setting
//
// about interpolation functions see more for details:
// https://github.com/nfnt/resize
func ResizeImage(imageData io.Reader, dst, filename string, width, height, interp uint, quality int) (save string, err error) {
	save = path.Clean(dst + "/" + filename)

	// make sure dst dir exist
	if err := os.MkdirAll(path.Dir(save), 0755); err != nil {
		return save, NewError(ErrResize, "mkdir dst", err.Error())
	}

	// decode image
	img, format, err := image.Decode(imageData)
	if err != nil {
		return save, NewError(ErrResize, "image decode", err.Error())
	}

	// resize image
	newImg := resize.Resize(width, height, img, resize.InterpolationFunction(interp))

	// encode image
	newFile, err := os.Create(save)
	if err != nil {
		return save, NewError(ErrResize, "image create", err.Error())
	}
	defer newFile.Close()

	switch format {
	case "png":
		err = png.Encode(newFile, newImg)
	case "jpeg":
		err = jpeg.Encode(newFile, newImg, &jpeg.Options{Quality: quality})
	case "gif":
		err = gif.Encode(newFile, newImg, nil)
	default:
		return save, NewError(ErrResize, "image format not support", format)
	}

	if err != nil {
		return "", NewError(ErrResize, "image encode", err.Error())
	}
	return save, nil
}
