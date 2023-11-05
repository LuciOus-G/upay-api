package models

//
//import (
//	"log"
//	"upay/api/v2/src/connections"
//)
//
//type Model struct {
//	Model interface{}
//}
//
//func Migrations() {
//	registeredModel := []Model{
//		{Model: UserModel{}},
//	}
//
//	connections.DatabaseInit()
//
//	for _, Models := range registeredModel {
//		err := connections.DB().Debug().AutoMigrate(Models.Model)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}
