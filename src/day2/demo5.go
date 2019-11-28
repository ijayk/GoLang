package main

import "fmt"

//inheritance with interface
type vehicle interface{
	engine()
	
	//doors()
}

type bmw struct{
	model string
	price int
}

type activa struct{
	model string
	price int
}
func main(){
	bmwobj:=bmw{"BMW01",100000}
	vehicledetails(bmwobj)
	activaobj:=activa{"Activa5G",7000}
	vehicledetails(activaobj)

}

func (obj bmw)engine(){
	fmt.Println("BMW has auto engine")

}

func vehicledetails(v vehicle){
	fmt.Println(v)
	v.engine()
	//type casting
	bmwobj,ok :=v.(bmw)
	if ok==true {
		bmwobj.doors()
	}
}

func (obj activa)engine(){
	fmt.Println("Activa engine is auto with no gears")
}

func (obj bmw)doors(){
	fmt.Println("bmw has 2 doors")
}




