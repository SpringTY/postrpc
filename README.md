### Post RPC


#### 总体目录结构
```text
idl: proto3的接口定义文件,分别应该定义四个模块
post_*: rpc各个模块服务Repo
sdk: rpc服务sdk的生成
sh: 更新、生成rpc服务和sdk的脚本
```
#### 依赖

> go 1.17.x \
> grpc \
> proto3 \
> grpcui

#### 运行RPC项目

编译
```bash
cd post_model_manage
sh build.sh
```
生成测试平台
```bash
grpcui -plaintext [host:port]
# 例如
grpcui -plaintext localhost:50051
```
