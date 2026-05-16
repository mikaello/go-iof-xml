package marshallers_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
	"github.com/mikaello/go-iof-xml/pkg/marshallers"
)

const sampleResultList = `<?xml version="1.0" encoding="UTF-8"?>
<ResultList xmlns="http://www.orienteering.org/datastandard/3.0" iofVersion="3.0" status="Complete">
  <Event>
    <Name>Spring Cup</Name>
    <StartTime>
      <Date>2026-04-12</Date>
      <Time>10:00:00+02:00</Time>
    </StartTime>
  </Event>
</ResultList>`

// Decode dispatches on the document's root element and returns a pointer
// to the matching iof_v3 type. Use a type switch or assertion to access
// the concrete value.
func ExampleDecode() {
	doc, err := marshallers.Decode([]byte(sampleResultList))
	if err != nil {
		log.Fatal(err)
	}

	switch d := doc.(type) {
	case *iof_v3.ResultList:
		fmt.Println("result list:", d.Event.Name)
	case *iof_v3.StartList:
		fmt.Println("start list:", d.Event.Name)
	}
	// Output: result list: Spring Cup
}

// DecodeAs is a typed variant that returns *T directly. It returns an
// error if the document's root element does not match T.
func ExampleDecodeAs() {
	list, err := marshallers.DecodeAs[iof_v3.ResultList]([]byte(sampleResultList))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list.Event.Name)
	fmt.Println(list.Event.StartTime.Date)
	// Output:
	// Spring Cup
	// 2026-04-12
}

// DecodeReader is like Decode but takes an io.Reader, useful when
// streaming from a file or HTTP response body.
func ExampleDecodeReader() {
	doc, err := marshallers.DecodeReader(bytes.NewReader([]byte(sampleResultList)))
	if err != nil {
		log.Fatal(err)
	}
	list := doc.(*iof_v3.ResultList)
	fmt.Println(list.Event.Name)
	// Output: Spring Cup
}

// EncodeXML round-trips: decode a document, mutate it, then serialise it
// back to XML with the standard declaration.
func ExampleEncodeXML() {
	list, err := marshallers.DecodeAs[iof_v3.ResultList]([]byte(sampleResultList))
	if err != nil {
		log.Fatal(err)
	}
	list.Event.Name = "Spring Cup (renamed)"

	out, err := marshallers.EncodeXML(list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.Contains(out, []byte("<Name>Spring Cup (renamed)</Name>")))
	// Output: true
}

// EncodeJSON serialises a document to JSON. The shape mirrors the Go
// structs and is not part of the IOF data standard.
func ExampleEncodeJSON() {
	list, err := marshallers.DecodeAs[iof_v3.ResultList]([]byte(sampleResultList))
	if err != nil {
		log.Fatal(err)
	}

	out, err := marshallers.EncodeJSON(list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.HasPrefix(bytes.TrimSpace(out), []byte("{")))
	// Output: true
}
