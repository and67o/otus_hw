package configuration

type Config struct {
	Logger LoggerConf
	Server HTTPConf
	DB     DBConf
	Memory Memory
}

type LoggerConf struct {
	Level   string `mapstructure:"log_level"`
	File    string `mapstructure:"log_file"`
	IsProd  bool   `mapstructure:"log_trace_on"`
	TraceOn bool   `mapstructure:"log_prod_on"`
}

type HTTPConf struct {
	Host string `mapstructure:"server_host"`
	Port string `mapstructure:"server_port"`
}

type Memory struct {
	Type string `mapstructure:"memory_type"`
}

type DBConf struct {
	User   string `mapstructure:"db_user"`
	Pass   string `mapstructure:"db_password"`
	DBName string `mapstructure:"db_database"`
	Host   string `mapstructure:"db_host"`
}
