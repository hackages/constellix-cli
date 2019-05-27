package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteNS(isValueId bool, isValueName bool, value string, c Config) {
	method := "DELETE"
	payload := strings.NewReader("")

	if isValueId {
		deleteNsById(value, method, payload, c)
	} else if isValueName {
		deleteNsByName(value, method, payload, c)
	} else {
		deleteNsByValue(value, method, payload, c)
	}
}

func deleteNsByValue(value string, method string, payload *strings.Reader, c Config) {
	var validValue = value
	var nsID = 0

	if !strings.HasSuffix(".", value) {
		validValue += "."
		nsID = GetId(validValue, c)
	} else {
		nsID = GetId(value, c)
	}

	url := BASE_URL + c.Constellix.Domain + "/records/ns/" + strconv.Itoa(nsID) + ""

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

func deleteNsByName(name string, method string, payload *strings.Reader, c Config) {
	nsID := GetIdByName(name, c)

	url := BASE_URL + c.Constellix.Domain + "/records/ns/" + strconv.Itoa(nsID) + ""

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

func deleteNsById(nsID string, method string, payload *strings.Reader, c Config) {
	url := BASE_URL + c.Constellix.Domain + "/records/ns/" + nsID

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

func CreateNS(value string, name string, c Config) {
	url := BASE_URL + c.Constellix.Domain + "/records/ns"
	method := "POST"

	if strings.HasSuffix(".", value) {
		payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + value + "\"\n }\n]}")
		send(url, method, payload, c)
	} else {
		validValue := value + "."
		payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + validValue + "\"\n }\n]}")
		send(url, method, payload, c)
	}
}
