[In Chinese 中文版](README.zh_cn.md)
# Gid
 [![Go Report Card](https://goreportcard.com/badge/github.com/zxgangandy/gid)](https://goreportcard.com/report/github.com/zxgangandy/gid)
 [![gitmoji](https://img.shields.io/badge/gitmoji-%20%F0%9F%98%9C%20%F0%9F%98%8D-FFDD67.svg?style=flat-square)](https://github.com/carloscuesta/gitmoji)
 [![GoDoc](https://godoc.org/github.com/zxgangandy/gid?status.svg)](http://godoc.org/github.com/zxgangandy/gid)
 [![License](https://img.shields.io/github/license/zxgangandy/gid?style=flat-square)](/LICENSE)
 
Gid is a distributed id generator tool implements by golang,
[Snowflake](https://github.com/twitter/snowflake) based unique ID generator. It
works as a component, and allows users to override workId bits and initialization strategy. As a result, it is much more
suitable for virtualization environment, such as [docker](https://www.docker.com/).

## Snowflake

**Snowflake algorithm：** 
An unique id consists of worker node, timestamp and sequence within that timestamp. Usually, it is a 64 bits number(long), and the default bits of that three fields are as follows:
```xml
+------+----------------------+----------------+-----------+
| sign |     delta seconds    | worker node id | sequence  |
+------+----------------------+----------------+-----------+
  1bit          30bits              20bits         13bits
```


sign(1bit)
The highest bit is always 0.

delta seconds (30 bits)
The next 30 bits, represents delta seconds since a customer epoch(2016-05-20). The maximum time will be 34 years.

worker id (20 bits)
The next 20 bits, represents the worker node id, maximum value will be 1.04 million. UidGenerator uses a build-in database based worker id assigner when startup by default, and it will reuse previous work node id after reboot.

sequence (13 bits)
the last 13 bits, represents sequence within the one second, maximum is 8192 per second（per server）by default.

## Dependency
- gorm


## Features
- light and easy to use 
- distributed id generator at local instead of by service or rpc
- worker id persistence solution (in database like mysql instead of cache storage)
- support clock moved backwards(can be disabled)
- support id length customization lower than 64 bits


## Design
- refer to baidu [uid-generator](https://github.com/baidu/uid-generator)


## Quick  Start

### Step1: Install golang, Mysql

### Step2: Create table worker_node

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
  UNIQUE KEY `UNIQ_IDX_HOST_PORT` (`host_name`,`port`) USING BTREE COMMENT 'unique index'
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='DB WorkerID Assigner for UID Generator';

```

### Step3: Install Lib

go get -u github.com/zxgangandy/gid 

### Step4: Usage

```golang

db := GetDB() //Your grom db
port := GetPort() //Your app port
c := config.New(db, port)
id := gid.New(c).GetUID() //Generate ID

```

## Customization

Change the TimeBits, WorkerBits, SeqBits of 'DefaultUidConfig' to get your customer uid, especially shorter uid.

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
Gid is [MIT licensed](./LICENSE).
