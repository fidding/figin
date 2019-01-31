package config

// Configuration 配置结构
type Configuration struct {
	Server server
	DB     database
	Cache  cache
	Logger logger
}

// server 基本配置结构体
type server struct {
	Name       string `mapstructure:"name"`
	Env        string `mapstructure:"env"`
	Port       int    `mapstructure:"port"`
	Author     string `mapstructure:"author"`
	Debug      bool   `mapstructure:"debug"`
	DebugLevel string `mapstructure:"debugLevel"`
}

// database 数据库结构体
type database struct {
	DBType   string `mapstructure:"dbType"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
}

// cache 缓存配置结构体
type cache struct {
	CacheType string `mapstructure:"cacheType"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Password  string `mapstructure:"password"`
	DB        int    `mapstructure:"db"`
}

// logger 日志配置结构体
type logger struct {
	LogType string `mapstructure:"logType"`
	Level   string `mapstructure:"level"`
}
