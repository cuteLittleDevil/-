package main

import "strconv"

type student struct {
	targetNum int
}

func NewStudent(num int) *student {
	return &student{targetNum: num}
}

func (s *student) Task(taskResources interface{}, num, taskSum int, leader ILeader) {
	if resource, ok := taskResources.([]int); ok {
		groupNum := len(resource) / taskSum
		if groupNum <= 1 {
			groupNum = 1
		}
		startNum := num * groupNum
		endNum := startNum + groupNum
		if endNum + groupNum > len(resource) {
			endNum = len(resource)
		}
		for i := startNum; i < endNum; i++ {
			if resource[i] == s.targetNum {
				leader.SetTaskResult("Found it is " + strconv.Itoa(i))
				return
			}
			if leader.IsCompleteTask() {
				return
			}
		}

	}
}
