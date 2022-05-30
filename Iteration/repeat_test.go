package iteration

import "testing"

//Unit Testing

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 8)
	expected := "aaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestClone(t *testing.T) {
	got := clone("myString")
	want := "myString"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func TestCompate(t *testing.T) {
	got := Compare("qwerty", "qwerty")
	want := 0

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

// Benchmark testing.

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}
