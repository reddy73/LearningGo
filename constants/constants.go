package constants

import "fmt"

const Pi float32 = 3.14
const SpeedOfLightInVaccum int64 = 299792458
const Organization string = "EPAM Systems"

//Example of an exported method
func MethodFromConstants() {
	fmt.Println("method from Constants")
}

//Example of an unexported method
func unaccessibleMethod() {
	fmt.Println("This is not accessible outside the package")
}
