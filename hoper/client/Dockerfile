
FROM scratch

# 将工作目录设置为 /hoper
WORKDIR /hoper

RUN mkdir /hoper/ghoper/hoper/bin

# 将当前目录内容复制到位于 /hoper 中的容器中
COPY  ./config/config.toml /hoper/config/
COPY  ./micro/hoper /hoper/ghoper/

# 使端口 8000 可供此容器外的环境使用
EXPOSE 8000

# 定义环境变量
ENV NAME World

# 在容器启动时运行
CMD ["./hoper/ghoper/hoper/bin/hoper"]
