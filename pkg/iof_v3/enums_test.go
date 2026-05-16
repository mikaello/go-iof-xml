package iof_v3_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
)

// TestEnumConstantsAreXMLValid verifies the enum constants round-trip
// through encoding/xml so they are usable as drop-in replacements for the
// underlying string literals.
func TestEnumConstantsAreXMLValid(t *testing.T) {
	type result struct {
		XMLName xml.Name            `xml:"PersonRaceResult"`
		Status  iof_v3.ResultStatus `xml:"Status"`
	}

	want := iof_v3.ResultStatusMissingPunch

	encoded, err := xml.Marshal(result{Status: want})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !bytes.Contains(encoded, []byte("<Status>MissingPunch</Status>")) {
		t.Errorf("marshalled XML = %q, want it to contain <Status>MissingPunch</Status>", encoded)
	}

	var got result
	if err := xml.Unmarshal(encoded, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Status != want {
		t.Errorf("Status = %q, want %q", got.Status, want)
	}
}
