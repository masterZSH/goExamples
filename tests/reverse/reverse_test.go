package reverse

import (
	"testing"
	
)


// 基准测试
func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		reverse(s)
	}
}