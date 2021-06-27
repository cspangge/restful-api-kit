package config

import (
	"encoding/json"
	"os"
	tools "restful-api-kit/utilities"
)

// Setup 配置
func Setup(filePath ...string) {
	var filePtr *os.File
	var err error

	mode := tools.GetEnv("mode")
	if len(filePath) != 0 {
		filePtr, err = os.Open(filePath[0]) //config的文件目录
	} else {
		if mode == "development" {
			filePtr, err = os.Open("config/dbStaging.json") //config的文件目录
		} else if mode == "production" {
			filePtr, err = os.Open("config/dbProduction.json") //config的文件目录
		} else {
			filePtr, err = os.Open("config/dbStaging.json") //config的文件目录
		}
	}

	if err != nil {
		tools.CheckErr(err)
		return
	}
	defer func(filePtr *os.File) {
		err := filePtr.Close()
		if err != nil {
			tools.CheckErr(err)
		}
	}(filePtr)
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database
}

// DatabaseSetting 数据库配置对象 实例
