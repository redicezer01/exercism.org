//Package lasagna helps you cook a briliant lasasgna from your favorite cooking book
package lasagna

// OvenTime returns how many minutes the lasagna should be in the oven
func OvenTime() int {
	return 40
}

// RemainingOvenTime calculates the remaining oven time in minutes
// t is actual minutes the lasagna has been in the oven
func RemainingOvenTime(t int) int {
	return OvenTime() - t
}

// PreparationTime calculates preparation time in minutes
// layerscnt is count of layers in lasagna
func PreparationTime(layerscnt int) int {
	return 2 * layerscnt
}

// ElapsedTime calculates the elapsed working time it minutes
// layerscnt is count of layers in lasagna
// t is number ofj minutes the lasagna has been in the oven
func ElapsedTime(layerscnt int, t int) int {
	return PreparationTime(layerscnt) + t
}
