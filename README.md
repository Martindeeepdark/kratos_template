# kratos-template

基于 [Kratos](https://github.com/go-kratos/kratos) 的 DDD 风格微服务模板，集成 [go-common](https://github.com/Martindeeepdark/go-common) 工具包。

## 架构

```
cmd/                        # 入口 & 依赖注入
  main.go                   # urfave/cli 总入口
  greeter/command.go        # greeter 服务子命令
configs/
  config.yaml               # 应用配置（Go struct，非 proto）
api/                        # Proto 定义的 API（生成代码）
internal/
  domain/
    greeter/
      entity/               # 领域实体 & Value Object
      repository/           # 仓储接口（interface）
      service/              # 领域服务
  application/
    greeter/
      service.go            # 应用服务（用例编排）
      dto.go                # DTO
  server/
    http.go                 # Kratos HTTP server
    grpc.go                 # Kratos gRPC server
  service/
    greeter.go              # Proto 接口实现 → 调用 application
  conf/
    config.go               # 配置结构体
infrastructure/
  data.go                   # gorm 初始化
  greeter_repo_impl.go      # 仓储实现
  model/                    # GORM model
```

依赖方向只有一个：**外 → 内**

```
service → application → domain ← infrastructure
```

`infrastructure` 实现 `domain/repository` 定义的接口，domain 层不感知任何基础设施细节。

## 集成

| 模块 | 用途 |
|------|------|
| go-common/errorx | 集中式错误码管理 |
| go-common/logs | zap 结构化日志 |
| go-kratos/kratos | HTTP + gRPC transport |
| gorm | MySQL ORM |
| urfave/cli | 多服务 CLI 入口 |

## 快速开始

### 环境要求

- Go 1.22+
- MySQL
- protoc（修改 proto 时需要）

### 安装

```bash
git clone https://github.com/Martindeeepdark/kratos_template.git
cd kratos_template
go mod tidy
```

### 配置

编辑 `configs/config.yaml`，修改数据库连接信息：

```yaml
data:
  database:
    driver: mysql
    source: root:your_password@tcp(127.0.0.1:3306)/your_db?parseTime=True&loc=Local
```

### 运行

```bash
# 编译并运行
make run

# 或手动
go build -o ./bin/kratos-template ./cmd
./bin/kratos-template greeter-service --config configs/config.yaml
```

启动后：
- HTTP: `http://localhost:8000/helloworld/your-name`
- gRPC: `localhost:9000`

### Proto 生成

```bash
# 安装 protoc 工具链
make init

# 重新生成 API 代码
make api
```

## 新增一个业务模块

以添加 `order` 模块为例：

### 1. 定义 Proto API

```bash
# 创建 proto 文件
mkdir -p api/order/v1
```

编写 `api/order/v1/order.proto`，然后生成代码：

```bash
make api
```

### 2. 创建 Domain 层

```
internal/domain/order/
├── entity/order.go           # 领域实体
├── repository/order_repo.go  # 仓储接口
└── service/order.go          # 领域服务
```

**仓储接口**（在 domain 层定义）：

```go
package repository

import "context"

type OrderRepository interface {
    Save(ctx context.Context, order *entity.Order) (*entity.Order, error)
    FindByID(ctx context.Context, id int64) (*entity.Order, error)
}
```

**领域服务**（业务规则放在这里）：

```go
package service

type OrderService struct {
    repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
    return &OrderService{repo: repo}
}
```

### 3. 创建 Application 层

```
internal/application/order/
├── service.go   # 应用服务（编排用例）
└── dto.go       # DTO（请求/响应转换）
```

### 4. 创建 Infrastructure 层

```
infrastructure/
├── model/order.go           # GORM model
└── order_repo_impl.go       # 仓储实现
```

### 5. 实现 Service 层

```go
// internal/service/order.go
type OrderService struct {
    v1.UnimplementedOrderServer
    app *application.AppService
}
```

在这里做 proto 类型 ↔ DTO 的转换，把 `errorx` 错误转成 Kratos 错误。

### 6. 注册 Server

在 `internal/server/http.go` 和 `grpc.go` 中注册新的 service。

### 7. 添加 CLI 子命令

```
cmd/order/command.go
```

手动组装依赖链：`repo → domain service → app service → proto service → server`。

## 错误处理

在 domain 层注册错误码，通过 `errorx` 传递到 service 层自动转换为 HTTP/gRPC 错误：

```go
// domain 层
errorx.Register(10001, "order not found")
return nil, errorx.New(10001, errorx.KV("order_id", "123"))

// service 层自动桥接
func toKratosError(err error) error {
    if se, ok := err.(errorx.StatusError); ok {
        return errors.New(int(se.Code()), "DOMAIN_ERROR", se.Msg())
    }
    return errors.InternalServer("INTERNAL_ERROR", err.Error())
}
```

## Docker

```bash
docker build -t kratos-template .
docker run --rm -p 8000:8000 -p 9000:9000 -v $(pwd)/configs:/data/conf kratos-template
```
