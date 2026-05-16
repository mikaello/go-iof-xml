package marshallers

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
)

const examplesDir = "../../test/v3-examples"

func TestDecodeAllExamples(t *testing.T) {
	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		t.Fatal(err)
	}

	var n int
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".xml") {
			continue
		}
		n++
		t.Run(e.Name(), func(t *testing.T) {
			data := readFile(t, filepath.Join(examplesDir, e.Name()))
			if _, err := Decode(data); err != nil {
				t.Fatalf("Decode: %v", err)
			}
		})
	}
	if n == 0 {
		t.Fatalf("no example files found in %s", examplesDir)
	}
}

func TestDecodeResultListFields(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "ResultList3.xml"))

	list, err := DecodeAs[iof_v3.ResultList](data)
	if err != nil {
		t.Fatalf("DecodeAs: %v", err)
	}

	if got, want := list.Event.Name, "Example event"; got != want {
		t.Errorf("Event.Name = %q, want %q", got, want)
	}
	if got, want := list.Event.StartTime.Date, "2011-07-30"; got != want {
		t.Errorf("Event.StartTime.Date = %q, want %q", got, want)
	}
	if list.Event.StartTime.Time == nil {
		t.Fatal("Event.StartTime.Time is nil")
	}
	if got, want := list.Event.StartTime.Time.Format(time.RFC3339), "2011-07-30T10:00:00+01:00"; got != want {
		t.Errorf("Event.StartTime.Time = %q, want %q", got, want)
	}
}

func TestDecodeReturnsCorrectType(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "ResultList3.xml"))

	got, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := got.(*iof_v3.ResultList); !ok {
		t.Fatalf("Decode returned %T, want *iof_v3.ResultList", got)
	}
}

func TestDecodeReaderEquivalentToDecode(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "StartList1.xml"))

	a, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := DecodeReader(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := a.(*iof_v3.StartList); !ok {
		t.Fatalf("Decode returned %T", a)
	}
	if _, ok := b.(*iof_v3.StartList); !ok {
		t.Fatalf("DecodeReader returned %T", b)
	}
}

func TestDecodeStripsUTF8BOM(t *testing.T) {
	raw := readFile(t, filepath.Join(examplesDir, "ResultList3.xml"))
	withBOM := append([]byte{0xEF, 0xBB, 0xBF}, raw...)

	if _, err := Decode(withBOM); err != nil {
		t.Fatalf("Decode with BOM: %v", err)
	}
}

func TestDecodeUnknownRootElement(t *testing.T) {
	_, err := Decode([]byte(`<?xml version="1.0"?><NotAnIofDocument/>`))
	if err == nil {
		t.Fatal("expected error for unknown root element")
	}
	if !strings.Contains(err.Error(), "unknown IOF v3 root element") {
		t.Errorf("error = %v, want unknown root element error", err)
	}
}

func TestDecodeEmptyInput(t *testing.T) {
	if _, err := Decode(nil); err == nil {
		t.Fatal("expected error for empty input")
	}
	if _, err := Decode([]byte{}); err == nil {
		t.Fatal("expected error for empty input")
	}
	if _, err := Decode([]byte("\xEF\xBB\xBF")); err == nil {
		t.Fatal("expected error for BOM-only input")
	}
}

func TestDecodeAsWrongType(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "ResultList3.xml"))
	_, err := DecodeAs[iof_v3.StartList](data)
	if err == nil {
		t.Fatal("expected error when decoding ResultList as StartList")
	}
}

func TestEncodeXMLRoundTrip(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "ResultList3.xml"))

	list, err := DecodeAs[iof_v3.ResultList](data)
	if err != nil {
		t.Fatal(err)
	}

	out, err := EncodeXML(list)
	if err != nil {
		t.Fatalf("EncodeXML: %v", err)
	}
	if !bytes.HasPrefix(out, []byte(`<?xml`)) {
		t.Errorf("EncodeXML output missing XML declaration: %q", out[:min(80, len(out))])
	}

	roundTripped, err := DecodeAs[iof_v3.ResultList](out)
	if err != nil {
		t.Fatalf("re-decode of encoded XML: %v", err)
	}
	if roundTripped.Event.Name != list.Event.Name {
		t.Errorf("Event.Name lost in round-trip: got %q, want %q", roundTripped.Event.Name, list.Event.Name)
	}
}

func TestEncodeXMLCourseData(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "CourseData_Individual_Step2.xml"))

	cd, err := DecodeAs[iof_v3.CourseData](data)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := EncodeXML(cd); err != nil {
		t.Fatalf("EncodeXML: %v", err)
	}
}

func TestEncodeJSON(t *testing.T) {
	data := readFile(t, filepath.Join(examplesDir, "StartList1.xml"))

	sl, err := DecodeAs[iof_v3.StartList](data)
	if err != nil {
		t.Fatal(err)
	}
	out, err := EncodeJSON(sl)
	if err != nil {
		t.Fatalf("EncodeJSON: %v", err)
	}
	if !bytes.HasPrefix(bytes.TrimSpace(out), []byte("{")) {
		t.Errorf("EncodeJSON did not produce a JSON object")
	}
}

func readFile(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return b
}
