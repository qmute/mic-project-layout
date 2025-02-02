package testdata

import "easyslip.cc/mic-project-layout/internal/domain"

func ClientInfo(ctype domain.ClientType, versions ...string) *domain.ClientInfo {
	ver := "v1.0.0"
	if len(versions) > 0 {
		ver = versions[0]
	}

	return &domain.ClientInfo{
		ClientType:    ctype,
		ClientVersion: ver,
		IP:            "127.0.0.1",
	}
}
