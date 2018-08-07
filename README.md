# whois - a go wrapper for whois command

[![go report card](https://goreportcard.com/badge/github.com/xellio/whois "go report card")](https://goreportcard.com/report/github.com/xellio/whois)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/xellio/whois?status.svg)](https://godoc.org/github.com/xellio/whois)

### Examples
Query by URL:
```
url := "http://github.com/xellio/whois"
res, err := whois.Query(url)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(strings.Join(res.Output["Registrant City"], ","))
//Output:
//San Francisco
```
Query by hostname:
```
host := "golang.org"
res, err := whois.QueryHost(host)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(strings.Join(res.Output["Registrant Country"], ","))
// Output:
// US
```
Query by IP:
```
ip := net.ParseIP("8.8.8.8")
res, err := whois.QueryIp(ip)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(strings.Join(res.Output["CIDR"], ","))
// Output:
// 8.0.0.0/8,8.8.8.0/24
```