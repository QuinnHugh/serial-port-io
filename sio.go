package main

import (
	"flag"
	"log"
	"os"

	"github.com/larspensjo/config"
	//can't use this
	//serial "github.com/tarm/goserial"

	"github.com/jacobsa/go-serial/serial"
)

var (
	conFile = flag.String("configfile", "/config.ini", "config file")
)

func main() {
	//获取当前路径
	file, _ := os.Getwd()
	cfg, err := config.ReadDefault(file + *conFile)
	//获取配置文件中的配置项
	id, err := cfg.String("COM", "COMID")
	//设置串口
	options := serial.OpenOptions{
		PortName:        id,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	//打开串口
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	//command, err := cfg.String("COM", "COMMAND")
	// 写入货柜串口命令
	// log.Printf("货柜打开指令 %s", command)
	// n, err := port.Write([]byte(command))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	buf := make([]byte, 128)
	n, err := port.Read(buf)
	log.Printf("读取窗口信息 %s", buf[:n])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
	defer port.Close()

}
