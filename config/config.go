package config

type Server struct {
	BizMySQL `json:"bizMysql" yml:"bizMysql"`
	Redis    `json:"redis" yml:"redis"`
	Log      `json:"log" yml:"log"`
	RoMySQL  `json:"roMysql" yml:"roMysql"`
	JWT      `json:"jwt" yml:"jwt"`
	RabbitMQ `json:"rabbitMQ" yml:"rabbitMQ"`
}

type BizMySQL struct {
	Username     string `json:"username" yml:"username"`
	Password     string `json:"password" yml:"password"`
	Path         string `json:"path" yml:"path"`
	Database     string `json:"dbName" yml:"database"`
	Config       string `json:"config" yml:"config"`
	MaxIdleConns int    `json:"maxIdleConns" yml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yml:"maxOpenConns"`
	LogMode      bool   `json:"logMode" yml:"logMode"`
}

type RoMySQL struct {
	Username     string `json:"username" yml:"username"`
	Password     string `json:"password" yml:"password"`
	Path         string `json:"path" yml:"path"`
	Database     string `json:"dbName" yml:"database"`
	Config       string `json:"config" yml:"config"`
	MaxIdleConns int    `json:"maxIdleConns" yml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yml:"maxOpenConns"`
	LogMode      bool   `json:"logMode" yml:"log-mode"`
}

type Redis struct {
	Addr     string `json:"addr" yml:"addr"`
	Password string `json:"password" yml:"password"`
	DB       int    `json:"db" yml:"db"`
}

type Log struct {
	Prefix  string `json:"prefix" yml:"prefix"`
	LogFile bool   `json:"logFile" yml:"logFile"`
	Stdout  string `json:"stdout" yml:"stdout"`
	File    string `json:"file" yml:"file"`
}

type JWT struct {
	SigningKey string `json:"signingKey" yml:"signingKey"`
}

type RabbitMQ struct {
	Path string `json:"path" yml:"path"`
}
