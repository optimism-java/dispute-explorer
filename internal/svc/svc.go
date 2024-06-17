package svc

import (
	"log"
	"time"

	"gorm.io/driver/postgres"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/optimism-java/dispute-explorer/internal/types"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var svc *ServiceContext

type ServiceContext struct {
	Config            *types.Config
	L1RPC             *ethclient.Client
	DB                *gorm.DB
	LatestBlockNumber int64
	SyncedBlockNumber int64
	SyncedBlockHash   common.Hash
}

func NewServiceContext(cfg *types.Config) *ServiceContext {
	storage, err := gorm.Open(postgres.Open(cfg.PostgresqlDataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Panicf("[svc]gorm get db panic: %s\n", err)
	}
	sqlDB, err := storage.DB()
	if err != nil {
		log.Panicf("[svc]gorm get sqlDB panic: %s\n", err)
	}
	// SetMaxIdleConns
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns)
	// SetMaxOpenConns
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns)
	// SetConnMaxLifetime
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MySQLConnMaxLifetime) * time.Second)

	rpc, err := ethclient.Dial(cfg.L1RPCUrl)
	if err != nil {
		log.Panicf("[svc] get eth client panic: %s\n", err)
	}

	svc = &ServiceContext{
		Config: cfg,
		L1RPC:  rpc,
		DB:     storage,
	}

	return svc
}
