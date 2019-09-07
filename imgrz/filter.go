package imgrz

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

// Filter is an filter pick up the image matching the condition
type Filter struct {
	pattern string // shell pattern name, eg */*.png
	limit   string // like +n, -n, eg +100k, filter larger than 100k image file
}

// Create an filter
func NewFilter(pattern, size string) *Filter {
	return &Filter{
		pattern,
		size,
	}
}

// DetectName check the input url or image file whether match the filter pattern
func (f *Filter) DetectName(image Image) (bool, error) {
	// no limit pattern filter
	if f.pattern == "*" {
		return true, nil
	}
	// image name
	var name string
	switch img := image.(type) {
	case *LocImage:
		name = img.Filename
	case *WebImage:
		name = img.Url
	default:
		return false, NewError(ErrDetectName, "image type", "unsupported image type")
	}
	// empty check
	if len(strings.Trim(name, " ")) == 0 {
		return false, NewError(ErrDetectName, "empty name", "image name is empty")
	}
	// shell pattern check
	if mch, err := path.Match(f.pattern, name); err != nil {
		return false, NewError(ErrDetectName, "pattern error", err.Error())
	} else if mch == false {
		return false, NewError(ErrDetectName, "not match", fmt.Sprintf("pattern:%s, name:%s", f.pattern, name))
	}

	return true, nil
}

// DetectSize check the input url
func (f *Filter) DetectSize(image Image) (bool, error) {
	// no limit size filter
	if len(strings.Trim(f.limit, " ")) == 0 {
		return true, nil
	}
	// filter check
	var size int64
	switch img := image.(type) {
	case *LocImage:
		fi, err := os.Stat(img.Filename)
		if err != nil {
			return false, NewError(ErrDetectSize, "detect local image size error", err.Error())
		}
		size = fi.Size()
	case *WebImage:
		resp, err := http.Head(img.Url)
		if err != nil {
			return false, NewError(ErrDetectSize, "detect http image size error", err.Error())
		}
		size = resp.ContentLength
	default:
		return false, NewError(ErrDetectName, "image type", "unsupported image type")
	}

	// human size detect
	ok, err := HumDSLimit(strconv.FormatInt(size, 10), f.limit)
	if err != nil {
		return false, NewError(ErrDetectSize, "detect size error", err.Error())
	}
	// not
	if ok == false {
		return false, NewError(ErrDetectSize, "not satisfy", fmt.Sprintf("limit size:%s, file size:%d", f.limit, size))
	}

	return true, nil
}
