# RPC速率限制配置建议

## 当前状态
✅ **已完成的迁移**
- `syncBlock.go` - 区块同步调用已迁移到RPC Manager
- `latestBlockNumber.go` - 最新区块号获取已迁移到RPC Manager  
- `logFilter.go` - 日志过滤和区块查询已迁移到RPC Manager
- `handler.go` - 已启用RPC监控

## 配置建议

### 开发环境配置
```bash
# 开发环境 - 保守的速率限制
export RPC_RATE_LIMIT=5     # 每秒5个请求
export RPC_RATE_BURST=2     # 允许突发2个请求
```

### 测试环境配置  
```bash
# 测试环境 - 中等速率限制
export RPC_RATE_LIMIT=10    # 每秒10个请求
export RPC_RATE_BURST=5     # 允许突发5个请求
```

### 生产环境配置
```bash
# 生产环境 - 根据RPC提供商限制调整
export RPC_RATE_LIMIT=25    # 每秒25个请求
export RPC_RATE_BURST=10    # 允许突发10个请求
```

## 监控和告警

### 查看RPC统计
启动应用后，你会在日志中看到类似这样的监控信息：
```
[RPC Stats] L1: 150 requests (0 limited), L2: 45 requests (0 limited), HTTP: 95
[RPC Limits] L1: 25.0/s (burst 10, tokens 8.50), L2: 25.0/s (burst 10, tokens 9.20)
```

### 关键指标说明
- **requests** - 总请求数
- **limited** - 被速率限制拒绝的请求数
- **tokens** - 当前可用的令牌数（越低表示使用越频繁）

### 告警条件
- ⚠️ `limited > 0` - 有请求被限制，需要关注
- 🚨 `tokens < 1.0` - 令牌即将耗尽，需要立即调整
- 📊 `requests增长过快` - 可能需要优化代码逻辑

## 优化建议

### 1. 根据使用模式调整限制
```bash
# 如果经常看到 "limited" 请求，可以适当提高限制
export RPC_RATE_LIMIT=30
export RPC_RATE_BURST=15

# 如果tokens经常很低，可以降低请求频率或提高限制
```

### 2. 分时段配置
```bash
# 可以在应用中实现动态调整
# 高峰期降低限制
ctx.RpcManager.UpdateRateLimit(15, 5, true)   # L1

# 低峰期提高限制  
ctx.RpcManager.UpdateRateLimit(35, 15, true)  # L1
```

### 3. 错误处理优化
当前迁移的代码已经包含了错误处理，但可以进一步优化：

```go
// 示例：在syncBlock.go中添加重试逻辑
for retries := 0; retries < 3; retries++ {
    blockJSON, err := ctx.RpcManager.HTTPPostJSON(context.Background(), requestBody, true)
    if err != nil {
        if strings.Contains(err.Error(), "rate limit exceeded") {
            // 指数退避
            time.Sleep(time.Duration(1<<retries) * time.Second)
            continue
        }
        return err
    }
    break // 成功，跳出重试循环
}
```

## 下一步行动

### 立即可以做的：
1. **调整配置** - 根据你的RPC提供商限制设置合适的值
2. **观察日志** - 启动应用并观察RPC监控输出
3. **调优限制** - 根据实际使用情况调整速率限制

### 接下来一周：
1. **迁移API层** - 迁移 `dispute_game_handler.go` 中的RPC调用
2. **添加缓存** - 为频繁查询的数据添加缓存层
3. **监控仪表板** - 考虑添加监控指标到现有的监控系统

### 长期优化：
1. **L1/L2分离配置** - 为L1和L2设置不同的速率限制
2. **请求优先级** - 为不同类型的请求设置优先级
3. **智能负载均衡** - 在多个RPC节点间分发请求

## 故障排除

### 常见问题
1. **频繁出现 "rate limit exceeded"**
   - 解决方案：提高 `RPC_RATE_LIMIT` 或 `RPC_RATE_BURST`
   
2. **tokens 经常接近 0**
   - 解决方案：检查请求频率，考虑添加缓存
   
3. **某些操作响应慢**
   - 解决方案：检查是否被速率限制影响，适当调整配置

### 调试命令
```bash
# 查看当前配置
echo "RPC_RATE_LIMIT: $RPC_RATE_LIMIT"
echo "RPC_RATE_BURST: $RPC_RATE_BURST"

# 查看应用日志中的RPC相关信息
tail -f logs/app.log | grep "RPC"
```

## 成果总结

通过这次迁移，你的项目现在具备了：
- ✅ **统一的速率限制** - 所有主要RPC调用都受到保护
- ✅ **实时监控** - 可以观察RPC使用情况和限制状态  
- ✅ **向后兼容** - 现有功能完全保持不变
- ✅ **可配置性** - 可以根据需要动态调整限制
- ✅ **故障预防** - 防止RPC配额耗尽导致的服务中断

这是一个重要的里程碑！🎉
