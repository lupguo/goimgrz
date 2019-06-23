package goimgrz

import "fmt"

const (
	// task
	ErrResImageType = 10

	// input extract
	ErrParse             = 20
	ErrGetLocalDirImages = 21
	ErrOpenLocalImage    = 22
	ErrOpenHttpImage     = 23

	// filter
	ErrDetectName = 30
	ErrDetectSize = 31

	// resize handle
	ErrResize = 40

	// save file
	ErrSaveFile = 60
)

type GirError struct {
	No   int
	Mark string
	Msg  string
}

func (e GirError) Error() string {
	return fmt.Sprintf("error(%d): %s, %s", e.No, e.Mark, e.Msg)
}

// new goimgrz error
func NewError(no int, mark, msg string) GirError {
	return GirError{no, mark, msg}
}

// error print out
func ErrorOut(no int, mark, msg string) {
	fmt.Println(NewError(no, mark, msg))
}
