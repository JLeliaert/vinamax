package vinamax

import (
//	"fmt"
	"math"
)

type node struct {
	//8 pointers to subnodes (if necessary)
	//top/bottom left/right front/back
	tlb *node
	tlf *node
	trb *node
	trf *node
	blb *node
	blf *node
	brb *node
	brf *node

	origin   vector      //the origin of the cube
	diameter float64     //the diameter
	com      vector      //centreofmagnetisation
	number   int64       //numberofparticles
	lijst    []*particle //lijst met alle particles
	volume   float64     //the total volume of all particles in the node
	m        vector      //magnetisation of the node
}

//adds particle to node
func (n *node) add(p *particle) {
	n.lijst = append(n.lijst, p)
}

//center of magnetisation
func (n *node) calculatecom() {
	comx := 0.
	comy := 0.
	comz := 0.
	total := 0.
	prefactor := 0.

	for i := range n.lijst {
		prefactor = n.lijst[i].msat * cube(n.lijst[i].r)
		comx += n.lijst[i].x * prefactor
		comy += n.lijst[i].y * prefactor
		comz += n.lijst[i].z * prefactor
		total += prefactor

	}
	n.com = vector{comx / total, comy / total, comz / total}
}

//descends into the tree, needed for Maketree()
func (w *node) descend() {
	w.calculatecom()
	if w.number > 1 {

		//initialiseert de 8 subnodes

		//tlb
		pos := vector{w.origin[0] - w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] + w.diameter/4}
		w.tlb = new(node)
		w.tlb.origin = pos
		w.tlb.diameter = w.diameter / 2

		//tlf
		pos = vector{w.origin[0] - w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] - w.diameter/4}
		w.tlf = new(node)
		w.tlf.origin = pos
		w.tlf.diameter = w.diameter / 2

		//trb
		pos = vector{w.origin[0] + w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] + w.diameter/4}
		w.trb = new(node)
		w.trb.origin = pos
		w.trb.diameter = w.diameter / 2

		//trf
		pos = vector{w.origin[0] + w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] - w.diameter/4}
		w.trf = new(node)
		w.trf.origin = pos
		w.trf.diameter = w.diameter / 2

		//blb
		pos = vector{w.origin[0] - w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] + w.diameter/4}
		w.blb = new(node)
		w.blb.origin = pos
		w.blb.diameter = w.diameter / 2

		//blf
		pos = vector{w.origin[0] - w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] - w.diameter/4}
		w.blf = new(node)
		w.blf.origin = pos
		w.blf.diameter = w.diameter / 2

		//brb
		pos = vector{w.origin[0] + w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] + w.diameter/4}
		w.brb = new(node)
		w.brb.origin = pos
		w.brb.diameter = w.diameter / 2

		//brf
		pos = vector{w.origin[0] + w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] - w.diameter/4}
		w.brf = new(node)
		w.brf.origin = pos
		w.brf.diameter = w.diameter / 2

		//for alle particles in node
		for i := range w.lijst {
			plaats := w.where(vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z})
			switch plaats {
			case 0:
				w.tlb.number += 1
				w.tlb.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at tlb")
			case 1:
				w.tlf.number += 1
				w.tlf.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at tlf")

			case 2:
				w.trb.number += 1
				w.trb.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at trb")

			case 3:
				w.trf.number += 1
				w.trf.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at trf")

			case 4:
				w.blb.number += 1
				w.blb.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at blb")

			case 5:
				w.blf.number += 1
				w.blf.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at blf")

			case 6:
				w.brb.number += 1
				w.brb.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at brb")

			case 7:
				w.brf.number += 1
				w.brf.add(w.lijst[i])
				//fmt.Println("particle at ",vector{w.lijst[i].x, w.lijst[i].y, w.lijst[i].z},"  was put at brf")
			}
		}
		//for iedere subnode
		//oeveel zitten derin?

		if w.tlb.number > 1 {
			w.tlb.descend()
		}
		if w.tlf.number > 1 {
			w.tlf.descend()
		}
		if w.trb.number > 1 {
			w.trb.descend()
		}
		if w.trf.number > 1 {
			w.trf.descend()
		}
		if w.blb.number > 1 {
			w.blb.descend()
		}
		if w.blf.number > 1 {
			w.blf.descend()
		}
		if w.brb.number > 1 {
			w.brb.descend()
		}
		if w.brf.number > 1 {
			w.brf.descend()
		}
	}
}

