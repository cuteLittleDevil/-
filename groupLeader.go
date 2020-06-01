package main

import "sync"

type groupLeader struct {
	result chan interface{}
	notice chan struct{}
	once sync.Once
}

func NewGroupLeader() *groupLeader {
	return &groupLeader{
		result: make(chan interface{}),
		notice: make(chan struct{}, 1),
	}
}

func (leader *groupLeader)GetTaskResult() chan interface{} {
	return leader.result
}

func (leader *groupLeader)SetTaskResult(result interface{})  {
	leader.once.Do(func() {
		leader.result <- result
	})
}

func (leader *groupLeader)NoticeTaskOver()  {
	leader.notice <- struct{}{}
	close(leader.notice)
	close(leader.result)
}

func (leader *groupLeader)IsCompleteTask() bool {
	if len(leader.notice) == 1{
		return true
	}else {
		return false
	}
}
