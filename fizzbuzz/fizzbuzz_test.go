package fizzbuzz_test

import (
	"fmt"
	"math/rand"
	"prospersof/fizzbuzz"
	"testing"
	"time"
)

func TestFizzBuzzGiven1(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven2(t *testing.T) {
	given := 2
	want := "2"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven3(t *testing.T) {
	given := 3
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven4(t *testing.T) {
	given := 4
	want := "4"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven5(t *testing.T) {
	given := 5
	want := "Buzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven6(t *testing.T) {
	given := 6
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven9(t *testing.T) {
	given := 9
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven10(t *testing.T) {
	given := 10
	want := "Buzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven15(t *testing.T) {
	given := 15
	want := "FizzBuzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven30(t *testing.T) {
	given := 30
	want := "FizzBuzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFormat(t *testing.T) {
	want := "Fizz-Fizz-Fizz-Fizz"
	given := func(int) int { return 3 }
	get := fizzbuzz.Format(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given(100), want, get)
	}
}

func closureRandom() fizzbuzz.RandFunc {
	i := 0
	store := []int{1, 3, 5, 15}
	return func(int) int {
		defer func() { i++ }()
		return store[i]
	}
}

func TestFormatByClosure(t *testing.T) {
	want := "1-Fizz-Buzz-FizzBuzz"
	closure := closureRandom

	get := fizzbuzz.Format(closure())

	if want != get {
		t.Errorf("want %q but get %q\n", want, get)
	}
}

func TestFormatInReal(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	fn := func(n int) int {
		return r.Intn(n) + 1
	}

	fmt.Println(fizzbuzz.Format(fn))
}

type Intn struct {
	i     int
	store []int
}

func (i *Intn) Intn(int) int {
	defer func() { i.i++ }()
	return i.store[i.i] - 1
}

func TestFormatGetInterface(t *testing.T) {
	want := "2-4-Fizz-16"
	// given := &Intn{i: 0, store: []int{1, 3, 5, 15}}
	given := closureRandom()
	get := fizzbuzz.FormatGetInterface(given)

	if want != get {
		t.Errorf("want %q but get %q\n", want, get)
	}
}
