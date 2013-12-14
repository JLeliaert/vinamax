package vinamax

import (
//	"fmt"
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

	origin   Vector      //the origin of the cube
	diameter float64     //the diameter
	com      Vector      //centreofmass
	number   float64     //numberofparticles
	lijst    []*Particle //lijst met alle particles
}

func (n *node) add(p *Particle) {
	n.lijst = append(n.lijst, p)
}

func (n *node) calculatecom() {
	//TODO: hiermee oppassen als de deeltjes een verschillende grootte kunnen hebben
	comx := 0.
	comy := 0.
	comz := 0.

	for i := range n.lijst {
		comx += n.lijst[i].X
		comy += n.lijst[i].Y
		comz += n.lijst[i].Z
	}
	n.com = Vector{comx / n.number, comy / n.number, comz / n.number}
}

func (w *node) descend() {
	w.calculatecom()
	if w.number > 1 {

		//initialiseert de 8 subnodes

		//tlb
		pos := Vector{w.origin[0] - w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] + w.diameter/4}
		w.tlb = new(node)
		w.tlb.origin = pos
		w.tlb.diameter = w.diameter / 2

		//tlf
		pos = Vector{w.origin[0] - w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] - w.diameter/4}
		w.tlf = new(node)
		w.tlf.origin = pos
		w.tlf.diameter = w.diameter / 2

		//trb
		pos = Vector{w.origin[0] + w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] + w.diameter/4}
		w.trb = new(node)
		w.trb.origin = pos
		w.trb.diameter = w.diameter / 2

		//trf
		pos = Vector{w.origin[0] + w.diameter/4, w.origin[1] + w.diameter/4, w.origin[2] - w.diameter/4}
		w.trf = new(node)
		w.trf.origin = pos
		w.trf.diameter = w.diameter / 2

		//blb
		pos = Vector{w.origin[0] - w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] + w.diameter/4}
		w.blb = new(node)
		w.blb.origin = pos
		w.blb.diameter = w.diameter / 2

		//blf
		pos = Vector{w.origin[0] - w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] - w.diameter/4}
		w.blf = new(node)
		w.blf.origin = pos
		w.blf.diameter = w.diameter / 2

		//brb
		pos = Vector{w.origin[0] + w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] + w.diameter/4}
		w.brb = new(node)
		w.brb.origin = pos
		w.brb.diameter = w.diameter / 2

		//brf
		pos = Vector{w.origin[0] + w.diameter/4, w.origin[1] - w.diameter/4, w.origin[2] - w.diameter/4}
		w.brf = new(node)
		w.brf.origin = pos
		w.brf.diameter = w.diameter / 2

		//for alle particles in node
		for i := range w.lijst {
			plaats := w.where(Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z})
			switch plaats {
			case 0:
				w.tlb.number += 1
				w.tlb.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at tlb")
			case 1:
				w.tlf.number += 1
				w.tlf.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at tlf")

			case 2:
				w.trb.number += 1
				w.trb.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at trb")

			case 3:
				w.trf.number += 1
				w.trf.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at trf")

			case 4:
				w.blb.number += 1
				w.blb.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at blb")

			case 5:
				w.blf.number += 1
				w.blf.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at blf")

			case 6:
				w.brb.number += 1
				w.brb.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at brb")

			case 7:
				w.brf.number += 1
				w.brf.add(w.lijst[i])
				//fmt.Println("particle at ",Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z},"  was put at brf")
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

func Maketree() {
	//werkt dit??
	for i := range Lijst {
		Universe.lijst = append(Universe.lijst, &Lijst[i])
		Universe.number += 1
	}
	Universe.descend()
}

func (n node) where(position Vector) int {
	//if not in node
	if position[0]> n.origin[0]+n.diameter/2 || position[0]< n.origin[0]-n.diameter/2 || position[1]> n.origin[1]+n.diameter/2 || position[1]< n.origin[1]-n.diameter/2  || position[2]> n.origin[2]+n.diameter/2 || position[2]< n.origin[2]-n.diameter/2 {

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
