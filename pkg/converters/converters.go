package converters

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mikaello/go-iof-xml/pkg/iof_v3"
)

func UnmarshalResultListV3(file string) iof_v3.ResultList {
	xmlFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var resultList iof_v3.ResultList
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &resultList)

	return resultList

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
