package basicpaxos

//NodeConfig defined all nodes.
// why nodes is a map ?
// beacuse a node must include a unqiue id
// i think  map key is a better idea to promise id is unique!
//example:
/**
{
  "nodes" : {

     1 : {
       "host" : "127.0.0.1",
       "port" : 34789
     },

     2 : {
       "host" : "127.0.0.1",
       "port" : 34790
     }
  }
}
**/

type NodeConfig struct {
	Nodes map[int]struct {
		Host string
		Port int
	} `json:"nodes"`
}
