package eventor_native

import "encoding/xml"

// DecodeOrganisation decodes a single Eventor native `<Organisation>`
// document.
func DecodeOrganisation(data []byte) (*Organisation, error) {
	var v Organisation
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// DecodeEvent decodes a single Eventor native `<Event>` document as
// returned by `GET /api/events/{id}`.
func DecodeEvent(data []byte) (*Event, error) {
	var v Event
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// DecodeEventClassList decodes an Eventor native `<EventClassList>`
// document as returned by `GET /api/eventclasses`.
func DecodeEventClassList(data []byte) (*EventClassList, error) {
	var v EventClassList
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// DecodeEntryFeeList decodes an Eventor native `<EntryFeeList>` document
// as returned by `GET /api/entryfees/event`.
func DecodeEntryFeeList(data []byte) (*EntryFeeList, error) {
	var v EntryFeeList
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// DecodeDocumentList decodes an Eventor native `<DocumentList>` document
// as returned by `GET /api/events/documents`.
func DecodeDocumentList(data []byte) (*DocumentList, error) {
	var v DocumentList
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// DecodeCompetitorCountList decodes an Eventor native
// `<CompetitorCountList>` document as returned by
// `GET /api/competitorcount`.
func DecodeCompetitorCountList(data []byte) (*CompetitorCountList, error) {
	var v CompetitorCountList
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// DecodeEntryList decodes an Eventor native `<EntryList>` document as
// returned by `GET /api/entries`. This is distinct from the IOF v3
// EntryList format.
func DecodeEntryList(data []byte) (*EntryList, error) {
	var v EntryList
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
