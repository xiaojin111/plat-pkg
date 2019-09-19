package whitelist

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type RulePattern string

type Preset map[string][]RulePattern

func (p Preset) GetRulePatterns(name string) []RulePattern {
	return p[name]
}

func LoadPreset(data []byte) (Preset, error) {
	p := Preset{}
	err := yaml.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// LoadPresetFromYamlFile 从单个 yaml 配置文件加载
func LoadPresetFromYamlFile(file string) (Preset, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadPreset(content)
}

// LoadPresetFormYamlDir 从文件目录中加载所有 yaml 文件定义的配置，重复定义将返回错误
func LoadPresetFormYamlDir(dir string) (Preset, error) {
	r := make(Preset)

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
			p, err := LoadPresetFromYamlFile(path)
			if err != nil {
				return nil, err
			}
			r, err = mergePreset(r, p)
			if err != nil {
				return nil, err
			}
		}
	}

	return r, nil
}

func mergePreset(a, b Preset) (Preset, error) {
	if a == nil || b == nil {
		return nil, errors.New("empty preset can't be merged")
	}

	r := make(Preset)

	// deep copy a to r
	for k, v := range a {
		r[k] = v
	}

	for k, v := range b {
		if _, exist := r[k]; exist {
			return nil, fmt.Errorf("duplicate preset: %q", k)
		}

		r[k] = v
	}

	return r, nil
}
