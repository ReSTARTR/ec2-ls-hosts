package main

import (
	"flag"
	"fmt"
	"github.com/ReSTARTR/ec2-ls-hosts/client"
	"log"
	"os"
	"strings"
)

var (
	version string
)

//
// parse filters option string
//  `-filters` pattern : key1:value1,key2:value2,...
func parseFilterString(s string) map[string]string {
	filters := make(map[string]string, 5)
	for _, kv := range strings.Split(s, ",") {
		a := strings.Split(kv, ":")
		if len(a) > 1 {
			filters[a[0]] = a[1]
		}
	}
	return filters
}

// parse columns option string
//  `-columns` pattern : c1,c2,c3,...
func parseColumnString(str string) []string {
	var columns []string
	for _, c := range strings.Split(str, ",") {
		columns = append(columns, c)
	}
	return columns
}

func main() {
	// parse options
	filterString := flag.String("filters", "", "key1:value1,key2:value2,...")
	columnString := flag.String("columns", "", "column1,column2,...")
	v := flag.Bool("v", false, "show version")
	flag.Parse()
	if *v {
		fmt.Println("version: " + version)
		os.Exit(0)
	}

	filters := parseFilterString(*filterString)
	columns := parseColumnString(*columnString)

	err := client.Describe(filters, columns)
	if err != nil {
		log.Fatal(err)
	}
}
