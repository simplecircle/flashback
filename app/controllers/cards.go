package controllers

import(
  "fmt"
  "github.com/revel/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "flashback/app/models"
  "reflect"
)

type Cards struct {
	*revel.Controller
  Helpers
}

func (c Cards) New() revel.Result {
	return c.Render()
}

func (c Cards) Index() revel.Result {
  currentUser := c.CurrentUser()
  coll2 := models.Db("cards")
  var cards []models.Card
fmt.Println(reflect.TypeOf(coll2))
fmt.Println(currentUser.Id)
  coll2.Find(bson.M{"userid": currentUser.Id}).All(&cards)

  //session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  //if err != nil {
          //panic(err)
  //}
  //defer session.Close()
	//session.SetMode(mgo.Monotonic, true)

  //coll := session.DB("flashbackDev").C("cards")
//fmt.Println(reflect.TypeOf(coll))
  //err = coll.Find(bson.M{"userid": currentUser.Id}).All(&cards)
  //if err != nil {
          //panic(err)
  //}
	return c.Render(currentUser, cards)
}

func (c Cards) Create(phrase string) revel.Result {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  coll := session.DB("flashbackDev").C("cards")
  err = coll.Insert(&models.Card{UserId: c.CurrentUser().Id, TargetLang: "buy", SourceLang: phrase})
  if err != nil {
    panic(err)
  }

  //result := Card{}
  var results []models.Card
  //err = coll.Find(bson.M{"german": phrase}).One(&result)
  err = coll.Find(bson.M{"english": "buy"}).All(&results)
  if err != nil {
          panic(err)
  }
  //revel.TRACE.Printf(phrase)
	return c.Redirect(Cards.Index)
}

