# Gid
Gid is a distributed id generator tool implements by golang 


## Dependency
- gorm


## Features
- light and easy to use 
- distributed id generator
- worker id persistence solution (in database instead of cache storage)


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
  UNIQUE KEY `UNIQ_IDX_HOST_PORT` (`host_name`,`port`) USING BTREE COMMENT 'host和端口的唯一索引'
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

## ChangeLog


## License
Gid is [MIT licensed](./LICENSE).K
