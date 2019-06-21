package main

import (
	"flag"
	"girls"
	"log"
)

// Cmdline variable store all the parameter where input from command line
var Cmdline struct {
	// http image
	url        string // the image's http(s) url which to be resize
	crawlerUrl string // the crawler url used by girls, download http images and resize them

	// local image
	dirname  string // the folder where image store which to be resize
	filename string // the image which to be resize

	// necessary parameter
	dst   string // resize image where store
	width uint   // resize image's width

	// advanced parameter
	height   uint   // resize image's height
	waterImg string // append an water image
	match    string // only matched image will to be resize
	quality  uint   // resize image quality percent
}

func init() {
	// http image
	flag.StringVar(&Cmdline.url, "img_url", "", "the http(s) url  of image which to be resize, image resource(url|dirname|filename) at least need set one")
	flag.StringVar(&Cmdline.crawlerUrl, "crawler_url", "", "the crawler url used by girls download the http images and resize only matched image files")

	// local image
	flag.StringVar(&Cmdline.dirname, "dir", "", "the dir path where image inside to be resize")
	flag.StringVar(&Cmdline.filename, "local_img", "", "the local image which to be resize")

	// necessary parameter
	flag.StringVar(&Cmdline.dst, "dst", "/tmp", "the output dir where image after  resize  store")
	flag.UintVar(&Cmdline.width, "width", 300, "set resize image's width")

	// advanced parameter
	flag.UintVar(&Cmdline.height, "height", 0, "set resize image's height")
	flag.StringVar(&Cmdline.waterImg, "water_img", "", "append water image")
	flag.StringVar(&Cmdline.match, "match", "*", "only matched image will to be resize")
	flag.UintVar(&Cmdline.quality, "quality", 75, "set resize image's quality percent")
}

func main() {
	flag.Parse()

	// create gir task
	gt := girls.NewGirTask()

	// resize http url image
	dst := Cmdline.dst
	w := Cmdline.width
	h := Cmdline.height
	if Cmdline.url != "" {
		gt.Add(girls.ResTypeHttp, []byte(Cmdline.url), dst, w, h)
	}

	// resize signal local image file
	if Cmdline.filename != "" {
		gt.Add(girls.ResTypeLocal, []byte(Cmdline.filename), dst, w, h)
	}

	// resize all image in directory and sub subdirectory
	if Cmdline.dirname != "" {
		imgs, _ := girls.GetImagesFromDir(Cmdline.dirname, Cmdline.match)

		for _, img := range imgs {
			gt.Add(girls.ResTypeLocal, []byte(img), dst, w, h)
		}
	}

	// check workload
	if gt.IsEmpty() {
		flag.PrintDefaults()
		log.Fatalln("resize image task is empty, check the parameters")
	}

	//  do resize task concurrently
	gt.DoResize()
}
