package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sort"
)

// Accounts struct for accounts and url objects
type Accounts struct {
	Accounts []Account `json:"accounts"`
	URL      string    `json:"url"`
}

// Account struct for account & secuirty group IDs
type Account struct {
	Account       string   `json:"account"`
	SecurityGroup []string `json:"security-group"`
}

// TODO: Add condition to check for existence of file, if not then create new security groups in each account
func parseConfig() (*Accounts, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	accounts := &Accounts{}
	if err2 := json.Unmarshal([]byte(byteValue), accounts); err2 != nil {
		panic(err2)
	}
	return accounts, nil
}

func getIps(url string) []string {
	var s []string
	addr, err := net.LookupIP(url)
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		for _, ip := range addr {
			s = append(s, ip.String()+"/32")
		}
	}
	sort.Strings(s)
	return s
}

func main() {
	config, _ := parseConfig()
	url := config.URL
	fmt.Println(getIps(url))
}
