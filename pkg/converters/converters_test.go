package converters

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestUnmarshalAllTypes(t *testing.T) {
	dirPath := "../../test/v3-examples"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		file := readFile(filepath.Join(dirPath, e.Name()))
		_ = GenericUnmarshalV3Xml(string(file))
	}
}

func TestUnmarshalResultList(t *testing.T) {
	file := readFile("../../test/v3-examples/ResultList4.xml")
	resultList := UnmarshalResultListV3(file)

	if resultList.Event.Name != "Example event" {
		t.Fatalf("Expected result list event name to be 'Example Event', found %s", resultList.Event.Name)
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
