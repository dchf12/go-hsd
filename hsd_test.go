package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{"test1", "foo", "foo", 0},
		{"test2", "foo", "bar", 3},
		{"test3", "foo", "foh", 1},
		{"test4", "foo", "boo", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := StringDistance(tt.lhs, tt.rhs)
			if got != tt.want {
				t.Errorf("StringDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
