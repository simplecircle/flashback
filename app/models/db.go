package models

import(
  "labix.org/v2/mgo"
  "reflect"
  "fmt"
)

func Db(collection string) *mgo.Collection{
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  //defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  fmt.Println(collection)
  coll := session.DB("flashbackDev").C(collection)
  fmt.Println("+++++++++++++++++++")
  fmt.Println(reflect.TypeOf(coll))
  fmt.Println(coll)
  return coll

}
