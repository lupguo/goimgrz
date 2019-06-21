package girls

import "flag"

// Cmdline variable store all the parameter where input from command line
var Cmdline struct {
	// http image
	url        string	// the image's http(s) url which to be resize
	crawlerUrl string	// the crawler url used by girls, download http images and resize them

	// local image
	filename string // the image which to be resize
	dirname  string // the folder where image store which to be resize

	// necessary parameter
	dst    string // resize image where store
	width  int    // resize image's width

	// advanced parameter
	height int    // resize image's height
	waterImg string // append an water image
	match    string // only matched image will to be resize
	quality  int    // resize image quality percent
}

func init() {
	// http image
	flag.StringVar(&Cmdline.url, "url", "", "the image's http(s) url which to be resize")
	flag.StringVar(&Cmdline.crawlerUrl, "crawler_url", "", "the crawler url used by girls, " +
		"download http images and resize them")

	// local image
	flag.StringVar(&Cmdline.filename, "img", "", "the image which to be resize")
	flag.StringVar(&Cmdline.dirname, "output_dir", "", "the folder where image store which to be resize")

	// necessary parameter
	flag.StringVar(&Cmdline.filename, "dst", "/tmp", "resize image where store")
	flag.IntVar(&Cmdline.width, "width", 300, "resize image's height")

	// advanced parameter
	flag.IntVar(&Cmdline.height, "height", 0, "resize image's width")
	flag.StringVar(&Cmdline.waterImg, "water_img", "", "append an water image")
	flag.StringVar(&Cmdline.match, "match", "*", "only matched image will to be resize")
	flag.IntVar(&Cmdline.quality, "quality", 75, "resize image quality percent")

	flag.Parse()
}
