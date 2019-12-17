package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintln(stderr, "usage: cidr CIDR")
		return 1
	}

	out, err := calc(args[1])
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
	}
	fmt.Fprintln(stdout, out)
	return 0
}

func calc(cidr string) (string, error) {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	bc := make(net.IP, 4, 4)
	for i, b := range ipnet.Mask {
		bc[i] = (0xff ^ b) + ipnet.IP[i]
	}

	return fmt.Sprint(ipnet.IP, " - ", bc), nil
}
