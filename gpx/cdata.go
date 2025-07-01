package gpx

import "encoding/xml"

// CDATA type to handle text with newline characters
type CDATA string

// MarshalXML method to wrap the text in a CDATA section
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}
