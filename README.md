# seata-golang

### 一个朴素的想法
作为一个刚入 Golang 坑的普通微服务开发者来讲，很容易产生一个朴素的想法，希望 Golang 微服务也有分布式事务解决方案。我们注意到阿里开源了 Java 版的分布式事务解决方案 Seata，本项目尝试将 Java 版的 Seata 改写一个 Golang 的版本。
在 Seata 没有 Golang 版本 client sdk 的情况下，Golang 版本的 TC Server 使用了和 Java 版 Seata 一样的通信协议，方便调试。
目前，基于内存的 TC Server 已实现。希望有同样朴素想法的开发者加入我们一起完善 Golang 版本的分布式事务解决方案。

### todo list
- [X] Memory Session Manager
- [ ] DB Session Manager  
- [ ] ETCD Session Manager  
- [ ] TC
- [ ] RM