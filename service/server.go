package service

import (
	"fmt"
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/command"
	"github.com/FlowerBirds/go-server-admin/config"
	"github.com/FlowerBirds/go-server-admin/handler"
	"github.com/FlowerBirds/go-server-admin/util"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type UdpServer struct {
	Port     int
	eventBus *base.LiveEventBus
	config   *config.UdpConfig
}

func NewUdpServer(serverConfig *config.UdpConfig) *UdpServer {
	return &UdpServer{
		Port:     serverConfig.Port,
		eventBus: base.NewEventBus(),
		config:   serverConfig,
	}
}

func (server *UdpServer) initCommand() {
	log.Println("init command")
	command.NewEchoCommand().AddToListener(server.eventBus)
	command.NewGetServerCommand(server.config).AddToListener(server.eventBus)
	command.NewSendCommand().AddToListener(server.eventBus)
	command.NewReceiveClientCommand(server.config).AddToListener(server.eventBus)
	command.NewDataBaseCommand().AddToListener(server.eventBus)
	command.NewEdgeMonitorCommand(server.config).AddToListener(server.eventBus)
	command.NewFeiQNotifyCommand(server.config).AddToListener(server.eventBus)
}

func (server *UdpServer) Start() {
	ip := util.GetIp()
	log.Printf("Try to start at %s:%d with version %s\n", ip, server.Port, server.config.Version)
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: server.Port,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Listen to udp server :" + strconv.Itoa(server.Port))
	defer listen.Close()
	server.config.InitDatabase()
	server.initCommand()
	server.initService()
	go server.initHttp()
	for {
		var data [1024 * 1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		receiveData := string(data[:n])
		if util.EnableDebug() {
			log.Println("receivedData data: " + receiveData + " at " + addr.String())
		}
		message, err := base.NewMessage(receiveData)
		if err != nil {
			log.Println(err)
		} else {
			evt := &base.Event{
				OpType:     message.OpType,
				InOut:      message.InOut,
				FromIp:     message.Ip,
				ToIP:       "",
				MessageStr: receiveData,
			}
			evt.MessageOri = message
			server.eventBus.PushEvent(evt)
		}
	}
}

/**
初始化服务相关功能
1.获取服务列表，间隔获取，获取到后则降低频率
2.
*/
func (server *UdpServer) initService() {
	log.Println("init service with server mode: " + strconv.FormatBool(server.config.IsServer))
	// 不管是监控端还是服务断，都需要寻找服务端
	// 当是服务端时，向其它服务器注册并发送相关信息，防止单点和监控信息丢失
	service := GetServerService{
		EventBus: *server.eventBus,
		Config:   server.config,
	}
	go service.Start()

	monitor := EdgeMonitorService{
		EventBus: *server.eventBus,
		Config:   server.config,
	}
	go monitor.Start()
}

/**
初始化Http服务，可以通过浏览器访问:2346展示当前服务收集的相关信息
1. ui页面，也是主页面
2. /list-client-message 获取监控端发送的消息列表
3. /list-servers 获取已经注册的服务器列表
4. /list-clients 获取已经注册到当前服务端的监控端列表
5.
*/
func (server *UdpServer) initHttp() {
	if server.config.IsServer {
		r := mux.NewRouter()
		staticDir := http.Dir(server.config.WorkDir + "/static/")
		log.Println("use static dir: " + staticDir)
		fs := http.FileServer(staticDir)
		uiHandler := http.StripPrefix("/ui", fs)
		r.PathPrefix("/ui").Handler(uiHandler).Methods(http.MethodGet)
		r.Handle("/", http.RedirectHandler("/ui/", http.StatusMovedPermanently)).Methods(http.MethodGet)
		r.Handle("/list-client-message", handler.MakeProxyHandler(handler.MakeListMessageHandler())).Methods(http.MethodPost)
		r.Handle("/list-servers", handler.MakeProxyHandler(handler.MakeListServerHandler(server.config))).Methods(http.MethodPost)
		r.Handle("/list-clients", handler.MakeProxyHandler(handler.MakeListClientHandler(server.config))).Methods(http.MethodPost)
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", server.config.HttpPort),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
			Handler:        r,
		}
		log.Printf("Listen to http server :%d\n", server.config.HttpPort)
		s.ListenAndServe()
	}
}
