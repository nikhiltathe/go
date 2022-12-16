package main

// Go offers built-in support for [regular expressions](https://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str1 := `Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
	118.99.81.204:8080 34 sec Indonesia - Tangerang Transparent 2.184.31.2 8080 58 sec 
	Iran Transparent 93.126.11.189 8080 1 min Iran - Esfahan Transparent 202.118.236.130 
	7777 1 min China - Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent`

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	submatchall := re.FindAllString(str1, -1)
	fmt.Println("Found :", len(submatchall))
	for _, element := range submatchall {
		fmt.Println(element)
	}

	pos := re.FindAllStringIndex(str1, -1)
	fmt.Println("Found :", len(pos))

	ini := 0
	newMsg := ""
	for _, v := range pos {
		fmt.Println(len(v))
		fmt.Println(v[0], v[1])
		newMsg += str1[ini:v[0]]
		ip, ok := HashoutIPorFQDN(str1[v[0]:v[1]])
		fmt.Println(ip, ok)
		newMsg += ip
		ini = v[1]
	}
	newMsg += str1[ini:]
	fmt.Println(newMsg)
}

func HashoutIPorFQDN(host string) (string, bool) {
	// log.Debug("Entering")
	// defer log.Debug("Exiting")

	hostIsIPAddress := true
	parts := strings.Split(host, ".")
	size := len(parts)
	if size == 4 {
		for _, x := range parts {
			// Check if threre is any character
			if i, err := strconv.Atoi(x); err == nil {
				if i < 0 || i > 255 {
					hostIsIPAddress = false
					break
				}
			} else {
				hostIsIPAddress = false
			}
		}
	} else {
		hostIsIPAddress = false
	}

	var output string
	if hostIsIPAddress {
		// Hashout all except last octet for IP address
		size := len(parts)
		output = "XXX.XXX.XXX." + parts[size-1]
	} else {
		// Hashout all except first octet for FQDN
		allXXXs := make([]string, 0)
		for i := 0; i < size-1; i++ {
			allXXXs = append(allXXXs, "XXX")
		}

		if size > 1 {
			output = strings.Join(allXXXs, ".")
			output = parts[0] + "." + output
		} else {
			output = parts[0]
		}
	}

	return output, hostIsIPAddress
}
