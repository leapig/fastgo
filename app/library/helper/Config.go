package helper

import (
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"strings"
)

// Config 从 config.yaml 文件中读取配置信息，并将这些信息设置为环境变量
func Config() {
	// 读取 config.yaml 文件
	yamlFile, err := os.ReadFile("fastgo.yaml")
	// 如果文件读取没有错误
	if err == nil {
		// 定义一个 map 用于存储解析后的 YAML 数据
		var conf map[string]interface{}
		// 解析 YAML 文件内容到 conf 变量中
		err = yaml.Unmarshal(yamlFile, &conf)
		// 如果解析没有错误
		if err == nil {
			// 遍历 conf 中的每个键值对
			for key, value := range conf {
				// 如果值是一个 map
				if isMap(value) {
					// 将该 map 中的键值对设置为环境变量
					setEnv(getKey(key), value.(map[interface{}]interface{}))
				} else {
					// 否则，直接将值设置为环境变量
					os.Setenv(getKey(key), value.(string))
				}
			}
		}
	}
}

// setEnv 将给定前缀和配置 map 中的键值对设置为环境变量
func setEnv(prefix string, conf map[interface{}]interface{}) {
	// 遍历 conf 中的每个键值对
	for key, value := range conf {
		// 如果值是一个 map
		if isMap(value) {
			// 递归地设置环境变量
			setEnv(getKey(prefix, key.(string)), value.(map[interface{}]interface{}))
		} else {
			// 否则，直接设置环境变量
			os.Setenv(getKey(prefix, key.(string)), value.(string))
		}
	}
}

// getKey 根据给定的字符串参数生成一个键
func getKey(args ...string) (key string) {
	// 遍历所有参数
	for index, value := range args {
		// 如果是第一个参数
		if index == 0 {
			// 将其作为键的初始值
			key = value
		} else {
			// 否则，将其拼接到键的后面，使用下划线分隔
			key = key + "_" + value
		}
	}
	// 将键转换为大写
	key = strings.ToUpper(key)
	return
}

// isMap 检查给定的值是否是一个 map
func isMap(v interface{}) bool {
	// 如果值为 nil，返回 false
	if v == nil {
		return false
	}
	// 获取值的反射类型
	rv := reflect.TypeOf(v)
	// 如果类型不是 map，返回 false
	if rv.Kind() != reflect.Map {
		return false
	}
	// 否则，返回 true
	return true
}
