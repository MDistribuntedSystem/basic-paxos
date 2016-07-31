package basicpaxos

type NetOpt int32

const (

	//proposer send prepare request
	NetOptPrepare NetOpt = 1
	//acceptor receive prepare request
	NetOptPrePareResponse NetOpt = 2
	//proposer send accept request
	NetOptAccept NetOpt = 3
	//acceptor receive accept request
	NetOptAcceptResponse NetOpt = 4

	//proposer checking acceptor
	NetOptPing NetOpt = 100
	//proposer sync data from other proposer
	NetOptSync NetOpt = 200
	//node dead
	NetOptShutdown NetOpt = 300

	NetOptRequest NetOpt = 9999
)

const (
	ProtocolVersion int32 = 1
)

/***
protocol bits layout
| Version (4bit) | NodeID (4 bit) | Opt (4 bit) | BodyLen (4 bit) | N Bit Body |
***/

type Protocol struct {
	Version int32
	NodeID  int32
	Opt     NetOpt
	BodyLen int32
}

func (protocol *Protocol) Init() {
	protocol.Version = ProtocolVersion
}

type Frame struct {
	protcol Protocol
	Body    []byte
}
