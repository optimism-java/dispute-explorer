package types

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat                string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource          string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3366)/dispute_explorer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns        int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns        int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime     int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	Blockchain               string `env:"BLOCKCHAIN" envDefault:"sepolia"`
	L1RPCUrl                 string `env:"L1_RPC_URL" envDefault:"https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30"`
	FromBlockNumber          int64  `env:"FROM_BLOCK_NUMBER" envDefault:"6026723"`
	FromBlockHash            string `env:"FROM_BLOCK_HASH" envDefault:"0x013a38a1a8b5073a852134716263c1e4b911f65bd7112d90cc0e93b1defcf3a8"`
	DisputeGameProxyContract string `env:"DISPUTE_GAME_PROXY_CONTRACT" envDefault:"0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		config = cfg
	}
	return config
}
