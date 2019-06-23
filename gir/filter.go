package gir

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

// Filter is an filter pick up the image matching the condition
type Filter struct {
	pattern string // shell pattern name
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
func (f *Filter) DetectName(name string) error {
	// empty check
	if len(strings.Trim(name, " ")) == 0 {
		return NewError(ErrDetectName, "task filter detect name pattern error", "input data is empty")
	}

	// shell pattern check
	if mch, err := path.Match(f.pattern, name); err != nil {
		return NewError(ErrDetectName, "task filter detect name pattern error", err.Error())
	} else if mch == false {
		return NewError(ErrDetectName, "task filter detect name pattern not match", fmt.Sprintf("pattern:%s, name:%s", f.pattern, name))
	}

	return nil
}

// DetectSize check the input url
func (f *Filter) DetectSize(rt ResourceType, data []byte) error {
	// filter not check when limit is empty
	if len(strings.Trim(f.limit, " ")) == 0 {
		return nil
	}

	// filter check
	var size int64
	switch rt {
	case ResTypeHttp:
		resp, err := http.Head(string(data))
		if err != nil {
			return err
		}
		size = resp.ContentLength
	case ResTypeLocal:
		fi, err := os.Stat(string(data))
		if err != nil {
			return NewError(ErrDetectSize, "task filter detect size error", err.Error())
		}
		size = fi.Size()
	}
	if ok, err := SatisfyHumanSize(string(size), f.limit); err != nil {
		return NewError(ErrDetectSize, "task filter detect size error", err.Error())
	} else if !ok {
		return NewError(ErrDetectSize, "task filter detect size not match", fmt.Sprintf("size:%s, file size:%d", f.limit, size))
	}
	return nil
}
