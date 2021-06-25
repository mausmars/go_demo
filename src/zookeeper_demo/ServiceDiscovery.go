package discovery

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

//PERSISTEN, PERSISTENT_SEQUENTAIL, EPHEMERAL, EPHEMERAL_SEQUENTAIL

/**
//flags有4种取值：
//0:永久，除非手动删除
//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
//zk.FlagSequence  = 2:会自动在节点后面添加序号
//3:Ephemeral和Sequence，即，短暂且自动添加序号
*/

type NodeData struct {
	NodeName    string `json:"nodeName"`
	Host        string `json:"host"`
	Port        int32  `json:"port"`
	ServiceName string `json:"serviceName"`

	Service interface{} `json:"-"`
	version int32       `json:"-"`
}

const ElectNodeName = "leader"

type IServiceDiscovery interface {
	Startup()
	Shutdown()
	Register(nodeData *NodeData)
	Unregister(path string)
	ModifyData(nodeData *NodeData)
	GetNodeData(nodeName string) *NodeData
	GetZKConn() *zk.Conn
}

type ServiceDiscoveryZK struct {
	RootPath       string
	Hosts          []string
	Acls           []zk.ACL
	SessionTimeout int

	leader     *NodeData
	nodes      *sync.Map
	localNodes *sync.Map
	conn       *zk.Conn
}

func (s *ServiceDiscoveryZK) Startup() {
	s.defaulConfig()

	//option := zk.WithEventCallback(s.callback)
	conn, _, err := zk.Connect(s.Hosts, time.Duration(s.SessionTimeout)*time.Millisecond)
	if err != nil {
		panic(err)
		return
	}
	s.conn = conn
	s.init()
}

func (s *ServiceDiscoveryZK) defaulConfig() {
	if s.RootPath == "" {
		s.RootPath = "/game"
	}
	if s.Acls == nil {
		s.Acls = zk.WorldACL(zk.PermAll)
	}
	if s.Hosts == nil {
		s.Hosts = []string{"127.0.0.1:2181"}
	}
	if s.SessionTimeout <= 0 {
		s.SessionTimeout = 5000
	}
	s.nodes = &sync.Map{}
	s.localNodes = &sync.Map{}
}

func (s *ServiceDiscoveryZK) init() {
	isExists, _, err := s.conn.Exists(s.RootPath)
	if err != nil {
		panic(err)
		return
	}
	if !isExists {
		//创建根节点
		_, err := s.conn.Create(s.RootPath, nil, 0, s.Acls)
		if err != nil {
			panic(err)
			return
		}
	}
	// 获取子节点并更新
	childrenNodes, _, ech, err := s.conn.ChildrenW(s.RootPath)
	if err != nil {
		panic(err)
		return
	}
	go s.watchCreataNode(ech)
	for _, node := range childrenNodes {
		if node == ElectNodeName {
			// 跳过leader结点
			continue
		}
		data, _, ech, err := s.conn.GetW(s.getPath(node))
		if err != nil {
			panic(err)
			return
		}
		go s.watchCreataNode(ech)
		nodeData := s.decode(data)
		s.nodes.Store(nodeData.NodeName, nodeData)
	}
	{
		//单独监听leader结点
		masterPath := s.getPath(ElectNodeName)
		s.setLeaderNode(masterPath)
	}
}

func (s *ServiceDiscoveryZK) setLeaderNode(masterPath string) {
	isExists, _, ech, err := s.conn.ExistsW(masterPath)
	if err != nil {
		panic(err)
		return
	}
	go s.watchCreataNode(ech)

	if isExists {
		data, stat, err := s.conn.Get(masterPath)
		if err != nil {
			return
		}
		if s.leader == nil {
			s.leader = s.decode(data)
			s.leader.version = stat.Version
		} else {
			if stat.Version > s.leader.version {
				s.leader = s.decode(data)
				s.leader.version = stat.Version
			}
		}
	}
}

func (s *ServiceDiscoveryZK) Shutdown() {
	s.conn.Close()
}

func (s *ServiceDiscoveryZK) Register(nodeData *NodeData) {
	s.localNodes.Store(nodeData.NodeName, nodeData)

	nodePath := s.getPath(nodeData.NodeName)
	isExists, _, err := s.conn.Exists(nodePath)
	if err != nil {
		panic(err)
		return
	}
	if isExists {
		s.conn.Delete(nodePath, -1)
	}
	data := s.encoded(nodeData)
	_, err = s.conn.Create(nodePath, data, zk.FlagEphemeral, s.Acls)
	if err != nil {
		panic(err)
		return
	}
	// 如果leader为空，就进行一下选举
	if s.leader == nil {
		masterPath := s.getPath(ElectNodeName)
		_, err = s.conn.Create(masterPath, data, zk.FlagEphemeral, s.Acls)
		if err != nil {
			return
		}
	}
}

func (s *ServiceDiscoveryZK) Unregister(nodeName string) {
	s.localNodes.Delete(nodeName)
	s.conn.Delete(s.getPath(nodeName), -1)

	if s.leader != nil && s.leader.NodeName == nodeName {
		masterPath := s.getPath(ElectNodeName)
		s.conn.Delete(masterPath, -1)
	}
}

