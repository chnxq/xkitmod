# Algs - 算法模块

Algs 是 XKitMod 框架中的算法模块，提供了微服务架构中常用的算法实现，包括熔断、限流、热点检测等功能。

## 概述

该模块包含了一系列用于微服务治理的核心算法，帮助开发者构建稳定、高效的分布式系统。

## 功能特性

- **熔断器 (Circuit Breaker)**：防止级联故障，提高系统稳定性
- **限流器 (Rate Limiter)**：控制请求流量，保护后端服务
- **热点键检测 (Hot Key Detection)**：识别和缓存热点数据
- **Top-K 算法**：统计最频繁出现的元素
- **子集选择算法**：智能选择服务节点子集

## 子模块

### circuitbreaker (熔断器)
实现了熔断机制，包含 SRE 熔断算法，能够自动检测服务异常并在适当时候开启熔断，避免故障扩散。

### ratelimit (限流器)
提供多种限流算法实现，包括 BBR 限流算法，可以根据系统负载情况动态调整限流阈值。

### hotkey (热点键检测)
实时检测热点键并将其缓存在本地，减少对后端存储的压力。

### topk (Top-K 算法)
实现 HeavyKeeper 算法，用于统计 Top-K 频次的数据项。

### subset (子集选择)
实现智能子集选择算法，用于在大规模集群中选择合适的服务节点子集。

## 使用示例

### 熔断器使用示例

```go
import "github.com/chnxq/xkitmod/algs/circuitbreaker"

// 创建熔断器组
group := circuitbreaker.NewGroup()

// 获取特定键的熔断器
cb := group.Get("my_key")

// 在请求前检查是否允许
if err := cb.Allow(); err != nil {
    // 熔断开启，直接返回错误或降级处理
    return err
}

// 执行业务逻辑
result, err := doSomething()

// 标记结果
if err != nil {
    cb.MarkFailed() // 标记失败
} else {
    cb.MarkSuccess() // 标记成功
}

return result, err
```

### 限流器使用示例

```go
import "github.com/chnxq/xkitmod/algs/ratelimit"

// 创建 BBR 限流器
limiter := ratelimit.NewBBR()

// 检查是否允许请求
if limiter.Allow() {
    // 执行业务逻辑
    doSomething()
} else {
    // 限流，执行降级逻辑
    handleFallback()
}
```

## 设计原则

1. **高性能**：算法实现经过优化，确保在高并发场景下的性能
2. **自适应**：算法能够根据系统状态自动调整参数
3. **易用性**：提供简洁的 API 接口
4. **可扩展**：支持自定义算法实现

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进 Algs 模块。