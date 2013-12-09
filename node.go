package vinamax

import (
	"fmt"
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

func (n node) add(p *Particle) {
	n.lijst = append(n.lijst, p)
}

func (n node) calculatecom() {
	//TODO
}

//COM van de nodes!!!!!!!!!!!!!!!!!!!!!!!

func (w node) descend() {
	w.calculatecom()

	if w.number > 1 {

		//initialiseert de 8 subnodes

		//tlb
		pos := Vector{w.origin[0] - w.diameter/2, w.origin[1] + w.diameter/2, w.origin[2] + w.diameter/2}
		k := node{origin: pos, diameter: w.diameter / 2}
		w.tlb = &k

		//tlf
		pos = Vector{w.origin[0] - w.diameter/2, w.origin[1] - w.diameter/2, w.origin[2] + w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.tlf = &k

		//trb
		pos = Vector{w.origin[0] + w.diameter/2, w.origin[1] + w.diameter/2, w.origin[2] + w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.trb = &k

		//trf
		pos = Vector{w.origin[0] + w.diameter/2, w.origin[1] - w.diameter/2, w.origin[2] + w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.trf = &k

		//blb
		pos = Vector{w.origin[0] - w.diameter/2, w.origin[1] + w.diameter/2, w.origin[2] - w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.blb = &k

		//blf
		pos = Vector{w.origin[0] - w.diameter/2, w.origin[1] - w.diameter/2, w.origin[2] - w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.blf = &k

		//brb
		pos = Vector{w.origin[0] + w.diameter/2, w.origin[1] + w.diameter/2, w.origin[2] - w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.brb = &k

		//brf
		pos = Vector{w.origin[0] + w.diameter/2, w.origin[1] - w.diameter/2, w.origin[2] - w.diameter/2}
		k = node{origin: pos, diameter: w.diameter / 2}
		w.brf = &k

		//for alle particles in node
		for i := range w.lijst {
			//kijkt waardanze zittn
			// steekt een pointer in de juste node
			plaats := w.where(Vector{w.lijst[i].X, w.lijst[i].Y, w.lijst[i].Z})
			switch plaats {
			case 0:
				w.tlb.number += 1
				w.tlb.add(&Lijst[i])
			case 1:
				w.tlf.number += 1
				w.tlf.add(&Lijst[i])

			case 2:
				w.trb.number += 1
				w.trb.add(&Lijst[i])

			case 3:
				w.trf.number += 1
				w.trf.add(&Lijst[i])

			case 4:
				w.blb.number += 1
				w.blb.add(&Lijst[i])

			case 5:
				w.blf.number += 1
				w.blf.add(&Lijst[i])

			case 6:
				w.brb.number += 1
				w.brb.add(&Lijst[i])

			case 7:
				w.brf.number += 1
				w.brf.add(&Lijst[i])
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

func maketree() {
	//werkt dit??
	for i := range Lijst {
		Universe.lijst = append(Universe.lijst, &Lijst[i])
	}
	Universe.descend()
}

func (n node) where(position Vector) int {

	if position[0] >= n.origin[0] && position[1] >= n.origin[1] && position[2] >= n.origin[2] {
		//trb
		return 2
	}

	if position[0] >= n.origin[0] && position[1] >= n.origin[1] && position[2] < n.origin[2] {
		//brb
		return 6
	}

	if position[0] >= n.origin[0] && position[1] < n.origin[1] && position[2] >= n.origin[2] {
		//trf
		return 3
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
		//blb
		return 4
	}

	if position[0] < n.origin[0] && position[1] < n.origin[1] && position[2] >= n.origin[2] {
		//tlf
		return 1
	}

	if position[0] < n.origin[0] && position[1] < n.origin[1] && position[2] < n.origin[2] {
		//blf
		return 5
	}

	fmt.Println("something went wrong")
	return -1
}
