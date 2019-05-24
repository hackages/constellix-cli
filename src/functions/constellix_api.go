package functions

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Tabs structure details all return values for CNAME records
type TabsWithoutValue struct {
	Id   int
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
	//Value string
	//RoundRobin []string
	//Geolocation int
	//RecordFailover string
	//Failover string
	//RoundRobinFailover []string
	//Pools []string
	//PoolsDetail []string
}

// Tabs structure details all return values for NS records
type TabsWithObjectValue struct {
	Id   int
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
	Value []NsValue
	//RoundRobin []string
}

type NsValue struct {
	Value string
	//DisableFlag bool
}

// Tabs structure details all return values for A records
type TabsWithArrayValue struct {
	Id   int
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

// Tab structure details all return values for domains
type DomainTabs struct {
	Id   int
	Name string
	// TypeId          int
	// HasGtdRegions   bool
	// HasGeoIP        bool
	// NameserverGroup int
	// Nameservers     []string
	// CreatedTs       string
	// DateTime        string
	// ModifiedTs      string
	// Note            string
	// Version         int
	// Status          string
	// Tags            []string
	// Soa             []string
}

type KeysConfig struct {
	ApiKey    string `mapstructure:"apiKey"`
	SecretKey string `mapstructure:"secretKey"`
	Domain    string `mapstructure:"domain"`
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
		fmt.Printf("Couldn't load config: %s\n", err)
		os.Exit(1)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("Couldn't read config: %s\n", err)
	}

	return c
}

func send(url string, method string, payload *strings.Reader, c Config) string {
	now := time.Now()
	timestamp := (now.UnixNano()) / 1000000
	myhmac := generateHmac(strconv.FormatInt(timestamp, 10), c.Constellix.ApiKey)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error:", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cns-security-token", c.Constellix.SecretKey+":"+myhmac+":"+strconv.FormatInt(timestamp, 10))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return string(body)
}

func generateHmac(timestamp string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(timestamp))
	expectedh := h.Sum(nil)
	c := base64.StdEncoding.EncodeToString([]byte(expectedh))

	return c
}
