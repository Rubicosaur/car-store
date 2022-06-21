package cars

type Vehicle interface {
	drive()
	stop()
}

type Car struct {
	make     string
	model    string
	color    string
	mileage  float64
	fuelTank float64
}
