package girls

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// check str whether in str list
func Inlist(str string, strList []string) bool {
	for _, s := range strList {
		if str == s {
			return true
		}
	}
	return false
}

// get image files from local path
func GetImagesFromDir(dirname, pattern string) ([]string, error) {
	var imgs []string
	exts := []string{".png", ".jpg", ".jpeg", ".gif", ".webp"}

	// extract local image which contain specified extension name
	walkFn := func(p string, info os.FileInfo, err error) error {
		// path match
		mtch, err := path.Match(pattern, info.Name())
		if err != nil || mtch == false{
			return err
		}
		// append specified extension
		if ext := strings.ToLower(filepath.Ext(p)); Inlist(ext, exts) {
			imgs = append(imgs, p)
		}
		return nil
	}
	err := filepath.Walk(dirname, walkFn)
	if err != nil {
		return nil, NewError(ErrGetLocalDirImages, "filepath walk error"+dirname, err.Error())
	}
	return imgs, nil
}