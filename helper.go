package goimgrz

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ByteSize
type ByteSize uint64

// Human readable size, like 10k
const (
	b ByteSize = 1 << (iota * 10)
	k
	M
)

var sizeMap = map[string]ByteSize{
	"b": b,
	"k": k,
	"M": M,
}

// inlist check str whether in list
func inlist(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}
	return false
}

// GetImagesFromDir gets slice of the image file and iterates through the selected directory and its subdirectories.
func GetImagesFromDir(dirname string) ([]string, error) {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		return nil, NewError(ErrGetLocalDirImages, "dir is not exist "+dirname, err.Error())
	}

	var imgs []string
	exts := []string{".png", ".jpg", ".jpeg", ".gif", ".webp"}

	// extract local image which contain specified extension name
	walkFn := func(p string, info os.FileInfo, err error) error {
		// append specified extension
		if ext := strings.ToLower(filepath.Ext(p)); inlist(ext, exts) {
			imgs = append(imgs, p)
		}
		return nil
	}
	err := filepath.Walk(dirname, walkFn)
	if err != nil {
		return nil, NewError(ErrGetLocalDirImages, "filepath walk error "+dirname, err.Error())
	}

	return imgs, nil
}

// ParseHumanDataSize parse human data size(like 10k,10M) to bytes
func ParseHumanDataSize(size string) (uint64, error) {
	unit := b
	var digits string
forloop:
	for _, c := range size {
		switch {
		case c >= '0' && c <= '9':
			digits += string(c)
		case c == 'k' || c == 'M' || c == 'G':
			unit = sizeMap[string(c)]
		default:
			break forloop
		}
	}
	if num, err := strconv.ParseInt(digits, 10, 32); err != nil {
		return 0, err
	} else {
		return uint64(num) * uint64(unit), nil
	}
}

// SatisfyHumanSize compare data sizes (bytes, kilobytes, megabytes), human readable sizes, parsing, compare
func SatisfyHumanSize(size string, limit string) (bool, error) {
	var compare func(a, b uint64) bool
	switch {
	case limit[0] == '+':
		limit = limit[1:]
		compare = func(a, b uint64) bool {
			return a >= b
		}
	case limit[0] == '-':
		limit = limit[1:]
		compare = func(a, b uint64) bool {
			return a < b
		}
	default:
		compare = func(a, b uint64) bool {
			return a == b
		}
	}

	nsize, err := ParseHumanDataSize(size)
	if err != nil {
		return false, err
	}
	nlimit, err := ParseHumanDataSize(limit)
	if err != nil {
		return false, err
	}

	return compare(nsize, nlimit), nil
}
