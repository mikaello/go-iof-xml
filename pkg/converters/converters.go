package converters

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"

	"github.com/mikaello/go-iof-xml/pkg/iof_v2"
	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
	"golang.org/x/net/html/charset"
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

func GenericUnmarshalV2Xml(dirtyXml string) interface{} {
	mainElementName := getMainElementName(dirtyXml)
	xmlData := removeUTF8BOM(dirtyXml, mainElementName)

	type temp struct {
		value interface{}
	}

	actual := temp{}

	switch mainElementName {
	case "PersonList":
		actual.value = new(iof_v2.PersonList)
	case "CompetitorList":
		actual.value = new(iof_v2.CompetitorList)
	case "RankList":
		actual.value = new(iof_v2.RankList)
	case "ClubList":
		actual.value = new(iof_v2.ClubList)
	case "EventList":
		actual.value = new(iof_v2.EventList)
	case "ServiceRequestList":
		actual.value = new(iof_v2.ServiceRequestList)
	case "EntryList":
		actual.value = new(iof_v2.EntryList)
	case "StartList":
		actual.value = new(iof_v2.StartList)
	case "ResultList":
		actual.value = new(iof_v2.ResultList)
	case "ClassData":
		actual.value = new(iof_v2.ClassData)
	case "CourseData":
		actual.value = new(iof_v2.CourseData)
	default:
		fmt.Printf("hei: %s.\n", mainElementName)

	}

	decoder := xml.NewDecoder(bytes.NewReader([]byte(xmlData)))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&actual.value)

	if err != nil {
		fmt.Printf("ERROR V3: Failed to unmarshal XML (%s): %s\n", mainElementName, err)
	}

	return actual.value
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

func UnmarshalIofV3OrganisationList(xml string) iof_v3.OrganisationList {
	result := iof_v3.OrganisationList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalIofV3EventList(xml string) iof_v3.EventList {
	result := iof_v3.EventList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalIofV3ClassList(xml string) iof_v3.ClassList {
	result := iof_v3.ClassList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalIofV3EntryList(xml string) iof_v3.EntryList {
	result := iof_v3.EntryList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalIofV3CourseData(xml string) iof_v3.CourseData {
	result := iof_v3.CourseData{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalStartListV3(xml string) iof_v3.StartList {
	result := iof_v3.StartList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalResultListV3(xml string) iof_v3.ResultList {
	result := iof_v3.ResultList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalServiceRequestListV3(xml string) iof_v3.ServiceRequestList {
	result := iof_v3.ServiceRequestList{}
	cleanAndUnmarshal(xml, &result)
	return result
}

func UnmarshalServiceControlCardListV3(xml string) iof_v3.ControlCardList {
	result := iof_v3.ControlCardList{}
	cleanAndUnmarshal(xml, &result)
	return result
}
