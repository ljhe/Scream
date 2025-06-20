package service

import (
	"fmt"
	"github.com/ljhe/scream/3rd/db/gorm"
	trdetcd "github.com/ljhe/scream/3rd/etcd"
	"github.com/ljhe/scream/3rd/logrus"
	"github.com/ljhe/scream/core/config"
	"github.com/ljhe/scream/core/iface"
	"os"
	"os/signal"
	"syscall"
)

func Init() error {
	// 加载系统配置文件
	config.Init()
	// 初始化日志模块
	logrus.Init(*config.ServerConfigPath)
	// 初始化内存池
	//mpool.MemoryPoolInit()
	// 初始化服务发现
	err := trdetcd.InitServiceDiscovery(config.SConf.Process.Node[0].Etcd)
	if err != nil {
		logrus.Errorf("InitServiceDiscovery err:%v", err)
		return err
	}
	// 初始化db
	config.OrmConnector = gorm.NewOrmConn()
	err = config.OrmConnector.Start("root:123456@(127.0.0.1:3306)/gamedb_9999?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		logrus.Errorf("init db err:%v", err)
		return err
	}
	return nil
}

func StartUp() {
	logrus.Infof(fmt.Sprintf("[ %s ] starting ...", config.SConf.Process.Node[0].Name))
	nodes := make([]iface.INetNode, 0)
	if config.SConf.Process.Node[0].Addr != "" {
		nodes = append(nodes, CreateAcceptor())
	}

	if config.SConf.Process.Node[0].WsAddr != "" {
		nodes = append(nodes, CreateWebSocketAcceptor())
	}

	for _, connect := range config.SConf.Process.Node[0].Connect {
		CreateConnector(connect)
	}
	logrus.Infof(fmt.Sprintf("[ %s ] start success ...", config.SConf.Process.Node[0].Name))

	WaitExitSignal()

	logrus.Infof(fmt.Sprintf("[ %s ] stoping ...", config.SConf.Process.Node[0].Name))
	for _, node := range nodes {
		Stop(node)
	}

	logrus.Infof(fmt.Sprintf("[ %s ] close ...", config.SConf.Process.Node[0].Name))
}

func WaitExitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	<-ch
}

func Stop(node iface.INetNode) {
	if node == nil {
		return
	}
	node.Stop()
	trdetcd.UnRegister(node)
}
