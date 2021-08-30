php调用go
思路：go生成c的so动态库，php调用动态库
//----------------------------------------------
创建 hello.go
注意：生成C可调用的so时，Go源代码需要以下几个注意。
必须导入"C"包
必须在可外部调用的函数前加上[//export 函数名]的注释
必须是main包，切含有main函数，main函数可以什么都不干
-----------------
go install -buildmode=shared -linkshared std
go build -buildmode=c-shared -o hello.so hello.go

php hello.php
php -d extension=./hello.so hello2.php
php hello2.php

php -d extension=/mnt/hgfs/vm_share/go_demo/src/php_demo/php_helloworld/hello.so hello2.php


修改extension加入最后
/software/php-7.3.29/lib/php.ini

extension=/mnt/hgfs/vm_share/go_demo/src/php_demo/php_helloworld/hello.so

//===============================================
go 结构体问题
https://pkg.go.dev/cmd/cgo#hdr-C_references_to_Go

大概思路 是都往c上靠