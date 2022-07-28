package models

type User struct {
	Name    string  `"json:name"`
	Age     string  `"json:age"`
	Address Address `"json:address"`
}

type Address struct {
	Street string `"json:street"`
	City   string `"json:city"`
	State  string `"json:state"`
}
