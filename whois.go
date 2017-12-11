package whois

import (
	"log"
	"net"
	"os/exec"
	"time"
)

type Result struct {
	Ip         net.IP
	Output     []byte
	GatherTime time.Duration
}

func Query(ip net.IP) (*Result, error) {
	r := &Result{
		Ip: ip,
	}
	start := time.Now()

	//	log.Println("whois", ip.String())

	//	return r, nil

	out, err := exec.Command("whois", ip.String()).Output()
	if err != nil {
		return r, err
	}
	r.GatherTime = time.Since(start)
	r.Output = out
	return r, nil
}
