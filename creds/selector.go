package creds

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	SHARED_CREDS_PROFILE = "default"
)

// iam "env" or "shared" or "ec2"
func SelectCredentials(iam string) *credentials.Credentials {
	var providers []credentials.Provider
	switch iam {
	case "env":
		providers = []credentials.Provider{
			&credentials.EnvProvider{},
		}
	case "shared":
		providers = []credentials.Provider{
			&credentials.SharedCredentialsProvider{Profile: SHARED_CREDS_PROFILE},
		}
	case "ec2":
		providers = []credentials.Provider{
			&ec2rolecreds.EC2RoleProvider{
				Client: ec2metadata.New(session.New()),
			},
		}
	default:
		providers = []credentials.Provider{
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{Profile: SHARED_CREDS_PROFILE},
			&ec2rolecreds.EC2RoleProvider{
				Client: ec2metadata.New(session.New()),
			}}
	}
	return credentials.NewChainCredentials(providers)
}
