package gir

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

// GirImage used for indicate the image resource which will being resize
type GirImage struct {
	resType ResourceType
	data    []byte
}

// GirImage.Resize used for various resource image type
func (gi *GirImage) ResizeTo(dst string, width, height uint) (string, error) {
	switch gi.resType {
	case ResTypeLocal:
		return ResizeLocImg(string(gi.data), dst, width, height)
	case ResTypeHttp:
		return ResizeHttpImg(string(gi.data), dst, width, height)
	default:
		return "", NewError(ErrResImageType, "resource type", "cannot recognize the resize image's resource type")
	}
}

// resizeLocalImage resize local image file, and save it to dst dirname
func ResizeLocImg(filename string, dst string, width, height uint) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", NewError(ErrOpenLocalImage, "open local image", err.Error())
	}
	defer file.Close()

	return ResizeImage(file, dst, filename, width, height)
}

// resize http url image file, and save it to dst dirname
func ResizeHttpImg(imgUrl string, dst string, width, height uint) (string, error) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		return "", NewError(ErrOpenHttpImage, "http get image", err.Error())
	}
	defer resp.Body.Close()

	// status check
	if resp.StatusCode != 200 {
		return "", NewError(ErrOpenHttpImage, "http error", fmt.Sprintf("Request Url(%s), StatusCode(%d), Status(%s)",
			imgUrl, resp.StatusCode, resp.Status))
	}

	// resize image
	return ResizeImage(resp.Body, dst, resp.Request.URL.Path, width, height)
}

// resize an img to dst dirname, return image path or error
func ResizeImage(imageData io.Reader, dst, filename string, width, height uint) (string, error) {
	// make sure dst dir exist
	if err := os.MkdirAll(dst, 0755); err != nil {
		return "", NewError(ErrResize, "mkdir dst", err.Error())
	}

	// decode image
	img, format, err := image.Decode(imageData)
	if err != nil {
		return "", NewError(ErrResize, "decode image", err.Error())
	}

	// resize image
	newImg := resize.Resize(width, height, img, resize.NearestNeighbor)

	// encode image
	save := path.Clean(dst + "/" + path.Base(filename))
	newFile, _ := os.Create(save)
	switch format {
	case "png":
		err = png.Encode(newFile, newImg)
	case "jpeg":
		err = jpeg.Encode(newFile, newImg, &jpeg.Options{Quality:85})
	case "gif":
		err = gif.Encode(newFile, newImg, nil)
	default:
		return "", NewError(ErrResize, "error encode image format", format)
	}

	if err != nil {
		return "", NewError(ErrResize, "resize & encode image", err.Error())
	}
	defer newFile.Close()

	return save, nil
}
