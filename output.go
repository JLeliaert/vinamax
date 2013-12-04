//Contains function to control the output of the program
package vinamax

import (
	"fmt"
	"os"
)

var F *os.File
var Err error
var Outputinterval float64
var twrite float64
var locations []Vector

func Output(interval float64) {
	F, Err = os.Create("./table.txt")
	check(Err)
	//	defer F.Close()
	Writeheader()
	Outputinterval = interval
	twrite = 0.
}

//checks the error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//calculates the average magnetisation components of all Particles
func averages(lijst Particles) Vector {
	avgs := Vector{0, 0, 0}
	for i := range lijst {
		avgs[0] += lijst[i].M[0]
		avgs[1] += lijst[i].M[1]
		avgs[2] += lijst[i].M[2]
	}
	return avgs.times(1./float64(len(lijst)))
}

//Writes the header in table.txt
func Writeheader() {
	header := fmt.Sprintf("#t\t<mx>\t<my>\t<mz>\n")
	_, Err = F.WriteString(header)
	check(Err)
}

func Tableadd_B_eff_at_location(a,b,c float64){
locations = append (locations,Vector{a,b,c})

}

//Writes the time and the vector of average magnetisation in the table
func Write(avg Vector) {
	if twrite >= Outputinterval && Outputinterval != 0 {
		string := fmt.Sprintf("%v\t%v\t%v\t%v", T, avg[0], avg[1], avg[2])
		_, Err = F.WriteString(string)
		check(Err)
		for i := range locations{
			string = fmt.Sprintf("\t%v\t%v\t%v", (B_ext[0]+demag(locations[i][0],locations[i][1],locations[i][2])[0]), (B_ext[1]+demag(locations[i][0],locations[i][1],locations[i][2])[1]), (B_ext[2]+demag(locations[i][0],locations[i][1],locations[i][2])[2]))
			_, Err = F.WriteString(string)
		check(Err)
		}
		_, Err = F.WriteString("\n")
		check(Err)
		twrite = 0.
	}
	twrite += Dt
}
