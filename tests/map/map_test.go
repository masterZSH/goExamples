package mp

import (
	"io"
	"strconv"
	"testing"
)

var testData map[string]int = make(map[string]int)
var testKeys []string = []string{"1", "2"}

func TestGetKeys(t *testing.T) {
	for _, v := range testKeys {
		testData[v], _ = strconv.Atoi(v)
	}
	keys := getKeys(testData)
	if len(keys) != len(testKeys) {
		t.Error("长度错误")
	}

	for _, v := range testKeys {
		if _, ok := testData[v]; !ok {
			t.Error("key不存在")
		}
	}
}
