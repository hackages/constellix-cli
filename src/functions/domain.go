package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get all domains
func GetAllDomain(c Config) []DomainTabs {
	url := BASE_URL + OPTIONS
	payload := strings.NewReader("")

	var tab []DomainTabs

	jsonBlob := send(url, GET, payload, c)
	err := json.Unmarshal([]byte(jsonBlob), &tab)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return tab
}

// Create a domain with one or more names
func CreateDomain(name []string, c Config) {
	stringReader := "{\n  \"names\": [\n"
	for _, i := range name {
		stringReader += "    \"" + i + "\",\n"
	}
	stringReader += "  ]\n}"
	payload := strings.NewReader(stringReader)

	send(BASE_URL, POST, payload, c)
}

func DeleteDomain(isId bool, value string, c Config) {
	payload := strings.NewReader("")

	if isId {
		deleteDomainById(value, payload, c)
	} else {
		deleteDomainByName(value, payload, c)
	}
}

// Delete domain by its name
func deleteDomainByName(name string, payload *strings.Reader, c Config) {
	id := name
	url := BASE_URL + id

	send(url, DELETE, payload, c)
}

// Delete domain by its id
func deleteDomainById(id string, payload *strings.Reader, c Config) {
	url := BASE_URL + id

	send(url, DELETE, payload, c)
}

// Get a specific domain
func GetDomain(name string, c Config) {
	domainId := GetDomainId(name, c)
	url := BASE_URL + strconv.Itoa(domainId)
	payload := strings.NewReader("")

	var domainTab []DomainTabs

	jsonBlob := send(url, GET, payload, c)

	err := json.Unmarshal([]byte(jsonBlob), &domainTab)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, i := range domainTab {
		fmt.Printf("%v\t %v\n", i.Id, i.Name)
	}
}

// Get a domain id by its name
func GetDomainId(name string, c Config) int {
	domains := GetAllDomain(c)

	domainId := 0
	for _, i := range domains {
		if i.Name == name {
			domainId = i.Id
		}
	}

	if domainId == 0 {
		fmt.Println("Error: Domain name does not exist")
		os.Exit(1)
	}
	return domainId
}
