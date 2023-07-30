package currency

type Currency interface {
	Times(float64) float64
	Equals(Money) bool
}

type Money struct {
	amount   float64
	currency string
}

func (m *Money) Times(multiplier float64) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m *Money) Equals(money Money) bool {
	return m.amount == money.amount
}

func (m *Money) Add(money float64) Money {
	return Money{amount: m.amount + money}
}

func (m *Money) GetAmount() float64 {
	return m.amount
}

type Dollar struct {
	Money
}

func NewDollar(amount float64) *Dollar {
	return &Dollar{
		Money{amount: amount, currency: "USD"},
	}
}

type Franc struct {
	Money
}

func NewFranc(amount float64) *Franc {
	return &Franc{
		Money{amount: amount, currency: "CHF"},
	}
}
