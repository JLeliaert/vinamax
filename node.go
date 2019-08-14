package vinamax

import (
	"log"
)

type node struct {
	origin   vector      //the origin of the cube
	diameter float64     //the diameter
	number   int64       //numberofparticles
	lijst    []*particle //lijst met alle particles
	volume   float64     //the total volume of all particles in the node
	m        vector      //magnetisation of the node
}

func ReturnParticle(num int) *particle {
	if num > len(Universe.lijst) {
		log.Fatal("there aren't that many particle")
	}

	return Universe.lijst[num]
}

//adds particle to node
func (n *node) add(p *particle) {
	n.lijst = append(n.lijst, p)
}
