package common

const (
	SocketTypTcpAcceptor   = "tcpAcceptor"
	SocketTypTcpConnector  = "tcpConnector"
	SocketTypTcpWSAcceptor = "tcpWebSocketAcceptor"
)

const (
	ContextSetEtcdKey = "etcd_node"
	ContextSetCtxKey  = "ctx"
)

const MsgMaxLen = 1024 * 40 // 40k(发送和接受字节最大数量)

const (
	MsgEncryptionRSA = 1
)

// 服务器节点类型枚举
// 服务器类型节点Type:[1 gate] [2 game]
const (
	ServiceNodeTypeGate    = 1
	ServiceNodeTypeGateStr = "gate"

	ServiceNodeTypeGame    = 2
	ServiceNodeTypeGameStr = "game"
)
