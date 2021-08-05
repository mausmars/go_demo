hello.go hello.c需要放一个目录下执行
# go build -o hello.so -buildmode=c-shared .

生成 hello.h hello.o 文件

编译main.c文件
# gcc -o main main.c hello/hello.so
