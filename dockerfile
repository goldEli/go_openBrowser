# FROM node:12-alpine
# FROM alpine:3.14
# FROM ops-harbor.staryuntech.com/sre/maven-openjdk8:20230717 AS builder
FROM centos:7

# 设置工作目录
WORKDIR /app

# 复制本地文件到容器
COPY . .


# RUN apk add chromium

RUN yum install chromium

