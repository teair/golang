package go_3

import (
	"testing"
)

func TestNoneRepeatingString(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		// Normal cases
		{"abcabcdabc", 4},
		{"bbbb", 1},
		{"pwwkew", 3},
		{"c", 1},
		{"abcdefg", 7},

		// Edge cases
		{"", 0},

		// Chinese support
		{"我是测试数据，哈哈哈哈哈哈", 8},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.n {
			t.Errorf("lengthOfNonRepeatingSubStr(%s); actual:%d; expected:%d", tt.s, actual, tt.n)
		}
	}
}

// BenchmarkNoneSubstr 性能测试
func BenchmarkNoneSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"

	for j := 0; j < 13; j++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))

	// 重置计时器【不算上边组合字符串的时间】
	b.ResetTimer()

	n := 8
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != n {
			b.Errorf("lengthOfNonRepeatingSubStr(%s); actual:%d; expected:%d", s, actual, n)
		}
	}
}
