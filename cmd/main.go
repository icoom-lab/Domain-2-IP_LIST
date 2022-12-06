package main

import (
	domain2list "github.com/icoom-lab/Domain-2-IP_LIST"
)

func main() {

	domainsList, err := domain2list.ReadLines("domains.txt")
	if err != nil {
		return
	}

	newListAAAARecords := make([]string, 0)

	for _, domain := range domainsList {
		AAAARecords, err := domain2list.GetAAAARecordsNTimes(domain, 4)
		if err == nil {
			newListAAAARecords = append(newListAAAARecords, AAAARecords...)
		}
	}

	// create file if not exists and read lines from file to slice
	nameFile := "iplist.txt"
	domain2list.CreateFileIfNotExists(nameFile)
	ipListFromFile, err := domain2list.ReadLines(nameFile)
	if err != nil {
		return
	}

	// merge 2 slices
	ipList := append(newListAAAARecords, ipListFromFile...)
	// remove duplicates
	ipList = domain2list.Unique(ipList)

	// write to file
	domain2list.WriteToFile(nameFile, ipList)
}
