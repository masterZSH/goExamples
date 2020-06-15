package split

import (
	"strconv"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests [][]string = [][]string{
		[]string{"a:b:c", ":", "3"},
		[]string{"a;c", ";", "2"},
		[]string{"a-b-c-d", "-", "4"},
	}

	for _, item := range tests {
		splits := strings.Split(item[0], item[1])
		want, _ := strconv.Atoi(item[2])
		if got := len(splits); got != want {
			t.Fatal("测试错误")
		}
	}

}
