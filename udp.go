package main

import (
	"log"
	"net"
)

func udpForward(forward ForwardStruct) {
	listenAddress, err := net.ResolveUDPAddr("udp", forward.From)
	if err != nil {
		log.Println("Error: failed to parse udp listen address")
	}
	src, err := net.ListenUDP(forward.Protocol, listenAddress)
	if err != nil {
		log.Printf("The connection failed: %v", err)
	}
	defer src.Close()

	var sliceDst []*net.UDPConn

	for _, to := range forward.To {
		log.Println("loop for from to to")
		dstAddr, err := net.ResolveUDPAddr(forward.Protocol, to)
		if err != nil {
			log.Printf("Error resolving destination address: %v\n", err)
		}

		dst, err := net.DialUDP(forward.Protocol, nil, dstAddr)
		if err != nil {
			log.Printf("The connection failed: %v", err)
		}
		defer dst.Close()

		sliceDst = append(sliceDst, dst)
	}

	for {
		log.Println("loop for to to from")
		buf := make([]byte, 65535)
		n, _, err := src.ReadFrom(buf)
		if err != nil {
			log.Printf("Error reading from UDP socket: %v\n", err)
		}

		for _, dst := range sliceDst {
			_, _ = dst.Write(buf[:n])
		}
	}

	log.Println("loops exited")
}
