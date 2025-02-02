package types

import "github.com/pkg/errors"

const (
	StatusUnknown  Status = iota // 未知
	StatusEnabled                // 启用
	StatusDisabled               // 禁用
)

// Status 状态，仅适用于开/关2种状态场景下
type Status int

func (p Status) Enabled() bool {
	return p == StatusEnabled
}

func (p Status) Disabled() bool {
	return p == StatusDisabled
}

func (p Status) Valid() error {
	if p == StatusEnabled || p == StatusDisabled {
		return nil
	}
	return errors.New("状态不合法")
}

func (p Status) Fmt() string {
	switch p {
	case StatusEnabled:
		return "启用"
	case StatusDisabled:
		return "禁用"
	default:
		return "未知"
	}
}
