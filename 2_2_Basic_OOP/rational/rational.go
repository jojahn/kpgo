package rational

import "strconv"

type Rational struct {
	nominator int
	denominator int
}

func New(nominator, denominator int) Rational {
	if denominator == 0 {
		panic("rational/New: Cannot divide by Zero")
	}
	rat := Rational{}
	rat.nominator = nominator
	rat.denominator = denominator
	rat.shorten()
	return rat
}

func (x Rational) Multiply(y Rational) Rational {
	c := New(x.nominator * y.nominator, x.denominator * y.denominator)
	return c
}

func (x Rational) Add(y Rational) Rational {
	c := New(
		x.nominator * y.denominator + y.nominator * x.denominator,
		x.denominator * y.denominator)
	return c
}

func (x Rational) String() string {
	return strconv.Itoa(x.nominator) + " / " + strconv.Itoa(x.denominator)
}

func (x *Rational) shorten() {
	divisor := gcd(x.nominator, x.denominator)
	x.denominator /= divisor
	x.nominator /= divisor
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

