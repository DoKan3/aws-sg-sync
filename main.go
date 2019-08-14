package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
func parseConfig() (accounts Accounts) {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	if err2 := json.Unmarshal([]byte(byteValue), &accounts); err2 != nil {
		panic(err2)
	}
	// fmt.Printf("URL: "+"%s\n", accounts.URL)
	return accounts
}

func getIps(url Accounts) string {

	// addr, err := net.LookupIP("nat.travisci.net")
	// addr, err := url
	// if err != nil {
	// 	fmt.Println("Unknown host")
	// } else {
	// 	for _, ip := range addr {
	// 		s = append(s, ip.String()+"/32")
	// 	}
	// }
	// sort.Strings(s)
	// return s
	return url.URL
	// fmt.Println(url.URL)
}

func main() {

}
