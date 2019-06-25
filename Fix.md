Already Fix Content:

#### Code
- [x] package name
- [x] Refactor one versions 

#### girls.go:
- [x] Traditionally you would move your "main" to cmd/girls and not have your package as "main". This makes it impossible to use for anyone.
- [x] Typically you would name the file with "main()" main.go.

#### resize.go:
- [x] No need to export var Cmdline.
- [x] newFile, _ := os.Create(save) Check your errors!
- [x] resize.NearestNeighbor if you are doing a resize library, do some basic research. This is without a doubt the worst option you could have chosen. (use bicubic/lanczos)
- [x] jpeg.Encode(newFile, newImg, &jpeg.Options{85}) - wasn't there supposed to be a flag for quality?
- [x] defer newFile.Close() you might as well close it here, this is way too late. Move it to right after you've opened it.

#### task.go:
- [x] // GirImage.Resize used for various resource image type... This documentation says nothing.
- [x] type GirImage struct - this should just be "Image". No need to prefix your types here.
- [x] type GirTask struct.. again, could just as well be Task.
- [x] fin: make(chan bool) - this might as well be a chan struct{} since you are not using the bool anyway.
- [x] GirTask So the data []byte is the input file name or URL? Why is this a []byte. And why is it called "data" - I assumed it was the image data.
- [x] type GirImage struct. => Refactor Image interface
- [x] func (gt *GirTask) IsEmpty() bool. I personally prefer Empty() bool since the 'Is' is implied by the returned bool.
- [x] log.Println("resize fail:", err): Consider whether this should be written to stderr instead. Also your program returns status code 0 even if something failed, maybe not the best idea.    