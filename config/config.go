package config

// GlobalConfigSetting 配置实例
var (
	GlobalConfigSetting = &Config{}
	DatabaseSetting     = &Database{}
)

// Config 配置对象
type (
	Config struct {
		Database *Database `json:"database"`
		Email    *Email    `json:"email"`
		Redis    *Redis    `json:"redis"`
	}

	Database struct {
		Type        string `json:"type"`
		User        string `json:"user"`
		Password    string `json:"password"`
		Host        string `json:"host"`
		Port        string `json:"port"`
		Name        string `json:"name"`
		TablePrefix string `json:"table_prefix"`
	}

	Redis struct {
		Db       string `json:"db"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	}

	Email struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	}
)
