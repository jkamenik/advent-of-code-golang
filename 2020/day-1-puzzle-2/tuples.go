package main

type tuple struct {
	A uint64
	B uint64
	C uint64

	MaxFactor uint64
}

func (t *tuple) OtherFactor() (factor uint64) {
	factor1 := t.A + t.B

	if factor1 >= t.MaxFactor {
		return 0
	}

	return t.MaxFactor - factor1
}

func (t *tuple) Multiple() uint64 {
	return t.A * t.B * t.C
}
