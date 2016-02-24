package client

import (
	"fmt"
	"github.com/ReSTARTR/ec2-ls-hosts/creds"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"regexp"
	"strings"
)

func Describe(filters map[string]string, columns []string) error {
	// build queries
	config := &aws.Config{Region: aws.String("ap-northeast-1")}
	config.Credentials = creds.SelectCredentials("") // TODO
	svc := ec2.New(session.New(), config)

	// call aws api
	options := &ec2.DescribeInstancesInput{}
	for k, v := range filters {
		options.Filters = append(options.Filters, &ec2.Filter{
			Name:   aws.String("tag:" + k),
			Values: []*string{aws.String(v)},
		})
	}

	// show info
	resp, err := svc.DescribeInstances(options)
	if err != nil {
		return err
	}
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println(formatInstance(inst, columns))
		}
	}
	return nil
}

func formatInstance(inst *ec2.Instance, columns []string) string {
	// fetch IPs
	var privateIps []string
	var publicIps []string
	for _, nic := range inst.NetworkInterfaces {
		for _, privateIp := range nic.PrivateIpAddresses {
			privateIps = append(privateIps, *privateIp.PrivateIpAddress)
			if privateIp.Association != nil {
				publicIps = append(publicIps, *privateIp.Association.PublicIp)
				break
			}
		}
	}

	// fetch tags
	tags := make(map[string]string, 5)
	for _, tag := range inst.Tags {
		tags[*tag.Key] = *tag.Value
	}

	var values []string
	for _, c := range columns {
		switch c {
		case "instance-id":
			values = append(values, *inst.InstanceId)
		case "private-ip":
			values = append(values, strings.Join(privateIps, ","))
		case "public-ip":
			values = append(values, strings.Join(publicIps, ","))
		default:
			// extract key-values as tag string
			matched, err := regexp.Match("tag:.+", []byte(c))
			if err == nil && matched {
				kv := strings.Split(c, ":")
				if v, ok := tags[kv[1]]; ok {
					values = append(values, v)
				}
			}
		}
	}

	return strings.Join(values, "\t")
}
