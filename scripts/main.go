package main

import "github.com/Arjun-P-J-WebomindApps/gobackend-prototype/scripts/typesense"

func main() {
	//ImportData()
	typesense.Setup()

	typesense.ExportData()
}
