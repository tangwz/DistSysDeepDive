# 《深入理解分布式系统》代码仓库和勘误

[实验: 使用 Go 语言实现 Paxos 共识算法](https://github.com/tangwz/DistSysDeepDive/tree/master/paxos)

# 已知勘误
* 60 页，

表 3-4
| 事务A  | 事务B   |
|--------|---------|
|        | Read(x) |
| x=1    |         |
| commit |         |
|        | Read(x) |
|        | commit  |

* 61 页，

表 3-5
| 事务A  | 事务B   |
|--------|---------|
|        | 查询用户总数为10（SELECT count(*) FROM users）|
| 插入一条新用户数据（INSERT INTO users VALUES ('Bob')）    |         |
| commit |         |
|        | 查询用户总数为11（SELECT count(*) FROM users） |
|        | commit  |

* 164 页“只有 <1，1> 和 ~~<2，2>~~ <2, 1> 可以安全地应用到状态机。”
* 164 页“因为它当选可能导致 ~~<2, 2>~~ <3, 2>、<4, 2> 和 <5, 2> 被删除，”
* 168 页“Raft 除了领导者选举和日志复制， 配置变更和 ~~领导者选举~~ 状态机部分都非常清晰，”
* 193 页三处“ ~~sage~~ saga 事务”
* 221 页“记录接收包的时间戳 ~~t4~~ t3，”
