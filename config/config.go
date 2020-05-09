package config


type Server struct {
	MySQL `json:"mysql" yml:"mysql"`
	Redis `json:"redis" yml:"redis"`
	Log   `json:"log" yml:"log"`
}

type MySQL struct {
	Username     string `json:"username" yml:"username"`
	Password     string `json:"password" yml:"password"`
	Path         string `json:"path" yml:"path"`
	Database       string `json:"dbName" yml:"database"`
	Config       string `json:"config" yml:"config"`
	MaxIdleConns int    `json:"maxIdleConns" yml:"max-idle-conns"`
	MaxOpenConns int    `json:"maxOpenConns" yml:"max-open-conns"`
	LogMode      bool   `json:"logMode" yml:"log-mode"`
}

type Redis struct {
	Addr     string `json:"addr" yml:"addr"`
	Password string `json:"password" yml:"password"`
	DB       int    `json:"db" yml:"db"`
}


type Log struct {
	Prefix  string `json:"prefix" yaml:"prefix"`
	LogFile bool   `json:"logFile" yaml:"log-file"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	File    string `json:"file" yaml:"file"`
}

