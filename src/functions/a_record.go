package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteA(isValueId bool, isValueName bool, value string, c Config) {
	method := "DELETE"
	payload := strings.NewReader("")

	if isValueId {
		deleteAById(value, method, payload, c)
	} else if isValueName {
		deleteAByName(value, method, payload, c)
	} else {
		deleteAByValue(value, method, payload, c)
	}
}

func deleteAByValue(value string, method string, payload *strings.Reader, c Config) {
	aID := GetId(value, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + strconv.Itoa(aID) + ""

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("A record deleted successfully ", url)
	}
}

func deleteAByName(name string, method string, payload *strings.Reader, c Config) {
	aID := GetId(name, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + strconv.Itoa(aID) + ""

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("A record deleted successfully ", url)
	}
}

func deleteAById(aID string, method string, payload *strings.Reader, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + aID

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("A record deleted successfully ", url)
	}
}

func CreateA(ip string, name string, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A"
	method := "POST"
	payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + ip + "\"\n }\n]}")

	send(url, method, payload, c)
}
