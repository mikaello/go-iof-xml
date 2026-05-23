package eventor_native

import "encoding/xml"

// DateClock is the Eventor date+time representation used throughout the
// native XML format. It maps to:
//
//	<SomeDate>
//	  <Date>2026-05-12</Date>
//	  <Clock>17:30:00</Clock>
//	</SomeDate>
type DateClock struct {
	Date  string `xml:"Date"`
	Clock string `xml:"Clock,omitempty"`
}

// ValueAttr is a single-attribute element pattern used by Eventor for
// enum-like fields, e.g. <EventClassStatus value="normal"/>.
type ValueAttr struct {
	Value string `xml:"value,attr"`
}

// -----------------------------------------------------------------------
// Organisation
// -----------------------------------------------------------------------

// CountryRef is the Eventor native XML representation of a country,
// referencing it by numeric id, IOC alpha-3 code, and localised name.
type CountryRef struct {
	CountryID ValueAttr `xml:"CountryId"`
	Alpha3    ValueAttr `xml:"Alpha3"`
	// Name may appear multiple times with different languageId attributes;
	// the first entry is used for the English name.
	Names []CountryName `xml:"Name"`
}

// EnglishName returns the first English language name, falling back to
// the first name of any language if no English name is present.
func (c CountryRef) EnglishName() string {
	for _, n := range c.Names {
		if n.LangID == "en" {
			return n.Value
		}
	}
	if len(c.Names) > 0 {
		return c.Names[0].Value
	}
	return ""
}

// CountryName is a localised country name element.
type CountryName struct {
	LangID string `xml:"languageId,attr"`
	Value  string `xml:",chardata"`
}

// Address is the physical address of an organisation in Eventor native
// XML. PII fields (careOf, street, city, zipCode) are XML attributes.
type Address struct {
	CareOf      string    `xml:"careOf,attr,omitempty"`
	Street      string    `xml:"street,attr,omitempty"`
	City        string    `xml:"city,attr,omitempty"`
	ZipCode     string    `xml:"zipCode,attr,omitempty"`
	AddressType ValueAttr `xml:"AddressType"`
	Country     CountryRef `xml:"Country"`
}

// Tele holds contact telephone/email data for an organisation.
type Tele struct {
	PhoneNumber string    `xml:"phoneNumber,attr,omitempty"`
	MailAddress string    `xml:"mailAddress,attr,omitempty"`
	TeleType    ValueAttr `xml:"TeleType"`
}

// ParentOrganisation is a reference to the parent organisation.
type ParentOrganisation struct {
	OrganisationID string `xml:"OrganisationId"`
}

// Organisation is the Eventor native XML representation of an
// organisation (club, federation, etc.) as returned by
// `GET /api/organisations/{id}` and embedded in other documents.
type Organisation struct {
	XMLName              xml.Name           `xml:"Organisation"`
	ID                   string             `xml:"OrganisationId"`
	Name                 string             `xml:"Name"`
	ShortName            string             `xml:"ShortName,omitempty"`
	MediaName            string             `xml:"MediaName,omitempty"`
	OrganisationTypeID   string             `xml:"OrganisationTypeId,omitempty"`
	Country              *CountryRef        `xml:"Country"`
	Address              *Address           `xml:"Address"`
	Tele                 *Tele              `xml:"Tele"`
	ParentOrganisation   *ParentOrganisation `xml:"ParentOrganisation"`
	OrganisationStatusID string             `xml:"OrganisationStatusId,omitempty"`
	ModifyDate           *DateClock         `xml:"ModifyDate"`
}

// -----------------------------------------------------------------------
// Event
// -----------------------------------------------------------------------

// EventCenterPosition is the map centre geographic position of a race.
type EventCenterPosition struct {
	X    string `xml:"x,attr"`
	Y    string `xml:"y,attr"`
	Unit string `xml:"unit,attr,omitempty"`
}

// WRSInfo contains World Ranking Series information for an event race.
type WRSInfo struct {
	EventRaceID string `xml:"EventRaceId"`
	EventDate   string `xml:"EventDate,omitempty"`
	Distance    string `xml:"Distance,omitempty"`
}

// EventRace is a single race within a (possibly multi-race) event.
type EventRace struct {
	ID          string               `xml:"EventRaceId"`
	EventID     string               `xml:"EventId,omitempty"`
	Name        string               `xml:"Name,omitempty"`
	RaceDate    *DateClock           `xml:"RaceDate"`
	Position    *EventCenterPosition `xml:"EventCenterPosition"`
	WRSInfo     *WRSInfo             `xml:"WRSInfo"`
}

