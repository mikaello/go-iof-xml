// Package eventor_extensions reads the Eventor-specific data carried
// in IOF XML 3.0 `<Extensions>` elements emitted by the Eventor REST
// API's `/.../iofxml` endpoints.
//
// Eventor writes these under the namespace
// `http://eventor.orientering.se/iofxmlextensions`, prefixed `eventor:`,
// inside the `<Extensions>` element on `<Event>` and on each `<Race>`:
//
//	<Extensions>
//	  <eventor:EventRaceId type="Eventor">24332</eventor:EventRaceId>
//	  <eventor:StartListExists>true</eventor:StartListExists>
//	  <eventor:ResultListExists>true</eventor:ResultListExists>
//	  <eventor:Discipline>Foot</eventor:Discipline>
//	  <eventor:Discipline>MountainBike</eventor:Discipline>
//	  <eventor:LightCondition>Day</eventor:LightCondition>
//	  <eventor:Attribute id="2">Ukas løype</eventor:Attribute>
//	</Extensions>
//
// `<eventor:Discipline>` and `<eventor:Attribute>` are both repeatable
// — Eventor lists each discipline / custom attribute as a separate
// element rather than wrapping them.
//
// The IOF XSD treats `<Extensions>` as an open catch-all
// (the generated [iof_v3.Extensions] struct is empty), so consumers who
// need to read the Eventor data must parse it separately. This package
// provides a parallel [EventList] / [Event] / [Race] hierarchy with the
// extension fields populated, so callers can decode a document twice —
// once with [iof_v3] for the standard data and once with this package
// for the Eventor extension data — and key them up by event id.
//
// The Norwegian, Swedish, Australian and International Eventor instances
// all expose the `/.../iofxml` endpoints, even when the per-instance
// `/api/documentation` page does not list them.
package eventor_extensions

import (
	"encoding/xml"
)

// Discipline is Eventor's discipline classification, distinct from the
// IOF v3 `Discipline` element. The values mirror the EventorDiscipline
// enum in eventor-api-openapi-spec.
type Discipline string

const (
	DisciplineFoot         Discipline = "Foot"
	DisciplineMountainBike Discipline = "MountainBike"
	DisciplineSki          Discipline = "Ski"
	DisciplineTrail        Discipline = "Trail"
	DisciplineIndoor       Discipline = "Indoor"
)

// LightCondition describes whether a race takes place during the day,
// at night, or both.
type LightCondition string

const (
	LightConditionDay         LightCondition = "Day"
	LightConditionNight       LightCondition = "Night"
	LightConditionDayAndNight LightCondition = "DayAndNight"
)

// Attribute is a custom event attribute attached to the parent element.
// The set of attributes is instance-specific (the Norwegian Eventor
// exposes e.g. "Ukas løype", "Paratilbud", "Flexoløp").
type Attribute struct {
	ID    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

// Extensions are the Eventor-namespaced fields seen inside an
// `<Extensions>` element on an `<Event>` or `<Race>`. Fields that
// were not present in the XML are left as their zero value.
type Extensions struct {
	XMLName          xml.Name       `xml:"Extensions"`
	EventRaceID      string         `xml:"EventRaceId"`
	StartListExists  bool           `xml:"StartListExists"`
	ResultListExists bool           `xml:"ResultListExists"`
	// Disciplines lists every discipline the event/race accommodates.
	// Eventor emits one <eventor:Discipline> element per discipline.
	Disciplines    []Discipline   `xml:"Discipline"`
	LightCondition LightCondition `xml:"LightCondition"`
	// Attributes lists the custom Eventor event attributes.
	Attributes []Attribute `xml:"Attribute"`
}

// Race is the slice of an IOF XML `<Race>` element that carries
// Eventor extensions.
type Race struct {
	RaceNumber int        `xml:"RaceNumber"`
	Name       string     `xml:"Name"`
	Extensions Extensions `xml:"Extensions"`
}

// Event is the slice of an IOF XML `<Event>` element that carries
// Eventor extensions, including any per-race extensions.
type Event struct {
	ID         string     `xml:"Id"`
	Name       string     `xml:"Name"`
	Races      []Race     `xml:"Race"`
	Extensions Extensions `xml:"Extensions"`
}

// EventList parses an IOF XML `<EventList>` document and surfaces only
// the bits needed to read Eventor extensions. Use [Decode] (or any of
// the [Decode*] helpers) to populate it from raw XML.
type EventList struct {
	XMLName xml.Name `xml:"EventList"`
	Events  []Event  `xml:"Event"`
}

// StartList parses an IOF XML `<StartList>` document for its Eventor
// extensions. The extensions live on the `<Event>` and `<Race>`
// elements at the top of the document.
type StartList struct {
	XMLName xml.Name `xml:"StartList"`
	Event   Event    `xml:"Event"`
}

// ResultList parses an IOF XML `<ResultList>` document for its Eventor
// extensions.
type ResultList struct {
	XMLName xml.Name `xml:"ResultList"`
	Event   Event    `xml:"Event"`
}

// DecodeEventList decodes an EventList XML document and returns the
// events with their Eventor extension fields populated.
func DecodeEventList(data []byte) (*EventList, error) {
	var el EventList
	if err := xml.Unmarshal(data, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// DecodeStartList decodes a StartList XML document for its Eventor
// extensions.
func DecodeStartList(data []byte) (*StartList, error) {
	var sl StartList
	if err := xml.Unmarshal(data, &sl); err != nil {
		return nil, err
	}
	return &sl, nil
}

// DecodeResultList decodes a ResultList XML document for its Eventor
// extensions.
func DecodeResultList(data []byte) (*ResultList, error) {
	var rl ResultList
	if err := xml.Unmarshal(data, &rl); err != nil {
		return nil, err
	}
	return &rl, nil
}
