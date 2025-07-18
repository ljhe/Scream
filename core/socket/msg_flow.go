package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ljhe/scream/core/iface"
	"github.com/ljhe/scream/def"
	"github.com/ljhe/scream/message"
	"io"
)

type MsgFlow struct {
	Packet iface.IDataPacket
}

func (m *MsgFlow) OnRcvMsg(s iface.ISession) (msg interface{}, err error) {
	opt := s.Node().(iface.IOption)
	opt.SocketReadTimeout(s, func() {
		msg, err = m.Packet.ReadMessage(s)
	})
	return
}

func (m *MsgFlow) OnSendMsg(s iface.ISession, msg interface{}) (err error) {
	opt := s.Node().(iface.IOption)
	opt.SocketWriteTimeout(s, func() {
		err = m.Packet.SendMessage(s, msg)
	})
	return err
}

type TcpDataPacket struct{}

func (t *TcpDataPacket) ReadMessage(s iface.ISession) (interface{}, error) {
	reader, ok := s.Conn().(io.Reader)
	if !ok || reader == nil {
		return nil, fmt.Errorf("TcpDataPacket ReadMessage get io.Reader err")
	}
	msg, msgId, err := message.RcvPackageData(reader)
	if err != nil {
		return nil, err
	}

	bt, err := message.DecodeMessage(msgId, msg)
	if err != nil {
		return nil, err
	}
	return bt, nil
}

func (t *TcpDataPacket) SendMessage(s iface.ISession, msg interface{}) (err error) {
	writer, ok := s.Conn().(io.Writer)
	if !ok || writer == nil {
		return fmt.Errorf("TcpDataPacket SendMessage get io.Writer err")
	}
	msgData, msgInfo, err := message.EncodeMessage(msg)
	if err != nil {
		return err
	}

	msgLen := len(msgData)
	mb := &message.MsgBase{
		MsgLen:    uint16(len(msgData)),
		MsgId:     msgInfo.ID,
		ChunkNum:  uint16(msgLen/def.MsgMaxLen + 1), // 计算分片数量
		ChunkId:   1,
		SendBytes: 0,
	}

	for mb.SendBytes < int(mb.MsgLen) {
		data := mb.Marshal(msgData)
		// 如果使用内存池 会导致每次发送的包里都会有空数据 所以写入的时候只写入有效数据的部分
		err = message.WriteFull(writer, data[:mb.ActualDataLen])
		if err != nil {
			return err
		}
		mb.SendBytes += mb.ChunkSize
		mb.ChunkId++
		mb.Release(data)
	}
	return nil
}

type WsDataPacket struct{}

func (w *WsDataPacket) ReadMessage(s iface.ISession) (interface{}, error) {
	conn, ok := s.Conn().(*websocket.Conn)
	if !ok || conn == nil {
		return nil, fmt.Errorf("WsDataPacket ReadMessage get websocket.Conn err")
	}
	typ, bt, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	switch typ {
	case websocket.BinaryMessage:
		msg, msgId, err := message.RcvPackageDataByByte(bt)
		if err != nil {
			return nil, err
		}
		bt, err := message.DecodeMessage(msgId, msg)
		if err != nil {
			return nil, err
		}
		return bt, nil
	default:
		return nil, fmt.Errorf("WsDataPacket ReadMessage type not binary message. typ:%v", typ)
	}

}

func (w *WsDataPacket) SendMessage(s iface.ISession, msg interface{}) (err error) {
	conn, ok := s.Conn().(*websocket.Conn)
	if !ok || conn == nil {
		return fmt.Errorf("WsDataPacket SendMessage get websocket.Conn err")
	}
	msgData, msgInfo, err := message.EncodeMessage(msg)
	if err != nil {
		return err
	}
	msgDataLen := len(msgData)
	// todo 注意上层发包不要超过最大值 之后这里可以改成如果超过最大值 就分片发送
	opt := s.Node().(iface.IOption)
	if msgDataLen > opt.GetMaxMsgLen() {
		return fmt.Errorf("ws sendMessage too big. msgId=%v msglen=%v maxlen=%v", 1, msgDataLen, opt.GetMaxMsgLen())
	}
	mb := &message.MsgBase{
		MsgId:  msgInfo.ID,
		FlagId: 1,
	}
	buf := mb.MarshalBytes(msgData)
	err = conn.WriteMessage(websocket.BinaryMessage, buf)
	return err
}
