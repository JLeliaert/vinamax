//This file contains all the constants

package vinamax

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (

	gamma0 = 1.7595e11          // Gyromagnetic ratio of electron, in rad/Ts
	mu0    = 4 * math.Pi * 1e-7 // Permeability of vacuum in Tm/A
	muB    = 9.2740091523E-24   // Bohr magneton in J/T
	kb     = 1.380650424E-23    // Boltzmann's constant in J/K
	qe     = 1.60217646E-19     // Electron charge in C
)

func init() {

	fmt.Println()
	fmt.Println("vinamax: a macrospin model to simulate magnetic nanoparticles")
	fmt.Println("Copyright (C) 2013  Jonathan Leliaert")
	fmt.Println()
	fmt.Println("This program is free software: you can redistribute it and/or modify")
	fmt.Println("it under the terms of the GNU General Public License as published by")
	fmt.Println("the Free Software Foundation, version 3 of the License")
	fmt.Println()
	fmt.Println("This program is distributed in the hope that it will be useful,")
	fmt.Println("but WITHOUT ANY WARRANTY; without even the implied warranty of")
	fmt.Println("MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the")
	fmt.Println("GNU General Public License for more details.")
	fmt.Println()
	fmt.Println("You should have received a copy of the GNU General Public License")
	fmt.Println("along with this program.  If not, see [http:fmt.Println(www.gnu.org/licenses/].")
	fmt.Println()
	fmt.Println("contact: jonathan.leliaert@gmail.com")
	fmt.Println()
	fmt.Println()
	fmt.Println()


fname := os.Args[0] 
f2name := strings.Split(fname, "/")
outdir =fmt.Sprint(f2name[len(f2name)-1],".out")
os.Mkdir(outdir, 0775)

}
