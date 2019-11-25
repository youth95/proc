package goproc_rpc

import (
	"log"
	"net"
	"net/rpc"
)

func StartRpcService(addr string) chan error {
	errCH := make(chan error)
	go func() {
		err := rpc.RegisterName("linux", new(Service))
		if err != nil {
			errCH <- err
			close(errCH)
			return
		}
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			errCH <- err
			close(errCH)
			return
		}
		errCH <- nil
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("RPC Server Accept Error: %s", err.Error())
				continue
			}
			rpc.ServeConn(conn)
		}
	}()
	return errCH
}
