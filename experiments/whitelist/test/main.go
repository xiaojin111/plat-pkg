package main

import (
	"fmt"

	"gitee.com/jt-heath/plat-pkg/v2/experiments/whitelist"
)

func main() {
	//    p,err:= whitelist.LoadPresetFromYamlFile("../testdata/preset/all-in-one.yml")
	p, err := whitelist.LoadPolicyFormYamlDir("../testdata/policy")
	fmt.Println(p)
	fmt.Println(err)
}
