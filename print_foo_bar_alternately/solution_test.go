package print_foo_bar_alternately

import (
	"testing"
)

func TestFooBar(t *testing.T) {
	n := 1000
	fb := NewFooBar(n)
	fb.RunTest()

	got := fb.String()
	want := "foobar"
	if wantLen := len(want) * n; len(got) != wantLen {
		t.Fatalf("got len %v, want len %v", len(got), wantLen)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			if gotIndex := i*6 + j; got[gotIndex] != want[j] {
				t.Fatalf("error at %v", gotIndex)
			}
		}
	}
}

func BenchmarkFooBar(b *testing.B) {
	n := 1000
	fb := NewFooBar(n)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fb.ResetBuffer()
		b.StartTimer()

		fb.RunBenchmark()
	}
}
