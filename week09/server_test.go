package week09

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/Terry-Mao/goim/api/protocol"
	bufio1 "github.com/Terry-Mao/goim/pkg/bufio"
)

/*
	粘包处理:
	1. fix length: 客户端和服务器约定每次发送请求的大小，适合每个包比较小的场景、因为固定大小多出来的Buf需要用\0来填充，存在浪费；
	2. delimiter based: 特殊字符来分割，比如\n，适合每个包能指定明确边界的场景;
	3. length field based frame decoder: 每个包根据特定格式编码解码，需要能解析出数据长度，类似简单序列化；
*/

// 服务端复用了goim的proto的解码函数
func TestServer(t *testing.T) {
	server := New(&Config{
		Host: "localhost",
		Port: "6789",
	})
	server.Run()
}

// 客户端复用goim的proto的编码函数
func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:6789")
	if err != nil {
		// handle error
	}
	// wr := bufio.NewWriter(conn)
	// num, err := wr.WriteString("GET / HTTP/1.0\n")
	// wr.Flush()

	wr := bufio1.NewWriter(conn)
	p := protocol.Proto{}
	p.Op = 1
	p.Seq = 2
	p.Ver = 1
	p.Body = []byte("hello world\n")
	p.WriteTCP(wr)
	wr.Flush()

	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("status: %v, err: %v", status, err)

	time.Sleep(time.Second)

	//output:
	//status: Message received: hello world
}
