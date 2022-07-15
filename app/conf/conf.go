package conf

import (
	"flag"
	"ginweb/common/conf"
	"ginweb/common/logz"
	"ginweb/common/mysql"
	"ginweb/common/redis"
)

var (
	confPath string
	Conf     = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "config.yaml", "指定配置文件 eg: -conf config.yaml")
}

type Config struct {
	Server *Server
	Data   *Data
	Log    *logz.Config
}

type Auth struct {
	Expire string
	Secret string
	Issuer string
}

type Data struct {
	Mysql *mysql.Config
	Redis *redis.Config
}

type Server struct {
	Addr    string
	Timeout int
}

type Client struct {
	Mgr struct {
		Endpoint string
	}
	Country struct {
		Endpoint string
	}
}

func LoadConfig() {
	conf.LoadFromYaml(confPath, &Conf)
}
