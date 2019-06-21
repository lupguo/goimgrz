package girls

import (
	"log"
	"sync"
)

// gir image data resource type maybe local image, http url, or stdin base64_encode
type ResourceType int

const (
	ResTypeLocal ResourceType = iota // local image
	ResTypeHttp                      // http url image
	ResTypeStdin                     // stdin base64_encode
)

// GirImage used for indicate the image resource which will being resize
type GirImage struct {
	resType ResourceType
	data    []byte
	dst     string
	width   uint
	height  uint
}

// GirImage.Resize used for various resource image type
func (gi *GirImage) Resize() (string, error) {
	switch gi.resType {
	case ResTypeLocal:
		return ResizeLocalImage(string(gi.data), gi.dst, gi.width, gi.height)
	case ResTypeHttp:
		return ResizeHttpImage(string(gi.data), gi.dst, gi.width, gi.height)
	default:
		return "", NewError(ErrResImageType, "resource type", "cannot recognize the resize image's resource type")
	}
}

// saveRs save resize result
type saveRs struct {
	save string
	w    uint
	h    uint
}

// GirTask used for collect image resize, dispatching resize image task, got the save result or fail info from channel
type GirTask struct {
	images []*GirImage
	chErr  chan error
	chSave chan saveRs
	fin    chan bool
}

// NewGirTask create an GirTas pointer
func NewGirTask() *GirTask {
	return &GirTask{
		images: []*GirImage{},
		chErr:  make(chan error),
		chSave: make(chan saveRs),
		fin:    make(chan bool),
	}
}

// Add specified image to gir task, waiting for resize
func (gt *GirTask) Add(rt ResourceType, data []byte, dst string, width, height uint) *GirTask {
	gi := &GirImage{
		resType: rt,
		data:    data,
		dst:     dst,
		width:   width,
		height:  height,
	}
	gt.images = append(gt.images, gi)

	return gt
}

// IsEmpty check girTask whether is empty
func (gt *GirTask) IsEmpty() bool {
	return len(gt.images) == 0
}

// Report synchronously report success or fail result in background, when gir task is finish
func (gt *GirTask) report() {
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
	// report in background
	gt.report()

	// concurrency task working
	wg := sync.WaitGroup{}
	for _, gti := range gt.images {
		wg.Add(1)
		go func(gi *GirImage) {
			defer wg.Done()
			save, err := gi.Resize()
			if err != nil {
				gt.chErr <- err
				return
			}
			gt.chSave <- saveRs{
				save,
				gi.width,
				gi.height,
			}
		}(gti)
	}

	// wait for all task finished
	wg.Wait()
	close(gt.chErr)
	close(gt.chSave)
	<-gt.fin
}
