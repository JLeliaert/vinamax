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
var output_B_ext = false
var output_Dt = false
var output_nrmzpos = false
var output_mdoth = false
var output_allmag = false
var output_u_anis = false
var output_u_anis_xy = false
var output_energy = false

//var timelastswitch =0.//EXTRA
//var updownswitch =true//EXTRA

//Sets the interval at which times the output table has to be written
func Output(interval float64) {
	//Print1 = false
	//Print0 = false
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

//Helemaal extra
//func plotswitchtime(){
//	if (updownswitch==true && Universe.lijst[0].m[2]<=-0.8){
//		updownswitch=false
//		fmt.Println(T-timelastswitch)
//		timelastswitch=T
//	}
//	if (updownswitch ==false && Universe.lijst[0].m[2]>=0.8){
//		updownswitch=true
//		fmt.Println(T-timelastswitch)
//		timelastswitch=T
//	}
//}

//checks the error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//calculates the average magnetisation components of all particles
func averages(lijst []*particle) vector {
	avgs := vector{0, 0, 0}
	for i := range lijst {
		avgs[0] += lijst[i].m[0]
		avgs[1] += lijst[i].m[1]
		avgs[2] += lijst[i].m[2]
	}
	return avgs.times(1. / float64(len(lijst)))
}

//calculates the average anisotropy components of all particles
func averages_u(lijst []*particle) vector {
	avgs := vector{0, 0, 0}
	for i := range lijst {
		if lijst[i].u_anis[0] < 0 {
			lijst[i].u_anis[0] = (-1) * lijst[i].u_anis[0]
			lijst[i].u_anis[1] = (-1) * lijst[i].u_anis[1]
			lijst[i].u_anis[2] = (-1) * lijst[i].u_anis[2]

		}
		avgs[0] += lijst[i].u_anis[0]
		avgs[1] += lijst[i].u_anis[1]
		avgs[2] += lijst[i].u_anis[2]
	}
	avgs = avgs.times(1. / float64(len(lijst)))
	return avgs
}

func averages_u_xy(lijst []*particle) vector {
	avgs := vector{0, 0, 0}
	for i := range lijst {
		if lijst[i].u_anis[0] < 0 {
			lijst[i].u_anis[0] = (-1) * lijst[i].u_anis[0]
			lijst[i].u_anis[1] = (-1) * lijst[i].u_anis[1]
			lijst[i].u_anis[2] = (-1) * lijst[i].u_anis[2]

		}
		avgs[0] += math.Sqrt(lijst[i].u_anis[0]*lijst[i].u_anis[0]) + (lijst[i].u_anis[1] * lijst[i].u_anis[1])
		avgs[2] += lijst[i].u_anis[2]
	}
	return avgs.times(1. / float64(len(lijst)))
}

//calculates the average moments of all particles
//TODO weigh with msat
func averagemoments(lijst []*particle) vector {
	avgs := vector{0, 0, 0}
	totalvolume := 0.
	for i := range lijst {
		radius := lijst[i].r
		volume := cube(radius) * 4. / 3. * math.Pi
		totalvolume += volume
		avgs[0] += lijst[i].m[0] * volume
		avgs[1] += lijst[i].m[1] * volume
		avgs[2] += lijst[i].m[2] * volume
	}
	//divide by total volume
	return avgs.times(1. / totalvolume)
}

//calculates the dotproduct of the average moments and the effective field of all particles
func averagemdoth(lijst []*particle) float64 {
	avg := 0.
	for i := range lijst {
		xcomp := lijst[i].m[0] * lijst[i].heff[0]
		ycomp := lijst[i].m[1] * lijst[i].heff[1]
		zcomp := lijst[i].m[2] * lijst[i].heff[2]
		avg = (xcomp + ycomp + zcomp) / mu0
	}
	return (avg)
}

//returns the number of particles with m_z larger than 0
func nrmzpositive(lijst []*particle) int {
	counter := 0
	for i := range lijst {
		if lijst[i].m[2] > 0. {
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
	if output_B_ext {
		header := fmt.Sprintf("\tB_ext_x\tB_ext_y\tB_ext_z")
		_, err = f.WriteString(header)
		check(err)
	}
	if output_Dt {
		header := fmt.Sprintf("\tDt")
		_, err = f.WriteString(header)
		check(err)
	}
	if output_nrmzpos {
		header := fmt.Sprintf("\tnrmzpos")
		_, err = f.WriteString(header)
		check(err)
	}
	if output_mdoth {
		header := fmt.Sprintf("\tmdotH")
		_, err = f.WriteString(header)
		check(err)
	}
	if output_allmag {
		for range Universe.lijst {
			header := fmt.Sprintf("\tm_x\tm_y\tm_z")
			_, err = f.WriteString(header)
			check(err)
		}
	}

	if output_u_anis {
		header := fmt.Sprintf("\tu_anis_x\tu_anis_y\tu_anis_z")
		_, err = f.WriteString(header)
		check(err)
	}
	if output_energy {
		header := fmt.Sprintf("\tE_zeeman\tE_demag\tE_anis\tE_therm\tE_total")
		_, err = f.WriteString(header)
		check(err)
	}
	if output_u_anis_xy {
		header := fmt.Sprintf("\tu_anis_xy\tu_anis_z")
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

////prints the suggested timestep for the simulation
//deprecated, is the responsibility of the user
//func printsuggestedtimestep() {
//	shouldbemaxerror := 5e-4
//	currentmaxerror := maxtauwitht //* Dt
//	fmt.Println("maxerr=", currentmaxerror)
//	fmt.Println("A good timestep would be: ", Dt*math.Pow(shouldbemaxerror/currentmaxerror, 1/float64(order)))
//}

//Adds the field at a specific location to the output table
func Tableadd_b_at_location(x, y, z float64) {
	tableaddcalled = true
	if outputinterval != 0 {
		log.Fatal("Output() should always come AFTER Tableadd_b_at_location()")
	}
	locations = append(locations, vector{x, y, z})

}

func Give_mz() float64 {
	return averagemoments(Universe.lijst)[2]
}

//Writes the time and the vector of average magnetisation in the table
func write(avg vector, forced bool) {
	if forced || (twrite >= outputinterval && outputinterval != 0) {
		string := fmt.Sprintf("%e\t%v\t%v\t%v", T.value, avg[0], avg[1], avg[2])
		_, err = f.WriteString(string)
		check(err)

		if output_B_ext {
			B_ext_x, B_ext_y, B_ext_z := B_ext(T.value)
			string = fmt.Sprintf("\t%v\t%v\t%v", B_ext_x, B_ext_y, B_ext_z)
			_, err = f.WriteString(string)
			check(err)
		}
		if output_Dt {
			string = fmt.Sprintf("\t%v", Dt.value)
			_, err = f.WriteString(string)
			check(err)
		}
		if output_nrmzpos {
			string = fmt.Sprintf("\t%v", nrmzpositive(Universe.lijst))
			_, err = f.WriteString(string)
			check(err)
		}
		if output_mdoth {
			string = fmt.Sprintf("\t%v", averagemdoth(Universe.lijst))
			_, err = f.WriteString(string)
			check(err)
		}
		if output_allmag {
			for _, i := range Universe.lijst {
				string = fmt.Sprintf("\t%v\t%v\t%v", i.m[0], i.m[1], i.m[2])
				_, err = f.WriteString(string)
				check(err)
			}
		}
		if output_u_anis {
			averaged_u_anis := averages_u(Universe.lijst)
			string = fmt.Sprintf("\t%v\t%v\t%v", averaged_u_anis[0], averaged_u_anis[1], averaged_u_anis[2])
			_, err = f.WriteString(string)
			check(err)
		}

		if output_energy {
			string = fmt.Sprintf("\t%v\t%v\t%v\t%v\t%v", E_zeeman(), E_demag(), E_anis(), E_therm(), E_total())
			_, err = f.WriteString(string)
			check(err)
		}

		if output_u_anis_xy {
			averaged_u_anis := averages_u_xy(Universe.lijst)
			string = fmt.Sprintf("\t%v\t%v", averaged_u_anis[0], averaged_u_anis[2])
			_, err = f.WriteString(string)
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

	case "phasediagram":
		{
			if Print1 {
				string := fmt.Sprintf("%d\n", 1)
				_, error = file.WriteString(string)
				check(error)
			}
			if Print0 {
				string := fmt.Sprintf("%d\n", 0)
				_, error = file.WriteString(string)
				check(error)
			}
			filecounter -= 1
			if (Print1 == false) && (Print0 == false) {
				fmt.Println("er is een simulatie niet uitgelopen, onduidelijk resultaat")
				string := fmt.Sprintf("%d\n", 2)
				_, error = file.WriteString(string)
				check(error)
			}
			if (Print1 == true) && (Print0 == true) {
				log.Fatal("dit kan niet")
			}
		}

	case "geometry":
		{
			// heel de lijst met particles aflopen en de locatie, straal en msat printen
			header := fmt.Sprintf("#position_x\tposition_y\tposition_z\tradius\tmsat\n")
			_, err = file.WriteString(header)
			check(err)

			for i := range Universe.lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", Universe.lijst[i].x, Universe.lijst[i].y, Universe.lijst[i].z, Universe.lijst[i].r, Universe.lijst[i].msat)
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

			for i := range Universe.lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", Universe.lijst[i].x, Universe.lijst[i].y, Universe.lijst[i].z, Universe.lijst[i].r, Universe.lijst[i].msat, Universe.lijst[i].m[0], Universe.lijst[i].m[1], Universe.lijst[i].m[2])
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

			for i := range Universe.lijst {
				string := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", Universe.lijst[i].x, Universe.lijst[i].y, Universe.lijst[i].z, Universe.lijst[i].r, Universe.lijst[i].msat, Universe.lijst[i].u_anis[0], Universe.lijst[i].u_anis[1], Universe.lijst[i].u_anis[2])
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

//adds a quantity to the output table, at the moment only "B_ext" is possible
func Tableadd(a string) {
	tableaddcalled = true
	if outputinterval != 0 {
		log.Fatal("Output() should always come AFTER Tableadd()")
	}
	switch a {
	case "B_ext":
		{
			output_B_ext = true
		}
	case "Dt":
		{
			output_Dt = true
		}
	case "nrmzpos":
		{
			output_nrmzpos = true
		}
	case "mdoth":
		{
			output_mdoth = true
		}
	case "allmag":
		{
			output_allmag = true
		}
	case "u_anis":
		{
			output_u_anis = true
		}
	case "energy":
		{
			output_energy = true
		}
	case "u_anis_xy":
		{
			output_u_anis_xy = true
		}

	default:
		{
			log.Fatal(a, " is currently not addable to table")
		}
	}
}

//returns a suggested timestep at the end of the simulation
//func Suggesttimestep() {
//	suggest_timestep = true
//}

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
	write(averagemoments(Universe.lijst), true)
}
