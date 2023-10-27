package arrowtests

import (
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/apache/arrow/go/v14/arrow/decimal128"
	"github.com/apache/arrow/go/v14/arrow/decimal256"
	"github.com/stretchr/testify/require"
)

func randDigit() string {
	return strconv.Itoa(int(rand.Int31() % 10))
}

func TestFromStringDecimal256(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	const maxprec = 76

	for {
		scale := int(rand.Int31()) % (maxprec + 1)

		s := ""
		if maxprec-scale > 0 {
			for i := 0; i < maxprec-scale; i++ {
				s += randDigit()
			}
		} else {
			s += "0"
		}

		if scale > 0 {
			s += "."

			for i := 0; i < scale; i++ {
				s += randDigit()
			}
		}

		num, err := decimal256.FromString(s, 76, int32(scale))
		rq.NoErrorf(err, "s=%s, scale=%d", s, scale)

		actualCoeff := num.BigInt()

		expectedCoeff, _ := (&big.Int{}).SetString(strings.Replace(s, ".", "", -1), 10)
		rq.Equal(expectedCoeff.Bytes(), actualCoeff.Bytes(), s)
	}
}

func TestFromStringDecimal128(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	const maxprec = 38

	for {
		scale := int(rand.Int31()) % (maxprec + 1)

		s := ""
		if maxprec-scale > 0 {
			for i := 0; i < maxprec-scale; i++ {
				s += randDigit()
			}
		} else {
			s += "0"
		}

		if scale > 0 {
			s += "."

			for i := 0; i < scale; i++ {
				s += randDigit()
			}
		}

		num, err := decimal128.FromString(s, maxprec, int32(scale))
		rq.NoErrorf(err, "s=%s, scale=%d", s, scale)

		actualCoeff := num.BigInt()

		expectedCoeff, _ := (&big.Int{}).SetString(strings.Replace(s, ".", "", -1), 10)
		rq.Equal(expectedCoeff.Bytes(), actualCoeff.Bytes(), s)
	}
}

func TestToStringDecimal256(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	const maxprec = 76

	for {
		scale := int(rand.Int31()) % (maxprec + 1)

		s := ""
		for {
			x := randDigit()
			if x == "0" {
				continue
			}
			s += x
			break
		}
		for i := 1; i < maxprec; i++ {
			s += randDigit()
		}
		integer, ok := (&big.Int{}).SetString(s, 10)
		rq.True(ok)
		dec := decimal256.FromBigInt(integer)

		var expected string

		if scale == maxprec {
			expected = "0." + s
		} else {
			for i := 0; i < maxprec-scale; i++ {
				expected += string(s[i])
			}
			if scale > 0 {
				expected += "."
				for i := maxprec - scale; i < maxprec; i++ {
					expected += string(s[i])
				}
			}
		}

		actual := dec.ToString(int32(scale))
		rq.Equal(expected, actual, "s=%s, scale=%d", s, scale)
	}
}

func TestToStringDecimal128(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	const maxprec = 38

	for {
		scale := int(rand.Int31()) % (maxprec + 1)

		s := ""
		for {
			x := randDigit()
			if x == "0" {
				continue
			}
			s += x
			break
		}
		for i := 1; i < maxprec; i++ {
			s += randDigit()
		}
		integer, ok := (&big.Int{}).SetString(s, 10)
		rq.True(ok)
		dec := decimal128.FromBigInt(integer)

		var expected string

		if scale == maxprec {
			expected = "0." + s
		} else {
			for i := 0; i < maxprec-scale; i++ {
				expected += string(s[i])
			}
			if scale > 0 {
				expected += "."
				for i := maxprec - scale; i < maxprec; i++ {
					expected += string(s[i])
				}
			}
		}

		actual := dec.ToString(int32(scale))
		rq.Equal(expected, actual, "s=%s, scale=%d", s, scale)
	}
}
