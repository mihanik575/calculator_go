package main


type Car struct {
	battery     int
	batteryDrain int
	speed       int
	distance    int
}


func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:     100,
		batteryDrain: batteryDrain,
		speed:       speed,
		distance:    0,
		}
	}

c := Car{60, 13}