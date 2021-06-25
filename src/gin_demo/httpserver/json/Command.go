package jsonmsg

type Command struct {
	CommandId int32  `json:"commandId"`
	Body      string `json:"body"`
}
