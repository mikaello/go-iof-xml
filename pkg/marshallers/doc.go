// Package marshallers encodes and decodes IOF Interface Standard v3 XML
// documents to and from the Go structs defined in
// github.com/mikaello/go-iof-xml/pkg/iof_v3.
//
// # Overview
//
// Decoding accepts XML with or without a UTF-8 BOM and supports non-UTF-8
// encodings declared in the XML prolog. The document type is identified
// by the root element. Use [Decode] when the type is not known statically,
// [DecodeAs] for a typed result, or [DecodeReader] when working with a
// stream.
//
// Encoding produces XML with the standard declaration and two-space
// indentation. [EncodeJSON] is provided for convenience and mirrors the
// XML structure; it is not part of the IOF data standard.
//
// # Related
//
//   - IOF v3 datastandard:
//     https://github.com/international-orienteering-federation/datastandard-v3
//   - Official XSD (mirrored under assets/xsd/iof_v3.xsd):
//     https://github.com/international-orienteering-federation/datastandard-v3/blob/master/IOF.xsd
//   - Java port: https://github.com/orienteering-oss/iof-xml
//   - C# port:
//     https://github.com/international-orienteering-federation/Dotnet-Client-IOF.XML.V3
package marshallers
