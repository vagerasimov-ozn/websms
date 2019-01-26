package types // import "gopkg.in/webnice/websms.v1/websms/types"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/xml"
	"time"
)

// TimeRfc1123z Time formating for XML
type TimeRfc1123z time.Time

// String Convert time to string
func (trz TimeRfc1123z) String() string { return time.Time(trz).Format(time.RFC1123Z) }

// Time Return time.Time object
func (trz TimeRfc1123z) Time() time.Time { return time.Time(trz) }

// MarshalXML Implementation of xml.Marshaller interface
func (trz TimeRfc1123z) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(trz).Format(time.RFC1123Z), start)
}

// MarshalXMLAttr Implementation of xml.MarshalerAttr interface
func (trz TimeRfc1123z) MarshalXMLAttr(name xml.Name) (attr xml.Attr, err error) {
	attr = xml.Attr{name, time.Time(trz).Format(time.RFC1123Z)}
	return
}

// Implementation of xml.Unmarshaler interface
func (trz *TimeRfc1123z) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t time.Time
	var s string

	if err = d.DecodeElement(&s, &start); err != nil {
		return
	}
	if t, err = time.ParseInLocation(time.RFC1123Z, s, time.Local); err != nil {
		return err
	}
	*trz = TimeRfc1123z(t)

	return
}

// Implementation of xml.UnmarshalerAttr interface
func (trz *TimeRfc1123z) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	var parse time.Time

	if parse, err = time.ParseInLocation(time.RFC1123Z, attr.Value, time.Local); err != nil {
		return
	}
	*trz = TimeRfc1123z(parse)

	return
}
