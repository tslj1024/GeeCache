### 1. LRU 缓存淘汰策略

- 介绍常用的三种缓存淘汰(失效)算法：FIFO，LFU 和 LRU
- 实现 LRU 缓存淘汰算法，**代码约80行**

### 2. 单机并发缓存

- 介绍 sync.Mutex 互斥锁的使用，并实现 LRU 缓存的并发控制。
- 实现 GeeCache 核心数据结构 Group，缓存不存在时，调用回调函数获取源数据，代码约150行

### 3. HTTP 服务端

- 介绍如何使用 Go 语言标准库 http 搭建 HTTP Server
- 并实现 main 函数启动 HTTP Server 测试 API，代码约60行
