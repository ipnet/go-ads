// Package ring
// @Description: 判断无向图是否有环
package ring

import "go-ads/pkg/graph"

type ring struct {
	g       graph.Graph
	hasRing bool
	debug   bool
}

func CreateRingAlgo(g graph.Graph) *ring {
	return &ring{
		g:     g,
		debug: false,
	}
}

func (r *ring) Calculate() {
	r.hasRing = false

	direct := r.g.IsDirect()
	if direct {
		dr := &directedRing{
			ring: r,
		}
		dr.calculate()
	} else {

	}
}

//
// HasRing
//  @Description: 是否有环
//  @receiver r
//  @return bool
//
func (r *ring) HasRing() bool {
	return r.hasRing
}

func (r *ring) SetDebug(debug bool) {
	r.debug = debug
}
