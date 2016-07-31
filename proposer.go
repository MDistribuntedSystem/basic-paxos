package basicpaxos

type Proposer struct{}

func (proposer *Proposer) Prepare() {

}

func (proposer *Proposer) OnPrepareResponse(status bool, ab int, av int) {

}

func (proposer *Proposer) OnAcceptResponse(status bool) {

}
