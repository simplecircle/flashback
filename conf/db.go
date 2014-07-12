package conf

import(
  "labix.org/v2/mgo"
  //"reflect"
  "fmt"
)

func Db() *mgo.Database {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  //defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  //db := session.DB("flashbackDev").C("cards")
  database := session.DB("flashbackDev")
  fmt.Println("inner+++++++++++++++++++")
  //fmt.Println(reflect.TypeOf(database))
  return database

}
