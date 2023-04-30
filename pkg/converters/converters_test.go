package converters

import (
	"testing"
)

func TestUnmarshalResultList(t *testing.T) {
	resultList := UnmarshalResultListV3("../../ResultList2.xml")

	if resultList.Event.Name != "Example event" {
		t.Fatalf("Expected result list event name to be 'Example Event', found %s", resultList.Event.Name)
	}
}
