package main

import (
	"flag"
	"log"
)

// Cmdline variable store all the parameter where input from command line
var Cmdline struct {
	// http image
	url        string // the image's http(s) url which to be resize
	urls       string // the image's http(s) urls
	crawlerUrl string // the crawler url used by girls, download http images and resize them

	// local image
	dirname  string // the folder where image store which to be resize
	filename string // the image which to be resize

	// necessary parameter
	dst   string // resize the image(s) where save to
	width uint   // resize image's width

	// advanced parameter
	height   uint   // resize image's height
	waterImg string // append an water image
	match    string // only matched image will to be resize
	quality  uint   // resize image quality percent
}

func init() {
	// http image
	flag.StringVar(&Cmdline.url, "url", "", "the http(s) url of image which to be resize, "+
		"image resource(url|urls|dirname|filename) at least need set one")
	flag.StringVar(&Cmdline.urls, "urls", "", "the http(s) urls of image which to be resize, separated by ','")
	flag.StringVar(&Cmdline.crawlerUrl, "crawler_url", "", "the crawler url used by girls download the http images and resize only matched image files")

	// local image
	flag.StringVar(&Cmdline.dirname, "scan_dir", "", "scan the dir where image inside to be resize")
	flag.StringVar(&Cmdline.filename, "img", "", "the local image which to be resize")

	// necessary parameter
	flag.StringVar(&Cmdline.dst, "dst", "/tmp", "the output dir where image after resize store")
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
	gt := NewGirTask()
	dst := Cmdline.dst
	w := Cmdline.width
	h := Cmdline.height

	// resize http url image
	if Cmdline.url != "" {
		gt.Add(ResTypeHttp, []byte(Cmdline.url), dst, w, h)
	}

	// resize http url image list
	if Cmdline.urls != "" {
		gt.AddUrls(Cmdline.urls, Cmdline.match, dst, w, h)
	}

	// resize signal local image file
	if Cmdline.filename != "" {
		gt.Add(ResTypeLocal, []byte(Cmdline.filename), dst, w, h)
	}

	// resize all image in directory and sub subdirectory
	if Cmdline.dirname != "" {
		gt.AddDirname(Cmdline.dirname, Cmdline.match, dst, w, h)
	}

	// check workload
	if gt.IsEmpty() {
		log.Fatalln("resize image task is empty, check the parameters")
	}

	//  do resize task concurrently
	gt.DoResize()
}
