// Package topology
// @Description: ζζζεΊ
package topology

import "go-ads/pkg/graph"

type topology struct {
	g graph.Graph
}

func CreateTopologyAlgo(g graph.Graph) *topology {
	return &topology{
		g: g,
	}
}
