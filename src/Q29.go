package src

import "math"

func Divide(dividend, divisor int) int {
	a, b := dividend, divisor
	if a == 0 || math.Abs(float64(a)) < math.Abs(float64(b)) {
		return 0
	}

	sign := false

	if a < 0 && b < 0 {
		a, b = -a, -b
	}
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		a, b = abs(a), abs(b)
		sign = true
	}

	d := 0

	for a >= b {
		c := 0
		b1 := b << 1
		for a >= b1 {
			c += 1
			b1 = b1 << 1
		}
		b1 = b1 >> 1
		d += int(math.Pow(2, float64(c)))
		a -= b1
	}

	if sign{
		d=-d
	}

	if d>math.MaxInt32 ||d<math.MinInt32{
		d = math.MaxInt32
	}
	return d

}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
