package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AddDeviceRequest struct {
	ProjectId    int32  `json:"project_id"`
	UserId       int64  `json:"user_id"`
	PlatformType string `json:"platform_type"`
	DeviceType   string `json:"device_type"`
	ChannelType  string `json:"channel_type"`
	DeviceToken  string `json:"device_token"`
	Timestamp    int64  `json:"timestamp"`
	Signature    string `json:"signature"`
}

type PushUserRequest struct {
	ProjectId int32   `json:"project_id"`
	UserIds   []int64 `json:"user_ids"`
	Msg       string  `json:"msg"`
	IsPlain   bool    `json:"is_plain"`
	Timestamp int64   `json:"timestamp"`
	Signature string  `json:"signature"`
}

type Message struct {
	Msg string `json:"message"`
}

type DataMessage struct {
	Data Message `json:"data"`
}

type GCMMessage struct {
	GCM DataMessage `json:"GCM"`
}

var projectId = int32(100027)
var key = "h1xbFlfgRWgw6YNBquNnUduFdhhlvNPED95JNxsNu27ao9NL"

var userId = int64(2701851355080)
var token = "eq5z4J3-AIA:APA91bHAW_87zJkfRuDBYvO3KK1Wz8UyRs_2569geStfcANKNC3dtsmMMTEGphQ0C3SaGxKsvvhaseB8GkH-C0FLPmi8tOIM1bv59BgHpfoC_Ag7aqdvkwv7b6I89lKqe9hhkd52A5hi"
//var platformType = "apns"
var platformType = "gcm"

func sendAddDeviceTest() {
	timestamp := time.Now().Unix()
	str := strconv.Itoa(int(projectId)) + ":" + strconv.FormatInt(timestamp, 10) + ":" + key

	c := md5.New()
	c.Write([]byte(str))
	signature := hex.EncodeToString(c.Sum(nil))
	fmt.Println(signature)

	url := "https://notification.puzzleplusgames.net/add-device"
	msg := &AddDeviceRequest{
		ProjectId:    projectId,
		UserId:       userId,
		PlatformType: platformType,
		DeviceType:   "phone",
		//ChannelType:  "mipush",
		DeviceToken: token,
		Timestamp:   timestamp,
		Signature:   signature,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal ", err)
		return
	}
	postjson := string(b)
	fmt.Println("msg= ", postjson)

	res, err := http.Post(url, "application/json", strings.NewReader(postjson))
	if err != nil {
		fmt.Println("Post ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("postJson Error StatusCode: %d postdata: %s", res.StatusCode, postjson)
		return
	}

	v, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ReadAll ", err)
		return
	}
	fmt.Println("AddDevice 平台返回  ", string(v))
}

func sendPushTest() {
	timestamp := time.Now().Unix()+60
	str := strconv.Itoa(int(projectId)) + ":" + strconv.FormatInt(timestamp, 10) + ":" + key

	c := md5.New()
	c.Write([]byte(str))
	signature := hex.EncodeToString(c.Sum(nil))
	fmt.Println(signature)

	url := "https://notification.puzzleplusgames.net/push"
	msg := &PushUserRequest{
		ProjectId: projectId,
		UserIds:   []int64{userId},
		Msg:       "test",
		IsPlain:   true,
		Timestamp: timestamp,
		Signature: signature,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal ", err)
		return
	}
	postjson := string(b)
	fmt.Println("msg= ", postjson)

	res, err := http.Post(url, "application/json", strings.NewReader(postjson))
	if err != nil {
		fmt.Println("Post ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("postJson Error StatusCode: %d postdata: %s", res.StatusCode, postjson)
		return
	}

	v, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ReadAll ", err)
		return
	}
	fmt.Println("Push 平台返回  ", string(v))
	// v = []byte("{\"code\":20014,\"msg\":\"20014\"}")
	// resqs := CheckSessionResponse{}
	//err = json.Unmarshal(v, resp)
}

func main() {
	sendAddDeviceTest()
	sendPushTest()
}
