package controllers

import(
  "github.com/revel/revel"
  "labix.org/v2/mgo/bson"
  "flashback/app/models"
  //"fmt"
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

  if currentUser.Email != "" {
    coll := models.Card{}.Coll()
    var cards []models.Card
    err := coll.Find(bson.M{"userid": currentUser.Id}).All(&cards)
      if err != nil {
              panic(err)
      }
    return c.Render(currentUser, cards)
  }
  return c.Redirect(Sessions.New)
}

func (c Cards) Create(phrase string) revel.Result {
  coll := models.Card{}.Coll()
  err := coll.Insert(&models.Card{UserId: c.CurrentUser().Id, TargetLang: "buy", SourceLang: phrase})
  if err != nil {
    panic(err)
  }
	return c.Redirect(Cards.Index)
}

