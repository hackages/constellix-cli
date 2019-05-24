package functions

import (
	"encoding/json"
	"fmt"
	"strings"
)

var tabWithoutValue []TabsWithoutValue
var tabWithArrayValue []TabsWithArrayValue
var tabWithObjectValue []TabsWithObjectValue

// Get id of record by its value(/name for CNAME records)
func GetId(value string, c Config) int {
	urlWithObjectValue := BASE_URL + c.Constellix.Domain + "/records/" + NS + OPTIONS
	urlWithArrayValue := BASE_URL + c.Constellix.Domain + "/records/" + A + OPTIONS
	urlWithoutValue := BASE_URL + c.Constellix.Domain + "/records/" + CNAME + OPTIONS
	payload := strings.NewReader("")

	valueID := 0

	getAllForType(urlWithoutValue, payload, c, "tabWithoutValue")
	getAllForType(urlWithObjectValue, payload, c, "tabWithObjectValue")
	getAllForType(urlWithArrayValue, payload, c, "tabWithArrayValue")

	for _, i := range tabWithArrayValue {
		if i.Value[0] == value {
			valueID = i.Id
		}
	}

	if valueID == 0 {
		for _, i := range tabWithObjectValue {
			if i.Value[0].Value == value {
				valueID = i.Id
			}
		}
	}

	if valueID == 0 {
		for _, i := range tabWithoutValue {
			if i.Name == value {
				valueID = i.Id
			}
		}
	}

	return valueID
}

// Get id of record by its name
func GetIdByName(name string, c Config) int {
	urlWithObjectValue := BASE_URL + c.Constellix.Domain + "/records/" + NS + OPTIONS
	urlWithArrayValue := BASE_URL + c.Constellix.Domain + "/records/" + A + OPTIONS

	payload := strings.NewReader("")

	valueID := 0

	getAllForType(urlWithObjectValue, payload, c, "tabWithObjectValue")
	getAllForType(urlWithArrayValue, payload, c, "tabWithArrayValue")

	for _, i := range tabWithArrayValue {
		if i.Name == name {
			valueID = i.Id
		}
	}

	if valueID == 0 {
		for _, i := range tabWithObjectValue {
			if i.Name == name {
				valueID = i.Id
			}
		}
	}

	return valueID
}

func GetAll(c Config) {
	urlWithoutValue := BASE_URL + c.Constellix.Domain + "/records/" + CNAME + OPTIONS
	urlWithObjectValue := BASE_URL + c.Constellix.Domain + "/records/" + NS + OPTIONS
	urlWithArrayValue := BASE_URL + c.Constellix.Domain + "/records/" + A + OPTIONS

	payload := strings.NewReader("")

	getAllForType(urlWithoutValue, payload, c, "tabWithoutValue")
	getAllForType(urlWithObjectValue, payload, c, "tabWithObjectValue")
	getAllForType(urlWithArrayValue, payload, c, "tabWithArrayValue")

	for _, i := range tabWithoutValue {
		fmt.Printf("%v\t %v\t %v\t\t %v\n", i.Id, i.Type, i.Name, i.ParentId)
	}

	for _, i := range tabWithObjectValue {
		fmt.Printf("%v\t %v\t %v\t %v\t %v\n", i.Id, i.Type, i.Name, i.Value[0].Value, i.ParentId)
	}

	for _, i := range tabWithArrayValue {
		fmt.Printf("%v\t %v\t %v\t %v\t %v\n", i.Id, i.Type, i.Name, i.Value, i.ParentId)
	}
}

func getAllForType(recordTypeUrl string, payload *strings.Reader, c Config, tab string) {
	jsonBlob := send(recordTypeUrl, GET, payload, c)
	var err error

	if tab == "tabWithoutValue" {
		err = json.Unmarshal([]byte(jsonBlob), &tabWithoutValue)
	} else if tab == "tabWithObjectValue" {
		err = json.Unmarshal([]byte(jsonBlob), &tabWithObjectValue)
	} else {
		err = json.Unmarshal([]byte(jsonBlob), &tabWithArrayValue)
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
}
