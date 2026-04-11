# Log - 日志模块

Log 是 XKitMod 框架中的日志模块，提供了结构化日志记录功能，支持多种日志级别和灵活的日志格式。

## 概述

该模块提供了一套完整的日志解决方案，包括日志级别管理、结构化输出、全局日志器等功能，适用于微服务架构中的日志记录需求。

## 功能特性

- **结构化日志**：支持键值对形式的日志输出
- **多级日志**：支持 Trace、Debug、Info、Warn、Error、Fatal 等多个日志级别
- **全局日志器**：提供默认的全局日志器实例
- **日志过滤**：支持按级别和内容进行日志过滤
- **上下文日志**：支持在日志中添加上下文信息

## 使用示例

### 基本日志记录

```go
import "github.com/chnxq/xkitmod/log"

// 使用默认日志器
_ = log.DefaultLogger.Log(log.InfoLevel, "msg", "application start", "service", "myservice")

// 或者创建自定义日志器
logger := log.NewStdLogger(os.Stdout)
_ = logger.Log(log.ErrorLevel, "error", "something went wrong", "code", 500)
```

### 带上下文的日志

```go
// 使用 With 函数添加固定上下文
ctxLogger := log.With(log.DefaultLogger, "service", "user-service", "version", "v1.0.0")
_ = ctxLogger.Log(log.InfoLevel, "action", "user_login", "user_id", 12345)

// 使用 WithContext 在日志中包含上下文信息
ctx := context.WithValue(context.Background(), "request_id", "req-12345")
ctxLogger = log.WithContext(ctxLogger, ctx)
_ = ctxLogger.Log(log.InfoLevel, "msg", "processing request")
```

### 日志级别控制

```go
// 设置日志级别过滤器
filter := log.LevelFilter(log.WarnLevel, log.DefaultLogger)
_ = filter.Log(log.InfoLevel, "msg", "this will be filtered out") // 不会被输出
_ = filter.Log(log.ErrorLevel, "msg", "this will be logged")     // 会被输出
```

## 设计原则

1. **接口抽象**：通过 Logger 接口定义，便于扩展和替换
2. **性能优先**：最小化日志记录对应用性能的影响
3. **灵活性**：支持多种日志格式和输出方式
4. **标准化**：遵循结构化日志的最佳实践

## 主要组件

- **Logger 接口**：定义了基本的日志记录方法
- **Level 类型**：定义了日志级别常量
- **DefaultLogger**：默认的标准输出日志器
- **With 函数**：用于添加日志上下文
- **全局函数**：提供全局日志器访问

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进 Log 模块。