package dataurl


import (
	"io/ioutil"

	"testing"
)


func TestParse(t *testing.T) {

	tests := []struct{
		DataURL              string
		ExpectedMediaType    string
		ExpectedContent      string
	}{
		// These 8 data URLs should all result in empty content with
		// the default media type (of "text/plain;charset=US-ASCI").
		{
			DataURL: `data:,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/plain,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/plain;charset=US-ASCII,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:;charset=US-ASCII,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:;base64,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/plain;base64,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/plain;charset=US-ASCII;base64,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:;charset=US-ASCII;base64,`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent:   ``,
		},



		// These 4 tests are from the data URL RFC -- from RFC 2397: http://tools.ietf.org/html/rfc2397
		{
			DataURL: `data:,A%20brief%20note`,
			ExpectedMediaType:    "text/plain;charset=US-ASCII",
			ExpectedContent:      `A brief note`,

		},
		{
			DataURL: `data:image/gif;base64,R0lGODdhMAAwAPAAAAAAAP///ywAAAAAMAAwAAAC8IyPqcvt3wCcDkiLc7C0qwyGHhSWpjQu5yqmCYsapyuvUUlvONmOZtfzgFzByTB10QgxOR0TqBQejhRNzOfkVJ+5YiUqrXF5Y5lKh/DeuNcP5yLWGsEbtLiOSpa/TPg7JpJHxyendzWTBfX0cxOnKPjgBzi4diinWGdkF8kjdfnycQZXZeYGejmJlZeGl9i2icVqaNVailT6F5iJ90m6mvuTS4OK05M0vDk0Q4XUtwvKOzrcd3iq9uisF81M1OIcR7lEewwcLp7tuNNkM3uNna3F2JQFo97Vriy/Xl4/f1cf5VWzXyym7PHhhx4dbgYKAAA7`,
			ExpectedMediaType:    "image/gif;charset=US-ASCII",
			ExpectedContent:      "GIF87a0\x000\x00\xf0\x00\x00\x00\x00\x00\xff\xff\xff,\x00\x00\x00\x000\x000\x00\x00\x02\xf0\x8c\x8f\xa9\xcb\xed\xdf\x00\x9c\x0eH\x8bs\xb0\xb4\xab\f\x86\x1e\x14\x96\xa64.\xe7*\xa6\t\x8b\x1a\xa7+\xafQIo8َf\xd7\xf3\x80\\\xc1\xc90u\xd1\b19\x1d\x13\xa8\x14\x1e\x8e\x14M\xcc\xe7\xe4T\x9f\xb9b%*\xadqyc\x99J\x87\xf0\u07b8\xd7\x0f\xe7\"\xd6\x1a\xc1\x1b\xb4\xb8\x8eJ\x96\xbfL\xf8;&\x92G\xc7'\xa7w5\x93\x05\xf5\xf4s\x13\xa7(\xf8\xe0\a8\xb8v(\xa7Xgd\x17\xc9#u\xf9\xf2q\x06We\xe6\x06z9\x89\x95\x97\x86\x97ض\x89\xc5jh\xd5Z\x8aT\xfa\x17\x98\x89\xf7I\xba\x9a\xfb\x93K\x83\x8aӓ4\xbc94C\x85Է\v\xca;:\xdcwx\xaa\xf6\xe8\xac\x17\xcdL\xd4\xe2\x1cG\xb9D{\f\x1c.\x9e\xed\xb8\xd3d3{\x8d\x9d\xad\xc5ؔ\x05\xa3\xdeծ,\xbf^^?\u007fW\x1f\xe5U\xb3_,\xa6\xec\xf1\xe1\x87\x1e\x1dn\x06\n\x00\x00;",

		},
		{
			DataURL: `data:text/plain;charset=iso-8859-7,%b8%f7%fe`,
			ExpectedMediaType:    "text/plain;charset=iso-8859-7",
			ExpectedContent:      "\xb8\xf7\xfe",

		},
		{
			DataURL: `data:application/vnd-xxx-query,select_vcount,fcol_from_fieldtable/local`,
			ExpectedMediaType:    "application/vnd-xxx-query;charset=US-ASCII",
			ExpectedContent:      "select_vcount,fcol_from_fieldtable/local",

		},



		// This 1 test is from here: http://www-archive.mozilla.org/quality/networking/testing/datatests.html
		{
			DataURL: `data:,;test`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent: `;test`,
		},



		// These 2 tests are from here: http://www-archive.mozilla.org/quality/networking/testing/datatests.html
		{
			DataURL: `data:text/plain,test`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent: `test`,
		},
		{
			DataURL: `data:text/plain;charset=US-ASCII,test`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent: `test`,
		},



		// This 1 test is from here: http://www-archive.mozilla.org/quality/networking/testing/datatests.html
		{
			DataURL: `data:,a,b`,
			ExpectedMediaType: "text/plain;charset=US-ASCII",
			ExpectedContent: `a,b`,
		},



		// This 1 test is from here: http://www-archive.mozilla.org/quality/networking/testing/datatests.html
		{
			DataURL: `data:application/vnd.mozilla.xul+xml,%3C?xml%20version=%221.0%22?%3E%3Cwindow%20xmlns=%22http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul%22%3E%3C?xml-stylesheet%20href=%22data:text/css,#a%7B-moz-box-flex:1;%7D%22?%3E%3Cbox%20id=%22a%22%3E%3Clabel%20value=%22This%20works%21%22/%3E%3C/box%3E%3Cbox/%3E%3C/window%3E`,
			ExpectedMediaType: "application/vnd.mozilla.xul+xml;charset=US-ASCII",
			ExpectedContent: `<?xml version="1.0"?><window xmlns="http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul"><?xml-stylesheet href="data:text/css,#a{-moz-box-flex:1;}"?><box id="a"><label value="This works!"/></box><box/></window>`,
		},



		// More "no data, no error" tests.
		{
			DataURL: `data:application/x-apple-banana-cherry,`,
			ExpectedMediaType: "application/x-apple-banana-cherry;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:application/x-apple-banana-cherry;base64,`,
			ExpectedMediaType: "application/x-apple-banana-cherry;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:application/x-apple-banana-cherry;charset=UTF-8,`,
			ExpectedMediaType: "application/x-apple-banana-cherry;charset=UTF-8",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:application/x-apple-banana-cherry;charset=UTF-8;base64,`,
			ExpectedMediaType: "application/x-apple-banana-cherry;charset=UTF-8",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/css,`,
			ExpectedMediaType: "text/css;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/csv,`,
			ExpectedMediaType: "text/csv;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/html,`,
			ExpectedMediaType: "text/html;charset=US-ASCII",
			ExpectedContent:   ``,
		},
		{
			DataURL: `data:text/x-asm,`,
			ExpectedMediaType: "text/x-asm;charset=US-ASCII",
			ExpectedContent:   ``,
		},




		// Some tests with content of "This is a test!".
		{
			DataURL: `data:text/plain;charset=utf-8,This%20is%20a%20test%21`,
			ExpectedMediaType:    "text/plain;charset=utf-8",
			ExpectedContent:      `This is a test!`,

		},
		{
			DataURL: `data:;charset=utf-8,This%20is%20a%20test%21`,
			ExpectedMediaType:    "text/plain;charset=utf-8",
			ExpectedContent:      `This is a test!`,

		},
		{
			DataURL: `data:text/plain,This%20is%20a%20test%21`,
			ExpectedMediaType:    "text/plain;charset=US-ASCII",
			ExpectedContent:      `This is a test!`,

		},
		{
			DataURL: `data:,This%20is%20a%20test%21`,
			ExpectedMediaType:    "text/plain;charset=US-ASCII",
			ExpectedContent:      `This is a test!`,

		},



		{
			DataURL: `data:text/plain;charset=utf-8;base64,VGhpcyBpcyBhIHRlc3Qh`,
			ExpectedMediaType:    "text/plain;charset=utf-8",
			ExpectedContent:      `This is a test!`,

		},
		{
			DataURL: `data:;charset=utf-8;base64,VGhpcyBpcyBhIHRlc3Qh`,
			ExpectedMediaType:    "text/plain;charset=utf-8",
			ExpectedContent:      `This is a test!`,

		},
		{
			DataURL: `data:text/plain;base64,VGhpcyBpcyBhIHRlc3Qh`,
			ExpectedMediaType:    "text/plain;charset=US-ASCII",
			ExpectedContent:      `This is a test!`,

		},
		{
			DataURL: `data:;base64,VGhpcyBpcyBhIHRlc3Qh`,
			ExpectedMediaType:    "text/plain;charset=US-ASCII",
			ExpectedContent:      `This is a test!`,

		},



		// Some tests with CSV data.
		{
			DataURL: `data:text/csv;charset=utf-8,first_name%2Clast_name%2Ccity%0D%0AJoe%2CBlow%2CVancouver%0D%0AJohn%2CDoe%2CWinnipeg%0D%0AJane%2CDoe%2CMontreal%0D%0ATommy%2CAtkins%2CLondon`,
			ExpectedMediaType:    "text/csv;charset=utf-8",
			ExpectedContent:
"first_name,last_name,city" + "\r\n" +
"Joe,Blow,Vancouver"        + "\r\n" +
"John,Doe,Winnipeg"         + "\r\n" +
"Jane,Doe,Montreal"         + "\r\n" +
"Tommy,Atkins,London",
		},
		{
			DataURL: `data:text/csv,first_name%2Clast_name%2Ccity%0D%0AJoe%2CBlow%2CVancouver%0D%0AJohn%2CDoe%2CWinnipeg%0D%0AJane%2CDoe%2CMontreal%0D%0ATommy%2CAtkins%2CLondon`,
			ExpectedMediaType:    "text/csv;charset=US-ASCII",
			ExpectedContent:
"first_name,last_name,city" + "\r\n" +
"Joe,Blow,Vancouver"        + "\r\n" +
"John,Doe,Winnipeg"         + "\r\n" +
"Jane,Doe,Montreal"         + "\r\n" +
"Tommy,Atkins,London",
		},
		{
			DataURL: `data:text/csv;charset=utf-8;base64,Zmlyc3RfbmFtZSxsYXN0X25hbWUsY2l0eQ0KSm9lLEJsb3csVmFuY291dmVyDQpKb2huLERvZSxXaW5uaXBlZw0KSmFuZSxEb2UsTW9udHJlYWwNClRvbW15LEF0a2lucyxMb25kb24=`,
			ExpectedMediaType:    "text/csv;charset=utf-8",
			ExpectedContent:
"first_name,last_name,city" + "\r\n" +
"Joe,Blow,Vancouver"        + "\r\n" +
"John,Doe,Winnipeg"         + "\r\n" +
"Jane,Doe,Montreal"         + "\r\n" +
"Tommy,Atkins,London",
		},
		{
			DataURL: `data:text/csv;base64,Zmlyc3RfbmFtZSxsYXN0X25hbWUsY2l0eQ0KSm9lLEJsb3csVmFuY291dmVyDQpKb2huLERvZSxXaW5uaXBlZw0KSmFuZSxEb2UsTW9udHJlYWwNClRvbW15LEF0a2lucyxMb25kb24=`,
			ExpectedMediaType:    "text/csv;charset=US-ASCII",
			ExpectedContent:
"first_name,last_name,city" + "\r\n" +
"Joe,Blow,Vancouver"        + "\r\n" +
"John,Doe,Winnipeg"         + "\r\n" +
"Jane,Doe,Montreal"         + "\r\n" +
"Tommy,Atkins,London",
		},



		// A test with a parameter other than "charset".
		// In this case "name".
		{
			DataURL: `data:text/plain;name=sample.txt;charset=utf-8,This%20is%20a%20test%21`,
			ExpectedMediaType:    "text/plain;name=sample.txt;charset=utf-8",
			ExpectedContent:      `This is a test!`,

		},



		{
			DataURL: `data:text/html;base64,PGh0bWw+PGhlYWQ+PHRpdGxlPlRlc3Q8L3RpdGxlPjwvaGVhZD48Ym9keT48cD5UaGlzIGlzIGEgdGVzdDwvYm9keT48L2h0bWw+Cg==`,
			ExpectedMediaType:    "text/html;charset=US-ASCII",
			ExpectedContent:      `<html><head><title>Test</title></head><body><p>This is a test</body></html>`+"\n",

		},



		// These 4 tests are from: https://bug161965.bmoattachments.org/attachment.cgi?id=94670
		{
			DataURL: `data:text/plain;charset=iso-8859-8-i;base64,+ezl7Q==`,
			ExpectedMediaType:    "text/plain;charset=iso-8859-8-i",
			ExpectedContent:      "\xf9\xec\xe5\xed",

		},
		{
			DataURL: `data:text/plain;charset=iso-8859-8-i,%f9%ec%e5%ed`,
			ExpectedMediaType:    "text/plain;charset=iso-8859-8-i",
			ExpectedContent:      "\xf9\xec\xe5\xed",

		},
		{
			DataURL: `data:text/plain;charset=UTF-8;base64,16nXnNeV150=`,
			ExpectedMediaType:    "text/plain;charset=UTF-8",
			ExpectedContent:      "שלום",

		},
		{
			DataURL: `data:text/plain;charset=UTF-8,%d7%a9%d7%9c%d7%95%d7%9d`,
			ExpectedMediaType:    "text/plain;charset=UTF-8",
			ExpectedContent:      "שלום",

		},
	}


	for testNumber, test := range tests {
		parcel, err := Parse(test.DataURL)
		if nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: %v\nData URL: %s", testNumber, err, test.DataURL)
			continue
		}
		if nil == parcel {
			t.Errorf("For test #%d, expected a non-nil parcel, but actually got nil: %v\nData URL: %s", testNumber, parcel, test.DataURL)
			continue
		}


		if expected, actual := test.ExpectedMediaType, parcel.MediaType(); expected != actual {
			t.Errorf("For test #%d, expected Media Type to be %q, but actually got %q.\nData URL: %s", testNumber, expected, actual, test.DataURL)
			continue
		}


		if expected, actual := test.ExpectedContent, parcel.String(); expected != actual {
			t.Errorf("For test #%d, expected content from .String() to be %q, but actually was %q.\nData URL: %s", testNumber, expected, actual, test.DataURL)
			continue
		}

		if expected, actual := []rune(test.ExpectedContent), parcel.Runes(); len(expected) != len(actual) {
			t.Errorf("For test #%d, expected %d runes from .Runes(), but actually got %d.\nData URL: %s", testNumber, len(expected), len(actual), test.DataURL)
			continue
		} else {
			for runeNumber, expectedRune := range expected {
				actualRune := actual[runeNumber]

				if expectedRune != actualRune {
					t.Errorf("For test #%d and rune #%d, expected rune to be %q, but actually was %q.\nData URL: %s", testNumber, runeNumber, expectedRune, actualRune, test.DataURL)
					continue
				}
			}
		}

		if readerActualBytes, err := ioutil.ReadAll(parcel.Reader()); nil != err {
			t.Errorf("For test #%d, did not expect an error when trying to read all from .Reader(), but actually got one: %v\nData URL: %s", testNumber, err, test.DataURL)
			continue
		} else {
			if expected, actual := test.ExpectedContent, string(readerActualBytes); expected != actual {
				t.Errorf("For test #%d, expected content content from .Reader() to be %q, but actually was %q.\nData URL: %s", testNumber, expected, actual, test.DataURL)
				continue
			}
		}

		if expected, actual := []byte(test.ExpectedContent), parcel.Bytes(); len(expected) != len(actual) {
			t.Errorf("For test #%d, expected %d bytes from .Bytes(), but actually got %d.\nData URL: %s", testNumber, len(expected), len(actual), test.DataURL)
			continue
		} else {
			for byteNumber, expectedByte := range expected {
				actualByte := actual[byteNumber]

				if expectedByte != actualByte {
					t.Errorf("For test #%d and byte #%d, expected byte to be %d, but actually was %d.\nData URL: %s", testNumber, byteNumber, expectedByte, actualByte, test.DataURL)
					continue
				}
			}
		}
	}
}


