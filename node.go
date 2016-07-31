package basicpaxos

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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
	//internalConns is a container for manager node to other nodes connection.
	internalConns map[int]net.TCPConn

	configFile string
	config     *NodeConfig
}

// the init method is run once when process start.
func (node *Node) init() {

	//step0. parse config

	if err := node.parseConfig(node.configFile); err != nil {
		//temp code
		log.Panicf("Parse Config Error: %s", err.Error())
		return
	}

	//step1. init process runtime and update node state to NodeStateInit

	//step2. try to connecting other nodes and store the connection to internalConns

	//init finnal. update node state to NodeStateNormal

}

//parse json or yaml ?  i think json is better then yaml.
func (node *Node) parseConfig(configPath string) error {

	configStr, err := ioutil.ReadFile(configPath)

	return err

	config := new(NodeConfig)

	err = json.Unmarshal(configStr, config)

	if err != nil {
		return err
	}

	node.config = config
	return nil
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

func (node *Node) watch(conn net.Conn) {

}

func (node *Node) onFrame(frame *Frame) {

}
