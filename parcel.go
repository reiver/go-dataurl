package dataurl


import (
	"bytes"
	"io"
	"strings"
)


var (
	emptyDefaultParcel = newParcel()
)


// Parcel is used to contain the result of parsing a Data URL.
//
// It provides the Bytes, Reader, Runes and String methods; used
// to retrieve the contents of a Data URL in []byte, io.Reader,
// []rune and string formats, respectively.
//
// It also provides the MediaType method; used to retrieve the
// (explicitly or implicitly) declared 'media type' of a Data URL.
//
// For example:
//
//	parcel, err := dataurl.Parse("data:,Hello")
//	if nil != err {
//		//@TODO
//	}
//	
//	contents := parcel.String() // == "Hello"
//	
//	mediaType := parcel.MediaType() // == "text/plain;charset=US-ASCII"
//
// Also, for example:
//
//	parcel, err := dataurl.Parse("data:application/x-apple-banana-cherry,Hello")
//	if nil != err {
//		//@TODO
//	}
//	
//	contents := parcel.String() // == "Hello"
//	
//	mediaType := parcel.MediaType() // == "application/x-apple-banana-cherry;charset=US-ASCII"
type Parcel interface {
	Bytes() []byte
	Reader() io.Reader
	Runes() []rune
	String() string

	MediaType() string
}


type internalParcel struct {
	buffer bytes.Buffer
	mediaType    string
}


func newParcel() *internalParcel {
	parcel := internalParcel{
		mediaType: "text/plain;charset=US-ASCII", // This is initialized to the default media type for a data URL.
	}

	return &parcel
}


func (parcel *internalParcel) Bytes() []byte {
	return parcel.buffer.Bytes()
}


func (parcel *internalParcel) Reader() io.Reader {
	return strings.NewReader( parcel.buffer.String() )
}


func (parcel *internalParcel) Runes() []rune {
	return []rune(parcel.buffer.String())
}


func (parcel *internalParcel) String() string {
	return parcel.buffer.String()
}


func (parcel *internalParcel) MediaType() string {
	return parcel.mediaType
}
