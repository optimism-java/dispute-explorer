package svc

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/optimism-java/dispute-explorer/internal/types"
	rpcpkg "github.com/optimism-java/dispute-explorer/pkg/rpc"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var svc *ServiceContext

type ServiceContext struct {
	Config            *types.Config
	L1RPC             *ethclient.Client
	L2RPC             *ethclient.Client
	L1RPCManager      *rpcpkg.ClientManager // L1 RPC client manager (primary)
	L2RPCManager      *rpcpkg.ClientManager // L2 RPC client manager (primary)
	DB                *gorm.DB
	LatestBlockNumber int64
	SyncedBlockNumber int64
	SyncedBlockHash   common.Hash
	Context           context.Context
}

func NewServiceContext(ctx context.Context, cfg *types.Config) *ServiceContext {
	storage, err := gorm.Open(mysql.Open(cfg.MySQLDataSource), &gorm.Config{
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

	// Create main RPC clients (maintain backward compatibility)
	rpc, err := ethclient.Dial(cfg.L1RPCUrl)
	if err != nil {
		log.Panicf("[svc] get L1 eth client panic: %s\n", err)
	}

	rpc2, err := ethclient.Dial(cfg.L2RPCUrl)
	if err != nil {
		log.Panicf("[svc] get L2 eth client panic: %s\n", err)
	}

	// Create L1 RPC client manager
	var l1RPCManager *rpcpkg.ClientManager
	if cfg.L1RPCUrls != "" {
		l1RPCManager, err = rpcpkg.NewClientManager(
			cfg.L1RPCUrls,
			cfg.RPCMaxRetries,
			time.Duration(cfg.RPCRetryDelay)*time.Second,
		)
		if err != nil {
			log.Printf("[svc] Failed to create L1 RPC manager: %v\n", err)
		}
	}

	// Create L2 RPC client manager
	var l2RPCManager *rpcpkg.ClientManager
	if cfg.L2RPCUrls != "" {
		l2RPCManager, err = rpcpkg.NewClientManager(
			cfg.L2RPCUrls,
			cfg.RPCMaxRetries,
			time.Duration(cfg.RPCRetryDelay)*time.Second,
		)
		if err != nil {
			log.Printf("[svc] Failed to create L2 RPC manager: %v\n", err)
		}
	}

	svc = &ServiceContext{
		Config:       cfg,
		L1RPC:        rpc,
		L2RPC:        rpc2,
		L1RPCManager: l1RPCManager,
		L2RPCManager: l2RPCManager,
		DB:           storage,
		Context:      ctx,
	}

	return svc
}
