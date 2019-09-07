package imgrz

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

type Category int

const (
	Local Category = iota
	Http
)

// Setting is image resize property setting like width,height,interp,quality,format
type Setting struct {
	Dst           string
	Width, Height uint
	Interp        uint
	Qty           int
	Format        string
}

// Image support using specified parameters resize image to dst
type Image interface {
	// SetResize set image resizing setting property, used for image resize
	SetResize(setting *Setting)

	// DoResize fit the specified resizing configuration resize image to destination
	// Fit the specified scaling configuration (such as size, interpolation function, path, quality),
	// perform image scaling, return the path of the scaled image, or any errors in the scaling process
	DoResize() (string, error)
}

// New create an image for a specific category
func NewImage(cate Category, ident string) Image {
	switch cate {
	case Local:
		return &LocImage{ident, nil}
	case Http:
		return &HttpImage{ident, nil}
	}
	return nil
}

// LocImage is an local image file whose image from local filesystem
type LocImage struct {
	Filename string
	Setting  *Setting
}

// HttpImage is an image from http(s) url
type HttpImage struct {
	Url     string
	Setting *Setting
}

// SetResize add image resizing property
func (img *LocImage) SetResize(setting *Setting) {
	img.Setting = setting
}

func (img *HttpImage) SetResize(setting *Setting) {
	img.Setting = setting
}

func (img *LocImage) DoResize() (save string, err error) {
	f, err := os.Open(img.Filename)
	if err != nil {
		return "", NewError(ErrOpenLocalImage, "open local image", err.Error())
	}
	defer f.Close()

	// resize local image
	s := img.Setting
	name := path.Base(img.Filename)
	return resizeImage(f, s.Dst, name, s.Format, s.Width, s.Height, s.Interp, s.Qty)
}

func (img *HttpImage) DoResize() (save string, err error) {
	resp, err := http.Get(img.Url)
	if err != nil {
		return "", NewError(ErrOpenHttpImage, "http get image", err.Error())
	}
	defer resp.Body.Close()

	// status check
	if resp.StatusCode != 200 {
		return "", NewError(ErrOpenHttpImage, "http error", fmt.Sprintf("Request Url(%s), StatusCode(%d), Status(%s)",
			img.Url, resp.StatusCode, resp.Status))
	}

	// resize url image
	s := img.Setting
	name := path.Base(resp.Request.URL.Path)
	return resizeImage(resp.Body, s.Dst, name, s.Format, s.Width, s.Height, s.Interp, s.Qty)
}

// resizeImage resize the image with the specified parameters, return the successful file or error
// support resize an image to dst with specified width and height, interpolation functions and quality setting
//
// about interpolation functions see more for details:
// https://github.com/nfnt/resize
func resizeImage(imageData io.Reader, dst, basename, format string, width, height, interp uint, quality int) (save string, err error) {
	// get new format name
	save = path.Clean(dst + "/" + GetFmtBasename(basename, format))

	// make sure dst dir exist
	if err := os.MkdirAll(path.Dir(save), 0755); err != nil {
		return save, NewError(ErrResize, "mkdir dst", err.Error())
	}

	// decode image
	img, origFmt, err := image.Decode(imageData)
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

	if format == "" {
		format = origFmt
	}
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
