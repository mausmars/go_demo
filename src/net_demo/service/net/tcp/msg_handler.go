package tcp

import "reflect"

type ChannelContext struct {
}

type IChannelHandler interface {
	channelActive(ctx *ChannelContext)                // 与客户端建立连接
	channelInactive(ctx *ChannelContext)              // 与客户端断开连接
	channelRead(ctx *ChannelContext, msg interface{}) // 接收到数据
	channelReadComplete(ctx *ChannelContext)
	exceptionCaught(ctx *ChannelContext, err error) // 异常
}

/**
消息关联
*/
type MsgRelation struct {
	cid     int32        //指令id
	msgType reflect.Type //消息类型
	handler IMsgHandler  //处理器
}

/**
消息处理器接口
*/
type IMsgHandler interface {
	handle(session ISession, msg interface{})
}

//默认处理器
type DefaultMsgHandler struct {
}

func (h *DefaultMsgHandler) handle(session ISession, msg interface{}) {
}
