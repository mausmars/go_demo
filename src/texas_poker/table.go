package texas_poker

type Card struct {
	n int // 2-14 A的特殊的也当做1
	t int // 1-4黑红梅方
}

type Seat struct {
	cards []*Card //底牌 2张
}

type ITable interface {
}

type Table struct {
	round int //回合数

	seats  []*Seat //座位
	indexs []*int  //索引 0庄 1小盲 2大盲
	cindex int     //当前索引

	//加注索引
	//加注值

	cards []*Card //底牌 5张

	observers map[int64]*Seat //旁观者
}
