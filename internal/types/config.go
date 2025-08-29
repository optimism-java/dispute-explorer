package types

import (
	"log"
	"strings"

	env "github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat                string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource          string `env:"MYSQL_DATA_SOURCE" envDefault:"root:123456@tcp(127.0.0.1:3306)/dispute_explorer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns        int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns        int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime     int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	Blockchain               string `env:"BLOCKCHAIN" envDefault:"sepolia"`
	L1RPCUrls                string `env:"L1_RPC_URLS" envDefault:"https://eth-sepolia.g.alchemy.com/v2/RT1mCGRyVMx1F-XlY4Es4Zz-Q8Jrasg6"`
	L2RPCUrls                string `env:"L2_RPC_URLS" envDefault:"https://opt-sepolia.g.alchemy.com/v2/RT1mCGRyVMx1F-XlY4Es4Zz-Q8Jrasg6"`
	NodeRPCUrls              string `env:"NODE_RPC_URLS" envDefault:"https://light-radial-slug.optimism-sepolia.quiknode.pro/e9329f699b371572a8cc5dd22d19d5940bb842a5/"`
	RPCRateLimit             int    `env:"RPC_RATE_LIMIT" envDefault:"5"`
	RPCRateBurst             int    `env:"RPC_RATE_BURST" envDefault:"2"`
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

// GetL1RPCUrls 获取L1 RPC URL列表
func (c *Config) GetL1RPCUrls() []string {
	urls := strings.Split(c.L1RPCUrls, ",")
	// 去除空格
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}
	return urls
}

// GetL2RPCUrls 获取L2 RPC URL列表
func (c *Config) GetL2RPCUrls() []string {
	urls := strings.Split(c.L2RPCUrls, ",")
	// 去除空格
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}
	return urls
}

// GetNodeRPCUrls 获取Node RPC URL列表
func (c *Config) GetNodeRPCUrls() []string {
	urls := strings.Split(c.NodeRPCUrls, ",")
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}
	return urls
}
