package main

import (
	"flag"
	"fmt"
	"github.com/ReSTARTR/ec2-ls-hosts/client"
	"github.com/ReSTARTR/ec2-ls-hosts/creds"
	"github.com/go-ini/ini"
	"os"
	"strings"
)

var (
	version string
)

func loadRegionInAwsConfig() string {
	cfg, err := creds.LoadAwsConfig()
	if err == nil {
		return cfg.Section("default").Key("region").Value()
	}
	return ""
}

func loadConfig() (cfg *ini.File, err error) {
	cfg, err = ini.LooseLoad(
		os.Getenv("HOME")+"/.ls-hosts",
		"/etc/ls-hosts.conf",
	)
	if err != nil {
		return
	}
	return
}

// string to map[string]string
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

// string to []string
func parseFieldsString(str string) []string {
	var fields []string
	for _, c := range strings.Split(str, ",") {
		fields = append(fields, c)
	}
	return fields
}

func optionsFromFile() *client.Options {
	opt := &client.Options{}
	if cfg, err := loadConfig(); err == nil {
		opt.Region = cfg.Section("options").Key("region").Value()
		opt.TagFilters = parseFilterString(cfg.Section("options").Key("tags").Value())
		opt.Fields = parseFieldsString(cfg.Section("options").Key("fields").Value())
		opt.Credentials = cfg.Section("options").Key("creds").Value()
	}
	return opt
}

func main() {
	// parse options
	filters := flag.String("filters", "", "key1:value1,key2:value2,...")
	tagFilters := flag.String("tags", "", "key1:value1,key2:value2,...")
	fields := flag.String("fields", "", "column1,column2,...")
	regionString := flag.String("region", "", "region name")
	credsString := flag.String("creds", "", "env, shared, iam")
	v := flag.Bool("v", false, "show version")
	flag.Parse()
	if *v {
		fmt.Println("version: " + version)
		os.Exit(0)
	}

	opt := optionsFromFile()
	opt.Region = loadRegionInAwsConfig()
	// merge optoins from cmdline
	if *filters != "" {
		for k, v := range parseFilterString(*filters) {
			opt.Filters[k] = v
		}
	}
	if *tagFilters != "" {
		for k, v := range parseFilterString(*tagFilters) {
			opt.TagFilters[k] = v
		}
	}
	if *fields != "" {
		opt.Fields = parseFieldsString(*fields)
	}
	if *regionString != "" {
		opt.Region = *regionString
	}
	if *credsString != "" {
		opt.Credentials = *credsString
	}

	// run
	err := client.Describe(opt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
