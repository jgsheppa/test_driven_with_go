package currency_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/test_driven_with_go/currency"
)

func TestCurrencyMultiplication(t *testing.T) {
	t.Parallel()
	dollar := currency.NewDollar(5)

	got := dollar.Times(2)
	want := currency.NewDollar(10).Money
	if !cmp.Equal(got, want, cmp.AllowUnexported(currency.Money{})) {
		t.Errorf("got %v want %v", got, want)
	}

	got = dollar.Times(3)
	want = currency.NewDollar(15).Money
	if !cmp.Equal(got, want, cmp.AllowUnexported(currency.Money{})) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCurrencyEquals_IsEqual(t *testing.T) {
	t.Parallel()

	money := currency.NewDollar(5).Money
	moreMoney := currency.NewDollar(5).Money

	got := money.Equals(moreMoney)
	want := true

	if !cmp.Equal(got, want, cmp.AllowUnexported(currency.Money{})) {
		t.Error(cmp.Diff(want, got))
	}

	money = currency.NewFranc(5).Money
	moreMoney = currency.NewFranc(5).Money

	got = money.Equals(moreMoney)
	want = true

	if !cmp.Equal(got, want, cmp.AllowUnexported(currency.Money{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCurrencyEquals_IsNotEqual(t *testing.T) {
	t.Parallel()

	money := currency.NewDollar(5).Money

	moreMoney := currency.NewFranc(6).Money
	got := money.Equals(moreMoney)
	want := false

	if !cmp.Equal(got, want, cmp.AllowUnexported(currency.Money{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestEqualityOfTwoCurrencies(t *testing.T) {
	t.Parallel()
	got := currency.NewDollar(5)
	want := currency.NewFranc(5)

	if cmp.Equal(got, want, cmp.AllowUnexported(currency.Money{})) {
		t.Error(cmp.Diff(want, got))
	}
}
