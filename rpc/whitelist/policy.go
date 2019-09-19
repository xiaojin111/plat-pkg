package whitelist

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type WhitelistPolicy struct {
	ID        string   `json:"id" yaml:"id"`
	Whitelist []string `json:"whitelist" yaml:"whitelist"`
}

type PolicySet map[string]*WhitelistPolicy

func LoadPolicy(data []byte) (*WhitelistPolicy, error) {
	w := WhitelistPolicy{}
	err := yaml.Unmarshal(data, &w)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

// LoadPolicyFromYamlFile 从单个 yaml 配置文件加载
func LoadPolicyFromYamlFile(file string) (*WhitelistPolicy, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadPolicy(content)
}

// LoadPolicyFormYamlDir 从文件目录中加载所有 yaml 文件定义的配置，重复定义将返回错误
func LoadPolicyFormYamlDir(dir string) (PolicySet, error) {
	r := make(PolicySet)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(dir, file.Name())
		ext := filepath.Ext(path)
		if ext == ".yml" || ext == ".yaml" {
			p, err := LoadPolicyFromYamlFile(path)
			if err != nil {
				return nil, err
			}
			r[p.ID] = p
		}
	}

	return r, nil
}
