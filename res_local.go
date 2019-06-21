package girls

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// get image files from local path
func GetLocalDirImages(dirname, pattern string) ([]string, error) {
	var imgs []string
	exts := []string{".png", ".jpg", ".jpeg", ".gif", ".webp"}

	// extract local image which contain specified extension name
	walkFn := func(p string, info os.FileInfo, err error) error {
		// path match
		mtch, err := path.Match(pattern, info.Name())
		if err != nil || mtch == false{
			return nil
		}
		// append specified extension
		if ext := strings.ToLower(filepath.Ext(p)); inlist(ext, exts) {
			imgs = append(imgs, info.Name())
		}
		return nil
	}
	err := filepath.Walk(dirname, walkFn)
	if err != nil {
		return nil, NewError(ErrGetLocalDirImages, "filepath walk error"+dirname, err.Error())
	}
	return imgs, nil
}
