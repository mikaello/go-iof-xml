// Package eventor_native provides Go types and decode helpers for the
// Eventor REST API's proprietary native XML format.
//
// The Eventor API returns two flavours of XML:
//   - IOF v3 XML — available on endpoints that accept an `iofVersion=3`
//     parameter (or a dedicated `/iofxml` path). Use the [iof_v3] package
//     and [marshallers] for those.
//   - Eventor native XML — the default XML returned by most Eventor API
//     endpoints. This package covers those document types.
//
// # Document types
//
// Each top-level Eventor native XML document type has a corresponding Go
// struct and a Decode helper:
//
//   - [Organisation] / [DecodeOrganisation] — single organisation from
//     `GET /api/organisations/{id}`
//   - [Event] / [DecodeEvent] — single event from `GET /api/events/{id}`
//   - [EventClassList] / [DecodeEventClassList] — event classes from
//     `GET /api/eventclasses`
//   - [EntryFeeList] / [DecodeEntryFeeList] — entry fees from
//     `GET /api/entryfees/event`
//   - [DocumentList] / [DecodeDocumentList] — event documents from
//     `GET /api/events/documents`
//   - [CompetitorCountList] / [DecodeCompetitorCountList] — entry/start
//     counts from `GET /api/competitorcount`
//   - [EntryList] / [DecodeEntryList] — entries from `GET /api/entries`
//
// # Usage
//
//	data, _ := os.ReadFile("organisation.xml")
//	org, err := eventor_native.DecodeOrganisation(data)
//	if err != nil { ... }
//	fmt.Println(org.Name)
package eventor_native
