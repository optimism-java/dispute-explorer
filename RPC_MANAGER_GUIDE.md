# RPC Manager 使用指南

## 概述

RPC Manager 是一个统一的RPC资源管理器，为项目中的所有RPC调用提供速率限制、监控和统一管理功能。

## 主要特性

### ✅ 已实现的功能
- **统一速率限制** - 所有RPC调用共享配置的速率限制
- **L1/L2分离管理** - 分别管理以太坊主网和Optimism网络的调用
- **实时监控** - 统计请求数量、速率限制状态等
- **向后兼容** - 保持现有代码可以正常工作
- **健康检查** - 提供系统健康状态检查

### 🔧 核心组件
1. **Manager** - 核心RPC管理器
2. **Monitor** - 监控和统计组件
3. **Factory** - 工厂函数，便于创建和配置

## 快速开始

### 1. 基本使用

```go
// 获取最新区块号（自动应用速率限制）
latest, err := ctx.RpcManager.GetLatestBlockNumber(context.Background(), true) // L1
latest, err := ctx.RpcManager.GetLatestBlockNumber(context.Background(), false) // L2

// 获取指定区块
block, err := ctx.RpcManager.GetBlockByNumber(context.Background(), big.NewInt(12345), true)

// HTTP JSON-RPC调用
response, err := ctx.RpcManager.HTTPPostJSON(context.Background(), requestBody, true)
```

### 2. 监控使用

```go
// 创建监控器
monitor := rpc.NewMonitor(ctx.RpcManager, 30*time.Second)

// 启动监控（在goroutine中）
go monitor.Start(context.Background())

// 获取健康检查信息
health := monitor.GetHealthCheck()
if !health.Healthy {
    log.Printf("RPC系统有问题: %v", health.Issues)
}
```

## 迁移指南

### 从现有代码迁移

#### 迁移前（没有速率限制）
```go
// ❌ 旧方式 - 没有速率限制
latest, err := ctx.L1RPC.BlockNumber(context.Background())
blockJSON, err := rpc.HTTPPostJSON("", ctx.Config.L1RPCUrl, requestBody)
```

#### 迁移后（有速率限制）
```go
// ✅ 新方式 - 自动应用速率限制
latest, err := ctx.RpcManager.GetLatestBlockNumber(context.Background(), true)
response, err := ctx.RpcManager.HTTPPostJSON(context.Background(), requestBody, true)
```

### 渐进式迁移策略

#### 阶段1：高频调用优先
优先迁移这些文件中的调用：
- `internal/handler/syncBlock.go`
- `internal/handler/latestBlockNumber.go` 
- `internal/handler/logFilter.go`

#### 阶段2：API层迁移
- `internal/api/dispute_game_handler.go`

#### 阶段3：完全替换
移除对原始 `ctx.L1RPC` 和 `ctx.L2RPC` 的依赖

## 配置说明

### 环境变量配置
```bash
# RPC速率限制配置
RPC_RATE_LIMIT=15    # 每秒允许的请求数
RPC_RATE_BURST=5     # 允许的突发请求数

# RPC节点URL
L1_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/your-key
L2_RPC_URL=https://opt-sepolia.g.alchemy.com/v2/your-key
```

### 动态调整限制
```go
// 运行时调整L1速率限制
ctx.RpcManager.UpdateRateLimit(20, 10, true)  // L1: 20 req/s, burst 10

// 运行时调整L2速率限制  
ctx.RpcManager.UpdateRateLimit(30, 15, false) // L2: 30 req/s, burst 15
```

## 监控和统计

### 基本统计信息
```go
stats := ctx.RpcManager.GetStats()
fmt.Printf("L1请求数: %d\n", stats.L1RequestCount)
fmt.Printf("L2请求数: %d\n", stats.L2RequestCount)
fmt.Printf("被限制的请求数: %d\n", stats.L1RateLimitedCount)
```

### 实时状态
```go
// 检查可用令牌数
l1Tokens := ctx.RpcManager.GetTokens(true)
l2Tokens := ctx.RpcManager.GetTokens(false)

// 检查当前限制配置
l1Rate, l1Burst := ctx.RpcManager.GetRateLimit(true)
l2Rate, l2Burst := ctx.RpcManager.GetRateLimit(false)
```

## 常见问题

### Q: 如何处理速率限制错误？
```go
latest, err := ctx.RpcManager.GetLatestBlockNumber(context.Background(), true)
if err != nil {
    if strings.Contains(err.Error(), "rate limit exceeded") {
        // 处理速率限制错误
        time.Sleep(1 * time.Second)
        // 重试逻辑
    }
    return err
}
```

### Q: 如何在不同环境使用不同的限制？
```go
// 开发环境
if config.Environment == "development" {
    ctx.RpcManager.UpdateRateLimit(5, 2, true)   // 较低的限制
} else {
    ctx.RpcManager.UpdateRateLimit(50, 20, true) // 生产环境较高的限制
}
```

### Q: 如何监控RPC使用情况？
```go
// 定期记录健康检查
go func() {
    ticker := time.NewTicker(1 * time.Minute)
    for range ticker.C {
        monitor.LogHealthCheck()
    }
}()
```

## 最佳实践

### 1. 合理设置速率限制
- **开发环境**: 低限制，避免意外消耗配额
- **测试环境**: 中等限制，模拟真实场景
- **生产环境**: 根据RPC提供商的限制设置

### 2. 监控设置
- 设置合理的监控间隔（建议30秒-5分钟）
- 在速率限制达到80%时发出警告
- 记录所有被限制的请求用于分析

### 3. 错误处理
- 对速率限制错误实现重试机制
- 使用指数退避策略
- 记录详细的错误日志用于调试

### 4. 性能优化
- 对频繁查询的数据实现缓存
- 合并可以批量处理的请求
- 优先使用L2网络（通常限制更宽松）

## 故障排除

### 检查速率限制状态
```bash
# 查看当前RPC统计
curl http://localhost:8088/health/rpc
```

### 常见错误信息
- `rate limit exceeded` - 速率限制已达到上限
- `context deadline exceeded` - 请求超时
- `connection refused` - RPC节点连接失败

### 调试技巧
1. 开启详细日志记录
2. 监控令牌消耗速度
3. 检查RPC节点响应时间
4. 分析请求模式

## 升级和维护

### 版本兼容性
- 当前版本保持与现有代码的完全向后兼容
- 新功能通过 `ctx.RpcManager` 访问
- 旧的 `ctx.L1RPC` 和 `ctx.L2RPC` 继续可用

### 未来计划
- [ ] 实现L1/L2独立的速率限制配置
- [ ] 添加请求优先级队列
- [ ] 实现智能负载均衡
- [ ] 添加缓存层减少RPC调用
