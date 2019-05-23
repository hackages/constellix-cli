package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteNS(value string, c Config) {
	var validValue = value
	var nsID = 0

	if !strings.HasSuffix(".", value) {
		validValue += "."
		nsID = GetId(validValue, c)
	} else {
		nsID = GetId(value, c)
	}

	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/ns/" + strconv.Itoa(nsID) + ""
	method := "DELETE"
	payload := strings.NewReader("")

	send(url, method, payload, c)

	if filepath.Base(url) == "0" {
		fmt.Println("Error: Wrong id -> ", url)
	} else {
		fmt.Println("Domain record deleted successfully ", url)
	}
}

func CreateNS(value string, name string, c Config) {
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/ns"
	method := "POST"

	if !strings.HasSuffix(".", value) {
		validValue := value + "."
		payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + validValue + "\"\n }\n]}")
		send(url, method, payload, c)
	} else {
		payload := strings.NewReader("{\n  \"name\": \"" + name + "\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \"" + value + "\"\n }\n]}")
		send(url, method, payload, c)
	}
}
