package goimgrz

import (
	"log"
	"strings"
	"sync"
)

// saveRs save resize result
type saveRs struct {
	save string
	w    uint
	h    uint
}

// GirTask used for collect image resize, dispatching resize image task, got the save result or fail info from channel
type GirTask struct {
	filter  *Filter
	images  []Image
	chErr   chan error
	chSave  chan saveRs
	fin     chan bool
	dst     string
	width   uint
	height  uint
	interp  uint
	quality int
	verbose bool
}

// NewGirTask create an GirTas pointer
func NewGirTask(dst string, w, h, interp uint) *GirTask {
	return &GirTask{
		chErr:  make(chan error),
		chSave: make(chan saveRs),
		fin:    make(chan bool),
		dst:    dst,
		width:  w,
		height: h,
		interp: interp,
	}
}

// SetVerbose setting the task show detail message
func (gt *GirTask) SetVerbose(v bool) *GirTask {
	gt.verbose = v
	return gt
}

// SetFilter setting the task filter, using for filter no match info
func (gt *GirTask) SetFilter(f *Filter) *GirTask {
	gt.filter = f
	return gt
}

// Filter filter specified image
func (gt *GirTask) Filter(rt ResourceType, data []byte) error {
	// detect name
	if err := gt.filter.DetectName(string(data)); err != nil {
		return err
	}
	// detect size
	if err := gt.filter.DetectSize(rt, data); err != nil {
		return err
	}
	return nil
}

// Add use filtering information to filter files, and add image to task for resizing
func (gt *GirTask) Add(image Image) *GirTask {
	// filter by name or size
	gt.images = append(gt.images, image)
	return gt
}

func (gt *GirTask) AddImg(img string) *GirTask {
	// filter by name or size
	gt.images = append(gt.images, &LocImage{img})
	return gt
}

func (gt *GirTask) AddImgs(imgs string) *GirTask {
	for _, img := range strings.Split(imgs, ",") {
		gt.Add(&LocImage{img})
	}
	return gt
}

func (gt *GirTask) AddUrl(url string) *GirTask {
	gt.images = append(gt.images, &HttpImage{url})
	return gt
}

func (gt *GirTask) AddUrls(urls string) *GirTask {
	for _, url := range strings.Split(urls, ",") {
		gt.Add(&HttpImage{url})
	}
	return gt
}

// AddDirname specified dirname, scan images and add it to gir task, waiting for resize
func (gt *GirTask) AddScanDir(dir string) *GirTask {
	// scan dir get images
	imgs, err := GetImagesFromDir(dir)
	if err != nil {
		gt.chErr <- err
		return gt
	}
	for _, img := range imgs {
		gt.Add(&LocImage{img})
	}
	return gt
}

// EmptyTask return girTask whether is empty
func (gt *GirTask) EmptyTask() bool {
	return len(gt.images) == 0
}

// Report synchronously report success or fail result in background, when gir task is finish
func (gt *GirTask) Report() {
	wg := sync.WaitGroup{}

	// report success
	wg.Add(1)
	go func() {
		defer wg.Done()
		for rs := range gt.chSave {
			log.Printf("resize ok: %s (inputW=%d,inputH=%d)\n", rs.save, rs.w, rs.h)
		}
	}()

	// report fail
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range gt.chErr {
			log.Println("resize fail:", err)
		}
	}()

	wg.Wait()
	gt.fin <- true
}

// ResizeImages concurrency resize image in it's GirImage slice
func (gt *GirTask) Run() {
	// concurrency task working
	wg := sync.WaitGroup{}

	// doResize resize an input image
	doResize := func(img Image) {
		defer wg.Done()
		save, err := img.ResizeTo(gt.dst, gt.width, gt.height, gt.interp, gt.quality)
		if err != nil {
			gt.chErr <- err
			return
		}
		gt.chSave <- saveRs{
			save,
			gt.width,
			gt.height,
		}
	}

	// dispatch task
	for _, img := range gt.images {
		wg.Add(1)
		go doResize(img)
	}

	// wait for all task finished, then close send chan to stop the report goroutine(like three handshake)
	wg.Wait()
	close(gt.chErr)
	close(gt.chSave)

	// wait fin msg
	<-gt.fin
}
