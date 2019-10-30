package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/bloomfilter"
)

func main() {
	filter := bloom.New(100, 4)
	fmt.Printf("%v\n", filter.Contains("gggg"))

	filter.Add("hhhh")
	filter.Add("gggg")

	fmt.Printf("%v\n", filter.Contains("aaaa"))
	fmt.Printf("%v\n", filter.Contains("bbbb"))
	fmt.Printf("%v\n", filter.Contains("cccc"))
	fmt.Printf("%v\n", filter.Contains("hhhh"))
	fmt.Printf("%v\n", filter.Contains("gggg"))
}
