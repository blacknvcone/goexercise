package database

import (
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
)

func Initconn() dbox.IConnection {

	ci := dbox.ConnectionInfo{
		"127.0.0.1",
		"goexercise",
		"",
		"",
		nil,
	}

	conn, err := dbox.NewConnection("mongo", &ci)
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}
	err = conn.Connect()
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}

	//fmt.Println("Connected into database !")

	return conn
}
