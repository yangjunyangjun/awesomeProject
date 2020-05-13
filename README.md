# awesomeProject
go-micro

make proto

protoc --proto_path=. --micro_out=. --go_out=. hellowolrd.proto


run consul 

consul agent -dev
