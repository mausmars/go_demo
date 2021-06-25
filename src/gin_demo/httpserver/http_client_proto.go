package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go_demo/src/component/statistics_service"
	protomsg "go_demo/src/gin_demo/httpserver/proto"
	"io/ioutil"
	"net/http"
	"time"
)

func sendMessage(client *http.Client, m string) {
	url := "http://localhost:8008/proto"

	msg := &protomsg.Command{
		CommandId: int32(1),
		Body:      []byte(m),
	}
	data, _ := proto.Marshal(msg)
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		println(err)
	}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	reMsg := &protomsg.Command{}
	proto.Unmarshal(body, reMsg)
	fmt.Println(resp.StatusCode)
	fmt.Println(reMsg)
	fmt.Println(resp.Header)
}

func sendPushMessage(client *http.Client, m string, statisticsService *statistics_service.StatisticsService) {
	url := "http://localhost:8008/proto"

	msg := &protomsg.Command{
		CommandId: int32(2),
		Body:      []byte(m),
	}
	data, _ := proto.Marshal(msg)
	buf := bytes.NewBuffer(data)

	commandInfo := &statistics_service.CommandInfo{
		SendTime: time.Now(),
		Command:  "http req",
		Flow:     int64(len(data)),
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		println(err.Error())
		statisticsService.Recorder(commandInfo)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		println(err.Error())
		statisticsService.Recorder(commandInfo)
		return
	}
	commandInfo.IsSuccess = true
	statisticsService.Recorder(commandInfo)

	body, _ := ioutil.ReadAll(resp.Body)
	reMsg := &protomsg.Command{}
	proto.Unmarshal(body, reMsg)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(reMsg)
	//fmt.Println(resp.Header)
	//fmt.Println(resp.Proto)
}

func main() {
	statisticsService := &statistics_service.StatisticsService{
	}
	statisticsService.StartUp()

	client := &http.Client{
		Timeout: 120 * time.Second,
		Transport: &http.Transport{
			// 不验证证书
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	var i int
	var explain = "----------------------------\n" +
		"输入指令序号\n" +
		"1.退出测试\n" +
		"2.发送消息\n" +
		"3.发送推送\n" +
		"4.查看测试信息"
	for {
		fmt.Println(explain)
		fmt.Scan(&i)
		var isOver = false
		switch i {
		case 1:
			isOver = true
			statisticsService.Shutdown()
			break
		case 2:
			go sendMessage(client, "abc")
			break
		case 3:
			for i := 0; i < 20000; i++ {
				go sendPushMessage(client, "push", statisticsService)
			}
			break
		case 4:
			statisticsService.Show()
			break
		}
		if isOver {
			break
		}
	}
}
