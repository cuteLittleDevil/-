package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	targetNum := 435
	s := GetRandomSlice(100000, 10000, targetNum)
	ch := GetTeacherInstance().StartTask(s, NewGroupLeader(), NewStudent(targetNum), 8, 5.0)
	fmt.Println(<-ch)
}

func GetRandomSlice(sliceLen int, randNum int, targetNum int) []int{
	tmp := make([]int, sliceLen)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sliceLen; i++ {
		randNum := rand.Intn(randNum)
		tmp[i] = randNum
		if randNum == targetNum {
			fmt.Println("randNum is ", i)
		}
	}
	return tmp
}

