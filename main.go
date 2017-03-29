package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/proc/sys/net/bridge/bridge-nf-call-iptables")
	if err != nil {
		log.Fatal(err)
	}
	lr := io.LimitReader(f, 1)
	defer f.Close()

	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, lr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
