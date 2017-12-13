package whois_test

import (
	"log"
	"net"

	"encoding/json"

	"github.com/xellio/whois"
)

//
// A basic example of usage
//
func ExampleQuery() {
	ip := net.ParseIP("8.8.8.8")
	res, err := whois.Query(ip)
	if err != nil {
		log.Println(err)
		return
	}

	jsonData, err := json.Marshal(res.Output)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(jsonData))
	// Output:
	// {"Address":["1025 Eldorado Blvd.","1600 Amphitheatre Parkway"],"CIDR":["8.0.0.0/8","8.8.8.0/24"],"City":["Broomfield","Mountain View"],"Comment":["ADDRESSES WITHIN THIS BLOCK ARE NON-PORTABLE"],"Country":["US","US"],"NetHandle":["NET-8-0-0-0-1","NET-8-8-8-0-1"],"NetName":["LVLT-ORG-8-8","LVLT-GOGL-8-8-8"],"NetRange":["8.0.0.0 - 8.255.255.255","8.8.8.0 - 8.8.8.255"],"NetType":["Direct Allocation","Reallocated"],"OrgAbuseEmail":["security@level3.com","network-abuse@google.com"],"OrgAbuseHandle":["APL8-ARIN","ABUSE5250-ARIN"],"OrgAbuseName":["Abuse POC LVLT","Abuse"],"OrgAbusePhone":["+1-877-453-8353","+1-650-253-0000"],"OrgAbuseRef":["https://whois.arin.net/rest/poc/APL8-ARIN","https://whois.arin.net/rest/poc/ABUSE5250-ARIN"],"OrgId":["LVLT","GOGL"],"OrgNOCEmail":["noc.coreip@level3.com"],"OrgNOCHandle":["NOCSU27-ARIN"],"OrgNOCName":["NOC Support"],"OrgNOCPhone":["+1-877-453-8353"],"OrgNOCRef":["https://whois.arin.net/rest/poc/NOCSU27-ARIN"],"OrgName":["Level 3 Communications, Inc.","Google LLC"],"OrgTechEmail":["ipaddressing@level3.com","arin-contact@google.com"],"OrgTechHandle":["IPADD5-ARIN","ZG39-ARIN"],"OrgTechName":["ipaddressing","Google LLC"],"OrgTechPhone":["+1-877-453-8353","+1-650-253-0000"],"OrgTechRef":["https://whois.arin.net/rest/poc/IPADD5-ARIN","https://whois.arin.net/rest/poc/ZG39-ARIN"],"Organization":["Level 3 Communications, Inc. (LVLT)","Google LLC (GOGL)"],"OriginAS":["",""],"Parent":["()","LVLT-ORG-8-8 (NET-8-0-0-0-1)"],"PostalCode":["80021","94043"],"Ref":["https://whois.arin.net/rest/net/NET-8-0-0-0-1","https://whois.arin.net/rest/org/LVLT","https://whois.arin.net/rest/net/NET-8-8-8-0-1","https://whois.arin.net/rest/org/GOGL"],"RegDate":["1992-12-01","1998-05-21","2014-03-14","2000-03-30"],"StateProv":["CO","CA"],"Updated":["2012-02-24","2017-01-28","2014-03-14","2017-10-16"]}
}
