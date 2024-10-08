### 1. LRU 缓存淘汰策略

- 介绍常用的三种缓存淘汰(失效)算法：FIFO，LFU 和 LRU
- 实现 LRU 缓存淘汰算法，**代码约80行**

### 2. 单机并发缓存

- 介绍 sync.Mutex 互斥锁的使用，并实现 LRU 缓存的并发控制。
- 实现 GeeCache 核心数据结构 Group，缓存不存在时，调用回调函数获取源数据，代码约150行

### 3. HTTP 服务端

- 介绍如何使用 Go 语言标准库 http 搭建 HTTP Server
- 并实现 main 函数启动 HTTP Server 测试 API，代码约60行

### 4. 一致性哈希

- 一致性哈希(consistent hashing)的原理以及为什么要使用一致性哈希。
- 实现一致性哈希代码，添加相应的测试用例，代码约60行

### 5. 分布式节点

- 注册节点(Register Peers)，借助一致性哈希算法选择节点。
- 实现 HTTP 客户端，与远程节点的服务端通信，代码约90行

### 6. 防止缓存击穿

- 缓存雪崩、缓存击穿与缓存穿透的概念简介。
- 使用 singleflight 防止缓存击穿，实现与测试。代码约70行

### 7. 使用 Protobuf 通信(网络原因未实现)

- 为什么要使用 protobuf？
- 使用 protobuf 进行节点间通信，编码报文，提高效率。代码约50行

