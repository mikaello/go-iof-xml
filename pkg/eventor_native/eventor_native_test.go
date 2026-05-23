package eventor_native_test

import (
	"os"
	"testing"

	"github.com/mikaello/go-iof-xml/pkg/eventor_native"
)

func fixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatalf("read fixture %q: %v", name, err)
	}
	return data
}

func TestDecodeOrganisation(t *testing.T) {
	org, err := eventor_native.DecodeOrganisation(fixture(t, "organisation.xml"))
	if err != nil {
		t.Fatalf("DecodeOrganisation: %v", err)
	}

	if org.ID != "273" {
		t.Errorf("ID = %q, want %q", org.ID, "273")
	}
	if org.Name == "" {
		t.Error("Name is empty")
	}
	if org.ShortName == "" {
		t.Error("ShortName is empty")
	}
	if org.Country == nil {
		t.Fatal("Country is nil")
	}
	if org.Country.Alpha3.Value != "NOR" {
		t.Errorf("Country.Alpha3 = %q, want %q", org.Country.Alpha3.Value, "NOR")
	}
	if org.Country.EnglishName() == "" {
		t.Error("Country.EnglishName() is empty")
	}
	if org.Address == nil {
		t.Fatal("Address is nil")
	}
	if org.Address.City == "" {
		t.Error("Address.City is empty")
	}
	if org.Address.ZipCode == "" {
		t.Error("Address.ZipCode is empty")
	}
	if org.Tele == nil {
		t.Fatal("Tele is nil")
	}
	if org.Tele.MailAddress == "" {
		t.Error("Tele.MailAddress is empty")
	}
	if org.ParentOrganisation == nil {
		t.Error("ParentOrganisation is nil")
	}
}

func TestDecodeEvent(t *testing.T) {
	ev, err := eventor_native.DecodeEvent(fixture(t, "event.xml"))
	if err != nil {
		t.Fatalf("DecodeEvent: %v", err)
	}

	if ev.ID == "" {
		t.Error("Event ID is empty")
	}
	if ev.Name == "" {
		t.Error("Event Name is empty")
	}
	if ev.StartDate == nil {
		t.Error("StartDate is nil")
	} else if ev.StartDate.Date == "" {
		t.Error("StartDate.Date is empty")
	}
	if len(ev.Organisers) == 0 {
		t.Error("no Organisers")
	} else if ev.Organisers[0].Organisation.Name == "" {
		t.Error("organiser Organisation.Name is empty")
	}
	if len(ev.EventRaces) == 0 {
		t.Error("no EventRaces")
	} else {
		race := ev.EventRaces[0]
		if race.ID == "" {
			t.Error("EventRace.ID is empty")
		}
		if race.RaceDate == nil || race.RaceDate.Date == "" {
			t.Error("EventRace.RaceDate is missing")
		}
	}
	if len(ev.HashTable) == 0 {
		t.Error("HashTable is empty")
	}
	// Verify Livelox config is present in hash table
	found := false
	for _, e := range ev.HashTable {
		if e.Key == "Eventor_LiveloxEventConfigurations" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Eventor_LiveloxEventConfigurations not found in HashTable")
	}
}

func TestDecodeEventClassList(t *testing.T) {
	ecl, err := eventor_native.DecodeEventClassList(fixture(t, "event_class_list.xml"))
	if err != nil {
		t.Fatalf("DecodeEventClassList: %v", err)
	}

	if len(ecl.EventClasses) == 0 {
		t.Fatal("EventClasses is empty")
	}
	first := ecl.EventClasses[0]
	if first.ID == "" {
		t.Error("EventClass.ID is empty")
	}
	if first.Name == "" {
		t.Error("EventClass.Name is empty")
	}
	if first.LowAge == 0 && first.HighAge == 0 {
		t.Error("EventClass age range not parsed")
	}
	if first.ClassType == nil {
		t.Error("EventClass.ClassType is nil")
	}
	if len(first.ClassRaceInfos) == 0 {
		t.Error("EventClass.ClassRaceInfos is empty")
	} else {
		info := first.ClassRaceInfos[0]
		if info.ID == "" {
			t.Error("ClassRaceInfo.ID is empty")
		}
		if info.EventRaceID == "" {
			t.Error("ClassRaceInfo.EventRaceID is empty")
		}
	}
}