//Build the tree needed for the FMM method, descends in the "universe" node
func Maketree() {
	treecalled = true
	universe.descend()
	universe.calculatevolume()
}

//returns the position of a particle in a node (in terms of subnode position), or -1 if the particle is not in the node
func (n node) where(position vector) int {
	//if not in node
	if position[0] > n.origin[0]+n.diameter/2 || position[0] < n.origin[0]-n.diameter/2 || position[1] > n.origin[1]+n.diameter/2 || position[1] < n.origin[1]-n.diameter/2 || position[2] > n.origin[2]+n.diameter/2 || position[2] < n.origin[2]-n.diameter/2 {

		return -1
	}

	if position[0] >= n.origin[0] && position[1] >= n.origin[1] && position[2] >= n.origin[2] {
		//trb
		return 2
	}

	if position[0] >= n.origin[0] && position[1] >= n.origin[1] && position[2] < n.origin[2] {
		//trf
		return 3
	}

	if position[0] >= n.origin[0] && position[1] < n.origin[1] && position[2] >= n.origin[2] {
		//brb
		return 6
	}

	if position[0] >= n.origin[0] && position[1] < n.origin[1] && position[2] < n.origin[2] {
		//brf
		return 7
	}

	if position[0] < n.origin[0] && position[1] >= n.origin[1] && position[2] >= n.origin[2] {
		//tlb
		return 0
	}

	if position[0] < n.origin[0] && position[1] >= n.origin[1] && position[2] < n.origin[2] {
		//tlf
		return 1
	}

	if position[0] < n.origin[0] && position[1] < n.origin[1] && position[2] >= n.origin[2] {
		//blb
		return 4
	}

	if position[0] < n.origin[0] && position[1] < n.origin[1] && position[2] < n.origin[2] {
		//blf
		return 5
	}
	return -1
}

//calculates the volumes
func (w *node) calculatevolume() {
	if w.number > 1 {
		//for every subnode

		w.tlb.calculatevolume()
		w.volume += w.tlb.volume
		w.tlf.calculatevolume()
		w.volume += w.tlf.volume
		w.trb.calculatevolume()
		w.volume += w.trb.volume
		w.trf.calculatevolume()
		w.volume += w.trf.volume
		w.blb.calculatevolume()
		w.volume += w.blb.volume
		w.blf.calculatevolume()
		w.volume += w.blf.volume
		w.brb.calculatevolume()
		w.volume += w.brb.volume
		w.brf.calculatevolume()
		w.volume += w.brf.volume

	}
	if w.number == 1 {
		w.volume = 4. / 3. * math.Pi * cube(w.lijst[0].r)
	}
	if w.number == 0 {
		w.volume = 0.
	}
}

//calculates the magnetisation of a node
func (w *node) calculatem() {
	switch w.number {
	case 0:
		w.m = vector{0., 0., 0.}
	case 1:
		w.m = w.lijst[0].m.times(w.lijst[0].msat * w.volume)
	default:
		w.m = vector{0., 0., 0.}
		//for every subnode
		if w.tlb.number >0{
		w.tlb.calculatem()
		w.m = w.m.add(w.tlb.m)
		}
		if w.tlb.number >0{
		w.tlf.calculatem()
		w.m = w.m.add(w.tlf.m)
		}
		if w.tlb.number >0{
		w.trb.calculatem()
		w.m = w.m.add(w.trb.m)
		}
		if w.tlb.number >0{
		w.trf.calculatem()
		w.m = w.m.add(w.trf.m)
		}
		if w.tlb.number >0{
		w.blb.calculatem()
		w.m = w.m.add(w.blb.m)
		}
		if w.tlb.number >0{
		w.blf.calculatem()
		w.m = w.m.add(w.blf.m)
		}
		if w.tlb.number >0{
		w.brb.calculatem()
		w.m = w.m.add(w.brb.m)
		}
		if w.tlb.number >0{
		w.brf.calculatem()
		w.m = w.m.add(w.brf.m)
		}
	}
}
