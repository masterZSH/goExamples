package main

import (
	"fmt"

	"github.com/jarcoal/httpmock"
)

func main() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// 注册
	httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles",
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))
	c := httpmock.GetTotalCallCount()
	fmt.Println(c)
	info := httpmock.GetCallCountInfo()
	fmt.Println(info)

}
