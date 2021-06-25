package model

type IData interface {
	Submit()
	Callback()
}

type IDataChanger interface {
	IData
	insert(data IData)
}

//--------------------------------------
type DataChanger struct {
	ChangeData map[string]IData
}

func (c *DataChanger) submit() {

}

func (c *DataChanger) callback() {

}

func (c *DataChanger) insert(data IData) {

}

//--------------------------------------
type UserData struct {
	UserId int64  `json:"userId" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Attr   string `json:"attr" form:"attr"`

	attrMap map[string]interface{}

	changeField map[string]interface{}
}

func NewUserData(UserId int64, Name string, Attr string) *UserData {
	return &UserData{
		UserId:      UserId,
		Name:        Name,
		Attr:        Attr,
		changeField: make(map[string]interface{}),
	}
}

func (data *UserData) SetUserId(userId int64) {
	if data.changeField["UserId"] == nil {
		data.changeField["UserId"] = data.UserId
	}
	data.UserId = userId
}

func (data *UserData) Submit() {
	data.reset()
}

func (data *UserData) Callback() {
	for k, v := range data.changeField {
		switch k {
		case "UserId":
			data.UserId = v.(int64)
			break
		case "Name":
			data.Name = v.(string)
			break
		case "Attr":
			data.Attr = v.(string)
			break
		}
	}
	data.reset()
}

func (data *UserData) reset() {
	if len(data.changeField) > 0 {
		data.changeField = make(map[string]interface{})
	}
}

//--------------------------------------
