package print_in_order

import "testing"

func TestFoo(t *testing.T) {
	f := NewFoo()
	f.RunTest()

	//goland:noinspection SpellCheckingInspection
	got, want := f.String(), "firstsecondthird"
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f := NewFoo()
		b.StartTimer()

		f.RunBenchmark()
	}
}
