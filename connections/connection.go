package connections

import db "github.com/eaciit/dbox"
import "log"

func PrepareConnection() (db.IConnection, error) {
	var connInfo = &db.ConnectionInfo{
		Host:     "localhost:27017",
		Database: "example",
		UserName: "",
		Password: "",
	}
	c, err := db.NewConnection("mongo", connInfo)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = c.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	return c, nil
}
