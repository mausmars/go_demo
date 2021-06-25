package net

type NetConfig struct {
	netProtocol string
	host        string

	buffSize  int32
	backlog   int32
	keepalive bool
	nodely    bool
	linger    int
	//----------------------

	acceptCoroutineNum int
}
