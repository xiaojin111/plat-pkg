package service

import "fmt"

// printFullVersionInfo 打印版本信息
func printFullVersionInfo(opts *Options) {
	fmt.Printf("%s %s\n", opts.FQDN(), opts.ProductVersion)
	fmt.Println("GitCommit:", opts.GitCommit)
	fmt.Println("BuildTime:", opts.BuildTime)
	fmt.Println("Go:", opts.GoVersion)
}

func logVersionInfo(opts *Options) {
	log.Infof("Starting service: %s", opts.FQDN())
	log.Infof("Product Version: %s", opts.ProductVersion)
	log.Infof("Git Commit: %s", opts.GitCommit)
	log.Infof("Build Time: %s", opts.BuildTime)
	log.Infof("Go Version: %s", opts.GoVersion)
}