// ContactTele is the contact telephone/email inside a ContactData block.
type ContactTele struct {
	PhoneNumber string    `xml:"phoneNumber,attr,omitempty"`
	MailAddress string    `xml:"mailAddress,attr,omitempty"`
	TeleType    ValueAttr `xml:"TeleType"`
}

// ContactData is the event-level contact information.
type ContactData struct {
	Tele *ContactTele `xml:"Tele"`
}

// EventOfficial links a person to a role in an event.
type EventOfficial struct {
	ID         string `xml:"EventOfficialId"`
	RoleTypeID string `xml:"RoleTypeId,omitempty"`
	PersonID   string `xml:"PersonId,omitempty"`
	EventID    string `xml:"EventId,omitempty"`
}

// EventOrganiser wraps the organisation that organises the event.
type EventOrganiser struct {
	Organisation Organisation `xml:"Organisation"`
}

// HashTableEntry is a key/value metadata pair attached to an event.
// Eventor uses these for miscellaneous settings such as maximum entries,
// Livelox configuration, and result timestamps.
type HashTableEntry struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// Event is the Eventor native XML representation of a single event as
// returned by `GET /api/events/{id}`.
type Event struct {
	XMLName              xml.Name         `xml:"Event"`
	ID                   string           `xml:"EventId"`
	Name                 string           `xml:"Name"`
	ClassificationID     string           `xml:"EventClassificationId,omitempty"`
	StatusID             string           `xml:"EventStatusId,omitempty"`
	DisciplineID         string           `xml:"DisciplineId,omitempty"`
	StartDate            *DateClock       `xml:"StartDate"`
	FinishDate           *DateClock       `xml:"FinishDate"`
	EventOfficials       []EventOfficial  `xml:"EventOfficial"`
	Organisers           []EventOrganiser `xml:"Organiser"`
	ClassTypeID          string           `xml:"ClassTypeId,omitempty"`
	EventRaces           []EventRace      `xml:"EventRace"`
	ContactData          *ContactData     `xml:"ContactData"`
	PunchingUnitType     *ValueAttr       `xml:"PunchingUnitType"`
	HashTable            []HashTableEntry `xml:"HashTableEntry"`
}

// -----------------------------------------------------------------------
// EventClassList
// -----------------------------------------------------------------------

// ClassType is the category/type of a competition class.
type ClassType struct {
	ID        string `xml:"ClassTypeId"`
	ShortName string `xml:"ShortName,omitempty"`
	Name      string `xml:"Name"`
}

// ClassRaceInfo carries per-race details for a competition class.
type ClassRaceInfo struct {
	MinRunners  int       `xml:"minRunners,attr,omitempty"`
	MaxRunners  int       `xml:"maxRunners,attr,omitempty"`
	NoOfEntries int       `xml:"noOfEntries,attr,omitempty"`
	NoOfStarts  int       `xml:"noOfStarts,attr,omitempty"`
	ID          string    `xml:"ClassRaceInfoId"`
	EventRaceID string    `xml:"EventRaceId"`
	Name        string    `xml:"Name,omitempty"`
	Status      ValueAttr `xml:"ClassRaceStatus"`
}

// EventClass is a competition class within an event.
type EventClass struct {
	LowAge          int             `xml:"lowAge,attr,omitempty"`
	HighAge         int             `xml:"highAge,attr,omitempty"`
	Sex             string          `xml:"sex,attr,omitempty"`
	NumberOfEntries int             `xml:"numberOfEntries,attr,omitempty"`
	ID              string          `xml:"EventClassId"`
	Name            string          `xml:"Name"`
	ShortName       string          `xml:"ClassShortName,omitempty"`
	Status          ValueAttr       `xml:"EventClassStatus"`
	ClassType       *ClassType      `xml:"ClassType"`
	ExternalID      string          `xml:"ExternalId,omitempty"`
	PunchingUnitType ValueAttr      `xml:"PunchingUnitType"`
	ClassRaceInfos  []ClassRaceInfo `xml:"ClassRaceInfo"`
}

