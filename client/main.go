package main

import (
	"github.com/xtaci/kcp-go/v5"
	"google.golang.org/protobuf/proto"
	"log"
	pb "supernova/proto"
	"time"
)

func main() {
	// wait for server to become ready
	time.Sleep(time.Second)

	com := make([]byte, 4096)

	// dial to the echo server
	if sess, err := kcp.DialWithOptions("127.0.0.1:30100", nil, 4, 2); err == nil {
		go func() {
			for {
				//data := time.Now().String()
				req := &pb.EchoReq{
					Ping: "hi hank.",
				}
				buf, _ := proto.Marshal(req)
				log.Println("sent:", req.Ping)

				if _, err := sess.Write(buf); err == nil {
				} else {
					log.Fatal(err)
				}
				time.Sleep(time.Second)
			}
		}()

		go func() {
			for {
				if n, rerr := sess.Read(com); rerr == nil && n > 0 {
					//log.Println("recv:", n)
					res := &pb.EchoRes{}
					proto.Unmarshal(com[:n], res)
					log.Println("recv:", res.Pong)
				}
			}
		}()
	} else {
		log.Fatal(err)
	}

	shutdown := make(chan bool)
	<-shutdown
}
