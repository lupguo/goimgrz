package main

import (
	"flag"
	"girls/gir"
	"log"
)

// cmd variable store all the parameter where input from command line
var cmd struct {
	// http image
	url        string // the image's http(s) url to be resize
	urls       string // the image's http(s) urls to be resize
	crawlerUrl string // the crawler url used by girls, download http images and resize them

	// local image
	img     string // the local image file which to be resize
	imgs    string // the local image files  which to be resize
	scanDir string // scan the dir where image inside to be resize

	// necessary parameter
	dst   string // resize the image(s) where save to
	width uint   // resize image's width

	// advanced parameter
	quality  uint   // resize image quality percent
	height   uint   // resize image's height
	waterImg string // append an water image
	verbose  bool   // show detailed message

	// file filter
	name string // use a shell pattern to detect a matching image file name
	size string // use image file size to filter, the size of file's unit bytes, support `k` (1024 bytes),`M`(1024k)
}

func init() {
	// http url image
	flag.StringVar(&cmd.url, "url", "", "the image's http(s) url to be resize, image resource(url|urls|img|imgs|dir) at least need set one")
	flag.StringVar(&cmd.urls, "urls", "", "image's http(s) urls to be resize, separated by ','")
	flag.StringVar(&cmd.crawlerUrl, "crawler_url", "", "the crawler url used by girls download the http images and resize only matched image files")

	// local image
	flag.StringVar(&cmd.img, "img", "", "the local image file which to be resize")
	flag.StringVar(&cmd.imgs, "imgs", "", "local image files which to be resize, separated by ','")
	flag.StringVar(&cmd.scanDir, "dir", "", "scan the dir where image inside to be resize")

	// necessary parameter
	flag.StringVar(&cmd.dst, "dst", "/tmp", "the output dir where image after resize store")
	flag.UintVar(&cmd.width, "width", 300, "set resize image's width")

	// advanced parameter
	flag.UintVar(&cmd.quality, "quality", 75, "set resize image's quality percent")
	flag.UintVar(&cmd.height, "height", 0, "set resize image's height")
	flag.StringVar(&cmd.waterImg, "water_img", "", "append water image")
	flag.BoolVar(&cmd.verbose, "verbose", false, "append water image")

	// filter
	flag.StringVar(&cmd.name, "name", "*", "using shell pattern to filter image, like *.png")
	flag.StringVar(&cmd.size, "size", "", "using file size to filter image, like +200k")
}

func main() {
	flag.Parse()

	// create gir task
	gt := gir.NewGirTask(cmd.dst, cmd.width, cmd.height)

	// setting gir filter && relative parameters
	gt.SetFilter(gir.NewFilter(cmd.name, cmd.size))
	gt.SetVerbose(cmd.verbose)

	// report in background synchronously
	go gt.Report()

	// add gir task
	if cmd.img != "" {
		gt.Add(gir.ResTypeLocal, []byte(cmd.img))
	}
	if cmd.url != "" {
		gt.Add(gir.ResTypeHttp, []byte(cmd.url))
	}
	if cmd.imgs != "" {
		gt.AddFiles(cmd.imgs)
	}
	if cmd.urls != "" {
		gt.AddUrls(cmd.urls)
	}
	if cmd.scanDir != "" {
		gt.AddScanDir(cmd.scanDir)
	}

	// check gir task
	if gt.IsEmptyTask() {
		flag.PrintDefaults()
		log.Fatalln("resize image task is empty, check the parameters")
	}

	//  do gir task concurrently
	gt.DoResize()
}
