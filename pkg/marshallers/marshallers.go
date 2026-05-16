package marshallers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"

	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
	"golang.org/x/net/html/charset"
)

// Document is the union of IOF v3 root document types. The constraint is
// used by [DecodeAs] to provide a typed unmarshalling helper.
type Document interface {
	iof_v3.CompetitorList |
		iof_v3.OrganisationList |
		iof_v3.EventList |
		iof_v3.ClassList |
		iof_v3.EntryList |
		iof_v3.CourseData |
		iof_v3.StartList |
		iof_v3.ResultList |
		iof_v3.ServiceRequestList |
		iof_v3.ControlCardList
}

var utf8BOM = []byte{0xEF, 0xBB, 0xBF}

// trimBOM strips a leading UTF-8 byte-order mark from data if present.
func trimBOM(data []byte) []byte {
	return bytes.TrimPrefix(data, utf8BOM)
}

// newDecoder returns an xml.Decoder with a permissive charset reader so
// non-UTF-8 documents (as declared in the XML prolog) can be decoded.
func newDecoder(r io.Reader) *xml.Decoder {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	return dec
}

// rootElementName returns the local name of the first XML start element in
// data. It is tolerant of leading whitespace, the XML declaration,
// comments and processing instructions.
func rootElementName(data []byte) (string, error) {
	dec := newDecoder(bytes.NewReader(data))
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return "", fmt.Errorf("no root element found")
		}
		if err != nil {
			return "", fmt.Errorf("read root element: %w", err)
		}
		if se, ok := tok.(xml.StartElement); ok {
			return se.Name.Local, nil
		}
	}
}

// Decode decodes an IOF v3 XML document. The returned value is a pointer
// to one of the iof_v3 root types (e.g. *iof_v3.ResultList) determined by
// the document's root element. Use a type switch or assertion to access it.
func Decode(data []byte) (any, error) {
	data = trimBOM(data)

	name, err := rootElementName(data)
	if err != nil {
		return nil, err
	}

	v, err := newDocument(name)
	if err != nil {
		return nil, err
	}

	if err := newDecoder(bytes.NewReader(data)).Decode(v); err != nil {
		return nil, fmt.Errorf("decode %s: %w", name, err)
	}
	return v, nil
}

// DecodeReader is like [Decode] but reads from r. r is fully consumed.
func DecodeReader(r io.Reader) (any, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read XML: %w", err)
	}
	return Decode(data)
}

// DecodeAs decodes an IOF v3 XML document into the specific type T.
// It returns an error if the document's root element does not match T.
//
//	list, err := marshallers.DecodeAs[iof_v3.ResultList](data)
func DecodeAs[T Document](data []byte) (*T, error) {
	data = trimBOM(data)

	var v T
	expected := typeName(&v)

	name, err := rootElementName(data)
	if err != nil {
		return nil, err
	}
	if name != expected {
		return nil, fmt.Errorf("root element is %q, want %q", name, expected)
	}

	if err := newDecoder(bytes.NewReader(data)).Decode(&v); err != nil {
		return nil, fmt.Errorf("decode %s: %w", expected, err)
	}
	return &v, nil
}

// typeName returns the short Go type name of *T (e.g. "ResultList"), which
// matches the XML root element name for IOF v3 documents.
func typeName(v any) string {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t.Name()
}

// EncodeXML serialises an IOF v3 document into XML with the standard
// declaration line and two-space indentation.
func EncodeXML(v any) ([]byte, error) {
	body, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encode XML: %w", err)
	}
	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	buf.Write(body)
	return buf.Bytes(), nil
}

// EncodeJSON serialises an IOF v3 document into indented JSON. The JSON
// shape mirrors the Go structs and is not part of the IOF data standard.
func EncodeJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// newDocument returns a pointer to a freshly allocated value of the iof_v3
// type matching the given root element name.
func newDocument(rootElement string) (any, error) {
	switch rootElement {
	case "CompetitorList":
		return new(iof_v3.CompetitorList), nil
	case "OrganisationList":
		return new(iof_v3.OrganisationList), nil
	case "EventList":
		return new(iof_v3.EventList), nil
	case "ClassList":
		return new(iof_v3.ClassList), nil
	case "EntryList":
		return new(iof_v3.EntryList), nil
	case "CourseData":
		return new(iof_v3.CourseData), nil
	case "StartList":
		return new(iof_v3.StartList), nil
	case "ResultList":
		return new(iof_v3.ResultList), nil
	case "ServiceRequestList":
		return new(iof_v3.ServiceRequestList), nil
	case "ControlCardList":
		return new(iof_v3.ControlCardList), nil
	default:
		return nil, fmt.Errorf("unknown IOF v3 root element: %q", rootElement)
	}
}
