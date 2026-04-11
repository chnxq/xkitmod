# Selector - 服务选择器模块

Selector 是 XKitMod 框架中的服务选择器模块，提供了多种负载均衡算法和服务节点选择策略，用于微服务架构中的服务发现和负载均衡。

## 概述

该模块实现了多种负载均衡算法，可以根据不同的业务场景选择合适的策略，同时提供了节点健康检查、服务过滤等功能。

## 功能特性

- **多种负载均衡算法**：支持随机、轮询、P2C、EWMA 等多种算法
- **节点管理**：动态管理服务节点列表
- **健康检查**：内置节点健康检查机制
- **服务过滤**：支持多种过滤策略（版本过滤等）
- **连接感知**：跟踪节点连接状态和负载情况

## 负载均衡算法

### Random (随机算法)
随机选择服务节点，实现简单，负载分布相对均匀。

### WRR (加权轮询算法)
根据节点权重进行轮询调度，权重高的节点被选中的概率更大。

### P2C (Power of Two Choices)
从所有节点中随机选择两个，然后选择其中负载较轻的一个，具有较好的负载均衡效果。

### EWMA (指数加权移动平均)
基于节点历史响应时间进行选择，倾向于选择响应更快的节点。

## 使用示例

### 基本使用

```go
import (
    "context"
    "github.com/chnxq/xkitmod/selector"
    "github.com/chnxq/xkitmod/selector/random"
)

// 创建选择器
builder := random.Builder{}
sel := builder.Build()

// 添加节点
nodes := []selector.Node{
    // 创建节点实例
}
sel.Apply(nodes)

// 选择节点
selected, done, err := sel.Select(context.Background())
if err != nil {
    // 处理错误
    return err
}
defer done(nil) // 请求完成后调用 done 回调

// 使用选中的节点
return callService(selected)
```

### 使用过滤器

```go
import "github.com/chnxq/xkitmod/selector/filter"

// 创建带过滤器的选择器
filters := []selector.Filter{
    filter.Version("v1.0.0"), // 只选择特定版本的节点
}
opts := []selector.Option{
    selector.WithFilter(filters...),
}

// 构建选择器时应用选项
sel := selector.NewSelector(opts...)
```

### 自定义节点

```go
import "github.com/chnxq/xkitmod/selector"

type CustomNode struct {
    addr   string
    weight int64
    version string
}

func (n *CustomNode) Address() string { return n.addr }
func (n *CustomNode) ServiceName() string { return "my-service" }
func (n *CustomNode) InitialWeight() *int64 { return &n.weight }
func (n *CustomNode) Version() string { return n.version }

// 使用自定义节点
node := &CustomNode{
    addr: "127.0.0.1:8080",
    weight: 100,
    version: "v1.0.0",
}
```

## 设计原则

1. **插件化架构**：支持自定义负载均衡算法和过滤器
2. **性能优先**：高效的节点选择算法
3. **容错性**：内置故障节点隔离机制
4. **可扩展性**：易于添加新的负载均衡策略

## 主要组件

- **Selector 接口**：定义了选择器的基本行为
- **Node 接口**：定义了服务节点的基本属性
- **Builder 接口**：用于构建选择器实例
- **Filter 接口**：用于过滤不合适的节点
- **Rebalancer 接口**：用于节点列表变更时的重新平衡

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进 Selector 模块。