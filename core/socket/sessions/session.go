package sessions

import (
	"github.com/gorilla/websocket"
	"github.com/ljhe/scream/3rd/logrus"
	"github.com/ljhe/scream/core/iface"
	"github.com/ljhe/scream/core/socket"
	"log"
	"net"
	"reflect"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

const SessionMainSendQueueLen = 2000

type Session struct {
	*socket.Processor // 事件处理相关
	socket.ContextSet // 记录session绑定信息
	iface.ISessionExtension
	node            iface.INetNode
	close           int64
	sendQueue       chan interface{}
	sendQueueMaxLen int
	exitWg          sync.WaitGroup
	id              uint64
	rcvPingNum      int
	children        sync.Map
	mu              sync.Mutex
}

func (s *Session) SetId(id uint64) {
	s.id = id
}

func (s *Session) GetId() uint64 {
	return s.id
}

func (s *Session) Node() iface.INetNode {
	return s.node
}

func (s *Session) GetProcessor() iface.IProcessor {
	return s.Processor
}

func (s *Session) IncRcvPingNum(inc int) {
	if inc <= 0 {
		s.rcvPingNum = inc
		return
	}
	s.rcvPingNum += inc
}

func (s *Session) RcvPingNum() int {
	return s.rcvPingNum
}

func (s *Session) Start() {
	atomic.StoreInt64(&s.close, 0)
	s.exitWg.Add(2)
	// 添加到session管理器中
	s.node.(iface.ISessionManager).Add(s)
	go func() {
		s.exitWg.Wait()
		close(s.sendQueue)
		s.node.(iface.ISessionManager).Remove(s)
	}()
	go s.RunRcv()
	go s.RunSend()
}

func (s *Session) Close() {
	// 已经关闭
	if ok := atomic.SwapInt64(&s.close, 1); ok != 0 {
		return
	}
	s.ConnClose()
}

func (s *Session) RunRcv() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("tcpSession Stack---::%v\n %s\n", err, string(debug.Stack()))
			debug.PrintStack()
		}
	}()

	for {
		msg, err := s.ReadMsg(s)
		if err != nil {
			logrus.Errorf("RunRcv ReadMsg err:%v sessionId:%d", err, s.GetId())
			// 做关闭处理 发送数据时已经无法发送
			atomic.StoreInt64(&s.close, 1)
			select {
			case s.sendQueue <- nil:
			default:
				logrus.Errorf("RunRcv sendQueue block len:%d sessionId:%d", len(s.sendQueue), s.GetId())
			}

			// 抛出关闭事件
			s.CloseEvent(err)
			break
		}

		s.ProcEvent(&socket.RcvProcEvent{Sess: s, Message: msg})
	}
	s.exitWg.Done()
}

func (s *Session) Send(msg interface{}) {
	if atomic.LoadInt64(&s.close) != 0 {
		return
	}
	select {
	case s.sendQueue <- msg:
	default:
	}
}

func (s *Session) RunSend() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Stack---::%v\n %s\n", err, string(debug.Stack()))
		}
	}()

	for data := range s.sendQueue {
		if atomic.LoadInt64(&s.close) == 1 {
			break
		}
		if data == nil {
			continue
		}
		err := s.SendMsg(&socket.SendProcEvent{Sess: s, Message: data})
		if err != nil {
			logrus.Errorf("session send msg err:%v. sessionId:%d dataT:%v data:%v", err, s.GetId(), reflect.TypeOf(data), data)
			break
		}
	}

	s.ConnClose()
	s.exitWg.Done()
}

func (s *Session) ConnClose() {
	conn := s.Conn()
	switch conn.(type) {
	case net.Conn:
		conn.(net.Conn).Close()
	case *websocket.Conn:
		conn.(*websocket.Conn).Close()
	}
}

func NewSession(node iface.INetNode) *Session {
	return &Session{
		node:      node,
		sendQueue: make(chan interface{}, SessionMainSendQueueLen),
		Processor: node.(interface {
			GetProc() *socket.Processor
		}).GetProc(),
	}
}
