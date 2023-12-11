package main

import "fmt"

//Ejemplo para creacion de motos, cuatrimotos, trimotos -- Todoterreno -- NighDrive -- Road
// Implementacion de Abstract Factory

type Motorcycle interface {
	Road()
	Jump()
}
type QuadBike interface {
	Skid()
}
type Trike interface {
	TakePeople()
}

// Motorcycles

type MotorcycleOffRoad struct{}

func (*MotorcycleOffRoad) Road() {
	fmt.Println("Motorcycle-OffRoad Driving on the stones...")
}
func (*MotorcycleOffRoad) Jump() {
	fmt.Println("Motorcycle-OffRoad Jumping too High over the ...")
}

type MotorcycleNighDrive struct{}

func (*MotorcycleNighDrive) Road() {
	fmt.Println("Motorcycle-NighDrive Driving very cool in the Night")
}
func (*MotorcycleNighDrive) Jump() {
	fmt.Println("Motorcycle-NighDrive Jumping too lower in the night")
}

type MotorcycleRoad struct{}

func (*MotorcycleRoad) Road() {
	fmt.Println("Motorcycle-Road driving far from the city")
}
func (*MotorcycleRoad) Jump() {
	fmt.Println("Motorcycle-Road jumping in the country")
}

//Quads

type QuadBikeOffRoad struct{}

func (*QuadBikeOffRoad) Skid() {}

type QuadBikeNighDrive struct{}

func (*QuadBikeNighDrive) Skid() {}

type QuadBikeRoad struct{}

func (*QuadBikeRoad) Skid() {}

//Trikes

type TrikeOffRoad struct{}

func (*TrikeOffRoad) TakePeople() {}

type TrikeNighDrive struct{}

func (*TrikeNighDrive) TakePeople() {}

type TrikeRoad struct{}

func (*TrikeRoad) TakePeople() {}

type IKTMFactory interface {
	createMotorCyle() Motorcycle
	createQuadBike() QuadBike
	createTrike() Trike
}

// type KTMFactory struct{}

// func (*KTMFactory) createMotorCyle() *Motorcycle {}
// func (*KTMFactory) createQuadBike() *QuadBike {}
// func (*KTMFactory) createTrike() *Trike       {}

type KTMFactoryOffRoad struct{}

func (*KTMFactoryOffRoad) createMotorCyle() Motorcycle {
	return &MotorcycleOffRoad{}
}
func (*KTMFactoryOffRoad) createQuadBike() QuadBike {
	return &QuadBikeOffRoad{}
}
func (*KTMFactoryOffRoad) createTrike() Trike {
	return &TrikeOffRoad{}
}

type KTMFactoryNighDrive struct{}

func (*KTMFactoryNighDrive) createMotorCyle() Motorcycle {
	return &MotorcycleNighDrive{}
}
func (*KTMFactoryNighDrive) createQuadBike() QuadBike {
	return &QuadBikeNighDrive{}
}
func (*KTMFactoryNighDrive) createTrike() Trike {
	return &TrikeNighDrive{}
}

type KTMFactoryRoad struct{}

func (*KTMFactoryRoad) createMotorCyle() Motorcycle {
	return &MotorcycleRoad{}
}
func (*KTMFactoryRoad) createQuadBike() QuadBike {
	return &QuadBikeRoad{}
}
func (*KTMFactoryRoad) createTrike() Trike {
	return &TrikeRoad{}
}

type Application struct {
	fac         IKTMFactory
	m           Motorcycle
	q           QuadBike
	t           Trike
	buildSetX10 func()
}

func NewApplication(params IKTMFactory) *Application {
	return &Application{
		fac: params,
		m:   params.createMotorCyle(),
		q:   params.createQuadBike(),
		t:   params.createTrike(),
	}
}

var configFileApp map[string]string = map[string]string{"type_ktm": "NighDrive"}

func mainx1() {
	var ktm_factory IKTMFactory
	var duke_200 Motorcycle
	// La aplicacion elige el tipo de fabrica dependiendo de la config
	// actual o de los ajustes del entorno y la crea - en tiempo de ejecución, normalmente
	// en la etapa de inicializacion

	if configFileApp["type_ktm"] == "OffRoad" {
		ktm_factory = &KTMFactoryOffRoad{}
		duke_200 = ktm_factory.createMotorCyle()
	}
	if configFileApp["type_ktm"] == "NighDrive" {
		ktm_factory = &KTMFactoryNighDrive{}
		duke_200 = ktm_factory.createMotorCyle()
	}
	if configFileApp["type_ktm"] == "Road" {
		ktm_factory = &KTMFactoryRoad{}
		duke_200 = ktm_factory.createMotorCyle()
	}

	duke_200.Jump()
	duke_200.Road()

}
