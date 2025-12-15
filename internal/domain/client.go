package domain

import (
	"github.com/cockroachdb/errors"
	"github.com/qmute/mic-project-layout/pkg/ut"
)

const (
	// ClientTypeNone 未知
	ClientTypeNone ClientType = ""
	// ClientTypeIOS iOS app
	ClientTypeIOS ClientType = "ios"
	// ClientTypeAndroid android app
	ClientTypeAndroid ClientType = "android"
	// ClientTypeHarmony harmonyOS app
	ClientTypeHarmony ClientType = "harmony"
)

// ClientType 客户端类型
type ClientType string

func (p ClientType) Valid() error {
	switch p {
	case ClientTypeNone, ClientTypeIOS, ClientTypeAndroid, ClientTypeHarmony:
		return nil
	default:
		return errors.Errorf("暂不支持客户端[%s]", p)
	}
}

func (p ClientType) String() string {
	return string(p)
}

func (p ClientType) IsIOSApp() bool {
	return p == ClientTypeIOS
}

func (p ClientType) IsAndroidApp() bool {
	return p == ClientTypeAndroid
}

func (p ClientType) IsHarmonyApp() bool {
	return p == ClientTypeHarmony
}

func (p ClientType) IsApp() bool {
	return p.IsIOSApp() || p.IsAndroidApp() || p.IsHarmonyApp()
}

type ClientInfo struct {
	ClientType    ClientType `validate:"required"` // 客户端类型
	ClientVersion string     // 客户端版本
	IP            string     `validate:"required"` // IP地址
}

func (p *ClientInfo) Valid() error {
	if err := ut.ValidStruct(p); err != nil {
		return err
	}

	if err := p.ClientType.Valid(); err != nil {
		return err
	}

	return nil
}
