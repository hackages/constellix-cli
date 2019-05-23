package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteA(name string, c Config) {
	aID := GetId(name, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/A/" + strconv.Itoa(aID) + ""
	method := "DELETE"
	payload := strings.NewReader("")

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
