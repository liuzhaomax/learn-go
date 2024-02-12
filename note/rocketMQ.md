# RocketMQ


## 1. 概述
### 1.1 用途
+ 限流削峰
+ 异步解耦
+ 数据收集

### 1.2 对比
|    MQ    |    开发语言    | 特点                            |
|:--------:|:----------:|:------------------------------|
| ActiveMQ |    Java    | 早期，活跃度小                       |
| RabbitMQ |   Erlang   | 吞吐量较低                         |
|  Kafka   | Scala/Java | 吞吐率高，日志收集，大数据实时计算，存在硬盘但吞吐率也很高 |
| RocketMQ |    Java    | 性能问题，吞吐率高，自研协议                |

### 1.3 协议
+ JMS, java messaging service，ActiveMQ实现
+ STOMP, streaming text oriented message protocol, ActiveMQ实现，RabbitMQ可通过插件支持
+ AMQP, advanced message queuing protocol, RocketMQ自研
+ MQTT, message queuing telemetry transport, IBM自研即时通信协议，二进制协议，用于服务器和低功耗的IoT设备的通信，RabbitMQ可通过插件支持

### 1.4 基本概念
+ 消息message：传输信息的物理载体，生产和消费数据的最消单位，每条消息必须属于一个主题
+ 主题topic：


## 2. 安装与启动
