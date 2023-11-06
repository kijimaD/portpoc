package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	// ランダムな空きポートを取得
	port, err := getFreePort()
	if err != nil {
		log.Fatal(err)
	}

	// HTTPハンドラを設定
	http.HandleFunc("/", helloHandler)

	addr := fmt.Sprintf(":%d", port)

	// ブラウザを開く
	listen := make(chan bool)
	go func() {
		<-listen
		open.Run(fmt.Sprintf("http://localhost%s", addr))
		fmt.Println("browser start")
	}()
	listen <- true

	// サーバを開始
	fmt.Printf("Server is running on port %d\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is working...")
}

func getFreePort() (int, error) {
	// 空きポートを検出
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	// 実際に使用されているポートを取得
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}
