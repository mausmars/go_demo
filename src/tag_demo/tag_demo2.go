package main

import (
	"encoding/json"
	"fmt"
)

type StealAttach struct {
	RType int32			`json:"rType"`	 	//獎勵類型
	Coin int32			`json:"coin"`	//coin 數量
	CardIds []int32	`json:"cardIds"`	//cardIds
}

type RewardAction struct {
	Fpid   		int64  	`json:"fpid" form:"fpid"`	//fpid
	Atype   	int32 	`json:"atype" db:"atype"`
	Bet   		int32 	`json:"bet" db:"bet"`
	Attach  	string 	`json:"attach" db:"attach"`
	StealAttach	[]StealAttach
}

func main(){
	rewardAction:=RewardAction{
		Fpid:1,
		Atype:2,
		Bet:2,
		Attach:"abc",
		StealAttach:[]StealAttach{StealAttach{
			RType:1, Coin:1,CardIds:[]int32{1,1},
		},StealAttach{
			RType:2, Coin:2,CardIds:[]int32{2,2},
		}},
	}
	data,err:=json.Marshal(rewardAction)
	if(err!=nil){
		fmt.Println("err ",err)
	}else{
		fmt.Println(string(data[:]))
	}

	var rewardAction2 RewardAction
	err  = json.Unmarshal(data, &rewardAction2)
	if(err!=nil){
		fmt.Println("err ",err)
	}else{
		fmt.Println(rewardAction2)
		fmt.Println(rewardAction2.StealAttach[0].CardIds)
	}
}
