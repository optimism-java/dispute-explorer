package svc

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/optimism-java/dispute-explorer/internal/types"
	"github.com/optimism-java/dispute-explorer/pkg/rpc"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var svc *ServiceContext

type ServiceContext struct {
	Config            *types.Config
	L1RPC             *ethclient.Client // 保留向后兼容
	L2RPC             *ethclient.Client // 保留向后兼容
	RpcManager        *rpc.Manager      // 新增统一RPC管理器
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

	// 创建原有的以太坊客户端（保持向后兼容）
	l1Client, err := ethclient.Dial(cfg.L1RPCUrl)
	if err != nil {
		log.Panicf("[svc] get L1 eth client panic: %s\n", err)
	}

	l2Client, err := ethclient.Dial(cfg.L2RPCUrl)
	if err != nil {
		log.Panicf("[svc] get L2 eth client panic: %s\n", err)
	}

	// 创建统一的RPC管理器
	rpcManager, err := rpc.CreateManagerFromConfig(cfg)
	if err != nil {
		log.Panicf("[svc] create RPC manager panic: %s\n", err)
	}

	svc = &ServiceContext{
		Config:     cfg,
		L1RPC:      l1Client,   // 保留向后兼容
		L2RPC:      l2Client,   // 保留向后兼容
		RpcManager: rpcManager, // 新的统一管理器
		DB:         storage,
		Context:    ctx,
	}
	return svc
}
