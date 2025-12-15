package types

import "github.com/cockroachdb/errors"

const (
	// GenderNone 未知
	GenderNone Gender = iota
	// GenderMale 男性
	GenderMale
	// GenderFemale 女性
	GenderFemale
)

type Gender uint

func (p Gender) Valid() error {
	switch p {
	case GenderNone, GenderMale, GenderFemale:
		return nil
	default:
		return errors.New("bad gender")
	}
}

func (p Gender) Int() int {
	return int(p)
}
