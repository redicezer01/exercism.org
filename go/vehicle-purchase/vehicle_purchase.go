package purchase

import "fmt"

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	better := option1
	if option1 > option2 {
		better = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", better)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if int(age) < 3 {
		return calcPrice(originalPrice, 80)
	} else if int(age) < 10 {
		return calcPrice(originalPrice, 70)
	}
	return calcPrice(originalPrice, 50)
}

// calcPrice func returns price recalculated as a percentage of original рrice
// percent accepts values from 0 to 100, e.g. 0, 10, 70
func calcPrice(price float64, percent uint8) float64 {
	return price * (float64(percent) / 100.0)
}
