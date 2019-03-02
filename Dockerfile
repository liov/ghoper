# 将官方 Python 运行时用作父镜像
FROM scratch

# 将工作目录设置为 /app
WORKDIR /hoper

# 将当前目录内容复制到位于 /app 中的容器中
COPY  ./config/config.toml /hoper/config/
COPY  ./micro/hoper /hoper/ghoper/

# 安装 requirements.txt 中指定的任何所需软件包
#RUN pip install -r requirements.txt

# 使端口 80 可供此容器外的环境使用
EXPOSE 8000

# 定义环境变量
ENV NAME World

# 在容器启动时运行 app.py
CMD ["./hoper/ghoper/hoper"]