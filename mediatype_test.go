package dataurl


import (
	"testing"
)


func TestSanitizeMediaType(t *testing.T) {

	tests := []struct{
		MediaType string
		Expected  string
	}{
		{
			MediaType: "",
			Expected:  "text/plain;charset=US-ASCII",
		},



		{
			MediaType: "text/plain",
			Expected:  "text/plain;charset=US-ASCII",
		},



		{
			MediaType: ";charset=US-ASCII",
			Expected:  "text/plain;charset=US-ASCII",
		},
		{
			MediaType: ";charset=utf-8",
			Expected:  "text/plain;charset=utf-8",
		},



		{
			MediaType: "text/csv",
			Expected:  "text/csv;charset=US-ASCII",
		},



		{
			MediaType: "image/png",
			Expected:  "image/png;charset=US-ASCII",
		},



		{
			MediaType: "text/html;charset=utf-8",
			Expected:  "text/html;charset=utf-8",
		},



		{
			MediaType: ";name=file.txt",
			Expected:  "text/plain;name=file.txt;charset=US-ASCII",
		},



		{
			MediaType: "image/png;name=file.png",
			Expected:  "image/png;name=file.png;charset=US-ASCII",
		},



		{
			MediaType: "image/png;name=file.png;charset=utf-8",
			Expected:  "image/png;name=file.png;charset=utf-8",
		},
		{
			MediaType: "image/png;charset=utf-8;name=file.png",
			Expected:  "image/png;charset=utf-8;name=file.png",
		},
	}


	for testNumber, test := range tests {

		actual, err := sanitizeMediaType(test.MediaType)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v", testNumber, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected media type to be %q, but actually was %q.", testNumber, expected, actual)
			continue
		}
	}
}


func TestSanitizeMediaTypeFail(t *testing.T) {

	tests := []struct{
		MediaType string
	}{
		{
			MediaType: "apple/banana/cherry",
		},



		{
			MediaType: "apple//banana",
		},



		{
			MediaType: "apple/banana;;cherry=grape",
		},
	}


	for testNumber, test := range tests {

		_, err := sanitizeMediaType(test.MediaType)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but actually did not get one: %v\nMedia Type: %q", testNumber, err, test.MediaType)
			continue
		}

	}
}
