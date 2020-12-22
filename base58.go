package base58

import (
	"fmt"
	"math/big"
)

var (
	bigZero  = big.NewInt(0)
	bigRadix = big.NewInt(58)
	alphabet = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "A",
		"B", "C", "D", "E", "F", "G", "H", "J", "K", "L",
		"M", "N", "P", "Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z",
	}
	b58table = [256]byte{
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 255, 255,
		255, 255, 255, 255, 255, 9, 10, 11, 12, 13,
		14, 15, 16, 255, 17, 18, 19, 20, 21, 255,
		22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		32, 255, 255, 255, 255, 255, 255, 33, 34, 35,
		36, 37, 38, 39, 40, 41, 42, 43, 255, 44,
		45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
		55, 56, 57, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255,
	}
)

// Encode takes a slice of bytes and encodes it to base58 string.
// Leading zero bytes are kept in place for precise decoding.
func Encode(input []byte) (output string) {
	num := new(big.Int).SetBytes(input)

	for num.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, bigRadix, mod)
		output = alphabet[mod.Int64()] + output
	}

	for _, i := range input {
		if i != 0 {
			break
		}
		output = alphabet[0] + output
	}
	return
}

// Decode takes string as an input and returns decoded string and error
// If provided string contains characters illegal for base58 the returned error will be <notnil>
func Decode(input string) (output []byte, err error) {
	result := big.NewInt(0)
	multi := big.NewInt(1)

	tmpBig := new(big.Int)

	for i := len(input) - 1; i >= 0; i-- {
		tmp := b58table[input[i]]
		if tmp == 255 {
			err = fmt.Errorf("invalid Base58 input string at character \"%s\", position %d", string(input[i]), i)
			return
		}
		tmpBig.SetInt64(int64(tmp))
		tmpBig.Mul(multi, tmpBig)
		result.Add(result, tmpBig)
		multi.Mul(multi, bigRadix)
	}

	tmpBytes := result.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(input); numZeros++ {
		if input[numZeros] != '1' {
			break
		}
	}
	length := numZeros + len(tmpBytes)
	output = make([]byte, length)
	copy(output[numZeros:], tmpBytes)

	return
}
