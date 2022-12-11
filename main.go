package domain2list

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/lixiangzhong/dnsutil"
	log "github.com/sirupsen/logrus"
)

// write to file slice string line by line
func WriteToFile(filename string, data []string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()

	for _, d := range data {
		if _, err := file.WriteString(d + "\n"); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

// from array get onley unique values
func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// read all lines from file and return a slice of string
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// create file txt if not exists
func CreateFileIfNotExists(filename string) error {
	if !fileExists(filename) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// Check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// get N times AAAA records for a domain
func GetAAAARecordsNTimes(domain string, n int) ([]string, error) {
	AAAARecord := make([]string, 0)

	listDns := []string{"8.8.8.8", "8.8.4.4", "2001:4860:4860::8888", "2001:4860:4860::8844"}
	for i := 0; i < n; i++ {

		rand.Seed(time.Now().Unix())
		n := rand.Int() % len(listDns)
		AAAA, err := AAAA(domain, listDns[n])
		if err == nil {
			AAAARecord = append(AAAARecord, AAAA...)
		}
	}

	return Unique(AAAARecord), nil
}

// get AAAA records for a domain
func AAAA(domain, dnsIp string) ([]string, error) {

	var dig dnsutil.Dig
	dig.SetDNS(dnsIp)          //or ns.xxx.com
	a, err := dig.AAAA(domain) // dig google.com @8.8.8.8
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
