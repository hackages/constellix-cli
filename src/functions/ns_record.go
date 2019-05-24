package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteNS(isValueId bool, isValueName bool, value string, c Config) {
	payload := strings.NewReader("")

	if isValueId {
		deleteNsById(value, payload, c)
	} else if isValueName {
		deleteNsByName(value, payload, c)
	} else {
		deleteNsByValue(value, payload, c)
	}
}

// Delete NS record by its DNS domain value
func deleteNsByValue(value string, payload *strings.Reader, c Config) {
	var validValue = value
	var nsID = 0

	if !strings.HasSuffix(".", value) {
		validValue += "."
		nsID = GetId(validValue, c)
	} else {
		nsID = GetId(value, c)
	}

	url := BASE_URL + c.Constellix.Domain + "/records/ns/" + strconv.Itoa(nsID) + ""

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

// Delete NS record by its name
func deleteNsByName(name string, payload *strings.Reader, c Config) {
	nsID := GetIdByName(name, c)

	url := BASE_URL + c.Constellix.Domain + "/records/ns/" + strconv.Itoa(nsID) + ""

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

// Delete NS record by its id
func deleteNsById(nsID string, payload *strings.Reader, c Config) {
	url := BASE_URL + c.Constellix.Domain + "/records/ns/" + nsID

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

// Create NS record with an optional name
func CreateNS(value string, name string, c Config) {
	url := BASE_URL + c.Constellix.Domain + "/records/ns"

	if strings.HasSuffix(".", value) {
		payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + value + "\"\n }\n]}")
		send(url, POST, payload, c)
	} else {
		validValue := value + "."
		payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + validValue + "\"\n }\n]}")
		send(url, POST, payload, c)
	}
}
