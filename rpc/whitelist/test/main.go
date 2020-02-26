package main

import (
	"fmt"

	"github.com/jinmukeji/plat-pkg/v2/rpc/whitelist"
)

func main() {
	//    p,err:= whitelist.LoadPresetFromYamlFile("../testdata/preset/all-in-one.yml")
	p, err := whitelist.LoadPolicyFormYamlDir("../testdata/policy")
	fmt.Println(p)
	fmt.Println(err)
}
