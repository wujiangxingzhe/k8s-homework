# module3

## 1. build httpserver
```
root@master:/opt# cd /opt/go_workspace/k8s-homework-master/
root@master:/opt/go_workspace/k8s-homework-master# go build -o httpserver main.go
```

## 2. Copy httpserver bin to docker build context
```
root@master:/opt/go_workspace/k8s-homework-master# cp httpserver /opt/docker/
```

## 3. Write dockerfile under docker build context
```
root@master:/opt/go_workspace/k8s-homework-master# cd /opt/docker/
root@master:/opt/docker# 
root@master:/opt/docker# tree
.
├── Dockerfile
└── httpserver

0 directories, 2 files
root@master:/opt/docker# cat Dockerfile 
#FROM alpine:3.16
FROM ubuntu:20.04 
MAINTAINER wujiangxingzhe

COPY ./httpserver /bin/httpserver
EXPOSE 8080

ENTRYPOINT ["/bin/httpserver"]
CMD []
```

## 4. Build httpserver image
```
root@master:/opt/docker# docker build -f /opt/docker/Dockerfile /opt/docker -t wujiangxingzhe/httpserver:v0.1
Sending build context to Docker daemon  6.599MB
Step 1/6 : FROM ubuntu:20.04
 ---> ba6acccedd29
Step 2/6 : MAINTAINER wujiangxingzhe
 ---> Running in feeebab61cdc
Removing intermediate container feeebab61cdc
 ---> d73c4a138ea4
Step 3/6 : COPY ./httpserver /bin/httpserver
 ---> 56b8adc88f5f
Step 4/6 : EXPOSE 8080
 ---> Running in 04003d33daf6
Removing intermediate container 04003d33daf6
 ---> 7d98260602d6
Step 5/6 : ENTRYPOINT ["/bin/httpserver"]
 ---> Running in 65e94a2bed73
Removing intermediate container 65e94a2bed73
 ---> 956a834f1e5f
Step 6/6 : CMD []
 ---> Running in 5db9d0e9332b
Removing intermediate container 5db9d0e9332b
 ---> 197874ff39b6
Successfully built 197874ff39b6
Successfully tagged wujiangxingzhe/httpserver:v0.1
```

## 5. Run container and test
```
root@master:/opt/docker# docker run --name httpserver -d wujiangxingzhe/httpserver:v0.1
b1f1feda8b8d621789b33d0c2032d8dab32674bf811470f507c59fe2a5771b64
root@master:/opt/docker# 
root@master:/opt/docker# docker ps | grep httpserver
b1f1feda8b8d   wujiangxingzhe/httpserver:v0.1                      "/bin/httpserver"        10 seconds ago   Up 7 seconds    8080/tcp   httpserver
root@master:/opt/docker# 
root@master:/opt/docker# ps -ef | grep httpserver
root       94005 1721659  0 09:10 pts/1    00:00:00 grep --color=auto httpserver
root     1760568 1712344  0 Oct15 pts/0    00:00:00 /opt/docker/httpserver
root     4173611 4173569  0 07:50 ?        00:00:00 /bin/httpserver
root@master:/opt/docker# 
root@master:/opt/docker# nsenter -t 4173611 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
90: eth0@if91: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
root@master:/opt/docker# 
root@master:/opt/docker# curl -H "Content-Type:application/json" -H "User:wujiangxingzhe" 172.17.0.2:8080 -I
HTTP/1.1 200 OK
Accept: */*
Content-Type: application/json
User: wujiangxingzhe
User-Agent: curl/7.68.0
Date: Sun, 16 Oct 2022 09:11:04 GMT
Content-Length: 2
```

## 6. Push the image to docker hub
```
root@master:/opt/docker# docker login
Authenticating with existing credentials...
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
root@master:/opt/docker# 
root@master:/opt/docker# docker push wujiangxingzhe/httpserver:v0.1
The push refers to repository [docker.io/wujiangxingzhe/httpserver]
1395528c9a08: Pushed 
9f54eef41275: Pushed 
v0.1: digest: sha256:e9ea0cb2bf48f3a07a7e466777643577c9062c54069ed27a5b621e768a4ea719 size: 740
```