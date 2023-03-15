# Yet Another ChatGPT Web

[English](README.md)

## 初始化

### 安装 Commit 前检查工具

#### 安装 `golangci-lint`

[Install tutorial](https://golangci-lint.run/usage/install/#local-installation)

#### 安装 `pre-commit`

```bash
pip install pre-commit
```

#### 运行 `pre-commit`

```bash
pre-commit install
```

### 安装 Go 依赖

#### 使用 `make`

```bash
make init
```

#### 使用命令

```bash
go mod download
```

## 运行

### 使用 `make`

```bash
make run
```

### 使用命令

```bash
go run main.go
```

## 项目结构

```
├── main.go # 程序入口文件
├── .env # 环境变量文件
├── controllers # 控制器目录
│   ├── user.go # 用户控制器
│   └── product.go # 产品控制器
├── models # 模型目录
│   ├── user.go # 用户模型
│   └── product.go # 产品模型
├── routes # 路由目录
│   └── routes.go # 路由定义文件
├── services # 服务目录
│   ├── user.go # 用户服务
│   └── product.go # 产品服务
├── tests # 测试目录
└── utils # 工具类目录
    ├── logger.go # 日志工具类文件
    └── db.go # 数据库连接工具类文件
```