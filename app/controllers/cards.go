package controllers

import(
  "github.com/revel/revel"
  "labix.org/v2/mgo"
 // "labix.org/v2/mgo/bson"
)

type Cards struct {
	*revel.Controller
}

func (c Cards) New() revel.Result {
	return c.Render()
}

func (c Cards) Create(phrase string) revel.Result {
  //session, err := mgo.Dial("localhost")
  //if err != nil {
          //panic(err)
  //}
  //defer session.Close()


  //c = session.DB("flashbackDev").C("test")
  //card := Card{"verteilen"}
  //err = c.Insert(card)
  //if err != nil {
          //panic(err)
  //}
  revel.TRACE.Printf(phrase)
	return c.Render()
}

