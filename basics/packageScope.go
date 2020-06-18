package basics

import (
	"constants"
	"datatypes"
	"fmt"
)

var a int = 10

const name string = "TEST"

func TryCrossPackage() {
	fmt.Println(a)
	fmt.Println(constants.Pi)
	fmt.Println(constants.SpeedOfLightInVaccum)
	constants.CallUnaccessibleMethod()

	datatypes.MaxUint8()
	datatypes.MaxMinInt8()
	datatypes.MaxUint16()
	datatypes.MaxMinInt16()

	datatypes.MaxMinFloat32()
	// var i uint8 = -8
	// fmt.Println(i)

}
