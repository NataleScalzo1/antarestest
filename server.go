package main

import (
	"fmt"
	"github.com/gookit/color"
	"net"
	"strings"
	"time"
)

var address string = "127.0.0.1"
var udpport string = "4420"
var tcpport string = "4421"

func WriteUDPMessage(msg string, localIp string) {

	start := time.Now()
	buffer := make([]byte, 1024)
	var reply string

	defer func() {

		fmt.Println("Time Elapsed: ", time.Since(start).String())
		fmt.Println("Message Sent: ", localIp, "  => ", msg)
		fmt.Println("Message Received: => ", reply)

		if r := recover(); r != nil {
			_, _ = color.Set(color.Red)
			fmt.Printf("UdpClient - Recovered from Panic: %v\n", r)
			_, _ = color.Reset()
		}
	}()

	remoteAddr := address + ":" + udpport
	localAddr := "127.0.0.1" + ":"

	remoteUdpAddr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("ResolveUDPAddr remote addr failed: ", err)
		_, _ = color.Reset()
		return
	}
	localUdpAddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("ResolveUDPAddr local addr failed: ", err)
		_, _ = color.Reset()
		return
	}
	conn, err := net.DialUDP("udp", localUdpAddr, remoteUdpAddr)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println(err)
		_, _ = color.Reset()
		return
	}

	if strings.HasPrefix(msg, "L1|") {
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	}
	if strings.HasPrefix(msg, "L7|") {
		conn.SetReadDeadline(time.Now().Add(1500 * time.Millisecond))
	}
	if strings.HasPrefix(msg, "L7B|") {
		conn.SetReadDeadline(time.Now().Add(1200 * time.Millisecond))
	}
	if strings.HasPrefix(msg, "L8|") {
		conn.SetReadDeadline(time.Now().Add(1200 * time.Millisecond))
	}

	defer conn.Close()

	data := []byte(msg + "\x00")
	conn.SetWriteBuffer(len(data))
	_, err = conn.Write(data)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println(err)
		_, _ = color.Reset()
		return
	}

	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println(err)
		_, _ = color.Reset()
		return
	}
	reply = string(buffer[0:n])
	if string(reply) == "STOP" {
		_, _ = color.Set(color.Red)
		fmt.Println(reply)
		_, _ = color.Reset()
	}
}
func WriteUDPMessage2(msg string, localIp string) {

	con, err := net.Dial("udp", address+":"+udpport)
	if err != nil {
		fmt.Println("Error connecting: " + err.Error())
	}

	defer con.Close()

	for {

		fmt.Println("UDP MESSAGE")

		con.Write([]byte("UDP MESSAGE"))

		return

	}
}

func WriteTCPMessage(msg string, localIp string) []byte {

	start := time.Now()
	reply := make([]byte, 1024)

	defer func() {

		fmt.Println("Time Elapsed: ", time.Since(start).String())
		fmt.Println("Message Sent: ", localIp, "  => ", msg)
		fmt.Println("Message Received: => ", string(reply))

		if r := recover(); r != nil {
			_, _ = color.Set(color.Red)
			fmt.Printf("TcpClient - Recovered from Panic: %v\n", r)
			_, _ = color.Reset()
		}
	}()

	remoteAddr := address + ":" + tcpport
	localAddr := "127.0.0.1" + ":"

	remoteTcpAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("ResolveTCPAddr remote addr failed: ", err)
		_, _ = color.Reset()
	}

	localTcpAddr, err := net.ResolveTCPAddr("tcp", localAddr)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("ResolveTCPAddr local addr failed: ", err)
		_, _ = color.Reset()
	}

	conn, err := net.DialTCP("tcp", localTcpAddr, remoteTcpAddr)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("Dial TCP failed: ", err)
		_, _ = color.Reset()
	}

	defer conn.Close()

	data := []byte(msg + "\x00")
	conn.SetNoDelay(true)
	conn.SetWriteBuffer(len(data))
	_, err = conn.Write(data)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("TCP Write to server failed: ", err)
		_, _ = color.Reset()
	}

	start = time.Now()

	if strings.HasPrefix(msg, "L6|") {
		conn.SetReadDeadline(time.Now().Add(1500 * time.Millisecond))
	}
	if strings.HasPrefix(msg, "L7|") {
		conn.SetReadDeadline(time.Now().Add(1500 * time.Millisecond))
	}

	_, err = conn.Read(reply)
	if err != nil {
		_, _ = color.Set(color.Red)
		fmt.Println("TCP Write to server failed: ", err)
		_, _ = color.Reset()
	}

	if string(reply) == "STOP" {
		_, _ = color.Set(color.Red)
		fmt.Println(reply)
		_, _ = color.Reset()
	}
	return reply
}

func SendCommandModbus(msg string, ip string) {

	con, err := net.Dial("tcp", "localhost:4421")
	if err != nil {
		fmt.Println("Error connecting: " + err.Error())
	}

	defer con.Close()

	for {

		fmt.Println(msg)
		return

	}
}
