package dataurl


import (
	"encoding/base64"
	"net/url"
	"strings"
)


const (
	dataColon               = "data:"
	semicolonBase64Comma    = ";base64,"
	comma                   = ","
)


var (
	errSyntaxErrorNoComma = newSyntaxErrorComplainer("Data URL does not contain a comma.")
)


// Parse parses a data URL contained in parameter 'dataURL', and if it
// contained a valid data URL, returns a Parcel, else returns an error.
//
// Example usage:
//
//	parcel, err := dataurl.Parse("data:,Hello%20world!")
//	if nil != err {
//		//@TODO
//	}
//	
//	fmt.Println(parcel.String()) // parcel.String() == "Hello world!"
func Parse(dataURL string) (Parcel, error) {

	// If it doesn't start with "data:", then it isn't a data URL.
	// If that's the case, then return the appropriate error.
	if !strings.HasPrefix(dataURL, dataColon)  {
		return nil, errNotADataUrl
	}


	// Deal with some particular cases, that all result in empty
	// content with the default media type (of "text/plain;charset=US-ASCII").
	//
	// We do this because for each of these we can just return
	// the ready-made global 'emptyDefaultParcel'. And not have
	// to allocate new memory.
	switch dataURL {
	case `data:,`,
	     `data:text/plain,`,
	     `data:text/plain;charset=US-ASCII,`,
	     `data:;charset=US-ASCII,`,
	     `data:;base64,`,
	     `data:text/plain;base64,`,
	     `data:text/plain;charset=US-ASCII;base64,`,
	     `data:;charset=US-ASCII;base64,`:
		return emptyDefaultParcel, nil
	}


	// These constants will be used later in this func to specify what
	// type of data URL we have.
	//
	// Is it a 'URL encoded' data URL?
	//
	// Or is it a 'base64 encoded' data URL?
	type encodingType int
	const (
		encodingUndefined encodingType = iota
		encodingUrl
		encodingBase64
	)


	// Validate the media type and figure out how the data URL
	// is encoded.
	//
	// Is it a URL encoded data URL?
	//
	// Or is it a base64 encoded data URL?
	encoding := encodingUndefined
	var mediaType string
	var encoded   string
	{
		indexOfSemicolonBase64Comma := strings.Index(dataURL, semicolonBase64Comma)
		indexOfComma                := strings.Index(dataURL, comma)

		index := -1
		if -1 != indexOfSemicolonBase64Comma {
			index = indexOfSemicolonBase64Comma
			encoding = encodingBase64
			encoded = dataURL[index+len(semicolonBase64Comma):]
		} else if -1 != indexOfComma {
			index = indexOfComma
			encoding = encodingUrl
			encoded = dataURL[index+len(comma):]
		}
		if -1 == index {
			return nil, errSyntaxErrorNoComma
		}

		var err error

		// The RFC for data URLs kind of suggests that there might be
		// some URL encoded bits in the media type, but does not really
		// seem clear about it.
		//
		// Also, "in the wild" I don't see others who support data URLs
		// actually implementing this!
		//
		// So here I don't try to do any URL decoding of the media type
		// part of the data URL.
		mediaType = dataURL[len(dataColon):index]
		if mediaType, err = sanitizeMediaType(mediaType); nil != err {
			return nil, err
		}
	}

	// Create a parcel.
	//
	// This will be used to return the result of parsing a data URL.
	parcel := newParcel()

	// Set the media type in the parcel.
	parcel.mediaType = mediaType

	// (Try to) set the contents in the parcel.
	switch encoding {
	case encodingBase64:
		bs, err := base64.StdEncoding.DecodeString(encoded)
		if nil != err {
//@TODO: Could this error be improved? Maybe even wrapped?
			return nil, newSyntaxErrorComplainer("%s", err.Error())
		}

		parcel.buffer.Write(bs)
	case encodingUrl:
		s, err := url.QueryUnescape(encoded)
		if nil != err {
//@TODO: Could this error be improved? Maybe even wrapped?
			return nil, newSyntaxErrorComplainer("%s", err.Error())
		}

		parcel.buffer.WriteString(s)
	default:
		// This should never happen.
		return nil, newInternalErrorComplainer("Something weird happened. It seems like there is an unknown encoding type for the data URL (other than either base64 encoded or URL encoded), but that shouldn't be possible.")
	}


	return parcel, nil
}


// MustParse is like dataurl.Parse(), expect it only returns a Parcel, and
// panic()s if there was an error parsing 'dataURL'.
//
// Example usage:
//
//	parcel := dataurl.MustParse("data:,Hello%20world!")
//	
//	fmt.Println(parcel.String()) // parcel.String() == "Hello world!"
func MustParse(dataURL string) Parcel {

	parcel, err := Parse(dataURL)
	if nil != err {
		panic(err)
	}

	return parcel
}
