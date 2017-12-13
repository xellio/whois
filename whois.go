package whois

import (
	"net"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

//
// Result struct
//
type Result struct {
	Ip         net.IP
	Host       string
	Raw        []byte
	Output     map[string][]string
	GatherTime time.Duration
}

//
// Query whois data for given url
//
func Query(urlToQuery string, args ...string) (r *Result, err error) {
	u, err := url.Parse(urlToQuery)
	if err != nil {
		return nil, err
	}
	return QueryHost(u.Host, args...)
}

//
// Query whois data for given host
//
func QueryHost(host string, args ...string) (r *Result, err error) {
	r = &Result{
		Host: host,
	}
	args = append([]string{host}, args...)
	err = r.execute(args)
	return
}

//
// Query whois data for given net.IP
//
func QueryIp(ip net.IP, args ...string) (r *Result, err error) {
	r = &Result{
		Ip: ip,
	}
	args = append([]string{ip.String()}, args...)
	err = r.execute(args)
	return
}

//
// Execute the whois command using the given args
//
func (r *Result) execute(args []string) error {
	start := time.Now()
	out, err := exec.Command("whois", args...).Output()
	if err != nil {
		return err
	}

	r.GatherTime = time.Since(start)
	r.Raw = out
	r.Output = make(map[string][]string)

	singleLines := strings.Split(string(out), "\n")
	for _, line := range singleLines {
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "%") || strings.HasPrefix(line, ">") {
			continue
		}
		lineParts := strings.Split(line, ": ")
		if len(lineParts) == 2 {
			tk := strings.TrimSpace(lineParts[0])
			r.Output[tk] = append(r.Output[tk], strings.TrimSpace(lineParts[1]))
		}
	}

	return nil
}
