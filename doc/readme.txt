vi /etc/pki/tls/openssl.cnf

/etc/pki/CA/private/cakey.pem
//-------------------------------------------------------------

CA 根证书
生成CA根证书私钥
openssl genrsa -out testdata/cakey.pem 1024
openssl genrsa -out testdata/cakey.pem

生成CA根证书
openssl req -new -x509 -key testdata/cakey.pem  -out testdata/cacert.pem

//----------------------------------------------------
SSL 服务器证书
生成服务器证书私钥
openssl genrsa -out testdata/server.key 1024
openssl genrsa -out testdata/server.key

openssl req -new -key testdata/server.key -out testdata/server.csr

//----------------------------------------------------
签署服务器证书
openssl ca -in testdata/server.csr -out testdata/server.crt

//直接執行上邊可能會報錯，根據錯誤提示作如下操作
cp testdata/cakey.pem /etc/pki/CA/private
cp testdata/cacert.pem /etc/pki/CA
touch /etc/pki/CA/index.txt
touch /etc/pki/CA/serial
echo "01" > /etc/pki/CA/serial

//----------------------------------------------------
創建server.pem
touch testdata/server.pem
testdata/server.crt>>testdata/server.pem
testdata/server.key>>testdata/server.pem
