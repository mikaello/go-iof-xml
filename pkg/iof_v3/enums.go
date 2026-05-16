package iof_v3

// Constants for the enumerated typed-string fields defined in the IOF v3
// XSD. Use these instead of string literals so the compiler catches typos
// and renames; refactoring tooling will follow them across the codebase.
//
// The values mirror the XSD's <xsd:enumeration> entries exactly. See the
// "Related" links in doc.go for the canonical schema.

// EventForm values (Event/Form).
const (
	EventFormIndividual EventForm = "Individual"
	EventFormTeam       EventForm = "Team"
	EventFormRelay      EventForm = "Relay"
)

// EventStatus values (Event/Status).
const (
	EventStatusPlanned     EventStatus = "Planned"
	EventStatusApplied     EventStatus = "Applied"
	EventStatusProposed    EventStatus = "Proposed"
	EventStatusSanctioned  EventStatus = "Sanctioned"
	EventStatusCanceled    EventStatus = "Canceled"
	EventStatusRescheduled EventStatus = "Rescheduled"
)

// EventClassification values (Event/Classification).
const (
	EventClassificationInternational EventClassification = "International"
	EventClassificationNational      EventClassification = "National"
	EventClassificationRegional      EventClassification = "Regional"
	EventClassificationLocal         EventClassification = "Local"
	EventClassificationClub          EventClassification = "Club"
)

// RaceDiscipline values (Race/Discipline).
const (
	RaceDisciplineSprint    RaceDiscipline = "Sprint"
	RaceDisciplineMiddle    RaceDiscipline = "Middle"
	RaceDisciplineLong      RaceDiscipline = "Long"
	RaceDisciplineUltralong RaceDiscipline = "Ultralong"
	RaceDisciplineOther     RaceDiscipline = "Other"
)

// EventClassStatus values (Class/Status).
const (
	EventClassStatusNormal           EventClassStatus = "Normal"
	EventClassStatusDivided          EventClassStatus = "Divided"
	EventClassStatusJoined           EventClassStatus = "Joined"
	EventClassStatusInvalidated      EventClassStatus = "Invalidated"
	EventClassStatusInvalidatedNoFee EventClassStatus = "InvalidatedNoFee"
)

// RaceClassStatus values (RaceClass/Status).
const (
	RaceClassStatusStartTimesNotAllocated RaceClassStatus = "StartTimesNotAllocated"
	RaceClassStatusStartTimesAllocated    RaceClassStatus = "StartTimesAllocated"
	RaceClassStatusNotUsed                RaceClassStatus = "NotUsed"
	RaceClassStatusCompleted              RaceClassStatus = "Completed"
	RaceClassStatusInvalidated            RaceClassStatus = "Invalidated"
	RaceClassStatusInvalidatedNoFee       RaceClassStatus = "InvalidatedNoFee"
)

// ResultStatus values (PersonRaceResult/Status, TeamMemberRaceResult/Status).
const (
	ResultStatusOK                 ResultStatus = "OK"
	ResultStatusFinished           ResultStatus = "Finished"
	ResultStatusMissingPunch       ResultStatus = "MissingPunch"
	ResultStatusDisqualified       ResultStatus = "Disqualified"
	ResultStatusDidNotFinish       ResultStatus = "DidNotFinish"
	ResultStatusActive             ResultStatus = "Active"
	ResultStatusInactive           ResultStatus = "Inactive"
	ResultStatusOverTime           ResultStatus = "OverTime"
	ResultStatusSportingWithdrawal ResultStatus = "SportingWithdrawal"
	ResultStatusNotCompeting       ResultStatus = "NotCompeting"
	ResultStatusMoved              ResultStatus = "Moved"
	ResultStatusMovedUp            ResultStatus = "MovedUp"
	ResultStatusDidNotStart        ResultStatus = "DidNotStart"
	ResultStatusDidNotEnter        ResultStatus = "DidNotEnter"
	ResultStatusCancelled          ResultStatus = "Cancelled"
)

// ControlType values (Control/type attribute).
const (
	ControlTypeControl          ControlType = "Control"
	ControlTypeStart            ControlType = "Start"
	ControlTypeFinish           ControlType = "Finish"
	ControlTypeCrossingPoint    ControlType = "CrossingPoint"
	ControlTypeEndOfMarkedRoute ControlType = "EndOfMarkedRoute"
)
