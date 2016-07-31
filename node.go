package basicpaxos

import (
	"io"
	"net"
)

type NodeState int

const (
	//NodeStateDead means node is dead
	NodeStateDead NodeState = 0
	//NodeStateNormal means node is working
	NodeStateNormal NodeState = 1
	//NodeStateInit means node is starting
	NodeStateInit NodeState = 2
	//NodeStateSync means node is sync data from other node
	NodeStateSync NodeState = 3
	//NodeStatePause means node is offline when  pasuse it.
	NodeStatePause NodeState = 4
)

type Node struct {
	id   int
	addr net.Addr
	io.Writer
	io.Reader
	backends []net.TCPConn
}

/**** State Machine Operator ********/

func (node *Node) GetStatus() NodeState {

	return NodeStateNormal
}

func (node *Node) UpdateState(state NodeState) {}

/**** Network Operator *********/

func (node *Node) Ping() bool {

	return true
}

func (node *Node) Write(data []byte) (int, error) {

	return 0, nil
}

func (node *Node) Read(dest []byte) (int, error) {
	return 0, nil
}

func (node *Node) watch(conn *net.Conn) {

	for {

	}

}

func (node *Node) onFrame(frame *Frame) {

}
