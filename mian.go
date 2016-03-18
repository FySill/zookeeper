package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

type ClientHeaderInfo struct {
	len    uint32
	code   uint32
	magic  uint16
	magic2 uint16
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
	/*buffer := bytes.NewBufferString("")*/
	var msg ClientHeaderInfo
	buf := setHeaderBuffer("ashycx", "xiang66", "532152", &msg)
	fmt.Println(buf.String())
	_, err = conn.Write([]byte(buf.String()))
	fmt.Println(conn)
	fmt.Println("1")
	if err != nil {
		fmt.Println(err)
	}

	buf = setJoinBuffer(&msg)
	fmt.Println(buf.String())
	_, err = conn.Write([]byte(buf.String()))
	fmt.Println(conn)
	fmt.Println("2")
	if err != nil {
		fmt.Println(err)
	}

	result, err := ioutil.ReadAll(conn)
	fmt.Println("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("aa" + string(result))
}

func setHeaderBuffer(user, password, roomid string, msg *ClientHeaderInfo) bytes.Buffer {
	var buffer bytes.Buffer
	str := ("type@=loginreq/username@=" + user + "/password@=" + password + "/roomid@=" + roomid + "/")
	//str := ("type@=loginreq/roomid@=" + roomid)
	l := len(str) + 4 + 4 + 4 + 1
	msg.len = uint32(l)
	msg.code = uint32(l)
	msg.magic = uint16(689)
	msg.magic2 = uint16(0)
	binary.Write(&buffer, binary.LittleEndian, msg)
	buffer.WriteString(str)
	buffer.WriteByte(0)
	fmt.Println(buffer)
	return buffer
	/*var icmp ICMP
	binary.Write(&buffer, binary.BigEndian, msg)
	fmt.Println(buffer)
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	fmt.Println(buffer)*/
}
func setJoinBuffer(msg *ClientHeaderInfo) bytes.Buffer {
	var buffer bytes.Buffer
	str := ("type@=joingroup/rid@=532152/gid@=91/")
	//str := ("type@=loginreq/roomid@=" + roomid)
	l := len(str) + 4 + 4 + 4 + 1
	msg.len = uint32(l)
	msg.code = uint32(l)
	msg.magic = uint16(689)
	msg.magic2 = uint16(0)
	binary.Write(&buffer, binary.LittleEndian, msg)
	buffer.WriteString(str)
	buffer.WriteByte(0)
	fmt.Println(buffer)
	return buffer
}
