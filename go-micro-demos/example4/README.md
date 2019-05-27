#编译
protoc --proto_path=$GOPATH/src:. --go_out=. example4/proto/model/*.proto 
protoc --proto_path=$GOPATH/src:. --micro_out=. example4/proto/rpcapi/*.proto 

#运行
##服务端
    在 srv/ 下
        go run main.go


##客户端
    在 cli/ 下
        go run main.go


#功能说明
	在example3的基础上
	- 增加自定义proto协议
	- stream例子，不抛出panic err，可正常执行
		最新go-micro已经修复了问题本身，尚未核验
	- bistream仍然不可用，待查