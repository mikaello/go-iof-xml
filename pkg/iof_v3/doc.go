// Package iof_v3 contains Go structs for the IOF Interface Standard
// version 3 XML schema, used by orienteering software for exchanging
// events, classes, entries, start lists, results and related data.
//
// The structs in iof_v3.go are generated from the official XSD using
// github.com/xuri/xgen, with two manual tweaks:
//
//   - omitempty is set on all non-pointer fields so empty values do not
//     round-trip as empty elements.
//   - [DateAndOptionalTime] has a hand-written UnmarshalXML/MarshalXML
//     pair so it parses and preserves both timezone-bearing and
//     timezone-less ISO 8601 Time strings.
//
// Use github.com/mikaello/go-iof-xml/pkg/marshallers to decode and encode
// documents into these types.
//
// # Related
//
//   - IOF v3 datastandard:
//     https://github.com/international-orienteering-federation/datastandard-v3
//   - Official XSD (mirrored under assets/xsd/iof_v3.xsd in this repo):
//     https://github.com/international-orienteering-federation/datastandard-v3/blob/master/IOF.xsd
package iof_v3
