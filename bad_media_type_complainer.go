package dataurl


import (
	"fmt"
)


type BadMediaTypeComplainer interface {
	BadRequestComplainer
	BadMediaTypeComplainer()
	WrappedError() error
}


type internalBadMediaTypeComplainer struct {
	wrappedErr error
}


func newBadMediaTypeComplainer(err error) BadMediaTypeComplainer {
	complainer := internalBadMediaTypeComplainer{
		wrappedErr:err,
	}

	return &complainer
}


func (complainer *internalBadMediaTypeComplainer) Error() string {
	return fmt.Sprintf("Bad Request: Bad Media Type: %s", complainer.wrappedErr.Error())
}


func (*internalBadMediaTypeComplainer) BadRequestComplainer() {
	// Nothing here.
}

func (*internalBadMediaTypeComplainer) BadMediaTypeComplainer() {
	// Nothing here.
}

func (complainer *internalBadMediaTypeComplainer) WrappedError() error {
	return complainer.wrappedErr
}
