package functions

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func DeleteCNAME(cname string, c Config) {
	cnameID := GetId(cname, c)
	url := "https://api.dns.constellix.com/v1/domains/" + c.Constellix.Domain + "/records/cname/" + strconv.Itoa(cnameID) + ""
	method := "DELETE"
	payload := strings.NewReader("")

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
