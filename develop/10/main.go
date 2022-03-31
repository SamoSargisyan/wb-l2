package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("tuc")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	client := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatalln("failed connecnt:", err)
	}
	defer client.Close()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	go func() {
		client.Send()
		cancel()
	}()
	go func() {
		client.Receive()
		cancel()
	}()

	<-ctx.Done()
}
