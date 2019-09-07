package imgrz

import "fmt"

// Resize Error
const (
	// input extract
	ErrGetLocalDirImages = 21
	ErrOpenLocalImage    = 22
	ErrOpenHttpImage     = 23
	// filter
	ErrDetectName = 30
	ErrDetectSize = 31
	// resize handle
	ErrResize = 40
)

type Error struct {
	No   int
	Mark string
	Msg  string
}

func (e Error) Error() string {
	return fmt.Sprintf("error(%d): %s, %s", e.No, e.Mark, e.Msg)
}

// new goimgrz error
func NewError(no int, mark, msg string) Error {
	return Error{no, mark, msg}
}