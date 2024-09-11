package node

import (
	"distributed-file-storage/internal/storage"
	"distributed-file-storage/internal/network"
)

type Node struct {
	Storage  storage.Storage
	Network  *network.NetworkService
	// 其他节点相关字段
}

func NewNode() *Node {
	// 初始化节点
}

func (n *Node) Start() error {
	// 启动节点服务
}