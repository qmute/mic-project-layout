package types

import "github.com/cockroachdb/errors"

const (
	SystemCatIOS     SystemCat = "ios"
	SystemCatAndroid SystemCat = "android"
	SystemCatHarmony SystemCat = "harmony"
)

type SystemCat string

func (p SystemCat) Valid() error {
	switch p {
	case SystemCatIOS, SystemCatAndroid, SystemCatHarmony:
		return nil
	default:
		return errors.Errorf("无效的分类:%s", p)
	}

}

func (p SystemCat) IsIOS() bool {
	return p == SystemCatIOS
}
