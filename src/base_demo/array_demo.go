package main

func main() {
	buf1 := []byte{1, 2, 3}
	modifySize1(buf1)
	println("---1改变后---")
	printBuf(buf1)
	println("==================")
	buf2 := &[]byte{1, 2, 3}
	modifySize2(buf2)
	println("---2改变后---")
	printBuf(*buf2)
}
func modifySize1(buf []byte) {
	buf = buf[:2]
	buf[0]=8
	buf[1]=9
	printBuf(buf)
}
func modifySize2(buf *[]byte) {
	//修改指针的值
	arr := *buf
	arr = arr[:2]
	*buf = arr
	printBuf(*buf)
}
func printBuf(buf []byte) {
	for _, b := range buf {
		print(b, ",")
	}
	println("")
}
