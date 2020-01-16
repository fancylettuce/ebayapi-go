package ebayapi

import (
	"bytes"
	"encoding/xml"
	"io"
)

// PrettPrintXML returns the XML data as a formatted string prepared for debugging output
func PrettPrintXML(data []byte) (string, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return string(b.Bytes()), nil
		}
		if err != nil {
			return "", err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return "", err
		}
	}
}
