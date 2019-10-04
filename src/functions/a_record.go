package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteA(isValueId bool, isValueName bool, value string, c Config) {
	payload := strings.NewReader("")

	if isValueId {
		deleteAById(value, payload, c)
	} else if isValueName {
		deleteAByName(value, payload, c)
	} else {
		deleteAByValue(value, payload, c)
	}
}

// Delete A record by its ip value
func deleteAByValue(value string, payload *strings.Reader, c Config) {
	aID := GetId(value, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + strconv.Itoa(aID) + ""

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("A record deleted successfully ", url)
	}
}

// Delete A record by its name
func deleteAByName(name string, payload *strings.Reader, c Config) {
	aID := GetId(name, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + strconv.Itoa(aID) + ""

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("A record deleted successfully ", url)
	}
}

// Delete A record by its id
func deleteAById(aID string, payload *strings.Reader, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + aID

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("A record deleted successfully ", url)
	}
}

// Create A record with an optional name
func CreateA(ip string, name string, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A"
	payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + ip + "\"\n }\n]}")

	send(url, POST, payload, c)
}
