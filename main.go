package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sts"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess = session.Must(session.NewSession())

// Accounts struct for accounts and url objects
type Accounts struct {
	Accounts []Account `json:"accounts"`
	URL      string    `json:"url"`
}

// Account struct for account & security group IDs
type Account struct {
	Account       string   `json:"account"`
	Role          string   `json:"role"`
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

func processAccount(acct Account, ips []string) {
	for _, a := range acct.Account {
		fmt.Println(a)
	}
}

func getCreds(role string) *sts.GetCallerIdentityOutput {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	creds := stscreds.NewCredentials(sess, role)
	svc := sts.New(sess, &aws.Config{Credentials: creds})
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	return result
}

func main() {
	config, _ := parseConfig()

	for _, acct := range config.Accounts {
		result := getCreds(acct.Role)
		// creds, err := getCreds(acct.Role)
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	//svc := ec2.New(sess, &aws.Config{Credentials: creds})
		fmt.Println(result)
		// }
	}

	// result, _ := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
	// 	GroupIds: aws.StringSlice(acct.SecurityGroup),
	// })
	// for _, group := range result.SecurityGroups {
	// 	fmt.Println(group)
	// }
	//}
}

// TODO:
// 1. Authentication with roles specified
// 2. Get SGs listed
// 3. For each SG compare IP rules against resolved IPs (set compare)
// 4. Drop all rules if comparison set compare differs and upload current
// 5. Get SGs again and output rules
