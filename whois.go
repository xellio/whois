package whois

import (
	"net"
	"os/exec"
	"strings"
	"time"
)

type Result struct {
	Ip         net.IP
	RawOutput  []byte
	Output     map[string]string
	GatherTime time.Duration
}

func Query(ip net.IP) (*Result, error) {
	r := &Result{
		Ip: ip,
	}
	start := time.Now()

	out, err := exec.Command("whois", ip.String()).Output()
	if err != nil {
		return r, err
	}
	r.GatherTime = time.Since(start)
	r.RawOutput = out
	r.Output = make(map[string]string)

	singleLines := strings.Split(string(out), "\n")

	for _, line := range singleLines {
		lineParts := strings.Split(line, ":")
		if len(lineParts) == 2 {
			r.Output[lineParts[0]] = lineParts[1]
		}
	}

	return r, nil
}
