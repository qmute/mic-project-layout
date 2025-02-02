package types

// Pager 分页参数
type Pager struct {
	Limit int `validate:"gt=0"`
	Page  int `validate:"gt=0"`
}

func (p Pager) Offset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p Pager) End() int {
	return (p.GetPage()) * p.GetLimit()
}

// GetPage 默认第1页
func (p Pager) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

// GetLimit 默认每页10条
func (p Pager) GetLimit() int {
	if p.Limit <= 0 {
		return 10
	}
	return p.Limit
}
