package types

import (
	"log"

	env "github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat                string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource          string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3367)/dispute_explorer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns        int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns        int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime     int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	Blockchain               string `env:"BLOCKCHAIN" envDefault:"sepolia"`
	L1RPCUrl                 string `env:"L1_RPC_URL" envDefault:"https://eth-sepolia.g.alchemy.com/v2/RT1mCGRyVMx1F-XlY4Es4Zz-Q8Jrasg6"`
	L2RPCUrl                 string `env:"L2_RPC_URL" envDefault:"https://opt-sepolia.g.alchemy.com/v2/RT1mCGRyVMx1F-XlY4Es4Zz-Q8Jrasg6"`
	RPCRateLimit             int    `env:"RPC_RATE_LIMIT" envDefault:"15"`
	RPCRateBurst             int    `env:"RPC_RATE_BURST" envDefault:"5"`
	FromBlockNumber          int64  `env:"FROM_BLOCK_NUMBER" envDefault:"5515562"`
	FromBlockHash            string `env:"FROM_BLOCK_HASH" envDefault:"0x5205c17557759edaef9120f56af802aeaa2827a60d674a0413e77e9c515bdfba"`
	DisputeGameProxyContract string `env:"DISPUTE_GAME_PROXY_CONTRACT" envDefault:"0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1"`
	APIPort                  string `env:"API_PORT" envDefault:"8088"`
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
