package tcp

import (
	"go_demo/src/net_demo/service/facade"
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"reflect"
)

/**
消息配置
*/
type MsgConfig struct {
	//{指令,映射}
	commandMap map[int32]*MsgRelation
	//{类型,映射}
	typeNameMap map[string]*MsgRelation
}

func (c *MsgConfig) AddMsgRelation(cid int32, msg interface{}, handler IMsgHandler) {
	msgType := reflect.TypeOf(msg)
	msgName := msgType.Elem().Name()
	fmt.Println("注册消息 ", msgName)
	var msgRelation = &MsgRelation{
		cid:     cid,
		msgType: msgType,
		handler: handler,
	}
	c.commandMap[cid] = msgRelation
	c.typeNameMap[msgName] = msgRelation
}

func (c *MsgConfig) getMsgRelationByMsgId(msgId int32) *MsgRelation {
	return c.commandMap[msgId]
}

func (c *MsgConfig) getMsgRelationByType(msg interface{}) *MsgRelation {
	msgType := reflect.TypeOf(msg)
	msgName := msgType.Elem().Name()
	return c.typeNameMap[msgName]
}

func NewMsgConfig() *MsgConfig {
	return &MsgConfig{
		commandMap:  make(map[int32]*MsgRelation),
		typeNameMap: make(map[string]*MsgRelation),
	}
}

/**
网络配置
*/
type NetConfig struct {
	Host         string `default:"0.0.0.0:12345" toml:"net_host"`
	BuffSize     int32  `default:"4096" toml:"net_buff_size"`
	Backlog      int32  `default:"1000" toml:"net_backlog"`
	Keepalive    bool   `default:"true" toml:"net_keepalive"`
	Nodely       bool   `default:"true" toml:"net_nodely"`
	Linger       int    `default:"0" toml:"net_linger"`
	ReadTimeout  int64  `default:"16000" toml:"net_read_timeout"`
	WriteTimeout int64  `default:"16000" toml:"net_write_timeout"`
	IdleTimeout  int64  `default:"600000" toml:"net_idle_timeout"`
}

func NewNetConfig(configPath string) *NetConfig {
	cfg := &NetConfig{}
	filePath, err := filepath.Abs(configPath)
	if err != nil {
		panic(err)
	}
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		panic(err)
	}
	return cfg
}

/**
网络服务配置
*/
type NetServiceConfig struct {
	netConfig     *NetConfig
	msgConfig     *MsgConfig
	serviceFacade *facade.ServiceFacade
}

func NewNetServiceConfig(configPath string, msgConfig *MsgConfig, serviceFacade *facade.ServiceFacade) *NetServiceConfig {
	return &NetServiceConfig{
		netConfig:     NewNetConfig(configPath),
		msgConfig:     msgConfig,
		serviceFacade: serviceFacade,
	}
}
