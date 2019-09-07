package main

import (
	"flag"
	"fmt"
	"github.com/tkstorm/goimgrz/imgrz"
	"log"
	"os"
)

const version = "0.0.1"

// cmd variable store all the parameter where input from command line
var cmd struct {
	// image source (url images or local images)
	url     string // the image's http(s) url to be resize
	urls    string // the image's http(s) urls to be resize
	img     string // the local image file which to be resize
	imgs    string // the local image files  which to be resize
	scanDir string // scan the dir where image inside to be resize
	// image filter
	name string // use a shell pattern to detect a matching image file name
	size string // use image file size to filter, the size of file's unit bytes, support `k` (1024 bytes),`M`(1024k)
	// resize parameter
	format  string // resize to image format
	quality int    // resize image quality percent
	interp  uint   // the provided interpolation functions support
	width   uint   // resize image's width
	height  uint   // resize image's height
	// image save
	dst string // resize the image(s) where save to
	// debug
	verbose bool // show detailed message
}

func init() {
	// http url image
	flag.StringVar(&cmd.url,"url", "", "")
	flag.StringVar(&cmd.urls, "urls", "", "")
	flag.StringVar(&cmd.img, "img", "", "")
	flag.StringVar(&cmd.imgs, "imgs", "", "")
	flag.StringVar(&cmd.scanDir, "scdir", "", "")
	// filter
	flag.StringVar(&cmd.name, "name", "*", "")
	flag.StringVar(&cmd.size, "size", "", "")
	// resize setting
	flag.UintVar(&cmd.width, "w", 0, "")
	flag.UintVar(&cmd.height, "h", 0, "")
	flag.StringVar(&cmd.format, "cfmt", "", "")
	flag.UintVar(&cmd.interp, "itp", 0, "")
	flag.IntVar(&cmd.quality, "qty", 75, "")
	// resize saving
	flag.StringVar(&cmd.dst, "dst", "/tmp/goimgrz", "")
	flag.BoolVar(&cmd.verbose, "v", false, "")
}

var usage = `Usage: goimgrz [options...]

Goimgrz version %s, a simple image resizing tool that supports web or local images.
https://github.com/tkstorm/goimgrz

Options:
Resize image source:
-url	Web images to be resize, image source(url|urls|img|imgs|dir) at least need set one
-urls	Multiple web images to be resize, separated by ','
-img	Local images to be resize	
-imgs	Multiple local images to be resize, separated by ','
-scdir	Scanned file image file directory

Filter:
-name	Using shell pattern to filter image, like *.png
-size	Using file size to filter image, like +200k

Resize Setting:
-w	Set resize image's width, default width is 0 represent origin image
-h	Set resize image's height, default height is 0 represent origin image 
-cfmt	Convert image output format(jpg|png|gif)
-itp	The provided interpolation functions support (from fast to slow execution time).
	0:NearestNeighbor,1:Bilinear,2:Bicubic,3:MitchellNetravali,4:Lanczos2,5:Lanczos3
-qty 	Set resize image's quality percent

Image Saving:
-dst	The output dir

Other:
-v	Verbose message
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, version))
	}
	flag.Parse()

	// create gir task
	setting := &imgrz.Setting{
		Dst:    cmd.dst,
		Qty:    cmd.quality,
		Width:  cmd.width,
		Height: cmd.height,
		Interp: cmd.interp,
		Format: cmd.format,
	}
	gt := imgrz.NewTask(setting)

	// setting gir filter && relative parameters
	gt.SetFilter(imgrz.NewFilter(cmd.name, cmd.size))
	gt.SetVerbose(cmd.verbose)

	// report in background synchronously
	go gt.Report()

	// add gir task
	if cmd.img != "" {
		gt.AddImg(cmd.img)
	}
	if cmd.imgs != "" {
		gt.AddImgs(cmd.imgs)
	}
	if cmd.url != "" {
		gt.AddUrl(cmd.url)
	}
	if cmd.urls != "" {
		gt.AddUrls(cmd.urls)
	}
	if cmd.scanDir != "" {
		gt.AddScanDir(cmd.scanDir)
	}

	// check gir task
	if gt.EmptyTask() {
		log.Fatalln("resize image task is empty, check the parameters or try option -h see more info.")
	}

	//  do gir task concurrently
	gt.Run()
}
