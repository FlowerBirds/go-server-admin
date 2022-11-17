package main

import (
	"flag"
	"fmt"
	"github.com/FlowerBirds/go-server-admin/config"
	"github.com/FlowerBirds/go-server-admin/service"
	"github.com/FlowerBirds/go-server-admin/util"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

var (
	udpServer  = false
	udpClient  = false
	toIp       = ""
	toPort     = 2345
	help       = false
	data       = ""
	dataDir    = ""
	debug      = false
	httpPort   = config.DEFAULT_HTTP_PORT
	configFile = ""
	devMode    = false
	workDir    = "."
	version    = "0.0.1"
)

func initCommandLine() {
	flag.BoolVar(&help, "help", false, "--help true or false, print usage")
	flag.BoolVar(&devMode, "dev.mode", false, "--dev.mode true or false, dev mode")
	flag.BoolVar(&udpServer, "udp.server", false, "--udp.server true or false, use server to collect data")
	flag.BoolVar(&udpClient, "udp.client", false, "--udp.client true or false, use send data by udp")
	flag.StringVar(&toIp, "to.ip", "127.0.0.1", "--to.ip string value, ip format, use send data to ip")
	flag.IntVar(&toPort, "to.port", config.DEFAULT_UDP_PORT, "--to.port int value, use send data to port")
	flag.StringVar(&data, "data", "", "--data string value, use send data")
	flag.StringVar(&dataDir, "data.dir", "data", "--data string value, store data")
	flag.BoolVar(&debug, "debug", false, "--debug bool value, enable debug and print info")
	flag.IntVar(&httpPort, "http.port", config.DEFAULT_HTTP_PORT, "--http.port int value, http port for server")
	flag.StringVar(&configFile, "config.file", "server.properties", "--config.file string value, config file")
	flag.Parse()
}

func initConfigFile() {
	viper.SetConfigName("server")
	viper.SetConfigType("properties")
	viper.AddConfigPath(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(fmt.Errorf("Fatal error config file: %w \n", err))
		return
	}
	log.Println("read http.port from config file: " + viper.Get("http.port").(string))
	port := viper.Get("http.port")
	if port != nil {
		p, err := strconv.Atoi(port.(string))
		if err == nil {
			httpPort = p
		}
	}
	configDebug := viper.Get("debug")
	log.Println("read debug from config file: " + viper.Get("debug").(string))
	if configDebug != nil {
		b, err := strconv.ParseBool(configDebug.(string))
		if err == nil {
			debug = b
		}
	}
	if !devMode {
		confDataDir := viper.Get("data.dir")
		log.Println("read data.dir from config file: " + confDataDir.(string))
		dataDir = confDataDir.(string)
	}
	server := viper.Get("udp.server")
	log.Println("read udp.server from config file: " + viper.Get("udp.server").(string))
	if server != nil {
		s, err := strconv.ParseBool(server.(string))
		if err == nil {
			udpServer = s
		}
	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate | log.Ltime)
	initCommandLine()
	if help {
		flag.PrintDefaults()
		return
	}
	util.UpdateDebug(debug)
	if udpClient {
		client := &service.UdpClient{
			ToIp:   toIp,
			ToPort: toPort,
			Data:   data,
		}
		client.Start()
		return
	}
	dir, err := util.GetCurrentPath()
	if err == nil && !devMode {
		workDir = dir
		log.Println("current work dir: " + dir)
	}
	// 只有在服务模式生效，client模式下不生效
	initConfigFile()
	util.UpdateDebug(debug)
	serverConfig := &config.UdpConfig{
		Port:     config.DEFAULT_UDP_PORT,
		IsServer: udpServer,
		Servers:  make(map[string]int64),
		DataDir:  dataDir,
		HttpPort: httpPort,
		WorkDir:  workDir,
		Version:  version,
	}
	server := service.NewUdpServer(serverConfig)
	server.Start()
}
