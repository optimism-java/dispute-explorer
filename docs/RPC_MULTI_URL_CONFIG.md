# 多 RPC URL 配置指南

## 概述

为了解决 RPC URL 流量不够的问题，我们实现了多 RPC URL 轮询和故障转移功能。系统会自动在多个 RPC 提供商之间切换，当某个 RPC 不可用时会自动切换到其他可用的 RPC。

## 配置方式

### 环境变量配置

在 `.env` 文件中添加以下配置：

```bash
# 单个 RPC URL (向后兼容)
L1_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/YOUR_ALCHEMY_KEY
L2_RPC_URL=https://opt-sepolia.g.alchemy.com/v2/YOUR_ALCHEMY_KEY

# 多个 RPC URLs (逗号分隔)
L1_RPC_URLS=https://eth-sepolia.g.alchemy.com/v2/KEY1,https://ethereum-sepolia.blockpi.network/v1/rpc/public,https://sepolia.infura.io/v3/KEY2
L2_RPC_URLS=https://opt-sepolia.g.alchemy.com/v2/KEY1,https://optimism-sepolia.blockpi.network/v1/rpc/public,https://optimism-sepolia.infura.io/v3/KEY2

# RPC 重试配置
RPC_RETRY_DELAY=1      # 重试延迟（秒）
RPC_MAX_RETRIES=3      # 最大重试次数
```

### 推荐的 RPC 提供商

#### Ethereum Sepolia (L1)
- Alchemy: `https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY`
- Infura: `https://sepolia.infura.io/v3/YOUR_PROJECT_ID`
- BlockPI: `https://ethereum-sepolia.blockpi.network/v1/rpc/public`
- Ankr: `https://rpc.ankr.com/eth_sepolia`
- QuickNode: `https://YOUR_ENDPOINT.ethereum-sepolia.quiknode.pro/YOUR_API_KEY`

#### Optimism Sepolia (L2)
- Alchemy: `https://opt-sepolia.g.alchemy.com/v2/YOUR_API_KEY`
- Infura: `https://optimism-sepolia.infura.io/v3/YOUR_PROJECT_ID`
- BlockPI: `https://optimism-sepolia.blockpi.network/v1/rpc/public`
- Ankr: `https://rpc.ankr.com/optimism_sepolia`

#### Ethereum Mainnet (生产环境)
- Alchemy: `https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY`
- Infura: `https://mainnet.infura.io/v3/YOUR_PROJECT_ID`
- BlockPI: `https://ethereum.blockpi.network/v1/rpc/public`
- Ankr: `https://rpc.ankr.com/eth`

#### Optimism Mainnet (生产环境)
- Alchemy: `https://opt-mainnet.g.alchemy.com/v2/YOUR_API_KEY`
- Infura: `https://optimism-mainnet.infura.io/v3/YOUR_PROJECT_ID`
- BlockPI: `https://optimism.blockpi.network/v1/rpc/public`
- Ankr: `https://rpc.ankr.com/optimism`

## 功能特性

### 1. 自动故障转移
- 当某个 RPC 端点不可用时，系统会自动切换到下一个可用的端点
- 支持 HTTP 和 WebSocket 连接的故障转移

### 2. 健康检查
- 每 30 秒自动检查所有 RPC 端点的健康状态
- 自动恢复已修复的 RPC 端点

### 3. 负载均衡
- 支持轮询和随机选择策略
- 均匀分布请求到不同的 RPC 提供商

### 4. 重试机制
- 可配置的重试次数和延迟
- 失败后自动尝试其他 RPC 端点

## 监控和状态

### RPC 状态 API

访问 `/disputegames/rpc-status` 查看所有 RPC 端点的状态：

```json
{
  "l1_rpc": {
    "primary_url": "https://eth-sepolia.g.alchemy.com/v2/...",
    "manager_status": {
      "https://eth-sepolia.g.alchemy.com/v2/...": true,
      "https://ethereum-sepolia.blockpi.network/v1/rpc/public": false,
      "https://sepolia.infura.io/v3/...": true
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
  },
  "l1_http": {
    "manager_status": {
      "https://eth-sepolia.g.alchemy.com/v2/...": true,
      "https://ethereum-sepolia.blockpi.network/v1/rpc/public": false
    },
    "healthy_count": 1
  },
  "l2_http": {
    "manager_status": {
      "https://opt-sepolia.g.alchemy.com/v2/...": true
    },
    "healthy_count": 1
  }
}
```

### 日志监控

系统会记录 RPC 切换和健康检查的详细日志：

```
[RPC.ClientManager] Successfully connected to https://eth-sepolia.g.alchemy.com/v2/...
[RPC.ClientManager] Marked RPC client https://ethereum-sepolia.blockpi.network/v1/rpc/public as unhealthy
[RPC.ClientManager] RPC client https://sepolia.infura.io/v3/... is back online
```

## 最佳实践

### 1. 多样化 RPC 提供商
- 使用至少 3 个不同的 RPC 提供商
- 包含免费和付费的提供商
- 选择不同地理位置的提供商

### 2. API 密钥管理
- 为每个提供商使用独立的 API 密钥
- 定期轮换 API 密钥
- 监控 API 配额使用情况

### 3. 配置优化
- 根据网络状况调整重试延迟
- 设置合适的重试次数（推荐 3-5 次）
- 监控系统性能并调整配置

### 4. 成本优化
- 将免费的公共 RPC 作为备选
- 主要使用付费的高性能 RPC
- 监控每个提供商的使用量

## 故障排除

### 常见问题

1. **所有 RPC 都不可用**
   - 检查网络连接
   - 验证 API 密钥是否有效
   - 检查是否达到配额限制

2. **某个 RPC 频繁失败**
   - 检查该提供商的服务状态
   - 验证 URL 格式是否正确
   - 检查是否需要特殊的认证头

3. **性能问题**
   - 调整重试延迟和次数
   - 检查健康检查频率
   - 监控网络延迟

### 调试命令

```bash
# 检查配置
curl http://localhost:8080/disputegames/rpc-status

# 查看日志
docker logs dispute-explorer-backend

# 测试 RPC 连接
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
  https://your-rpc-url
```

## 升级指南

### 从单 RPC 配置升级

1. 保留现有的 `L1_RPC_URL` 和 `L2_RPC_URL` 配置
2. 添加新的 `L1_RPC_URLS` 和 `L2_RPC_URLS` 配置
3. 添加重试配置参数
4. 重启服务

系统会自动检测多 RPC 配置并启用故障转移功能，同时保持向后兼容性。
