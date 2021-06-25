package main

import (
	"fmt"
	model2 "go_demo/src/dao_demo/model"

	"encoding/json"
)

func main() {
	userData := model2.NewUserData(1, "test", "{}")
	userData.SetUserId(111)
	fmt.Println(*userData)
	userData.SetUserId(222)
	fmt.Println(*userData)
	userData.Callback()
	fmt.Println(*userData)

	data, _ := json.Marshal(userData)
	fmt.Println(string(data))
	m := &model2.UserData{}
	json.Unmarshal(data, m)
	fmt.Println(m)
}
