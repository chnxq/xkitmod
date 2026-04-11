# XKitMod - 微服务工具包

XKitMod 是一个 Go 微服务开发工具包，提供了微服务架构中常用的各种组件和功能模块。

## 概述

XKitMod 是一个模块化的微服务开发框架，旨在提供一套完整的微服务解决方案。它包含了配置管理、日志记录、负载均衡、熔断器、限流等微服务架构中的核心组件。

## 特性

- **配置管理**：支持多种配置源（环境变量、文件等），提供动态配置更新功能
- **日志系统**：结构化日志记录，支持多种日志级别
- **服务发现与负载均衡**：提供多种负载均衡策略（随机、轮询、P2C 等）
- **熔断器**：实现熔断机制，提高系统容错能力
- **限流器**：支持多种限流算法，保护系统稳定性
- **编码支持**：支持 JSON、YAML、XML、Protocol Buffers 等多种数据格式
- **错误处理**：统一的错误处理机制
- **元数据管理**：灵活的元数据管理功能

## 模块介绍

### algs (算法模块)
- **circuitbreaker**：熔断器实现，包含 SRE 熔断算法
- **ratelimit**：限流器实现，包含 BBR 限流算法
- **hotkey**：热点键检测与缓存
- **topk**：Top-K 算法实现（HeavyKeeper）
- **subset**：子集选择算法

### config (配置管理)
- 支持多种配置源
- 动态配置更新
- 配置解析与验证
- 多种格式支持（JSON、YAML、Protobuf 等）

### encoding (数据编码)
提供多种数据格式的编解码功能，支持微服务间不同类型数据的序列化和反序列化：
- **JSON 编码/解码**：标准 JSON 格式支持
- **YAML 编码/解码**：YAML 格式支持，适合配置文件
- **XML 编码/解码**：XML 格式支持
- **Protocol Buffers 支持**：高效的二进制序列化格式
- **Form 表单编码**：URL 编码和表单数据处理
- **注册机制**：支持自定义编解码器的动态注册和获取

### errors (错误处理)
提供统一的错误处理机制，支持错误链和详细错误信息管理：
- **统一错误类型定义**：标准化的错误结构，包含错误码、原因和消息
- **gRPC 错误映射**：与 gRPC 状态码的无缝转换
- **错误包装与追踪**：支持错误链和根本原因追踪
- **错误详情**：支持附加错误详情信息
- **兼容性**：与 Go 1.13+ 错误链兼容

### log (日志系统)
提供结构化日志记录功能，支持多种日志级别和灵活的日志格式：
- 结构化日志
- 多级日志支持
- 日志过滤与格式化

### selector (服务选择器)
提供多种负载均衡算法和服务节点选择策略：
- 负载均衡算法（随机、轮询、P2C、EWMA 等）
- 节点健康检查
- 服务发现集成
- 过滤器支持

### metadata (元数据)
提供请求/响应元数据管理功能，支持跨服务的上下文传递：
- **键值对存储**：支持多值键的元数据存储
- **上下文传递**：在服务调用链中传递元数据
- **大小写处理**：自动处理键名的大小写转换
- **灵活操作**：支持添加、获取、设置和删除元数据项
- **高效传输**：轻量级的元数据传输机制

## 安装

```bash
go mod tidy
```

## 使用示例

### 配置管理使用示例

```go
import "github.com/chnxq/xkitmod/config"
import "github.com/chnxq/xkitmod/config/file"

// 创建配置实例
cfg := config.New(
    config.WithSource(file.NewSource("./config.yaml")),
)

// 加载配置
if err := cfg.Load(); err != nil {
    panic(err)
}

// 使用配置
var conf AppConf
if err := cfg.Scan(&conf); err != nil {
    panic(err)
}
```

### 日志使用示例

```go
import "github.com/chnxq/xkitmod/log"

logger := log.DefaultLogger
_ = logger.Log(log.InfoLevel, "msg", "application start")
```

### 负载均衡使用示例

```go
import "github.com/chnxq/xkitmod/selector"

// 构建选择器
sel := selector.NewSelector(selector.WithFilter(filter.Version))

// 添加节点
nodes := []selector.Node{/* ... */}
sel.Apply(nodes)

// 选择节点
selected, done, err := sel.Select(context.Background())
if err != nil {
    // 处理错误
}
defer done(nil) // 标记请求完成
```

## 架构设计

XKitMod 遵循以下设计原则：

1. **模块化**：各功能模块相互独立，可单独使用
2. **接口抽象**：通过接口定义，便于扩展和替换
3. **性能优先**：在保证功能的同时注重性能优化
4. **易用性**：提供简洁的 API 接口

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进 XKitMod。

## 许可证

本项目采用 [LICENSE 文件中指定的许可证] - 请参阅 LICENSE 文件获取详细信息。