package goimgrz

import "testing"

func TestGirTask_DoResize(t *testing.T) {
	// http image
	urlImgs := []string{
		"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png",
		"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false",
		"https://uidesign.gbtcdn.com/GB/image/2019/20190612_10650/1190x420.gif?imbypass=false",
	}
	// local file
	localImgs  := []string{
		"./testdata/IMG_2489.JPG",
	}

	gt := NewGirTask("/tmp", 400, 0)

	// new http gir task
	for _, u := range urlImgs {
		gt.Add(ResTypeHttp, []byte(u))
	}
	// new local gir task
	for _, l := range localImgs {
		gt.Add(ResTypeLocal, []byte(l))
	}

	// gir do resize task
	gt.DoResize()
}