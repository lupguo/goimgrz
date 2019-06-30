package goimgrz

import (
	"log"
	"strings"
	"sync"
)

// saveRs save resize result
type saveRs struct {
	save   string
	w      uint
	h      uint
	interp uint
	qty    int
	format string
}

// Task used for collect image resize, dispatching resize image task, got the save result or fail info from channel
type Task struct {
	filter  *Filter
	images  []Image
	chErr   chan error
	chSave  chan saveRs
	fin     chan struct{}
	setting *Setting
	verbose bool
}

// NewTask create an GirTas pointer
func NewTask(setting *Setting) *Task {
	return &Task{
		filter:  new(Filter),
		chErr:   make(chan error),
		chSave:  make(chan saveRs),
		fin:     make(chan struct{}),
		setting: setting,
	}
}

// SetVerbose setting the task show detail message
func (gt *Task) SetVerbose(v bool) *Task {
	gt.verbose = v
	return gt
}

// SetFilter setting the task filter, using for filter no match info
func (gt *Task) SetFilter(f *Filter) *Task {
	gt.filter = f
	return gt
}

// Filter filter specified image
func (gt *Task) Filter(image Image) error {
	// detect name
	if ok, err := gt.filter.DetectName(image); !ok {
		return err
	}
	// detect size
	if ok, err := gt.filter.DetectSize(image); !ok {
		return err
	}

	return nil
}

// Add use filtering information to filter files, and add image to task for resizing
func (gt *Task) Add(image Image) *Task {
	// filter by name or size
	if ok, err := gt.filter.DetectName(image); !ok {
		if gt.verbose {
			gt.chErr <- err
		}
		return gt
	}
	if ok, err := gt.filter.DetectSize(image); !ok {
		if gt.verbose {
			gt.chErr <- err
		}
		return gt
	}

	// task image resize setting
	image.SetResize(gt.setting)

	gt.images = append(gt.images, image)
	return gt
}

func (gt *Task) AddImg(img string) *Task {
	gt.Add(NewImage(Local, img))
	return gt
}

func (gt *Task) AddImgs(imgs string) *Task {
	for _, img := range strings.Split(imgs, ",") {
		gt.Add(NewImage(Local, img))
	}
	return gt
}

func (gt *Task) AddUrl(url string) *Task {
	gt.Add(NewImage(Http, url))
	return gt
}

func (gt *Task) AddUrls(urls string) *Task {
	for _, url := range strings.Split(urls, ",") {
		gt.Add(NewImage(Http, url))
	}
	return gt
}

// AddDirname specified dirname, scan images and add it to gir task, waiting for resize
func (gt *Task) AddScanDir(dir string) *Task {
	// scan dir get images
	imgs, err := GetImagesFromDir(dir)
	if err != nil {
		gt.chErr <- err
		return gt
	}
	for _, img := range imgs {
		gt.Add(NewImage(Local, img))
	}
	return gt
}

// EmptyTask return Task whether is empty
func (gt *Task) EmptyTask() bool {
	return len(gt.images) == 0
}

// Report synchronously report success or fail result in background, when gir task is finish
func (gt *Task) Report() {
	wg := sync.WaitGroup{}

	// report success
	wg.Add(1)
	go func() {
		defer wg.Done()
		for rs := range gt.chSave {
			log.Printf("resize ok: %s (inputW=%d,inputH=%d,qty=%d,interp=%d,format='%s')\n",
				rs.save, rs.w, rs.h, rs.qty, rs.interp, rs.format)
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
	gt.fin <- struct{}{}
}

// ResizeImages concurrency resize image in it's GirImage slice
func (gt *Task) Run() {
	// concurrency task working
	wg := sync.WaitGroup{}

	// doResize resize an input image
	doResize := func(img Image) {
		defer wg.Done()
		save, err := img.DoResize()
		if err != nil {
			gt.chErr <- err
			return
		}
		gt.chSave <- saveRs{
			save,
			gt.setting.Width,
			gt.setting.Height,
			gt.setting.Interp,
			gt.setting.Qty,
			gt.setting.Format,
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
