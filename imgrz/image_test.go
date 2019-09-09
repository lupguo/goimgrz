package imgrz

import (
	"os"
	"testing"
)

func TestLocImage_ResizeTo(t *testing.T) {
	dst := os.TempDir() + "/goimgrz/imgs"
	test := []struct {
		img     string
		Setting *Setting
		want    bool
	}{
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 0, Qty: 75}, true},
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 1, Qty: 75}, true},
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 2, Qty: 75, Width: 800, Height: 200}, true},
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 3, Qty: 75, Width: 400}, true},
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 4, Qty: 75}, true},
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 5, Qty: 75, Format: "jpeg"}, true},
		{"./testdata/gopher2018.png", &Setting{Dst: dst, Interp: 0, Qty: 75, Format: "gif"}, true},
		{"./testdata/IMG_2489.JPG", &Setting{Dst: dst, Interp: 0, Format: "png"}, true},
		{"./testdata/not_exits.jpg", &Setting{}, false},
	}

	for _, tt := range test {
		img := &LocImage{tt.img, tt.Setting}
		save, err := img.DoResize()
		if err != nil {
			t.Logf("img=%s, setting=%+v, got=%t, want=%t, err: %s", tt.img, *tt.Setting, false, tt.want, err)
			if tt.want == true {
				t.Fail()
			}
			continue
		}
		t.Log("resize ok:", save)
	}
}

func TestHttpImage_ResizeTo(t *testing.T) {
	dst := os.TempDir() + "/goimgrz/imgs"
	test := []struct {
		url     string
		Setting *Setting
		want    bool
	}{
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 0, Qty: 75}, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 1, Qty: 75}, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 2, Qty: 75, Width: 800, Height: 200}, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 3, Qty: 75, Width: 400}, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 4, Qty: 75}, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 5, Qty: 75, Format: "jpeg"}, true},
		{"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png", &Setting{Dst: dst, Interp: 0, Qty: 75, Format: "gif"}, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", &Setting{Dst: dst, Interp: 0, Format: "png"}, true},
		{"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false", &Setting{}, false},
	}

	for _, tt := range test {
		img := &LocImage{tt.url, tt.Setting}
		save, err := img.DoResize()
		if err != nil {
			t.Logf("img=%s, setting=%+v, got=%t, want=%t, err: %s", tt.url, *tt.Setting, false, tt.want, err)
			if tt.want == true {
				t.Fail()
			}
			continue
		}
		t.Log("resize ok:", save)
	}
}
