package mathutil

import (
	"fmt"
	"math"
	"testing"
)

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func TestSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		want int
	}{
		{name: "empty", in: nil, want: 0},
		{name: "one", in: []int{5}, want: 5},
		{name: "many", in: []int{1, 2, 3, 4}, want: 10},
		{name: "negatives", in: []int{-2, 4, -1}, want: 1},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := sum(tc.in...); got != tc.want {
				t.Fatalf("sum(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	t.Parallel()

	cases := map[int]bool{
		-4: false,
		0:  false,
		1:  false,
		2:  true,
		3:  true,
		4:  false,
		5:  true,
		9:  false,
		13: true,
	}

	for n, want := range cases {
		n := n
		want := want
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			if got := isPrime(n); got != want {
				t.Fatalf("isPrime(%d) = %v, want %v", n, got, want)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b int
		want int
	}{
		{a: 54, b: 24, want: 6},
		{a: 24, b: 54, want: 6},
		{a: -24, b: 54, want: 6},
		{a: 17, b: 13, want: 1},
		{a: 0, b: 12, want: 12},
	}

	for _, tc := range tests {
		if got := gcd(tc.a, tc.b); got != tc.want {
			t.Fatalf("gcd(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sum(input...)
	}
}

func FuzzGCD(f *testing.F) {
	seeds := [][2]int{
		{54, 24},
		{17, 13},
		{100, 10},
	}
	for _, s := range seeds {
		f.Add(s[0], s[1])
	}

	f.Fuzz(func(t *testing.T, a, b int) {
		got := gcd(a, b)
		if got < 0 {
			t.Fatalf("gcd(%d, %d) must be non-negative, got %d", a, b, got)
		}
		if a != 0 && got != 0 && a%got != 0 {
			t.Fatalf("gcd(%d, %d) = %d does not divide %d", a, b, got, a)
		}
		if b != 0 && got != 0 && b%got != 0 {
			t.Fatalf("gcd(%d, %d) = %d does not divide %d", a, b, got, b)
		}
	})
}
