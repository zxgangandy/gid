# Gid
Gid是一个用golang开发的，
基于[Snowflake](https://github.com/twitter/snowflake) 分布式id生成器工具。

Gid以组件形式工作在应用项目中, 支持自定义workerId位数和初始化策略, 从而适用于docker等虚拟化环境下实例自动重启、漂移等场景。

## Snowflake

**Snowflake algorithm：** 
指定机器 & 同一时刻 & 某一并发序列，是唯一的。据此可生成一个64 bits的唯一ID（long）。

+------+----------------------+----------------+-----------+
| sign |     delta seconds    | worker node id | sequence  |
+------+----------------------+----------------+-----------+
  1bit          30bits              7bits         13bits
  
sign(1bit)
固定1bit符号标识，即生成的UID为正数。

delta seconds (30 bits)
当前时间，相对于时间基点"2016-05-20"的增量值，单位：秒，最多可支持约34年

worker id (20 bits)
机器id，最多可支持约104w次机器启动。内置实现为在启动时由数据库分配，重启还可复用相同ip和端口号的worker ID。

sequence (13 bits)
每秒下的并发序列，13 bits可支持每秒8192个并发。

## 依赖
- gorm


## 功能
- light and easy to use 
- distributed id generator
- worker id persistence solution (in database instead of cache storage)
- support clock moved backwards(can be disabled)
- support id length customer lower than 64 bits


## 设计
- 参看百度 [uid-generator](https://github.com/baidu/uid-generator)


## 快速开始

### Step1: 安装 golang, Mysql

### Step2: 创建 worker_node表

```sql
DROP TABLE IF EXISTS `worker_node`;
CREATE TABLE `worker_node` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment id',
  `host_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'host name',
  `port` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'port',
  `type` int NOT NULL COMMENT 'node type: CONTAINER(1), ACTUAL(2), FAKE(3)',
  `launch_date` date NOT NULL COMMENT 'launch date',
  `modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'modified time',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQ_IDX_HOST_PORT` (`host_name`,`port`) USING BTREE COMMENT 'host和端口的唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='DB WorkerID Assigner for UID Generator';

```

### Step3: 安装库

go get -u github.com/zxgangandy/gid 

### Step4: 使用

```golang

db := GetDB() //Your grom db
port := GetPort() //Your app port
c := config.New(db, port)
id := gid.New(c).GetUID() //Generate ID

```

## 定制化

通过修改DefaultUidConfig的TimeBits, WorkerBits, SeqBits等字段可以定制化你自己的uid, 特别是长度更短的uid.

```golang

db := GetDB() //Your grom db
port := GetPort() //Your app port
c := config.New(db, port)
c.WorkerBits = 8 // set new WorkerBits value
c.SeqBits = 15 // set new SeqBits value
id := gid.New(c).GetUID() //Generate ID

```

## ChangeLog


## License
Gid 的 [MIT licensed](./LICENSE).
