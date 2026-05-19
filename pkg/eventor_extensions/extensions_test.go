package eventor_extensions_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mikaello/go-iof-xml/pkg/eventor_extensions"
)

func loadFixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}
	return data
}

func TestDecodeEventList(t *testing.T) {
	data := loadFixture(t, "EventList_with_eventor_extensions.xml")
	el, err := eventor_extensions.DecodeEventList(data)
	if err != nil {
		t.Fatalf("DecodeEventList: %v", err)
	}
	if len(el.Events) != 1 {
		t.Fatalf("got %d events, want 1", len(el.Events))
	}
	ev := el.Events[0]
	if got := ev.ID; got != "23514" {
		t.Errorf("event id = %q, want 23514", got)
	}
	if got := ev.Extensions.StartListExists; !got {
		t.Errorf("event StartListExists = %v, want true", got)
	}
	if got := ev.Extensions.ResultListExists; got {
		t.Errorf("event ResultListExists = %v, want false", got)
	}
	if got := ev.Extensions.Discipline; got != eventor_extensions.DisciplineFoot {
		t.Errorf("event Discipline = %q, want Foot", got)
	}
	// Light condition not present on the event-level extensions.
	if got := ev.Extensions.LightCondition; got != "" {
		t.Errorf("event LightCondition = %q, want empty", got)
	}
}

func TestDecodeEventListRaceExtensions(t *testing.T) {
	data := loadFixture(t, "EventList_with_eventor_extensions.xml")
	el, err := eventor_extensions.DecodeEventList(data)
	if err != nil {
		t.Fatalf("DecodeEventList: %v", err)
	}
	if len(el.Events) != 1 || len(el.Events[0].Races) != 1 {
		t.Fatalf("unexpected event/race count")
	}
	r := el.Events[0].Races[0]
	if got := r.Extensions.EventRaceID; got != "24332" {
		t.Errorf("race EventRaceID = %q, want 24332", got)
	}
	if got := r.Extensions.StartListExists; !got {
		t.Errorf("race StartListExists = %v, want true", got)
	}
	if got := r.Extensions.ResultListExists; !got {
		t.Errorf("race ResultListExists = %v, want true", got)
	}
	if got := r.Extensions.Discipline; got != eventor_extensions.DisciplineFoot {
		t.Errorf("race Discipline = %q, want Foot", got)
	}
	if got := r.Extensions.LightCondition; got != eventor_extensions.LightConditionDay {
		t.Errorf("race LightCondition = %q, want Day", got)
	}
}

func TestDecodeEventListNoExtensions(t *testing.T) {
	xml := []byte(`<?xml version="1.0" encoding="utf-8"?>
<EventList xmlns="http://www.orienteering.org/datastandard/3.0" iofVersion="3.0">
  <Event>
    <Id>1</Id>
    <Name>Plain event</Name>
  </Event>
</EventList>`)
	el, err := eventor_extensions.DecodeEventList(xml)
	if err != nil {
		t.Fatalf("DecodeEventList: %v", err)
	}
	if len(el.Events) != 1 {
		t.Fatalf("got %d events, want 1", len(el.Events))
	}
	ev := el.Events[0]
	if ev.Extensions.EventRaceID != "" || ev.Extensions.StartListExists || ev.Extensions.Discipline != "" {
		t.Errorf("expected zero-valued extensions on event with no <Extensions>, got %+v", ev.Extensions)
	}
}
