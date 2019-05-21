package functions

import (
  "path/filepath"
  "fmt"
  "strings"
  "strconv"
  "net/http"
  "io/ioutil"
  "time"
  "os"
  "crypto/hmac"
  "crypto/sha1"
  "encoding/base64"
  "encoding/json"
  "github.com/spf13/viper"
)

type Tabs struct {
        Id int
        Type string
        //RecordType string
        Name string
        //RecordOption string
        //NoAnswer bool
        //Note string
        //Ttl int
        //GtdRegion int
        ParentId int
        //Parent string
        //Source string
        //ModifiedTs uint
        Value string
        //RoundRobin []string
        //Geolocation int
        //RecordFailover string
        //Failover string
        //RoundRobinFailover []string
        //Pools []string
        //PoolsDetail []string
}

type ArecordTabs struct {
        Id int
        Type string
        //RecordType string
        Name string
        //RecordOption string
        //NoAnswer bool
        //Note string
        //Ttl int
        //GtdRegion int
        ParentId int
        //Parent string
        //Source string
        //ModifiedTs uint
        Value []string
        //RoundRobin []string
        //Geolocation int
        //RecordFailover string
        //Failover string
        //RoundRobinFailover []string
        //Pools []string
        //PoolsDetail []string
}

type KeysConfig struct {
    	ApiKey string `mapstructure:"apiKey"`
   	SecretKey string `mapstructure:"secretKey"`
	Domain string `mapstructure:"domain"`
}

type Config struct {
        Constellix KeysConfig `mapstructure:"constellix"`
}

func InitConfig(configFile string) Config {
	v := viper.New()
	
	configName := filepath.Base(configFile)
	name := strings.TrimSuffix(configName, filepath.Ext(configName))
	configPath := filepath.Dir(configFile)
	
	v.SetConfigName(name)
	v.AddConfigPath(configPath)
	
	if err := v.ReadInConfig(); err != nil {
        	fmt.Printf("couldn't load config: %s\n", err)
        	os.Exit(1)
	}	
	
	var c Config
	if err := v.Unmarshal(&c); err != nil {
        	fmt.Printf("1couldn't read config: %s\n", err)
        }
	
	return c
}

func DeleteCNAME(cname string, c Config){
        cnameID := GetId(cname,c)
        url := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/cname/"+strconv.Itoa(cnameID)+""
        method := "DELETE"
        payload := strings.NewReader("")
        send(url, method, payload, c)
        fmt.Println(url)
}

func DeleteA(name string, c Config){
        aID := GetId(name, c)
        url := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/A/"+strconv.Itoa(aID)+""
        method := "DELETE"
        payload := strings.NewReader("")
        send(url, method, payload, c)
        fmt.Println(url)
}

func GetId(value string, c Config)int{
        urlC := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/cname?%20&offset=0&max=10&sort=name&order=asc"
        urlA := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/A?%20&offset=0&max=10&sort=name&order=asc"
        
	method := "GET"
        payload := strings.NewReader("")
        
        var tabC []Tabs
        var tabA []ArecordTabs
	
        valueID := 0 

        jsonBlobC := send(urlC,method,payload,c)
        jsonBlobA := send(urlA,method,payload,c)

	errC  := json.Unmarshal([]byte(jsonBlobC),&tabC)
	errA  := json.Unmarshal([]byte(jsonBlobA),&tabA)
        if errC != nil {
		fmt.Println("error:", errC)
        }
        if errA != nil {
		fmt.Println("error:", errA)
        }
	  
        for i:=0; i < len(tabA); i++ {
                if tabA[i].Value[0] == value {
                	valueID = tabA[i].Id
                }
	}
	
	if valueID == 0 {
		for i:=0; i < len(tabC); i++ {
                	if tabC[i].Name == value {
                        	valueID = tabC[i].Id
                	}
        	}
	}
	
        return valueID
}

func CreateCNAME(cname string, c Config){
        url := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/cname"
        method := "POST"
        payload := strings.NewReader("{\n  \"name\": \""+cname+"\",\n  \"ttl\": \"1800\"\n }")
        send(url,method,payload,c)
}

func CreateA(ip string, name string, c Config){
        url := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/A"
        method := "POST"
        
	payload := strings.NewReader("{\n  \"name\": \""+name+"\",\n  \"ttl\": \"1800\"\n,\n  \"roundRobin\": [\n{\n  \"value\": \""+ip+"\"\n }\n]}")
	
	send(url,method,payload,c)
}

        
func GetAll(c Config){
        urlC := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/cname?%20&offset=0&max=10&sort=name&order=asc"
        urlA := "https://api.dns.constellix.com/v1/domains/"+c.Constellix.Domain+"/records/A?%20&offset=0&max=10&sort=name&order=asc"
        
	method := "GET"
        payload := strings.NewReader("")

        var tabC []Tabs
        var tabA []ArecordTabs

	jsonBlobC := send(urlC,method,payload,c)
	jsonBlobA := send(urlA,method,payload,c)
	
	errC  := json.Unmarshal([]byte(jsonBlobC),&tabC)
	errA  := json.Unmarshal([]byte(jsonBlobA),&tabA)
        if errC != nil {
		fmt.Println("error:", errC)
	} else if errA != nil {
                fmt.Println("error:", errA)
	}

	for i:=0; i < len(tabA); i++ {
                fmt.Printf("%v %v %v %v %v\n", tabA[i].Id, tabA[i].Type, tabA[i].Name, tabA[i].Value, tabA[i].ParentId)
        }

        for i:=0; i < len(tabC); i++ {
                fmt.Printf("%v %v %v %v\n", tabC[i].Id, tabC[i].Type, tabC[i].Name, tabC[i].ParentId)
        }
}

func GetAllDomain(c Config){
	url := "https://api.dns.constellix.com/v1/domains?offset=0&max=10&sort=name&order=asc"
  	method := "GET"
	payload := strings.NewReader("")

	var tab []Tabs

        jsonBlob := send(url,method,payload,c)
	err  := json.Unmarshal([]byte(jsonBlob),&tab)
	if err != nil {
                fmt.Println("error:", err)
        }
  
        for i:=0; i < len(tab); i++ {
                fmt.Printf("%v %v %v %v \n", tab[i].Id, tab[i].Type, tab[i].Name, tab[i].ParentId)
        }
}

func send(url string, method string, payload *strings.Reader, c Config) string{
	now := time.Now()
        timestamp := (now.UnixNano())/1000000
        myhmac := generateHmac(strconv.FormatInt(timestamp,10), c.Constellix.ApiKey)
  	client := &http.Client {
      		CheckRedirect: func(req *http.Request, via []*http.Request) error {
                	return http.ErrUseLastResponse
        	},
        }
        req, err := http.NewRequest(method, url, payload)
	if err != nil {
        	fmt.Println(err)
        }
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("x-cns-security-token",c.Constellix.SecretKey+":" +myhmac+":"+strconv.FormatInt(timestamp,10))
        res, err := client.Do(req)
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
	
	return string(body)
}

func generateHmac(timestamp string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(timestamp ))
	expectedh := h.Sum(nil)
	c := base64.StdEncoding.EncodeToString([]byte(expectedh))
	
	return c
}
