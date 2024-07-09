package utils

import "math/big"

func AttoFILToFIL(attoFIL string) string {
	value := new(big.Int)
	value.SetString(attoFIL, 10) // attoFIL 값 파싱
	fil := new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(1e18))
	return fil.Text('f', 18) // 18 소수점 자릿수로 출력
}
