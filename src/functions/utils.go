package functions

import (
	"encoding/json"
	"fmt"
	"strings"
)

const BASE_URL = "https://api.dns.constellix.com/v1/domains/"
const OPTIONS = "?%20&offset=0&max=10&sort=name&order=asc"

const CNAME = "cname"
const A = "A"
const NS = "ns"

var tabWithoutValue []TabsWithoutValue
var tabWithArrayValue []TabsWithArrayValue
var tabWithObjectValue []TabsWithObjectValue

func GetId(value string, c Config) int {
	urlWithObjectValue := BASE_URL + c.Constellix.Domain + "/records/" + NS + OPTIONS
	urlWithArrayValue := BASE_URL + c.Constellix.Domain + "/records/" + A + OPTIONS
	urlWithoutValue := BASE_URL + c.Constellix.Domain + "/records/" + CNAME + OPTIONS

	method := "GET"
	payload := strings.NewReader("")

	valueID := 0

	getAllForType(urlWithoutValue, method, payload, c, "tabWithoutValue")
	getAllForType(urlWithObjectValue, method, payload, c, "tabWithObjectValue")
	getAllForType(urlWithArrayValue, method, payload, c, "tabWithArrayValue")

	for i := 0; i < len(tabWithArrayValue); i++ {
		if tabWithArrayValue[i].Value[0] == value {
			valueID = tabWithArrayValue[i].Id
		}
	}

	if valueID == 0 {
		for i := 0; i < len(tabWithObjectValue); i++ {
			if tabWithObjectValue[i].Value[0].Value == value {
				valueID = tabWithObjectValue[i].Id
			}
		}
	}

	if valueID == 0 {
		for i := 0; i < len(tabWithoutValue); i++ {
			if tabWithoutValue[i].Name == value {
				valueID = tabWithoutValue[i].Id
			}
		}
	}

	return valueID
}

func GetAll(c Config) {
	urlWithoutValue := BASE_URL + c.Constellix.Domain + "/records/" + CNAME + OPTIONS
	urlWithObjectValue := BASE_URL + c.Constellix.Domain + "/records/" + NS + OPTIONS
	urlWithArrayValue := BASE_URL + c.Constellix.Domain + "/records/" + A + OPTIONS

	method := "GET"
	payload := strings.NewReader("")

	getAllForType(urlWithoutValue, method, payload, c, "tabWithoutValue")
	getAllForType(urlWithObjectValue, method, payload, c, "tabWithObjectValue")
	getAllForType(urlWithArrayValue, method, payload, c, "tabWithArrayValue")

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

func getAllForType(recordTypeUrl string, method string, payload *strings.Reader, c Config, tab string) {
	jsonBlob := send(recordTypeUrl, method, payload, c)
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
