package main

import (
	"github.com/AlphaFoxz/hot-deploy-go-example/domain/api"
	"github.com/AlphaFoxz/hot-deploy-go-example/domain/declare"
	"github.com/AlphaFoxz/hot-deploy-go-example/generator"
	"time"
)

type RepoImpl struct {
}

func (RepoImpl) FindById(id string) declare.ParkingAgg {
	return declare.ParkingAgg{
		//TODO EventQueue_未实现
		EventQueue_: nil,
		Id:          declare.Plate{Number: id},
	}
}

func main() {
	domainApi := api.New(RepoImpl{})
	command := domainApi.NewCheckInCommand(declare.Plate{Number: "1"}, time.Now())
	println(command.Handle())

	//启动代码生成守护进程
	generator.Listen("./domain")
}
