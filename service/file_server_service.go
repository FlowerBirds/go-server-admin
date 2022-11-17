package service

type FileServerService struct {
}

/**
 * 启动文件接收服务，并反馈接收的TCP端口
 */
func (s FileServerService) Start() int {

	go s.receiveFile()
	return 0
}

func (s FileServerService) receiveFile() {

}
