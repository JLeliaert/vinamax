//This file contains all the constants

package vinamax

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
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

	fmt.Println(`
vinamax: a macrospin model to simulate magnetic nanoparticles
Copyright (C) 2013  Jonathan Leliaert

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, version 3 of the License

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see [http:fmt.Println(www.gnu.org/licenses/].

contact: jonathan.leliaert@gmail.com


`)

	fname := os.Args[0]
	f2name := strings.Split(fname, "/")
	outdir = fmt.Sprint(f2name[len(f2name)-1], ".out")
	os.Mkdir(outdir, 0775)
	f, err3 := os.Open(outdir)
	files, _ := f.Readdir(1)
	// clean output dir, copied from mumax
	if len(files) != 0 {
		filepath.Walk(outdir, func(path string, i os.FileInfo, err error) error {
			if path != outdir {
				os.RemoveAll(path)
			}
			return nil
		})
	}

	check(err3)
}
