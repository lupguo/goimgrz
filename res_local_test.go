package girls

import "testing"

func TestGetLocalImages(t *testing.T) {
	imgs, err := GetLocalDirImages("./testdata", "*")
	if err != nil {
		t.Error(err)
	}
	t.Log(imgs)
}
