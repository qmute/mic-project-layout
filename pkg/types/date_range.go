package types

type DateRange struct {
	A Date `validate:"required"`
	B Date `validate:"required"`
}
