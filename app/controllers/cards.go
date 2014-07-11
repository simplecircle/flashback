package controllers

import(
  //"fmt"
  "github.com/revel/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "flashback/app/models"
)

type Cards struct {
	*revel.Controller
}
type Card struct {
  English string
  German string
}

func (c Cards) CurrentUser() models.User {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  defer session.Close()
  cookie, _ := c.Request.Cookie("authToken")
  cookieAuthToken := cookie.Value
  currentUser := models.User{}
  coll := session.DB("flashbackDev").C("users")
  coll.Find(bson.M{"authtoken" : cookieAuthToken}).One(&currentUser)
  if err != nil {
          panic(err)
  }
  //fmt.Println(currentUser.Email)
  return currentUser
}

func (c Cards) New() revel.Result {
	return c.Render()
}
func (c Cards) Index() revel.Result {
  currentUser := c.CurrentUser()
	return c.Render(currentUser)
}

func (c Cards) Create(phrase string) revel.Result {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  coll := session.DB("flashbackDev").C("cards")
  err = coll.Insert(&Card{"buy", phrase})
  if err != nil {
          panic(err)
  }

  //result := Card{}
  var results []Card
  //err = coll.Find(bson.M{"german": phrase}).One(&result)
  err = coll.Find(bson.M{"english": "buy"}).All(&results)
  if err != nil {
          panic(err)
  }
  hell := []string{"ass blood", "tit blood", "face blood"}
  //fmt.Println("Card:", result.English)
  //revel.TRACE.Printf(phrase)
	return c.Render(results, hell)
}

