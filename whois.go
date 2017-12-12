package whois

import (
	"net"
	"os/exec"
	"strings"
	"time"
)

//
// Result struct
//
type Result struct {
	Ip         net.IP
	Raw        []byte
	Output     map[string]string
	GatherTime time.Duration
}

//
// Query whois command for given net.IP
//
func Query(ip net.IP, args ...string) (*Result, error) {
	r := &Result{
		Ip: ip,
	}
	args = append([]string{ip.String()}, args...)

	start := time.Now()

	out, err := exec.Command("whois", args...).Output()
	if err != nil {
		return r, err
	}

	r.GatherTime = time.Since(start)
	r.Raw = out
	r.Output = make(map[string]string)

	singleLines := strings.Split(string(out), "\n")
	for _, line := range singleLines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		lineParts := strings.Split(line, ": ")
		if len(lineParts) == 2 {
			r.Output[lineParts[0]] = strings.TrimSpace(lineParts[1])
		}
	}

	return r, nil
}
