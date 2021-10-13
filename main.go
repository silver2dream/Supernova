package main

import (
	"fmt"
	"github.com/xtaci/kcp-go/v5"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("This is supernova, the super project will be start.")
	if lisener, err := kcp.ListenWithOptions("0.0.0.0:30100", nil, 4, 2); err == nil {
		go client()

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
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		n, err = s.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func client() {
	// wait for server to become ready
	time.Sleep(time.Second)

	// dial to the echo server
	if sess, err := kcp.DialWithOptions("127.0.0.1:30100", nil, 4, 2); err == nil {
		for {
			data := time.Now().String()
			msg:=proto.
			buf := make([]byte, len(data))
			log.Println("sent:", data)
			if _, err := sess.Write([]byte(data)); err == nil {
				// read back the data
				if _, err := io.ReadFull(sess, buf); err == nil {
					log.Println("recv:", string(buf))
				} else {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	} else {
		log.Fatal(err)
	}
}
