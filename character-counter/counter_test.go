package character_counter

import "testing"

func Test_viaSharedMem(t *testing.T) {
	type args struct {
		ii   []string
		what string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "looking for 2 a",
			args: struct {
				ii   []string
				what string
			}{ii: []string{"aabbcc", "ddee"}, what: "a"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := viaSharedMem(tt.args.ii, tt.args.what); got != tt.want {
				t.Errorf("viaSharedMem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_viaChan(t *testing.T) {
	type args struct {
		ii   []string
		what string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "looking for 2 a",
			args: struct {
				ii   []string
				what string
			}{ii: []string{"aabbcc", "ddee"}, what: "a"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := viaChan(tt.args.ii, tt.args.what); got != tt.want {
				t.Errorf("viaSharedMem() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_viaSharedMem-8   	     344	   3262262 ns/op
func Benchmark_viaSharedMem(b *testing.B) {
	ii := make([]string, 10000)
	for i := 0; i < len(ii); i++ {
		ii[i] = "aabbcc"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		viaSharedMem(ii, "a")
	}
}

func Benchmark_viaChan(b *testing.B) {
	ii := make([]string, 10000)
	for i := 0; i < len(ii); i++ {
		ii[i] = "aabbcc"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		viaChan(ii, "a")
	}
}