// EventClassList is the document returned by `GET /api/eventclasses`.
type EventClassList struct {
	XMLName      xml.Name     `xml:"EventClassList"`
	EventClasses []EventClass `xml:"EventClass"`
}

// -----------------------------------------------------------------------
// EntryFeeList
// -----------------------------------------------------------------------

// Amount is a monetary value with an optional ISO 4217 currency code.
type Amount struct {
	Currency string `xml:"currency,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// EntryFee is a single fee option that competitors can be charged.
type EntryFee struct {
	TaxIncluded   string    `xml:"taxIncluded,attr,omitempty"`
	EntryFeeType  string    `xml:"entryFeeType,attr,omitempty"`
	Type          string    `xml:"type,attr,omitempty"`
	ID            string    `xml:"EntryFeeId"`
	Name          string    `xml:"Name"`
	Amount        *Amount   `xml:"Amount"`
	ExternalFee   *Amount   `xml:"ExternalFee"`
	ValidToDate   *DateClock `xml:"ValidToDate"`
	FromDateOfBirth *DateClock `xml:"FromDateOfBirth"`
	ToDateOfBirth   *DateClock `xml:"ToDateOfBirth"`
}

// EntryFeeList is the document returned by `GET /api/entryfees/event`.
type EntryFeeList struct {
	XMLName    xml.Name   `xml:"EntryFeeList"`
	EntryFees  []EntryFee `xml:"EntryFee"`
}

// -----------------------------------------------------------------------
// DocumentList
// -----------------------------------------------------------------------

// Document is a file attachment for an event (e.g. invitation, results).
type Document struct {
	ID          string `xml:"id,attr"`
	ReferenceID string `xml:"referenceId,attr,omitempty"`
	Name        string `xml:"name,attr"`
	URL         string `xml:"url,attr"`
	ModifyDate  string `xml:"modifyDate,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
}

// DocumentList is the document returned by `GET /api/events/documents`.
type DocumentList struct {
	XMLName   xml.Name   `xml:"DocumentList"`
	Documents []Document `xml:"Document"`
}

// -----------------------------------------------------------------------
// CompetitorCountList
// -----------------------------------------------------------------------

// CompetitorCount holds entry and start counts for a single event.
type CompetitorCount struct {
	EventID         string `xml:"eventId,attr"`
	NumberOfEntries int    `xml:"numberOfEntries,attr,omitempty"`
	NumberOfStarts  int    `xml:"numberOfStarts,attr,omitempty"`
}

// CompetitorCountList is the document returned by
// `GET /api/competitorcount`.
type CompetitorCountList struct {
	XMLName xml.Name          `xml:"CompetitorCountList"`
	Counts  []CompetitorCount `xml:"CompetitorCount"`
}

// -----------------------------------------------------------------------
// EntryList (Eventor native)
// -----------------------------------------------------------------------

// CCard is a control card (e-punch) reference in an entry.
type CCard struct {
	ID              string    `xml:"CCardId"`
	PunchingUnitType ValueAttr `xml:"PunchingUnitType"`
}

// EntryCompetitor links a person and organisation to an entry.
type EntryCompetitor struct {
	ID             string     `xml:"CompetitorId"`
	PersonID       string     `xml:"PersonId,omitempty"`
	OrganisationID string     `xml:"OrganisationId,omitempty"`
	CCard          *CCard     `xml:"CCard"`
	ModifyDate     *DateClock `xml:"ModifyDate"`
}

// EntryClass links an entry to a competition class.
type EntryClass struct {
	Sequence    int    `xml:"sequence,attr,omitempty"`
	EventClassID string `xml:"EventClassId"`
}

// Entry is a single competitor entry for an event.
type Entry struct {
	ID          string           `xml:"EntryId"`
	Competitor  *EntryCompetitor `xml:"Competitor"`
	EntryClass  *EntryClass      `xml:"EntryClass"`
	EventID     string           `xml:"EventId,omitempty"`
	EventRaceID string           `xml:"EventRaceId,omitempty"`
	BibNumber   string           `xml:"BibNumber,omitempty"`
	EntryDate   *DateClock       `xml:"EntryDate"`
	ModifyDate  *DateClock       `xml:"ModifyDate"`
}

// EntryList is the Eventor native entry list returned by
// `GET /api/entries`. This is distinct from the IOF v3 EntryList.
type EntryList struct {
	XMLName xml.Name `xml:"EntryList"`
	Entries []Entry  `xml:"Entry"`
}
