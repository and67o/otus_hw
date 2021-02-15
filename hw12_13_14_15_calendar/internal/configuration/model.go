package configuration

type MemoryType string

type Config struct {
	Logger LoggerConf
	Rest   HTTPConf
	DB     DBConf
	Memory Memory
	GRPC   GRPCConf
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

type GRPCConf struct {
	Host string `mapstructure:"server_host"`
	Port string `mapstructure:"server_port"`
}

type Memory struct {
	Type MemoryType `mapstructure:"memory_type"`
}

type DBConf struct {
	User   string `mapstructure:"db_user"`
	Pass   string `mapstructure:"db_password"`
	DBName string `mapstructure:"db_database"`
	Host   string `mapstructure:"db_host"`
	Port   int    `mapstructure:"db_port"`
}
