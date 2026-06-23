package config

// ServerConfig 服务配置
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// CorsConfig 跨域配置
type CorsConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
}

// AuthConfig 认证配置（Admin 和 Api 共用字段）
type AuthConfig struct {
	RefreshTokenTtl int    `mapstructure:"refresh_token_ttl"`
	AccessTokenTtl  int    `mapstructure:"access_token_ttl"`
	HmacSecret      string `mapstructure:"hmac_secret"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level   string        `mapstructure:"level"`
	Writers []string      `mapstructure:"writers"`
	File    FileLogConfig `mapstructure:"file"`
}

// FileLogConfig 文件日志配置
type FileLogConfig struct {
	Path       string `mapstructure:"path"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
	LocalTime  bool   `mapstructure:"local_time"`
}

// MySQLConfig MySQL配置
type MySQLConfig struct {
	Dsn    string `mapstructure:"dsn"`
	Prefix string `mapstructure:"prefix"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// MongoDBConfig MongoDB配置
type MongoDBConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
	Enabled  bool   `mapstructure:"enabled"`
}
