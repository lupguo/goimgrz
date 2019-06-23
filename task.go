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
	images  []*GirImage
	chErr   chan error
	chSave  chan saveRs
	fin     chan bool
	dst     string
	width   uint
	height  uint
	verbose bool
}

// NewGirTask create an GirTas pointer
func NewGirTask(dst string, w, h uint) *GirTask {
	return &GirTask{
		images: []*GirImage{},
		chErr:  make(chan error),
		chSave: make(chan saveRs),
		fin:    make(chan bool),
		dst:    dst,
		width:  w,
		height: h,
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

// Add using filter pickup specified image to gir task, waiting for resize
func (gt *GirTask) Add(rt ResourceType, data []byte) *GirTask {
	if err := gt.filter.DetectName(string(data)); err != nil {
		if gt.verbose {
			gt.chErr <- err
		}
	} else if err := gt.filter.DetectSize(rt, data); err != nil {
		if gt.verbose {
			gt.chErr <- err
		}
	} else {
		gt.images = append(gt.images, &GirImage{
			resType: rt,
			data:    data,
		})
	}
	return gt
}

// AddUrls specified urls, add url image to gir task, waiting for resize
func (gt *GirTask) AddUrls(urls string) *GirTask {
	for _, url := range strings.Split(urls, ",") {
		gt.Add(ResTypeHttp, []byte(url))
	}
	return gt
}

// AddFiles specified local files, add to gir task, waiting for resize
func (gt *GirTask) AddFiles(imgs string) *GirTask {
	for _, img := range strings.Split(imgs, ",") {
		gt.Add(ResTypeLocal, []byte(img))
	}
	return gt
}

// AddDirname specified dirname, scan images and add it to gir task, waiting for resize
func (gt *GirTask) AddScanDir(dir string) *GirTask {
	if imgs, err := GetImagesFromDir(dir); err != nil {
		gt.chErr <- err
	} else {
		for _, img := range imgs {
			gt.Add(ResTypeLocal, []byte(img))
		}
	}
	return gt
}

// IsEmpty check girTask whether is empty
func (gt *GirTask) IsEmptyTask() bool {
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
func (gt *GirTask) DoResize() {
	// concurrency task working
	wg := sync.WaitGroup{}

	// do resize function
	doResize := func(gi *GirImage) {
		defer wg.Done()
		save, err := gi.ResizeTo(gt.dst, gt.width, gt.height)
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
	for _, gti := range gt.images {
		wg.Add(1)
		go doResize(gti)
	}

	// wait for all task finished, then close send chan to stop the report goroutine(like three handshake)
	wg.Wait()
	close(gt.chErr)
	close(gt.chSave)

	// wait fin msg
	<-gt.fin
}
