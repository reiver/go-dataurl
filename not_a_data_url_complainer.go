package dataurl


var (
	errNotADataUrl = newNotADataUrlComplainer()
)


type NotADataUrlComplainer interface {
	BadRequestComplainer
	NotADataUrlComplainer()
}


type internalNotADataUrlComplainer struct{}


func newNotADataUrlComplainer() NotADataUrlComplainer {
	return new(internalNotADataUrlComplainer)
}


func (*internalNotADataUrlComplainer) Error() string {
	return "Bad Request: not a data URL."
}


func (*internalNotADataUrlComplainer) BadRequestComplainer() {
	// Nothing here.
}

func (*internalNotADataUrlComplainer) NotADataUrlComplainer() {
	// Nothing here.
}
