#背景
	基于例子

#编译
protoc --proto_path=$GOPATH/src:. --go_out=. example3/proto/model/*.proto 
protoc --proto_path=$GOPATH/src:. --micro_out=. example3/proto/rpcapi/*.proto 

#运行
##服务端
    在 srv/ 下
        go run main.go


##客户端
    在 cli/ 下
        go run main.go

#功能说明
	在example2的基础上
	- 增加自定义handler：handler/say2hello.go