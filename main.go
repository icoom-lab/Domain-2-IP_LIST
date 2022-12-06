package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lixiangzhong/dnsutil"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(getAAAARecordsNTimes("youtube.com", 4))
}

// get N times AAAA records for a domain
func getAAAARecordsNTimes(domain string, n int) ([]string, error) {
	AAAARecord := make([]string, 0)

	listDns := []string{"1.1.1.1", "8.8.8.8", "9.9.9.9", "208.67.222.222"}
	for i := 0; i < n; i++ {

		rand.Seed(time.Now().Unix())
		n := rand.Int() % len(listDns)
		AAAA, err := AAAA(domain, listDns[n])
		if err == nil {
			AAAARecord = append(AAAARecord, AAAA...)
		}
	}

	// remove duplicate ip
	uniqueIp := make(map[string]bool)
	var cleanListIp = []string{}

	for _, val := range AAAARecord {
		if _, ok := uniqueIp[val]; !ok {
			uniqueIp[val] = true
			cleanListIp = append(cleanListIp, val)
		}
	}

	return cleanListIp, nil
}

// get AAAA records for a domain
func AAAA(domain, dnsIp string) ([]string, error) {

	var dig dnsutil.Dig
	dig.SetDNS(dnsIp)                 //or ns.xxx.com
	a, err := dig.AAAA("youtube.com") // dig google.com @8.8.8.8
	if err != nil {
		log.Error(err)
		return nil, err
	} else {

		ipListAAAA := make([]string, 0)

		for _, val := range a {
			ipListAAAA = append(ipListAAAA, val.AAAA.String())
		}
		return ipListAAAA, err
	}
}
