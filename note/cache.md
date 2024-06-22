# 缓存
redis mongodb 雪崩 穿透 击穿 布隆过滤器 多级缓存架构
红锁Redlock

> https://www.bilibili.com/video/BV14x421k7hJ/?spm_id_from=333.999.0.0&vd_source=7c58b2999811337c63c69bda6933f026

> https://www.bilibili.com/video/BV1Fd4y1T7pD/?spm_id_from=333.999.0.0&vd_source=7c58b2999811337c63c69bda6933f026

> https://www.bilibili.com/video/BV1df421S7bZ/?spm_id_from=333.999.0.0&vd_source=7c58b2999811337c63c69bda6933f026

> https://www.bilibili.com/video/BV1KS421K78d/?spm_id_from=333.999.0.0&vd_source=7c58b2999811337c63c69bda6933f026

## 1. 数据类型

+ String：Redis 中最基本的数据类型，可以是任何类型的字符串，如文本或二进制数据（如图片）。
+ Hash：类似于一个具有键值对的字典，特别适合用于存储对象。
+ List：一个按插入顺序排序的字符串列表，可以添加到列表的头部或尾部。
+ Set：一个无序的字符串集合，集合成员是唯一的（不能重复）
+ ZSet（有序集合）：类似于 Set，但每个成员都会关联一个分数，Redis 会根据分数自动排序。
+ Bitmap（位图）：使用字符串存储位值（0 或 1），可以对其进行位操作。
+ HyperLogLog：一种用于基数估算的数据结构，适合用于大数据去重计数，且占用空间非常小。
+ Geo（地理位置）：用于存储地理位置信息并执行地理位置操作。
+ Stream（流）：用于处理消息流的高级数据类型，类似于消息队列。
+ Pub/Sub（发布/订阅）：虽然严格来说不是一种数据类型，但它是一种消息传递模式，允许发送者发布消息，接收者订阅消息。
