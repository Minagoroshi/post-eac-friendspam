package accounts

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var Accounts []string
var accountRegex = "^[a-zA-Z0-9]{3,20}:[a-zA-Z0-9]{3,20}$"

// Load is a function to load each line of filepath into a slice of strings
func Load(filepath string) []string {
	var accounts []string
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		accounts = append(accounts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	Accounts = accounts
	return accounts
}

// Check is a function to check if the account is valid using the accountRegex regex
func Check(account string) bool {
	valid, _ := regexp.MatchString(accountRegex, account)
	return valid
}
