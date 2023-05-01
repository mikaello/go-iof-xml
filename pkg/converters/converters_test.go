package converters

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mikaello/go-iof-xml/pkg/iof_v2"
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

func TestUnmarshalAllV2Types(t *testing.T) {
	dirPath := "../../test/v2-examples"
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
		_, err = GenericUnmarshalV2Xml(string(file))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUnmarshalV2ResultList(t *testing.T) {
	file := readFile("../../test/v2-examples/ResultList_example3.xml")
	result, err := GenericUnmarshalV2Xml(string(file))
	if err != nil {
		t.Fatal(err)
	}
	resultList := result.(*iof_v2.ResultList)

	if resultList.ClassResult[0].TeamResult[0].TeamStatus.ValueAttr != "OK" {
		t.Errorf("Expected result list event first team status to be 'OK', found %s", resultList.ClassResult[0].TeamResult[0].TeamStatus.ValueAttr)
	}
	if resultList.ClassResult[0].TeamResult[1].TeamStatus.ValueAttr != "Active" {
		t.Errorf("Expected result list event second team status to be 'Active', found %s", resultList.ClassResult[0].TeamResult[1].TeamStatus.ValueAttr)
	}
	if resultList.ClassResult[0].TeamResult[1].Time.Value != "55:03" {
		t.Errorf("Expected result list event second team result time to be '55:03', found %s", resultList.ClassResult[0].TeamResult[1].Time.Value)
	}
}

func TestUnmarshalV2ClubList(t *testing.T) {
	file := readFile("../../test/v2-examples/ClubList_example.xml")
	result, err := GenericUnmarshalV2Xml(string(file))
	if err != nil {
		t.Fatal(err)
	}
	resultList := result.(*iof_v2.ClubList)

	if resultList.Club[0].Address[0].CityAttr != "Højbjerg" {
		t.Errorf("Expected first club in club list to have address city 'Højbjerg', found %s", resultList.Club[0].Address[0].CityAttr)
	}
}

func readFile(fileName string) []byte {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return byteValue
}
