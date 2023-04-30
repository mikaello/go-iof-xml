package converters

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"

	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
)

func getMainElementName(xml string) string {
	re := regexp.MustCompile("<([A-Z][a-zA-Z]+)")
	match := re.FindStringSubmatch(xml)
	if match != nil {
		return match[1]
	}
	return ""
}

const UTF8_BOM = "\uFEFF"

func removeUTF8BOM(s string, mainElement string) string {
	if len(s) > 0 && s[0:len(UTF8_BOM)] == UTF8_BOM {
		fmt.Printf("WARNING: removing BOM from XML of type %s\n", mainElement)
		s = s[len(UTF8_BOM):]
	}
	return s
}

func cleanAndUnmarshal(dirtyXml string, xmlType interface{}) interface{} {
	xmlTypeName := reflect.TypeOf(xmlType).Elem().Name()
	xmlData := removeUTF8BOM(dirtyXml, xmlTypeName)

	fmt.Printf("Unmarshalling %s.\n", xmlTypeName)

	err := xml.Unmarshal([]byte(xmlData), xmlType)
	if err != nil {
		fmt.Printf("ERROR V3: Failed to unmarshal XML (%s): %s\n", xmlTypeName, err)
	}

	return xmlType
}

func GenericUnmarshalV3Xml(dirtyXml string) interface{} {
	mainElementName := getMainElementName(dirtyXml)
	xmlData := removeUTF8BOM(dirtyXml, mainElementName)

	type temp struct {
		value interface{}
	}

	actual := temp{}

	switch mainElementName {
	case "CompetitorList":
		actual.value = new(iof_v3.CompetitorList)
	case "OrganisationList":
		actual.value = new(iof_v3.OrganisationList)
	case "EventList":
		actual.value = new(iof_v3.EventList)
	case "ClassList":
		actual.value = new(iof_v3.ClassList)
	case "EntryList":
		actual.value = new(iof_v3.EntryList)
	case "CourseData":
		actual.value = new(iof_v3.CourseData)
	case "StartList":
		actual.value = new(iof_v3.StartList)
	case "ResultList":
		actual.value = new(iof_v3.ResultList)
	case "ServiceRequestList":
		actual.value = new(iof_v3.ServiceRequestList)
	case "ControlCardList":
		actual.value = new(iof_v3.ControlCardList)
	default:
		fmt.Printf("hei: %s.\n", mainElementName)
	}

	err := xml.Unmarshal([]byte(xmlData), &actual.value)
	if err != nil {
		fmt.Printf("ERROR V3: Failed to unmarshal XML (%s): %s\n", mainElementName, err)
	}

	return actual.value
}

func UnmarshalIofV3CompetitorList(xml string) iof_v3.CompetitorList {
	result := iof_v3.CompetitorList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalResultListV3(xmlobj []byte) iof_v3.ResultList {
	result := iof_v3.ResultList{}
	err := xml.Unmarshal(xmlobj, &result)
	fmt.Println(err)

	/*
		cleanAndUnmarshal(xml, &result)
	*/
	return result
}

func main() {
	// Open our xmlFile
	xmlFile, err := os.Open("users.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var resultList iof_v3.ResultList
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &resultList)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(resultList.ClassResult); i++ {
		fmt.Println("User Name: " + resultList.ClassResult[i].Class.ModifyTimeAttr)
	}
}
