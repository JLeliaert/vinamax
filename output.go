//Contains function to control the output of the program
package vinamax

import (
	"fmt"
	"log"
	"math"
	"os"
)

var f *os.File
var err error
var twrite float64

var locations []vector
var filecounter int = 0

var output_u_anis = false
var output_energy = false

var outputList []outputQuantity

type outputQuantity interface {
	header() string
	value() string
}

//Sets the interval at which times the output table has to be written
func Output(interval float64) {
	if interval != 0 {
		outputcalled = true
		if Test == false {
			f, err = os.Create(outdir + "/table.txt")
			check(err)
			//	defer f.Close()
		}
		if Test == true {
			name := fmt.Sprintf("table%d.txt", Counter)
			f, err = os.OpenFile(outdir+"/"+name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			check(err)
			//defer file.Close()
			Counter += 1
		}
		writeheader()
	}
	outputinterval = interval
	twrite = interval

}

//checks the error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Writeintable(a string) {
	string := fmt.Sprintf("%v\n", a)
	_, err = f.WriteString(string)
	check(err)
}

func Tablesave() {
	if outputcalled == false {
		outputcalled = true
		f, err = os.Create(outdir + "/table.txt")
		check(err)
		writeheader()
	}
	write(averagemoments(), true)
}

//print position and magnitisation of a particle
func (p particle) string() string {
	return fmt.Sprintf("particle@(%v, %v, %v), %v %v %v", p.x, p.y, p.z, p.m[0], p.m[1], p.m[2])
}

//calculates the average magnetisation components of all particles
func averages() vector {
	avgs := vector{0, 0, 0}
	for _, p := range lijst {
		for i := 0; i < 3; i++ {
			avgs[i] += p.m[i]
		}
	}
	return avgs.times(1. / float64(len(lijst)))
}

//calculates the average anisotropy components of all particles
func averages_u() vector {
	avgs := vector{0, 0, 0}
	for _, p := range lijst {
		for i := 0; i < 3; i++ {
			if p.u_anis[0] < 0 {
				p.u_anis[i] = (-1) * p.u_anis[i]

			}
			avgs[i] += p.u_anis[i]
		}
	}
	avgs = avgs.times(1. / float64(len(lijst)))
	return avgs
}

//calculates the average moments of all particles
func averagemoments() vector {
	avgs := vector{0, 0, 0}
	totalmoment := 0.
	for _, p := range lijst {
		radius := p.rc
		volume := cube(radius) * 4. / 3. * math.Pi
		totalmoment += volume * p.msat
		for i := 0; i < 3; i++ {
			avgs[i] += p.msat * p.m[i] * volume
		}
	}
	//divide by total volume
	return avgs.times(1. / totalmoment)
}

//returns the number of particles with m_z larger than 0
func nrmzpositive() int {
	counter := 0
	for _, p := range lijst {
		if p.m[2] > 0. {
			counter++
		}
	}
	return counter
}

//Writes the header in table.txt
func writeheader() {
	header := fmt.Sprintf("#t\t<mx>\t<my>\t<mz>")
	_, err = f.WriteString(header)
	check(err)
	for _, o := range outputList {
		header := fmt.Sprintf(o.header())
		_, err = f.WriteString(header)
		check(err)
	}

	for i := range locations {

		header = fmt.Sprintf("\t(B_x\tB_y\tB_z)@(%v,%v,%v)", locations[i][0], locations[i][1], locations[i][2])
		_, err = f.WriteString(header)
		check(err)
	}

	header = fmt.Sprintf("\n")
	_, err = f.WriteString(header)
	check(err)

}

//Adds the field at a specific location to the output table
func Tableadd_b_at_location(x, y, z float64) {
	tableaddcalled = true
	if outputinterval != 0 {
		log.Fatal("Output() should always come AFTER Tableadd_b_at_location()")
	}
	locations = append(locations, vector{x, y, z})

}

func Give_mz() float64 {
	return averagemoments()[2]
}

//Writes the time and the vector of average magnetisation in the table
func write(avg vector, forced bool) {
	if forced || (twrite >= outputinterval && outputinterval != 0) {
		string := fmt.Sprintf("%e\t%v\t%v\t%v", T.value, avg[0], avg[1], avg[2])
		_, err = f.WriteString(string)
		check(err)

		for _, o := range outputList {
			_, err = f.WriteString(o.value())
			check(err)
		}

		for i := range locations {

			string = fmt.Sprintf("\t%v\t%v\t%v", (demag(locations[i][0], locations[i][1], locations[i][2])[0]), (demag(locations[i][0], locations[i][1], locations[i][2])[1]), (demag(locations[i][0], locations[i][1], locations[i][2])[2]))
			_, err = f.WriteString(string)
			check(err)
		}
		if !forced {
			_, err = f.WriteString("\n")
			check(err)
		}
		twrite = 0.
	}
	twrite += Dt.value
}

//Saves different quantities. At the moment only "geometry" and "m" are possible
func Save(a string) {
	//een file openen met unieke naam (counter voor gebruiken)
	name := fmt.Sprintf("%v%06v.txt", a, filecounter)
	file, error := os.OpenFile(outdir+"/"+name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(error)
	defer file.Close()
	filecounter += 1
	switch a {

	case "geometry":
		{
			// heel de lijst met particles aflopen en de locatie, straal en msat printen
			header := fmt.Sprintf("#position_x\tposition_y\tposition_z\tradius\tmsat\n")
			_, err = file.WriteString(header)
			check(err)

			for _, p := range lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", p.x, p.y, p.z, p.rc, p.msat)
				_, error = file.WriteString(string)
				check(error)
			}
		}
	case "m":
		{
			// loop over entire list with particles and print location, radius, msat and mag
			header := fmt.Sprintf("#t= %v\n#position_x\tposition_y\tposition_z\tradius\tmsat\tm_x\tm_y\tm_z\n", T.value)
			_, err = file.WriteString(header)
			check(err)

			for _, p := range lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", p.x, p.y, p.z, p.rc, p.msat, p.m[0], p.m[1], p.m[2])
				_, error = file.WriteString(string)
				check(error)
			}
		}
	case "anis":
		{
			// loop over entire list with particles and print location, radius, msat and mag
			header := fmt.Sprintf("#t= %v\n#position_x\tposition_y\tposition_z\tradius\tmsat\tu_anis_x\tu_anis_y\tu_anis_z\n", T.value)
			_, err = file.WriteString(header)
			check(err)

			for _, p := range lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", p.x, p.y, p.z, p.rc, p.msat, p.u_anis[0], p.u_anis[1], p.u_anis[2])
				_, error = file.WriteString(string)
				check(error)
			}
		}
	default:
		{
			log.Fatal(a, " is not a quantitity that can be saved")
		}
	}
}
