package test2

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
	"github.com/highras/rtm-server-sdk-go/src/rtm"
)

type PrintLocker struct {
	mutex sync.Mutex
}

func (locker *PrintLocker) P2PMessage(fromUid int64, toUid int64, mtype int8, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive P2P msg: from:%d -> to:%d mtype:%d, mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, toUid, mtype, mid, mtime, message, attrs)
}
func (locker *PrintLocker) GroupMessage(fromUid int64, groupId int64, mtype int8, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive group msg: from:%d -> group:%d mtype:%d, mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, groupId, mtype, mid, mtime, message, attrs)
}
func (locker *PrintLocker) RoomMessage(fromUid int64, roomId int64, mtype int8, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive room msg: from:%d -> room:%d mtype:%d, mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, roomId, mtype, mid, mtime, message, attrs)
}
func (locker *PrintLocker) Event(pid int32, event string, uid int64, eventTime int32, endpoint string, data string) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Println("[Server Push] Receive event: %s: user:%d, time:%d, endpoint:%s, data:%s\n", event, uid, eventTime, endpoint, data)
}
func (locker *PrintLocker) P2PChat(fromUid int64, toUid int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive P2P msg: from:%d -> to:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, toUid, mid, mtime, message, attrs)
}
func (locker *PrintLocker) GroupChat(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive group msg: from:%d -> group:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, groupId, mid, mtime, message, attrs)
}
func (locker *PrintLocker) RoomChat(fromUid int64, roomId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive room msg: from:%d -> room:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, roomId, mid, mtime, message, attrs)
}
func (locker *PrintLocker) P2PAudio(fromUid int64, roomId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive P2P audio: from:%d -> to:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, toUid, mid, mtime, message, attrs)
}
func (locker *PrintLocker) P2PCmd(fromUid int64, roomId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive P2P cmd: from:%d -> to:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, toUid, mid, mtime, message, attrs)
}
func (locker *PrintLocker) GroupAudio(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive group audio: from:%d -> group:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, groupId, mid, mtime, message, attrs)
}
func (locker *PrintLocker) GroupCmd(fromUid int64, groupId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive group cmd: from:%d -> group:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, groupId, mid, mtime, message, attrs)
}
func (locker *PrintLocker) RoomAudio(fromUid int64, roomId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive room audio: from:%d -> room:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, roomId, mid, mtime, message, attrs)
}
func (locker *PrintLocker) RoomCmd(fromUid int64, roomId int64, mid int64, message string, attrs string, mtime int64) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	fmt.Printf("[Server Push] Receive room cmd: from:%d -> room:%d mid:%d mtime: %d\nmessage: %s\nattrs: %s\n", fromUid, roomId, mid, mtime, message, attrs)
}
func (locker *PrintLocker) print(proc func()) {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()

	proc()
}

var locker PrintLocker = PrintLocker{}

var (
	adminUid int64   = 111
	fromUid  int64   = 102456
	toUid    int64   = 102457
	toUids   []int64 = []int64{102458, 102459, 102460, 102461, 102462, 102463, 102464, 102465, 102466, 102467, 102468}
	groupId  int64   = 12345
	roomId   int64   = 9981
	mtype    int8    = 127
)

func addListen(client *rtm.RTMServerClient) {
	//-- sync mode
	err := client.AddListen([]int64{groupId}, []int64{roomId}, []int64{toUid}, []string{"login"},func (a int,b string){
		fmt.Println("?????????????????? ",a,"   ",b)
	})
	locker.print(func() {
		if err == nil {
			fmt.Printf("AddListen in sync mode is fine.\n")
		} else {
			fmt.Printf("AddListen in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	//err = client.AddListen([]int64{groupId}, []int64{roomId}, toUids, []string{"logout"}, func(errorCode int, errInfo string) {
	//	locker.print(func() {
	//		if errorCode == fpnn.FPNN_EC_OK {
	//			fmt.Printf("AddListen in async mode is fine.\n")
	//		} else {
	//			fmt.Printf("AddListen in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
	//		}
	//	})
	//})
	//
	//if err != nil {
	//	locker.print(func() {
	//		fmt.Printf("AddListen in async mode error, err: %v\n", err)
	//	})
	//}
}

func removeListen(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.RemoveListen([]int64{groupId}, []int64{roomId}, toUids, []string{"login"})
	locker.print(func() {
		if err == nil {
			fmt.Printf("RemoveListen in sync mode is fine.\n")
		} else {
			fmt.Printf("RemoveListen in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.RemoveListen([]int64{groupId}, []int64{roomId}, []int64{toUid}, []string{"logout"}, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("RemoveListen in async mode is fine.\n")
			} else {
				fmt.Printf("RemoveListen in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("RemoveListen in async mode error, err: %v\n", err)
		})
	}
}

func setListen(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.SetListen([]int64{groupId}, []int64{roomId}, []int64{toUid}, []string{"login"})
	locker.print(func() {
		if err == nil {
			fmt.Printf("SetListen in sync mode is fine.\n")
		} else {
			fmt.Printf("SetListen in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.SetListen([]int64{groupId}, []int64{roomId}, toUids, []string{"logout"}, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("SetListen in async mode is fine.\n")
			} else {
				fmt.Printf("SetListen in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("SetListen in async mode error, err: %v\n", err)
		})
	}
}

func setListenStatus(client *rtm.RTMServerClient) {

	//-- sync mode
	err := client.SetListenStatus(true, true, true, false)
	locker.print(func() {
		if err == nil {
			fmt.Printf("SetListenStatus in sync mode is fine.\n")
		} else {
			fmt.Printf("SetListenStatus in sync mode error, err: %v\n", err)
		}
	})

	//-- async mode
	err = client.SetListenStatus(true, true, true, false, func(errorCode int, errInfo string) {
		locker.print(func() {
			if errorCode == fpnn.FPNN_EC_OK {
				fmt.Printf("SetListenStatus in async mode is fine.\n")
			} else {
				fmt.Printf("SetListenStatus in async mode error, error code: %d, error info:%s\n", errorCode, errInfo)
			}
		})
	})

	if err != nil {
		locker.print(func() {
			fmt.Printf("SetListenStatus in async mode error, err: %v\n", err)
		})
	}
}

func Test() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pid, err := strconv.Atoi("80000017")
	if err != nil {
		fmt.Println("Pid is invalid. Error:", err)
		return
	}
	client := rtm.NewRTMServerClient(int32(pid), "c249ebf3-eee9-452f-b4d3-9f3e5dcb793c", "rtm-intl-frontgate.ilivedata.com:13325")

	locker := &PrintLocker{}
	client.SetMonitor(locker)

	addListen(client)
	locker.print(func() {
		fmt.Println("Add listen, waiting 20 second for client send messages")
	})
	time.Sleep(20 * time.Second)

	//removeListen(client)
	//locker.print(func() {
	//	fmt.Println("Remove listen, waiting 20 second for client send messages")
	//})
	//time.Sleep(20 * time.Second)
	//
	//setListen(client)
	//locker.print(func() {
	//	fmt.Println("Set listen, waiting 20 second for client send messages")
	//})
	//time.Sleep(20 * time.Second)
	//
	//setListenStatus(client)
	//locker.print(func() {
	//	fmt.Println("Set listen status, waiting 20 second for client send messages")
	//})
	time.Sleep(20 * time.Second)

	time.Sleep(time.Second) //-- Waiting for the async callback printed.
}
