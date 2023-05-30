package boomer

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var tempPath = "temp.yaml"

type Temp struct {
	Node *Node `mapstructure:"node" json:"node" yaml:"node"`
}

type Node struct {
	ID string `mapstructure:"id" json:"id" yaml:"id"`
}

type tempConfig struct {
	*Temp
}

var TempConfig *tempConfig

func NewTempConfig() *tempConfig {
	if TempConfig != nil {
		return TempConfig
	}
	var cf Temp
	bytes, err := os.ReadFile(tempPath)
	if err != nil {
		fmt.Sprintln("临时配置文件不存在:", tempPath)
		cf.Node = new(Node)
		cf.Node.ID = ""
		configOutput, err := yaml.Marshal(cf)
		if err != nil {
			panic(fmt.Errorf("config 转换失败：%w", err))
		}
		if err := os.WriteFile(tempPath, configOutput, 0644); err != nil {
			panic(fmt.Errorf("配置文件写入失败：%w", err))
		}
	} else {
		err = yaml.Unmarshal(bytes, &cf)
		if err != nil {
			panic(fmt.Errorf("读取配置文件失败：%w", err))
		}
	}
	TempConfig = &tempConfig{Temp: &cf}
	return TempConfig
}

func (tc *tempConfig) GetNodeID() string {
	return tc.Node.ID
}

func (tc *tempConfig) SetNodeID(id string) {
	tc.Node.ID = id
	configOutput, err := yaml.Marshal(tc.Temp)
	if err != nil {
		panic(fmt.Errorf("config 转换失败：%w", err))
		return
	}
	if err := os.WriteFile(tempPath, configOutput, 0644); err != nil {
		panic(fmt.Errorf("配置文件写入失败：%w", err))
		return
	}
	return
}
