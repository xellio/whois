// !build +testing

package whois_test

import (
	"fmt"
	"net"
	"strings"

	"github.com/xellio/whois"
)

//
// ExampleQuery - An example of using the Query function
//
func ExampleQuery() {
	url := "http://github.com/xellio/whois"
	res, err := whois.Query(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Join(res.Output["Registrant City"], ","))
	//Output:
	//San Francisco
}

//
// ExampleQueryHost - An example of using the QueryHost function
//
func ExampleQueryHost() {
	host := "golang.org"
	res, err := whois.QueryHost(host)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Join(res.Output["Registrant Country"], ","))
	// Output:
	// US
}

//
// ExampleQueryIp - An example of using the QueryIp function
//
func ExampleQueryIP() {
	ip := net.ParseIP("8.8.8.8")
	res, err := whois.QueryIP(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Join(res.Output["CIDR"], ","))
	// Output:
	// 8.0.0.0/8,8.8.8.0/24
}