func (s *ServiceDiscoveryZK) ModifyData(nodeData *NodeData) {
	s.localNodes.Store(nodeData.NodeName, nodeData)

	nodePath := s.getPath(nodeData.NodeName)
	isExists, _, err := s.conn.Exists(nodePath)
	if err != nil {
		panic(err)
		return
	}
	data := s.encoded(nodeData)
	if isExists {
		s.conn.Set(nodePath, data, -1)
	}
	// 修改leader数据
	if s.leader != nil && s.leader.NodeName == nodeData.NodeName {
		masterPath := s.getPath(ElectNodeName)
		s.conn.Set(masterPath, data, -1)
	}
}

func (s *ServiceDiscoveryZK) GetNodeData(nodeName string) *NodeData {
	nodeData, ok := s.nodes.Load(nodeName)
	if ok {
		return nodeData.(*NodeData)
	} else {
		return nil
	}
}

func (s *ServiceDiscoveryZK) GetZKConn() *zk.Conn {
	return s.conn
}

func (s *ServiceDiscoveryZK) decode(data []byte) *NodeData {
	nodeData := &NodeData{}
	json.Unmarshal(data, nodeData)
	return nodeData
}

func (s *ServiceDiscoveryZK) encoded(nodeData *NodeData) []byte {
	data, _ := json.Marshal(nodeData)
	return data
}

func (s *ServiceDiscoveryZK) getPath(nodeName string) string {
	return s.RootPath + "/" + nodeName
}

func (s *ServiceDiscoveryZK) getNodeName(path string) string {
	return path[len(s.RootPath+"/"):]
}

func (s *ServiceDiscoveryZK) callback(event zk.Event) {
	fmt.Println("path:", event.Path, "type:", event.Type.String(), "state:", event.State.String())
}

func (s *ServiceDiscoveryZK) watchCreataNode(ech <-chan zk.Event) {
	select {
	case event := <-ech:
		fmt.Println("path:", event.Path, "type:", event.Type.String(), "state:", event.State.String(), "server", event.Server)

		if event.Type == zk.EventNodeChildrenChanged {
			childrenNodes, _, ech, err := s.conn.ChildrenW(s.RootPath)
			if err != nil {
				panic(err)
				return
			}
			go s.watchCreataNode(ech)

			for _, nodeName := range childrenNodes {
				if nodeName == ElectNodeName {
					// 跳过leader结点
					continue
				}
				if _, ok := s.nodes.Load(nodeName); !ok {
					//新产生的
					fmt.Println("发现新节点 node=", nodeName)
					data, _, ech, err := s.conn.GetW(s.getPath(nodeName))
					if err != nil {
						panic(err)
						return
					}
					go s.watchCreataNode(ech)
					nodeData := s.decode(data)
					s.nodes.Store(nodeData.NodeName, nodeData)
				}
			}
		} else if event.Type == zk.EventNodeDeleted {
			nodeName := s.getNodeName(event.Path)
			if nodeName == ElectNodeName {
				// 清除leader
				s.leader = nil
				// 注册选举结点的监听
				fmt.Println("刪除leader結點，并注册 node=", nodeName)
				_, _, ech, err := s.conn.ExistsW(event.Path)
				if err != nil {
					panic(err)
					return
				}
				go s.watchCreataNode(ech)
				// 选举
				localNode := s.getLocalNode()
				if localNode == nil {
					return
				}
				// 如果被删除了就选一下
				data := s.encoded(localNode)
				_, err = s.conn.Create(event.Path, data, zk.FlagEphemeral, s.Acls)
				if err != nil {
					panic(err)
					return
				}
			} else {
				s.nodes.Delete(nodeName)
				fmt.Println("刪除結點 node=", nodeName)
			}
		} else if event.Type == zk.EventNodeDataChanged {
			fmt.Println("结点数据变更 node=", event.Path)
			_, _, ech, err := s.conn.ExistsW(event.Path)
			if err != nil {
				panic(err)
				return
			}
			go s.watchCreataNode(ech)

			data, _, err := s.conn.Get(event.Path)
			if err != nil {
				panic(err)
				return
			}
			nodeData := s.decode(data)
			s.nodes.Store(nodeData.NodeName, nodeData)
		} else if event.Type == zk.EventNodeCreated {
			fmt.Println("新加結點 node=", event.Path)
			nodeName := s.getNodeName(event.Path)
			if nodeName == ElectNodeName {
				//设置leader节点
				s.setLeaderNode(event.Path)
				fmt.Println("新加leader結點 node=", nodeName)
			}
		}
	}
}

func (s *ServiceDiscoveryZK) getLocalNode() *NodeData {
	var v *NodeData = nil
	s.localNodes.Range(func(key, value interface{}) bool {
		v = value.(*NodeData)
		return false
	})
	return v
}

func (s *ServiceDiscoveryZK) PrintNodes() {
	fmt.Println("=== leaderNodes ===")
	if s.leader != nil {
		fmt.Println(*s.leader)
	}

	fmt.Println("=== localNodes ===")
	s.localNodes.Range(func(key, value interface{}) bool {
		fmt.Println("{", key, " ", *value.(*NodeData), "}")
		return true
	})
	fmt.Println("=== nodes ===")
	s.nodes.Range(func(key, value interface{}) bool {
		fmt.Println("{", key, " ", *value.(*NodeData), "}")
		return true
	})
}
