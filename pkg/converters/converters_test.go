package converters

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnmarshalAllTypes(t *testing.T) {
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
		_ = GenericUnmarshalV3Xml(string(file))
	}
}

func TestUnmarshalResultList(t *testing.T) {
	file := readFile("../../test/v3-examples/ResultList3.xml")
	resultList := UnmarshalResultListV3(file)

	if resultList.Event.Name != "Example event" {
		t.Fatalf("Expected result list event name to be 'Example Event', found %s", resultList.Event.Name)
	}
	if resultList.Event.StartTime.Time.Format("15:04:05-07:00") != "10:00:00+01:00" {
		t.Fatalf("Expected result list event start time to be '10:00:00+01:00', found %s", resultList.Event.StartTime.Time.Format("15:04:05-07:00"))
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
