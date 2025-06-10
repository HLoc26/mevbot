package helper

import (
	"errors"
	"math/big"
)

type BigIntHelper struct{}

// FromString tạo big.Int từ chuỗi số thập phân
func (b *BigIntHelper) FromString(s string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil, errors.New("invalid number string")
	}
	return n, nil
}

// FromExp tạo big.Int từ 10^exp, ví dụ exp=18 tương đương 1e18
func (b *BigIntHelper) FromExp(exp int64) *big.Int {
	base := big.NewInt(10)
	return new(big.Int).Exp(base, big.NewInt(exp), nil)
}

func (h *BigIntHelper) MulExp(base int64, exp int) *big.Int {
	baseBig := big.NewInt(base)
	tenPow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exp)), nil) // 10^exp
	return new(big.Int).Mul(baseBig, tenPow)                                // base * 10^exp
}
