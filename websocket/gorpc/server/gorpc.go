package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A / args.B
	return nil
}

func ServerMain() {

	arith := new(Arith)

	// 注册一个Arith的RPC服务
	rpc.Register(arith)

	// 通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上
	rpc.HandleHTTP()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println(err.Error())
	}
}
