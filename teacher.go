package main

import (
	"errors"
	"sync"
	"time"
)

var once sync.Once
var instance *teacher

// 老师布置任务
type teacher struct {}

func GetTeacherInstance() *teacher {
	once.Do(func() {
		instance = &teacher{}
	})
	return instance
}

// 小组长沟通任务
type ILeader interface {
	SetTaskResult(result interface{})
	GetTaskResult() chan interface{}
	NoticeTaskOver()
	IsCompleteTask() bool
}

// 学生完成任务
type ITask interface {
	Task(taskResources interface{}, num, taskSum int, leader ILeader)
}

func (*teacher)StartTask(taskResources interface{}, leader ILeader, student ITask, taskSum int, timeOutSecond float64) chan interface{}{

	// 1 开始任务
	for i := 0; i < taskSum; i++ {
		go student.Task(taskResources, i, taskSum, leader)
	}

	// 2 任务完成 或者 超时
	var result interface{}
	select {
	case <- time.After(time.Second * time.Duration(timeOutSecond)):
		result = errors.New("Timeout! Not Found")
	case tmp := <- leader.GetTaskResult():
		result = tmp
	}

	// 3 通知任务结束
	leader.NoticeTaskOver()
	ch := make(chan interface{}, 1)
	ch <- result
	return ch
}

