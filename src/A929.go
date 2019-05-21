package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
	prefix := make(map[int]string)
	for k, v := range emails {
		temp := strings.Split(v, "@")
		prefix[k] = strings.Split(temp[0], "+")[0]
		prefix[k] = strings.Replace(prefix[k], ".", "", -1)
		prefix[k] = prefix[k] + "@" + temp[1]
	}
	transfer := make(map[string]int)
	for k, v := range prefix {
		transfer[v] = k
	}
	return len(transfer)
}