func TestParseFail(t *testing.T) {

	tests := []struct{
		DataURL string
	}{
		{
			DataURL: ``,

		},
		{
			DataURL: `http://example.com/robots.txt`,

		},
		{
			DataURL: `ftp://example.net/some/path/to/a/file.txt`,

		},
		{
			DataURL: `data::text/plain;charset=utf-8,This%20is%20a%20test%21`,

		},
		{
			DataURL: `datatext/plain;charset=utf-8,This%20is%20a%20test%21`,

		},
		{
			DataURL: `data:text/plaincharset=utf-8,This%20is%20a%20test%21`,

		},
		{
			DataURL: `data:text/plain;charset=utf-8This%20is%20a%20test%21`,

		},
		{
			DataURL: `data:`,
		},
		{
			DataURL: `datum:,`,
		},
//		{
//			DataURL: `data:;,test`,
//
//		},
	}


	for testNumber, test := range tests {
		parcel, err := Parse(test.DataURL)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %v\nData URL: %q\nParcel Media Type: %q\nParcel Content: %q", testNumber, err, test.DataURL, parcel.MediaType(), parcel.String())
			continue
		}
		if nil != parcel {
			t.Errorf("For test #%d, expected a returned parcel to be nil, but actually got: %v\nData URL: %s", testNumber, parcel, test.DataURL)
			continue
		}

	}
}
