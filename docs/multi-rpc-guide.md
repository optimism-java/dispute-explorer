# 多 RPC URL 配置指南

## 概述

现在项目支持多个 RPC URL 配置，提供自动故障转移和负载均衡功能。当某个 RPC 提供商出现问题时，系统会自动切换到其他可用的提供商。

## 为什么需要统一的 ClientManager？

### 原来的问题
最初设计了两种管理器：
- **ClientManager**: 管理 `ethclient.Client` 实例，用于合约调用
- **HTTPClientManager**: 管理原始 HTTP JSON-RPC 调用，用于简单查询

### 优化后的方案
现在只使用一个 **ClientManager**，它内置了 HTTP 功能：
- 统一管理所有 RPC 连接
- 减少代码复杂性
- 保持相同的故障转移能力

## 配置方法

### 1. 基本配置（.env 文件）

```bash
# 单个 RPC URL（向后兼容）
L1_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/YOUR_KEY

# 多个 RPC URLs（新功能）
L1_RPC_URLS=https://eth-sepolia.g.alchemy.com/v2/KEY1,https://ethereum-sepolia.blockpi.network/v1/rpc/public,https://sepolia.infura.io/v3/KEY2

L2_RPC_URL=https://opt-sepolia.g.alchemy.com/v2/YOUR_KEY
L2_RPC_URLS=https://opt-sepolia.g.alchemy.com/v2/KEY1,https://optimism-sepolia.blockpi.network/v1/rpc/public

# 重试配置
RPC_RETRY_DELAY=1        # 重试延迟（秒）
RPC_MAX_RETRIES=3        # 最大重试次数
```

### 2. 推荐的 RPC 提供商组合

```bash
# L1 以太坊 Sepolia 测试网
L1_RPC_URLS=https://eth-sepolia.g.alchemy.com/v2/YOUR_ALCHEMY_KEY,https://ethereum-sepolia.blockpi.network/v1/rpc/public,https://sepolia.infura.io/v3/YOUR_INFURA_KEY,https://eth-sepolia.public.blastapi.io

# L2 Optimism Sepolia 测试网  
L2_RPC_URLS=https://opt-sepolia.g.alchemy.com/v2/YOUR_ALCHEMY_KEY,https://optimism-sepolia.blockpi.network/v1/rpc/public,https://optimism-sepolia.infura.io/v3/YOUR_INFURA_KEY
```

## 使用场景

### 1. ethclient 调用（合约交互）
```go
// 自动使用多 RPC 故障转移
client := ctx.L1RPCManager.GetClient()
if client != nil {
    result, err := client.BlockNumber(context.Background())
}

// 或使用重试机制
err := ctx.L1RPCManager.ExecuteWithRetry(func(client *ethclient.Client) error {
    result, err := client.BlockNumber(context.Background())
    return err
})
```

### 2. HTTP JSON-RPC 调用（简单查询）
```go
// 通过 ClientManager 进行 HTTP 调用
blockJSON, err := ctx.L1RPCManager.HTTPPostJSONWithFailover("", 
    `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x1", true],"id":1}`)
```

## 监控和状态

### 查看 RPC 状态
访问 API 端点：
```
GET http://localhost:8080/disputegames/rpc-status
```

返回示例：
```json
{
  "l1_rpc": {
    "primary_url": "https://eth-sepolia.g.alchemy.com/v2/...",
    "manager_status": {
      "https://eth-sepolia.g.alchemy.com/v2/...": true,
      "https://ethereum-sepolia.blockpi.network/v1/rpc/public": true,
      "https://sepolia.infura.io/v3/...": false
    },
    "healthy_count": 2
  },
  "l2_rpc": {
    "primary_url": "https://opt-sepolia.g.alchemy.com/v2/...",
    "manager_status": {
      "https://opt-sepolia.g.alchemy.com/v2/...": true,
      "https://optimism-sepolia.blockpi.network/v1/rpc/public": true
    },
    "healthy_count": 2
  }
}
```

## 特性

### ✅ 自动故障转移
- 当 RPC 调用失败时，自动切换到下一个健康的端点
- 支持多次重试，可配置重试次数和延迟

### ✅ 健康检查
- 每 30 秒自动检查所有 RPC 端点
- 自动恢复已修复的端点
- 实时更新端点健康状态

### ✅ 负载均衡
- 轮询方式分发请求
- 随机选择可用端点
- 避免单点过载

### ✅ 混合提供商策略
- 付费提供商（Alchemy, Infura）：稳定性高，速度快
- 免费提供商（BlockPI, 公共端点）：作为备份
- 降低整体成本，提高可用性

## 最佳实践

1. **优先级排序**: 将最稳定的 RPC 放在前面
2. **混合使用**: 结合付费和免费提供商
3. **监控告警**: 定期检查 `/disputegames/rpc-status` 端点
4. **适当重试**: 根据网络情况调整 `RPC_MAX_RETRIES`
5. **日志分析**: 关注日志中的故障转移信息

## 向后兼容

如果只配置了单个 RPC URL，系统会正常工作：
- `L1_RPC_URL` 和 `L2_RPC_URL` 仍然有效
- 不配置 `L1_RPC_URLS` 时会回退到单 URL 模式
- 现有代码无需修改
