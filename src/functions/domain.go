package functions

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetAllDomain(c Config) {
	url := "https://api.dns.constellix.com/v1/domains?offset=0&max=10&sort=name&order=asc"
	method := "GET"
	payload := strings.NewReader("")

	var tab []TabsWithoutValue

	jsonBlob := send(url, method, payload, c)
	err := json.Unmarshal([]byte(jsonBlob), &tab)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := 0; i < len(tab); i++ {
		fmt.Printf("%v %v %v %v \n", tab[i].Id, tab[i].Type, tab[i].Name, tab[i].ParentId)
	}
}
