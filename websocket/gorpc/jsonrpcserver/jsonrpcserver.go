package jsonrpcserver

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
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

func (t *Arith) Divide(args *Args, quot *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quot.Quo = args.A / args.B
	quot.Rem = args.A % args.B
	return nil
}

func JsonServerMain() {
	arith := new(Arith)
	_ = rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func checkError(e error) {
	if e != nil {
		fmt.Println("Fatal error:", e)
		os.Exit(1)
	}
}
