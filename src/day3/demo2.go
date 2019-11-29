package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

funct main(){
	upass:="Admin"
	encrypass,_:=bcrypt.GenerateFromPassword([]byte(upass),1)
	fmt.Println(encrypass)
}