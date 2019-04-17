package config

import (
	"github.com/BurntSushi/toml"
	"strings"
)

type Config struct {
	Server []Server `toml:"server"`
	Php    Php
}

type Server struct {
	Port       string
	WebRoot    string `toml:"web_root"`
	Index      string
	SslEnabled bool    `toml:"ssl_enabled"`
	Key        string  `toml:"ssl_key"`
	Cert       string  `toml:"ssl_cert"`
	Vhosts     []Vhost `toml:"vhost"`
	Proxies    []Proxy `toml:"proxy"`
}

type Vhost struct {
	Name    string `toml:"server_name"`
	Port    string
	WebRoot string `toml:"web_root"`
	Index   string
	Proxies []Proxy `toml:"proxy"`
}

type Php struct {
	Enabled bool
	FpmSock string `toml:"fpm_sock"`
}

type Proxy struct {
	Path string
	Url  string
}

func (p *Proxy) IsWebsocketProxy() bool {
	proto := p.Url[0:strings.Index(p.Url, "://")]
	if proto == "ws" || proto == "wss" {
		return true
	}
	return false
}

func LoadConfig() Config {
	var conf Config
	_, err := toml.DecodeFile("config.tml", &conf)
	if err != nil {
		panic(err)
	}
	for i := range conf.Server {
		for j := range conf.Server[i].Vhosts {
			if conf.Server[i].Vhosts[j].Port == "" {
				conf.Server[i].Vhosts[j].Port = conf.Server[i].Port
			}
			if conf.Server[i].Vhosts[j].WebRoot == "" {
				conf.Server[i].Vhosts[j].WebRoot = conf.Server[i].WebRoot
			}
			if conf.Server[i].Vhosts[j].Index == "" {
				conf.Server[i].Vhosts[j].Index = conf.Server[i].Index
			}
		}
	}
	return conf
}
