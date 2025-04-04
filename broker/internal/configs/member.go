package configs

import (
	"os"
)

type MemberServiceConfig struct {
	Domain string
	Port   string
}

var MemberService MemberServiceConfig

func LoadMemberServiceConfig() {
	MemberService = MemberServiceConfig{
		Domain: os.Getenv("MEMBER_SERVICE_DOMAIN"),
		Port:   os.Getenv("MEMBER_SERVICE_PORT"),
	}
}
