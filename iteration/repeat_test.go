package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("b")
	want := "bbbbb"
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	got := Repeat("a")
	fmt.Println(got)
	// Output: aaaaa
}
