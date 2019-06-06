FROM scratch

# 将工作目录设置为 /hoper
WORKDIR /manage

# 将当前目录内容复制到位于 /hoper 中的容器中

COPY  ./manage /manage/manage

#Run go get -u github.com/gobuffalo/packr/packr

# 使端口 8001 可供此容器外的环境使用
EXPOSE 8001

# 在容器启动时运行 app.py
CMD ["./manage/manage"]
