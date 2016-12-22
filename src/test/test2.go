package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func Handle(conn net.Conn) {

	cmd := exec.Command("nslookup")
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	conn.(*io.Reader) = stdout
	//	content, err := ioutil.ReadAll(stdout)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(string(content))

	//	cmd := exec.Command("nslookup")
	//	stdin, _ := cmd.StdinPipe()
	//	stdout, _ := cmd.StdoutPipe()
	//	okChan := make(chan bool, 1)
	//	go func(w io.Writer, r io.Reader) {
	//		log.Println("go1")
	//		okChan <- true
	//		io.Copy(w, r)
	//	}(conn, stdout)
	//	//	cmd.Stdout = os.Stdout
	//	go func(w io.Writer, r io.Reader) {
	//		log.Println("go2")
	//		okChan <- true
	//		io.Copy(w, r)
	//	}(stdin, conn)
	//	<-okChan
	//	<-okChan
	//	log.Println("go3")
	//	if err != nil {
	//		fmt.Println("err2")
	//		fmt.Println(err)
	//	}
	//	cmd.Stdout = os.Stdout
	cmd.Start()
	//	fmt.Println("111")
	//	content, _ := ioutil.ReadAll(stdout)
	//	fmt.Println("222")
	//	fmt.Println(string(content))
	//	for {
	//		buf := make([]byte, 2048)
	//		n, err := conn.Read(buf)
	//		if err != nil {
	//			fmt.Println("err3")
	//			fmt.Println(err)
	//			break
	//		}
	//		fmt.Print(string(buf[:n]))
	//		_, err = stdin.Write(buf[:n])
	//		if err != nil {
	//			fmt.Println("err4")
	//			fmt.Println(err)
	//			break
	//		}
	//	}
	//	conn.Close()
	//	stdin.Close()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start server...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err1")
			log.Fatal(err)
		}
		go Handle(conn) // 每次建立一个连接就放到单独的线程内做处理
	}
}
