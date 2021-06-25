package net

const (
	ChannelType_In  = 1
	ChannelType_Out = 2
)

type ChannelHandlerContext struct {
}

type IChannelHandler interface {
	channelType() int
}

type IChannelInboundHandler interface {
	IChannelHandler
	channelActive(ctx *ChannelHandlerContext)                 // 与客户端建立连接
	channelInactive(ctx *ChannelHandlerContext)               // 与客户端断开连接
	channelRead(ctx *ChannelHandlerContext, var2 interface{}) // 接收到数据
	channelReadComplete(ctx *ChannelHandlerContext)
	exceptionCaught(ctx *ChannelHandlerContext, var2 error) // 异常
}

type IChannelOutboundHandler interface {
	IChannelHandler
	write(ctx *ChannelHandlerContext, var2 interface{})
}
