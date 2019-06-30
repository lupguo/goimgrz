package goimgrz

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// The byteSize is the smallest unit of storage, representing a byte
type byteSize uint64

// Growth based on human-readable unit formats, 1k=1024b,1M=1024k..
const (
	b byteSize = 1 << (iota * 10)
	k
	M
	G
	T
	P
)

// Human-readable unit formats, 1k=1024b
var sizeMap = map[string]byteSize{
	"c": b,
	"k": k,
	"M": M,
	"G": G,
	"T": T,
	"P": P,
}

// inlist checks whether s is in l
func inlist(s string, l []string) bool {
	for _, ss := range l {
		if ss == s {
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

// HumDS2Bytes parse human data size(like 10k,10M) to bytes unit length, like 1k = 1024bytes
func HumDS2Bytes(size string) (len uint64, err error) {
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

// HumDSLimit compare human data sizes and limit setting (like find -size).
// Param size and limit both support human data size format(bytes, kilobytes, megabytes), like 100k, 1M
//
// HumDSLimit("2000", "+1k") represent whether 1000 bytes >= 1*1024 bytes
// 	+/-: >= or < , 1k: 1k=1024bytes
func HumDSLimit(size string, limit string) (bool, error) {
	if len(size) ==0 || len(limit) == 0 {
		return false, errors.New("size or size limit is empty")
	}

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

	// compare size and limit length
	SizeLen, err := HumDS2Bytes(size)
	if err != nil {
		return false, err
	}
	limitLen, err := HumDS2Bytes(limit)
	if err != nil {
		return false, err
	}

	return compare(SizeLen, limitLen), nil
}

// GetFmtBasename get the new image basename based on the specified target image format and the name of the original image
func GetFmtBasename(basename, format string) string  {
	// assigned format
	switch format {
	case "png":
		fallthrough
	case "gif":
		basename = strings.Replace(basename, path.Ext(basename),"."+format, -1)
	case "jpeg", "jpg":
		basename = strings.Replace(basename, path.Ext(basename),".jpg", -1)
	}
	return basename
}
