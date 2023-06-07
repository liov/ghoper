
FROM scratch

# 将工作目录设置为 /hoper
WORKDIR /ghoper

# 使端口 8000 可供此容器外的环境使用
EXPOSE 8000

ADD ghoper /ghoper

# 在容器启动时运行
CMD ["./ghoper","-conf", "ghoper.toml"]
