package paxos

import "fmt"

type Proposer struct {
	// 服务器 id
	id int
	// 当前提议者已知的最大轮次
	round int
	// 提案编号 = (轮次, 服务器 id)
	number int
	// 接受者 id 列表
	acceptors []int
}

func (p *Proposer) propose(v interface{}) interface{} {
	p.round++
	p.number = p.proposalNumber()

	// 第一阶段(phase 1)
	prepareCount := 0
	maxNumber := 0
	for _, aid := range p.acceptors {
		args := MsgArgs{
			Number: p.number,
			From:   p.id,
			To:     aid,
		}
		reply := new(MsgReply)
		err := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Prepare", args, reply)
		if !err {
			continue
		}
		
		if reply.Ok {
			prepareCount++
			if reply.Number > maxNumber {
				maxNumber = reply.Number
				v = reply.Value
			}
		}
		
		if prepareCount == p.majority() {
			break
		}
	}

	// 第二阶段(phase 2)
	acceptCount := 0
	if prepareCount >= p.majority() {
		for _, aid := range p.acceptors {
			args := MsgArgs{
				Number: p.number,
				Value: v,
				From: p.id,
				To: aid,
			}
			reply := new(MsgReply)
			ok := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Accept", args, reply)
			if !ok {
				continue
			}
			
			if reply.Ok {
				acceptCount++
			}
		}
	}
	
	if acceptCount >= p.majority() {
		// 选择的提案的值
		return v
	}
	return nil
}

func (p *Proposer) majority() int {
	return len(p.acceptors) / 2 + 1
}

// 提案编号 = (轮次, 服务器 id)
func (p *Proposer) proposalNumber() int {
	return p.round << 16 | p.id
}
