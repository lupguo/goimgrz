package main

import (
	"os"
	"testing"
)

func TestResizeLocalImage(t *testing.T) {
	localImg := "./testdata/IMG_2489.JPG"
	t.Log(localImg)

	// storage resize to user cache dir
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(os.TempDir(), cacheDir)

	// resize test
	dst := os.TempDir() + "/resizeCache"
	save, err := ResizeLocalImage(localImg, dst, 800, 0)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(save)
}

func TestResizeHttpImage(t *testing.T) {
	urlImgs := []string{
		"https://cdn-images-1.medium.com/max/2400/1*pV0ZUbW1dURx-_YOWu1mzQ.png",
		"https://uidesign.gbtcdn.com/GB/image/2019/20190617_10732/New_B.jpg?imbypass=false",
		"https://uidesign.gbtcdn.com/GB/image/2019/20190612_10650/1190x420.gif?imbypass=false",
	}
	t.Log(urlImgs)

	// storage resize to user cache dir
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(os.TempDir(), cacheDir)

	// resize test
	dst := os.TempDir() + "/resizeCache"

	for _, urlImg := range urlImgs {
		save, err := ResizeHttpImage(urlImg, dst, 400, 0)
		if err != nil {
			t.Error("resize error: ",err)
			continue
		}
		t.Log("resize ok,",save)
	}

}