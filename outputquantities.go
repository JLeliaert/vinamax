package vinamax

import (
	"fmt"
	"log"
)

//adds a quantity to the output table
func Tableadd(a string) {
	switch a {
	case "B_ext":
		{
			b := output_B_ext{}
			outputList = append(outputList, &b)
		}
	case "Dt":
		{
			b := output_Dt{}
			outputList = append(outputList, &b)
		}
	case "NrMzPos":
		{
			b := output_NrMzPos{}
			outputList = append(outputList, &b)
		}
	case "AllMag":
		{
			b := output_AllMag{}
			outputList = append(outputList, &b)
		}
	case "U_anis":
		{
			b := output_U_anis{}
			outputList = append(outputList, &b)
		}
	case "Energy":
		{
			b := output_Energy{}
			outputList = append(outputList, &b)
		}

	default:
		{
			log.Fatal(a, " is currently not addable to table")
		}
	}
}

//OutputQuantity B_ext
type output_B_ext struct {
}

func (o output_B_ext) header() string {
	return "\tB_ext_x\tB_ext_y\tB_ext_z"
}

func (o output_B_ext) value() string {
	B_ext_x, B_ext_y, B_ext_z := B_ext(T.value)
	return fmt.Sprintf("\t%v\t%v\t%v", B_ext_x, B_ext_y, B_ext_z)
}

//OutputQuantity Dt
type output_Dt struct {
}

func (o output_Dt) header() string {
	return "\tDt"
}
func (o output_Dt) value() string {
	return fmt.Sprintf("\t%v", Dt.value)
}

//OutputQuantity NrMzPos
type output_NrMzPos struct {
}

func (o output_NrMzPos) header() string {
	return "\tnrmzpos"
}

func (o output_NrMzPos) value() string {
	return fmt.Sprintf("\t%v", nrmzpositive())
}

//OutputQuantity AllMag
type output_AllMag struct {
}

func (o output_AllMag) header() string {
	header := ""
	for range lijst {
		header += fmt.Sprintf("\tm_x\tm_y\tm_z")
	}
	return header
}

func (o output_AllMag) value() string {
	string := ""
	for _, i := range lijst {
		string += fmt.Sprintf("\t%v\t%v\t%v", i.m[0], i.m[1], i.m[2])
	}
	return string
}

//OutputQuantity U_anis
type output_U_anis struct {
}

func (o output_U_anis) header() string {
	return "\tu_x\tu_y\tu_z"
}

func (o output_U_anis) value() string {
	averaged_u := averages_u()
	return fmt.Sprintf("\t%v\t%v\t%v", averaged_u[0], averaged_u[1], averaged_u[2])
}

//OutputQuantity Energy
type output_Energy struct {
}

func (o output_Energy) header() string {
	return "\tE_zeeman\tE_demag\tE_anis\tE_therm\tE_total"
}

func (o output_Energy) value() string {
	return fmt.Sprintf("\t%v\t%v\t%v\t%v\t%v", E_zeeman(), E_demag(), E_anis(), E_therm(), E_total())
}
