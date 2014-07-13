package db

import(
  "labix.org/v2/mgo"
  "reflect"
  "fmt"
)

func Connect() *mgo.Database {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  //defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  database := session.DB("flashbackDev")
  fmt.Println("connect +++++++++++++++++++")
  fmt.Println(reflect.TypeOf(database))
  return database

}
