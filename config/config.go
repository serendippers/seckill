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
	DBName       string `json:"db-name" yml:"db-name"`
	Config       string `json:"config" yml:"config"`
	MaxIdleConns int    `json:"max_idle_conns" yml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yml:"max_open_conns"`
	LogMode      bool   `json:"log_mode" yml:"log_mode"`
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

