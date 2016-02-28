package dataurl


import (
	"fmt"
	"mime"
	"strings"
)


const (
	defaultMediaType = "text/plain;charset=US-ASCII"
)


// sanitizeMediaType turns a media type found in a data URL, into
// a "full" media type.
//
// A data URL does not have to explicitly specify a media type.
// For example: "data:,hello". In cases like this, the "full"
// media type (implicitly) is: "text/plain;charset=US-ASCII".
//
// Also a data URL may specify the MIME type without specifying
// the charset. For example: "data:text/csv,hello". In cases like
// this, the charset (implicitly) is: "charset=US-ASCII". And thus,
// in this example, the "full" media type (implicltly) is:
// "text/csv;charset=US-ASCII".
//
// Also a data URL may specify the charset without specifying the
// MIME type. For example: "data:;charset=utf-8,hello". In cases
// like this, the MIME type (implicitly) is: "text/plain". And thus,
// the charset (implicitly) is: "text/plain;charset=utf-8"
func sanitizeMediaType(mediaType string) (string, error) {

	if "" == mediaType {
		return defaultMediaType, nil
	}

	if strings.HasPrefix(mediaType, ";") {
		mediaType = fmt.Sprintf("text/plain%s", mediaType)
	}

	_, params, err := mime.ParseMediaType(mediaType)
	if nil != err {
		return "", newBadMediaTypeComplainer(err)
	}

	if _, ok := params["charset"]; !ok {
		mediaType = fmt.Sprintf("%s;charset=US-ASCII", mediaType)
	}

	return mediaType, nil
}
