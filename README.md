# whois - a go wrapper for whois command

[Documentation](https://godoc.org/github.com/xellio/whois)

### Example
```
import (
	"log"
	"net"

	"github.com/xellio/whois"
)
func main() {
	ip := net.ParseIP("8.8.8.8")
	res, err := whois.Query(ip)
	if err != nil {
		log.Println(err)
	}

	for key, value := range res.Output {
		log.Println(key, value)
	}
}
```