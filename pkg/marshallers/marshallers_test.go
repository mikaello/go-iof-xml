package marshallers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestUnmarshalAllV3Types(t *testing.T) {
	dirPath := "../../test/v3-examples"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".xml") {
			continue
		}
		file := readFile(filepath.Join(dirPath, e.Name()))
		fmt.Println("Unmarshalling: " + e.Name())
		_, err = GenericUnmarshalV3Xml(string(file))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUnmarshalV3ResultList(t *testing.T) {
	file := readFile("../../test/v3-examples/ResultList3.xml")
	resultList := UnmarshalIofV3ResultList(string(file))

	if resultList.Event.Name != "Example event" {
		t.Errorf("Expected result list event name to be 'Example Event', found %s", resultList.Event.Name)
	}
	if resultList.Event.StartTime.Date != "2011-07-30" {
		t.Errorf("Expected result list event start date to be '2011-07-30', found %s", resultList.Event.StartTime.Date)
	}
	if resultList.Event.StartTime.Time == nil {
		t.Fatal("Expected result list event start time to be defined, was 'nil'")
	}
	if resultList.Event.StartTime.Time.Format("15:04:05-07:00") != "10:00:00+01:00" {
		t.Errorf("Expected result list event start time to be '10:00:00+01:00', found %s", resultList.Event.StartTime.Time.Format("15:04:05-07:00"))
	}
	if resultList.Event.StartTime.Time.Format(time.RFC3339) != "2011-07-30T10:00:00+01:00" {
		t.Errorf("Expected result list event start time to be '2011-07-30T10:00:00+01:00', found %s", resultList.Event.StartTime.Time.Format(time.RFC3339))
	}
}
func TestMarshalV3ResultList(t *testing.T) {
	file := readFile("../../test/v3-examples/ResultList3.xml")
	resultList := UnmarshalIofV3ResultList(string(file))
	_, err := MarshallIofXml(resultList)

	if err != nil {
		t.Fatalf("Could not marshal result list: %s", err)
	}
}
func TestMarshalV3CourseData(t *testing.T) {
	file := readFile("../../test/v3-examples/CourseData_Individual_Step2.xml")
	courseData := UnmarshalIofV3CourseData(string(file))
	_, err := MarshallIofXml(courseData)

	if err != nil {
		t.Fatalf("Could not marshal course data list: %s", err)
	}
}

func TestMarshalV3CourseDataToJson(t *testing.T) {
	file := readFile("../../test/v3-examples/StartList1.xml")
	courseData := UnmarshalIofV3StartList(string(file))
	_, err := MarshallToIofUnofficialJson(courseData)

	if err != nil {
		t.Fatalf("Could not marshal course data list: %s", err)
	}
}

func readFile(fileName string) []byte {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return byteValue
}
