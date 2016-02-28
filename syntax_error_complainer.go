package dataurl


import (
	"fmt"
)


// SyntaxErrorComplainer is used to represent a specific kind of BadRequestComplainer error.
// Specifically, it represents a syntax error in a data URL passed to the dataurl.Parse() func
// or the dataurl.MustParse() func.
//
// Example usage is as follows:
//
//	parcel, err := dataurl.Parse("data:,Hello%20world!")
//	if nil != err {
//		switch err.(type) {
//	
//		case dataurl.SyntaxErrorComplainer: // ← Here we are detecting if the error returned was due to a syntax error, in the data URL. Also note that it comes BEFORE the 'dataurl.BadRequestComplainer' case; THAT IS IMPORTANT!
//	
//			fmt.Printf("The data URL passed to dataurl.Parse() had a syntax error in it. The error message describing the syntax error is....\n%s\n", err.Error())
//			return
//	
//		case dataurl.BadRequestComplainer:
//	
//			fmt.Printf("Something you did when you called dataurl.Parse() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case dataurl.InternalErrorComplainer:
//	
//			fmt.Printf("It's not your fault; it's my fault. Something bad happened internally when dataurl.Parse() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
type SyntaxErrorComplainer interface {
	BadRequestComplainer
	SyntaxErrorComplainer()
}


// internalSyntaxErrorComplainer is the only underlying implementation that fits the
// SyntaxErrorComplainer interface, in this library.
type internalSyntaxErrorComplainer struct {
	msg string
}


// newSyntaxErrorComplainer creates a new internalSyntaxErrorComplainer (struct) and
// returns it as a SyntaxErrorComplainer (interface).
func newSyntaxErrorComplainer(format string, a ...interface{}) SyntaxErrorComplainer {
	msg := fmt.Sprintf(format, a...)

	err := internalSyntaxErrorComplainer{
		msg:msg,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// SyntaxErrorComplainer interface).
func (err *internalSyntaxErrorComplainer) Error() string {
	s := fmt.Sprintf("Bad Request: Syntax Error: %s", err.msg)
	return s
}


// BadRequestComplainer method is necessary to satisfy the 'BadRequestComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalSyntaxErrorComplainer) BadRequestComplainer() {
	// Nothing here.
}


// SyntaxErrorComplainer method is necessary to satisfy the 'SyntaxErrorComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalSyntaxErrorComplainer) SyntaxErrorComplainer() {
	// Nothing here.
}
