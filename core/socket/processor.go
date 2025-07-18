package socket

import (
	"fmt"
	"github.com/ljhe/scream/core/iface"
	"log"
	"reflect"
	"time"
)

type Processor struct {
	MsgFlow   iface.IMsgFlow      // 处理消息读写
	Hooker    iface.IHookEvent    // 不进入主消息队列 直接操作
	MsgHandle iface.IMsgHandle    // 消息主队列
	MsgRouter iface.EventCallBack // 路由
}

func (n *Processor) SetMsgFlow(v iface.IMsgFlow) {
	n.MsgFlow = v
}

func (n *Processor) SetHooker(v iface.IHookEvent) {
	n.Hooker = v
}

func (n *Processor) SetMsgHandle(v iface.IMsgHandle) {
	n.MsgHandle = v
}

func (n *Processor) SetMsgRouter(msgr iface.EventCallBack) {
	n.MsgRouter = msgr
}

func (n *Processor) GetMsgRouter() iface.EventCallBack {
	return n.MsgRouter
}

func (n *Processor) GetProc() *Processor {
	return n
}

func (n *Processor) ProcEvent(e iface.IProcEvent) {
	if n.Hooker != nil {
		e = n.Hooker.InEvent(e)
	}
	if e != nil {
		switch n.Hooker.(type) {
		case *SessionChildHookEvent:
			e.Session().GetProcessor().GetMsgRouter()(e)
		default:
			if n.MsgHandle != nil {
				n.MsgHandle.PostCb(func() {
					start := time.Now()
					n.MsgRouter(e)
					duration := time.Since(start)
					log.Printf("%+v 方法 耗时: %s (%dμs / %dns)\n", reflect.TypeOf(e.Msg()), duration, duration.Microseconds(), duration.Nanoseconds())
				})
			}
		}
	}
}

func (n *Processor) ReadMsg(s iface.ISession) (interface{}, error) {
	if n.MsgFlow != nil {
		return n.MsgFlow.OnRcvMsg(s)
	}
	return nil, fmt.Errorf("message rpc is nil")
}

func (n *Processor) SendMsg(e iface.IProcEvent) error {
	if n.Hooker != nil {
		e = n.Hooker.OutEvent(e)
	}
	if n.MsgFlow != nil {
		return n.MsgFlow.OnSendMsg(e.Session(), e.Msg())
	}
	return nil
}
