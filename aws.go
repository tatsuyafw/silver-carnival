package main

import (
	"os"
	"strings"
)

const DEFAULT_REGION = "us-east-1"

type AWS struct{}

func (*AWS) Url(s string) string {
	m := urlmap()
	r := region()
	return "https://" + strings.Replace(m[s], "REGION", r, -1)
}

// TODO: return with err
func (*AWS) Validate(service string) bool {
	m := urlmap()
	_, ok := m[service]
	return ok
}

func urlmap() map[string]string {
	return map[string]string{
		"ec2":         "REGION.console.aws.amazon.com/ec2/v2/home?REGION&region=REGION",
		"ecs":         "REGION.console.aws.amazon.com/ecs/home?region=REGION",
		"elasticache": "REGION.console.aws.amazon.com/elasticache/home?region=REGION",
		"iam":         "console.aws.amazon.com/iam/home?region=REGION",
		"rds":         "REGION.console.aws.amazon.com/rds/home?region=REGION",
		"route53":     "console.aws.amazon.com/route53/home?region=REGION",
		"s3":          "console.aws.amazon.com/s3/home?region=REGION",
		"vpc":         "REGION.console.aws.amazon.com/vpc/home?region=REGION",
	}
}

func (*AWS) supported() []string {
	m := urlmap()
	s := []string{}
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}

func region() string {
	r := os.Getenv("AWS_REGION")
	if r == "" {
		r = DEFAULT_REGION
	}
	return r
}
