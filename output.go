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
var locations []Vector
var filecounter int = 0
var output_B_ext = false

//Sets the interval at which times the output table has to be written
func Output(interval float64) {
	outputcalled = true
	f, err = os.Create(outdir + "/table.txt")
	check(err)
	//	defer f.Close()
	writeheader()
	Outputinterval = interval
	twrite = interval
}

//checks the error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//calculates the average magnetisation components of all Particles
func averages(lijst []*Particle) Vector {
	avgs := Vector{0, 0, 0}
	for i := range lijst {
		avgs[0] += lijst[i].m[0]
		avgs[1] += lijst[i].m[1]
		avgs[2] += lijst[i].m[2]
	}
	return avgs.times(1. / float64(len(lijst)))
}

//Writes the header in table.txt
func writeheader() {
	header := fmt.Sprintf("#t\t<mx>\t<my>\t<mz>")
	_, err = f.WriteString(header)
	check(err)
	if output_B_ext {
		header := fmt.Sprintf("\tB_ext_x\tB_ext_y\t<B_ext_z")
		_, err = f.WriteString(header)
		check(err)
	}
	header = fmt.Sprintf("\n")
	_, err = f.WriteString(header)
	check(err)

}

//prints the suggested timestep for the simulation
func printsuggestedtimestep() {
	shouldbemaxerror := 1e-5
	currentmaxerror := Dt * maxtauwitht
	fmt.Println("A good timestep would be: ", Dt*math.Pow(shouldbemaxerror/currentmaxerror, 1/2.))
}

//Adds the effective field at a specific location to the output table
func Tableadd_B_eff_at_location(a, b, c float64) {
	if Outputinterval != 0 {
		log.Fatal("Output() should always come AFTER Tableadd_B_eff_at_location()")
	}
	if universe.inworld(Vector{a, b, c}) {
		locations = append(locations, Vector{a, b, c})
	} else {
		fmt.Println("error: not in universe")
	}
}

//Writes the time and the vector of average magnetisation in the table
func write(avg Vector) {
	if twrite >= Outputinterval && Outputinterval != 0 {
		string := fmt.Sprintf("%v\t%v\t%v\t%v", T, avg[0], avg[1], avg[2])
		_, err = f.WriteString(string)
		check(err)

		if output_B_ext {
			B_ext_x, B_ext_y, B_ext_z := B_ext(T)
			string = fmt.Sprintf("\t%v\t%v\t%v", B_ext_x, B_ext_y, B_ext_z)
			_, err = f.WriteString(string)
			check(err)
		}

		for i := range locations {
			B_ext_x, B_ext_y, B_ext_z := B_ext(T)

			string = fmt.Sprintf("\t%v\t%v\t%v", (B_ext_x + demag(locations[i][0], locations[i][1], locations[i][2])[0]), (B_ext_y + demag(locations[i][0], locations[i][1], locations[i][2])[1]), (B_ext_z + demag(locations[i][0], locations[i][1], locations[i][2])[2]))
			_, err = f.WriteString(string)
			check(err)
		}
		_, err = f.WriteString("\n")
		check(err)
		twrite = 0.
	}
	twrite += Dt
}

//Saves different quantities. At the moment only "geometry" and "m" are possible
func Save(a string) {
	//een file openen met unieke naam (counter voor gebruiken)
	name := fmt.Sprintf("%v%06v.txt", a, filecounter)
	file, error := os.Create(outdir + "/" + name)
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

			for i := range universe.lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", universe.lijst[i].X, universe.lijst[i].Y, universe.lijst[i].Z, universe.lijst[i].r, universe.lijst[i].msat)
				_, error = file.WriteString(string)
				check(error)
			}
		}
	case "m":
		{
			// heel de lijst met particles aflopen en de locatie straal,msat en magnetisatie printen
			header := fmt.Sprintf("#t= %v\n#position_x\tposition_y\tposition_z\tradius\tmsat\tm_x\tm_y\tm_z\n", T)
			_, err = file.WriteString(header)
			check(err)

			for i := range universe.lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", universe.lijst[i].X, universe.lijst[i].Y, universe.lijst[i].Z, universe.lijst[i].r, universe.lijst[i].msat, universe.lijst[i].m[0], universe.lijst[i].m[1], universe.lijst[i].m[2])
				_, error = file.WriteString(string)
				check(error)
			}
		}
	default:
		{
			fmt.Println(a, " is not a quantitity that can be saved")
		}
	}
}

//adds a quantity to the output table, at the moment only "B_ext" is possible
func Tableadd(a string) {
	if Outputinterval != 0 {
		log.Fatal("Output() should always come AFTER Tableadd()")
	}
	switch a {
	case "B_ext":
		{
			output_B_ext = true
		}
	default:
		{
			log.Fatal(a, " is currently not addable to table")
		}
	}
}

//returns a suggested timestep at the end of the simulation
func SuggestTimestep() {
	suggest_timestep = true
}
