package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteCNAME(isValueId bool, value string, c Config) {
	payload := strings.NewReader("")

	if isValueId {
		deleteCnameById(value, payload, c)
	} else {
		deleteCnameByName(value, payload, c)
	}
}

// Delete a CNAME record by its name
func deleteCnameByName(name string, payload *strings.Reader, c Config) {
	cnameID := GetId(name, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname/" + strconv.Itoa(cnameID) + ""

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong name -> ", url)
	} else {
		fmt.Println("CNAME record deleted successfully ", url)
	}
}

// Delete a CNAME record by its id
func deleteCnameById(cnameID string, payload *strings.Reader, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname/" + cnameID

	send(url, DELETE, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("CNAME record deleted successfully ", url)
	}
}

// Create a CNAME record
func CreateCNAME(cname string, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname"
	payload := strings.NewReader("{\n  \"name\": \"" + cname + "\",\n  \"ttl\": \"1800\"\n }")

	send(url, POST, payload, c)
}
