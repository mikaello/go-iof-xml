package converters

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

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
	resultList := UnmarshalResultListV3(string(file))

	if resultList.Event.Name != "Example event" {
		t.Fatalf("Expected result list event name to be 'Example Event', found %s", resultList.Event.Name)
	}
	if resultList.Event.StartTime.Time.Format("15:04:05-07:00") != "10:00:00+01:00" {
		t.Fatalf("Expected result list event start time to be '10:00:00+01:00', found %s", resultList.Event.StartTime.Time.Format("15:04:05-07:00"))
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
		t.Fatalf("Expected result list event first team status to be 'OK', found %s", resultList.ClassResult[0].TeamResult[0].TeamStatus.ValueAttr)
	}
	if resultList.ClassResult[0].TeamResult[1].TeamStatus.ValueAttr != "Active" {
		t.Fatalf("Expected result list event second team status to be 'Active', found %s", resultList.ClassResult[0].TeamResult[1].TeamStatus.ValueAttr)
	}
	if resultList.ClassResult[0].TeamResult[1].Time.Value != "55:03" {
		t.Fatalf("Expected result list event second team result time to be '55:03', found %s", resultList.ClassResult[0].TeamResult[1].Time.Value)
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
