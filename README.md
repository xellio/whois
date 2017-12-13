# whois - a go wrapper for whois command

[Documentation](https://godoc.org/github.com/xellio/whois)

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