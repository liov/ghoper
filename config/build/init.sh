#!/usr/bin/env bash
#ps -aux | grep main

#发现并没有8080端口的Tomcat进程。

#使用命令：

#netstat -apn

#kill -9 [PID]
git checkout init.sh
git checkout ../micro/client/hoper
git pull
chmod 777 init.sh
chmod 777 ../micro/client/hoper
nohup  ./../micro/client/hoper &


#git status
#git checkout
