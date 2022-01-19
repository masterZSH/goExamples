package main

import (
	"fmt"

	cuckoo "github.com/seiflotfy/cuckoofilter"
)

func main() {
	cf := cuckoo.NewFilter(1000)
	cf.InsertUnique([]byte("123"))
	cf.Insert([]byte("foo"))
	cf.Insert([]byte("foo"))

	// 查询
	r := cf.Lookup([]byte("123"))
	fmt.Println(r)

	// 数量
	ct := cf.Count()
	fmt.Println(ct)

	// 删除
	cf.Delete([]byte("123"))

	// 重置
	cf.Reset()
}
