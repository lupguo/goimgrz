package girls

import "testing"

func TestGetLocalImages(t *testing.T) {
	imgs, err := GetImagesFromDir("./testdata", "*")
	if err != nil {
		t.Error(err)
	}
	t.Log(imgs)
}
