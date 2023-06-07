FROM frolvlad/alpine-glibc

# 将工作目录设置为 /hoper
WORKDIR /app

ADD ghoper /app

# 在容器启动时运行
CMD ["./ghoper","-conf", "./config/ghoper.toml"]
