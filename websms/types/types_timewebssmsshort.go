package types // import "gopkg.in/webnice/websms.v1/websms/types"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/xml"
	"time"
)

const timeWebsmsShortOneFormat = `02.01.2006 15:04:05`

// TimeWebsmsShortOne Time formating for XML
type TimeWebsmsShortOne time.Time

// String Convert time to string
func (tso TimeWebsmsShortOne) String() string { return time.Time(tso).Format(timeWebsmsShortOneFormat) }

// Time Return time.Time object
func (tso TimeWebsmsShortOne) Time() time.Time { return time.Time(tso) }

// MarshalXML Implementation of xml.Marshaller interface
func (tso TimeWebsmsShortOne) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(tso).Format(timeWebsmsShortOneFormat), start)
}

// MarshalXMLAttr Implementation of xml.MarshalerAttr interface
func (tso TimeWebsmsShortOne) MarshalXMLAttr(name xml.Name) (attr xml.Attr, err error) {
	attr = xml.Attr{name, time.Time(tso).Format(timeWebsmsShortOneFormat)}
	return
}

// Implementation of xml.Unmarshaler interface
func (tso *TimeWebsmsShortOne) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t time.Time
	var s string

	if err = d.DecodeElement(&s, &start); err != nil {
		return
	}
	if t, err = time.ParseInLocation(timeWebsmsShortOneFormat, s, time.Local); err != nil {
		return err
	}
	*tso = TimeWebsmsShortOne(t)

	return
}

// Implementation of xml.UnmarshalerAttr interface
func (tso *TimeWebsmsShortOne) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	var parse time.Time

	if parse, err = time.ParseInLocation(timeWebsmsShortOneFormat, attr.Value, time.Local); err != nil {
		return
	}
	*tso = TimeWebsmsShortOne(parse)

	return
}
