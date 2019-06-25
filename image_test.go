package goimgrz

import (
	"os"
	"testing"
)

func TestLocImage_ResizeTo(t *testing.T) {
	test := []struct {
		img    string
		interp uint
		qty    int
		want   bool
	}{
		{"./testdata/gopher2018.png", 0, 75, true},
		{"./testdata/gopher2018.png", 1, 60, true},
		{"./testdata/gopher2018.png", 2, 75, true},
		{"./testdata/gopher2018.png", 3, 60, true},
		{"./testdata/gopher2018.png", 4, 75, true},
		{"./testdata/gopher2018.png", 5, 60, true},
		{"./testdata/IMG_2489.JPG", 0, 60, true},
		{"./testdata/IMG_2489.JPG", 1, 75, true},
		{"./testdata/IMG_2489.JPG", 2, 60, true},
		{"./testdata/IMG_2489.JPG", 3, 75, true},
		{"./testdata/IMG_2489.JPG", 4, 60, true},
		{"./testdata/IMG_2489.JPG", 5, 75, true},
		{"./testdata/not_exits.jpg", 0, 75, false},
	}

	dst := os.TempDir() + "/goimgrz/imgs"
	for _, tt := range test {
		img := &LocImage{tt.img}
		save, err := img.ResizeTo(dst, 300, 0, 0, 60)
		if err != nil {
			t.Logf("img=%s, interp=%d, qty=%d, got=%t, want=%t, err: %s", tt.img, tt.interp, tt.qty, false, tt.want,err)
			if tt.want == true {
				t.Fail()
			}
			continue
		}
		t.Log("resize ok:", save)
	}
}

func TestHttpImage_ResizeTo(t *testing.T) {
	test := []struct {
		url    string
		interp uint
		qty    int
		want   bool
	}{
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", 0, 75, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", 1, 60, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", 2, 75, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", 3, 60, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", 4, 75, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", 5, 60, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", 0, 60, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", 1, 75, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", 2, 60, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", 3, 75, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", 4, 60, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", 5, 75, true},
		{"https://notexisturl.com/testdata/not_exits.jpg", 0, 75, false},
	}

	dst := os.TempDir() + "/goimgrz/urls"
	for _, tt := range test {
		img := &HttpImage{tt.url}
		save, err := img.ResizeTo(dst, 300, 0, 0, 60)
		if err != nil {
			t.Logf("img=%s, interp=%d, qty=%d, got=%t, want=%t, err: %s", tt.url, tt.interp, tt.qty, false, tt.want,err)
			if tt.want != false {
				t.Fail()
			}
			continue
		}
		t.Log("resize ok:", save)
	}
}
