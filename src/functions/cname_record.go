package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteCNAME(isValueId bool, value string, c Config) {
	method := "DELETE"
	payload := strings.NewReader("")

	if isValueId {
		deleteCnameById(value, method, payload, c)
	} else {
		deleteCnameByName(value, method, payload, c)
	}
}

func deleteCnameByName(name string, method string, payload *strings.Reader, c Config) {
	cnameID := GetId(name, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname/" + strconv.Itoa(cnameID) + ""

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong name -> ", url)
	} else {
		fmt.Println("CNAME record deleted successfully ", url)
	}
}

func deleteCnameById(cnameID string, method string, payload *strings.Reader, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname/" + cnameID

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("CNAME record deleted successfully ", url)
	}
}

func CreateCNAME(cname string, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname"
	method := "POST"
	payload := strings.NewReader("{\n  \"name\": \"" + cname + "\",\n  \"ttl\": \"1800\"\n }")

	send(url, method, payload, c)
}
