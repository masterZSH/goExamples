package main

import (
	"fmt"
	"strings"
)

var co string = "We should vet packages that are part of the registry to make sure that they don't refer to other hosts, outside of deno.land/x/. We have had reports of some packages referring to denopkg.com, which is potentially a useful community feature but we should either include that capability (ability to refer easily to repos on GitHub) into deno.land/x/ for the use such packages, or we expect everything to only leverage other packages on deno.land/x/."

func main() {
	s1 := strings.Split(co, " ")
	mp := make(map[string]int)
	for _, v := range s1 {
		_, ok := mp[v]
		if !ok {
			mp[v] = 0
		}
		mp[v]++
	}
	fmt.Printf("%q\n", s1)
	fmt.Printf("%+v\n", mp)
}
