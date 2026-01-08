package speed

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
        distance: 0,
    }
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    var newDistance = 0
    var newBattery = 0
    if car.batteryDrain > car.battery {
       newDistance = car.distance
       newBattery = car.battery
    } else { 
       newDistance = car.distance + car.speed
       newBattery = car.battery - car.batteryDrain 
    }
	newCar := Car {
        battery: newBattery,
        batteryDrain: car.batteryDrain,
        speed: car.speed,
        distance: newDistance,
    }
    return newCar
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    var distanceToTravel = track.distance - car.distance
    var finalDistance = car.distance + distanceToTravel
    var timesToDrain = distanceToTravel / car.speed // 10 / 5 == drain 2 times
     //note: in order to go 1 speed, the battery is drained by batteryDrain 1 time. So with speed = 5, and batteryDrain = 2, to go a distance of 10, the battery = 100 would be drained twice, with battery = 96
    var fullBatteryDrain = car.battery - (car.batteryDrain * timesToDrain)
   
	var canFinish = false
    var finalBatteryTooLow = fullBatteryDrain < 0
    canFinish = !finalBatteryTooLow && finalDistance >= track.distance
    return canFinish

}