func TestDecodeEntryFeeList(t *testing.T) {
	efl, err := eventor_native.DecodeEntryFeeList(fixture(t, "entry_fee_list.xml"))
	if err != nil {
		t.Fatalf("DecodeEntryFeeList: %v", err)
	}

	if len(efl.EntryFees) == 0 {
		t.Fatal("EntryFees is empty")
	}
	for _, fee := range efl.EntryFees {
		if fee.ID == "" {
			t.Errorf("EntryFee.ID is empty for fee %q", fee.Name)
		}
		if fee.Name == "" {
			t.Error("EntryFee.Name is empty")
		}
		if fee.Amount == nil {
			t.Errorf("EntryFee.Amount is nil for fee %q", fee.Name)
		} else if fee.Amount.Currency == "" {
			t.Errorf("EntryFee.Amount.Currency is empty for fee %q", fee.Name)
		}
		if fee.ValidToDate == nil {
			t.Errorf("EntryFee.ValidToDate is nil for fee %q", fee.Name)
		}
	}
	// Verify age-bounded fees parse correctly
	var hasFrom, hasTo bool
	for _, fee := range efl.EntryFees {
		if fee.FromDateOfBirth != nil {
			hasFrom = true
		}
		if fee.ToDateOfBirth != nil {
			hasTo = true
		}
	}
	if !hasFrom {
		t.Error("no EntryFee with FromDateOfBirth found")
	}
	if !hasTo {
		t.Error("no EntryFee with ToDateOfBirth found")
	}
}

func TestDecodeDocumentList(t *testing.T) {
	dl, err := eventor_native.DecodeDocumentList(fixture(t, "document_list.xml"))
	if err != nil {
		t.Fatalf("DecodeDocumentList: %v", err)
	}

	if len(dl.Documents) == 0 {
		t.Fatal("Documents is empty")
	}
	for _, doc := range dl.Documents {
		if doc.ID == "" {
			t.Errorf("Document.ID is empty for %q", doc.Name)
		}
		if doc.Name == "" {
			t.Error("Document.Name is empty")
		}
		if doc.URL == "" {
			t.Errorf("Document.URL is empty for %q", doc.Name)
		}
		if doc.Type == "" {
			t.Errorf("Document.Type is empty for %q", doc.Name)
		}
	}
}

func TestDecodeCompetitorCountList(t *testing.T) {
	ccl, err := eventor_native.DecodeCompetitorCountList(fixture(t, "competitor_count_list.xml"))
	if err != nil {
		t.Fatalf("DecodeCompetitorCountList: %v", err)
	}

	if len(ccl.Counts) == 0 {
		t.Fatal("Counts is empty")
	}
	count := ccl.Counts[0]
	if count.EventID == "" {
		t.Error("CompetitorCount.EventID is empty")
	}
	if count.NumberOfEntries == 0 {
		t.Error("CompetitorCount.NumberOfEntries is 0")
	}
}

func TestDecodeEntryList(t *testing.T) {
	el, err := eventor_native.DecodeEntryList(fixture(t, "entry_list.xml"))
	if err != nil {
		t.Fatalf("DecodeEntryList: %v", err)
	}

	if len(el.Entries) == 0 {
		t.Fatal("Entries is empty")
	}
	for _, entry := range el.Entries {
		if entry.ID == "" {
			t.Error("Entry.ID is empty")
		}
		if entry.Competitor == nil {
			t.Errorf("Entry %q: Competitor is nil", entry.ID)
			continue
		}
		if entry.Competitor.ID == "" {
			t.Errorf("Entry %q: Competitor.ID is empty", entry.ID)
		}
		if entry.EntryClass == nil {
			t.Errorf("Entry %q: EntryClass is nil", entry.ID)
		} else if entry.EntryClass.EventClassID == "" {
			t.Errorf("Entry %q: EntryClass.EventClassID is empty", entry.ID)
		}
		if entry.EventID == "" {
			t.Errorf("Entry %q: EventID is empty", entry.ID)
		}
	}
}
