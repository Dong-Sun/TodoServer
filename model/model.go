package model

import (
	"log"
	"sort"
)

type todo struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
	IsDone  bool   `json:"isdone"`
}

var todoList map[int]todo
var size int

func Create(msg string) {
	if msg != "" {
		temp := todo{
			Message: msg,
			Id:      size,
			IsDone:  false,
		}

		todoList[size] = temp
		size++
	}
}

func Read() []todo {
	temp := []todo{}
	for _, r := range todoList {
		temp = append(temp, r)
	}
	sort.Slice(temp, func(i, j int) bool { return temp[i].Id < temp[j].Id })
	log.Println(temp)
	return temp
}

func Update(id int) {
	for _, r := range todoList {
		if r.Id == id {
			r.IsDone = !r.IsDone
		}
		break
	}
}

func Delete(id int) {
	for _, r := range todoList {
		if r.Id == id {
			delete(todoList, r.Id)
		}
		break
	}
}

func Init() {
	todoList = make(map[int]todo)
	size = 0
}
