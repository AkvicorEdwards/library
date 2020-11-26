# Library

- **Complex** 默认的书本，用于记录偶然遇见的句子
- 可添加其他书本，方便整理
- 为书本添加标签
- 记录开始和完成阅读的时间
- 可为每个记录添加感悟

# Usage

```shell script
go mod download
# 编译
go build
# 初始化数据库
./library init
# 运行（端口8080）
./library
```

## 初始化数据库

数据库使用sqlite3，文件名为`library.db`

`./library init`