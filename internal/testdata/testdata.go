package testdata

import "easyslip.cc/mic-project-layout/pkg/types"

func getId[T any](id ...T) T {
	if len(id) == 0 {
		var t T
		return t
	}

	return id[0]
}

func Pager() types.Pager {
	return types.Pager{
		Limit: 10,
		Page:  1,
	}
}
