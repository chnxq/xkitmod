# Config - 配置管理模块

Config 是 XKitMod 框架中的配置管理模块，提供了灵活的配置加载、解析和动态更新功能，支持多种配置源和数据格式。

## 概述

该模块提供了一套完整的配置管理解决方案，支持从不同来源加载配置、动态监听配置变化以及多种数据格式解析。

## 功能特性

- **多源配置**：支持从文件、环境变量等多种来源加载配置
- **动态更新**：支持配置热更新，无需重启应用
- **格式支持**：支持 JSON、YAML、XML、Protobuf 等多种数据格式
- **配置合并**：支持多个配置源的智能合并
- **类型转换**：自动将配置值转换为目标类型
- **观察模式**：支持配置变更监听

## 配置源

### File (文件源)
支持从本地文件加载配置，支持多种文件格式（JSON、YAML、XML 等）。

### Env (环境变量源)
支持从环境变量加载配置，便于容器化部署。

## 使用示例

### 基本使用

```go
import (
    "github.com/chnxq/xkitmod/config"
    "github.com/chnxq/xkitmod/config/file"
)

// 创建配置实例
cfg := config.New(
    config.WithSource(
        file.NewSource("./config.yaml"), // 从文件加载
    ),
)

// 加载配置
if err := cfg.Load(); err != nil {
    panic(err)
}

// 使用配置
var appConf AppConf
if err := cfg.Scan(&appConf); err != nil {
    panic(err)
}

// 获取特定键的值
dbHost := cfg.Value("database.host").String()
dbPort := cfg.Value("database.port").Int()
```

### 使用多个配置源

```go
import (
    "github.com/chnxq/xkitmod/config"
    "github.com/chnxq/xkitmod/config/file"
    "github.com/chnxq/xkitmod/config/env"
)

// 同时从文件和环境变量加载配置
cfg := config.New(
    config.WithSource(
        file.NewSource("./config.yaml"),
        env.NewSource(), // 从环境变量加载
    ),
)

if err := cfg.Load(); err != nil {
    panic(err)
}
```

### 监听配置变化

```go
// 监听特定键的变化
err := cfg.Watch("database", func(key string, value config.Value) {
    // 当数据库相关配置发生变化时执行回调
    log.Printf("Config %s changed to %v", key, value)
})
if err != nil {
    panic(err)
}
```

### 高级配置扫描

```go
type ServerConfig struct {
    Name    string `json:"name"`
    Addr    string `json:"addr"`
    Timeout int    `json:"timeout"`
}

// 将配置扫描到结构体
var serverConf ServerConfig
if err := cfg.Scan(&serverConf); err != nil {
    panic(err)
}
```

## 设计原则

1. **接口抽象**：通过接口定义，便于扩展和替换配置源
2. **性能优先**：高效的配置解析和缓存机制
3. **灵活性**：支持多种配置源和格式
4. **类型安全**：提供类型安全的值访问方法

## 主要组件

- **Config 接口**：定义了配置管理的基本方法
- **Value 接口**：定义了配置值的操作方法
- **Source 接口**：定义了配置源的基本行为
- **Reader 接口**：负责配置的读取和解析
- **Watcher 接口**：负责监听配置变化

## 配置格式支持

- **JSON**：标准 JSON 格式
- **YAML**：人类友好的数据序列化格式
- **XML**：可扩展标记语言
- **Protobuf**：Google 的序列化协议
- **Properties**：Java 风格的键值对格式

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进 Config 模块。