package girls

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

// resize local image file, and save it to dst dirname
func ResizeLocalImage(filename string, dst string, width, height uint) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", NewError(ErrOpenLocalImage, "open local image", err.Error())
	}
	defer file.Close()

	return resizeImage(file, dst, filename, width, height)
}

// resize http url image file, and save it to dst dirname
func ResizeHttpImage(imgUrl string, dst string, width, height uint) (string, error) {
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
	return resizeImage(resp.Body, dst, resp.Request.URL.Path, width, height)
}

// resize an img to dst dirname, return image path or error
func resizeImage(imageData io.Reader, dst, filename string, width, height uint) (string, error) {
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
		err = jpeg.Encode(newFile, newImg, &jpeg.Options{85})
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
