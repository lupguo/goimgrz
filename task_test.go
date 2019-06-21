package girls

import "testing"

func TestGirTask_ResizeImages(t *testing.T) {
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

	girTask := NewGirTask()
	outDir := "/tmp/girls"

	// new http gir task
	for _, u := range urlImgs {
		girTask.Add(ResTypeHttp, []byte(u), outDir, 300, 0)
	}
	// new local gir task
	for _, l := range localImgs {
		girTask.Add(ResTypeLocal, []byte(l), outDir, 400, 0)
	}

	// gir do resize task
	girTask.DoResize()
}