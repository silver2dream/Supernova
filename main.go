package main

import (
	"fmt"
	"github.com/xtaci/kcp-go/v5"
	"google.golang.org/protobuf/proto"
	"log"
	"supernova/proto"
)

func main() {
	fmt.Println("This is supernova, the super project will be start.")
	if lisener, err := kcp.ListenWithOptions("0.0.0.0:30100", nil, 4, 2); err == nil {
		for {
			s, err := lisener.AcceptKCP()
			if err != nil {
				log.Fatal(err)
			}
			go handlerEcho(s)
		}
	} else {
		log.Fatal(err)
	}
}

func handlerEcho(s *kcp.UDPSession) {
	buf := make([]byte, 4096)
	msg := make(chan []byte, 1000)

	go func() {
		for {
			n, err := s.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}
			msg <- buf[:n]
		}
	}()

	go func() {
		for {
			select {
			case pack := <-msg:
				req := &pb.EchoReq{}
				proto.Unmarshal(pack, req)
				res := &pb.EchoRes{
					Pong: req.Ping,
				}
				log.Println("server:", res.Pong)
				data, _ := proto.Marshal(res)

				go func() {
					n, err := s.Write(data)
					if err == nil {
						log.Println(n)
					}
				}()

			}

		}
	}()
}
