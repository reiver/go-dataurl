package dataurl


type BadRequestComplainer interface {
	error
	BadRequestComplainer()
}
