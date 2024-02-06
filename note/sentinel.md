# Sentinel - 限流 熔断 降级

## 1. 对比

|        |      Sentinel      |   Hystrix   |
|:------:|:------------------:|:-----------:|
|  规则配置  |       支持多数据源       |   支持多数据源    |
|   限流   | 基于QPS，可以基于调用关系来做限流 |    基本支持     |
|  实时指标  |        滑动窗口        |    滑动窗口     |
|  隔离策略  |       信号量隔离        | 信号量隔离/线程池隔离 |
|  流量整形  |    支持慢启动、匀速器模式     |     不支持     |
| 系统负载保护 |         支持         |     不支持     |

## 2. 限流 

> https://sentinelguard.io/zh-cn/docs/golang/flow-control.html

## 3. 熔断 降级

> https://sentinelguard.io/zh-cn/docs/golang/circuit-breaking.html
