# 使用官方的 Go 作为基础镜像
FROM golang:1.19-alpine

# 设置工作目录
WORKDIR /app

# 将当前目录内容复制到 /app 目录中
COPY . .

# 下载依赖
RUN go mod tidy

# 编译 Go 程序
RUN go build -o web .

# 运行编译好的二进制文件
CMD ["./web"]
