package controllers
 
import (
	"github.com/stunti/bloggo/app/models"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"time"
  "fmt"
)

type Blog struct {
	Application
}

func (c Blog) Index() revel.Result {
	article := new(models.Article)
	articles := article.All(c.MongoSession)

	return c.Render(articles)
}

func (c Blog) Show(IdBlog string) revel.Result {
	article := new(models.Article)
  fmt.Println("idBlog: ", IdBlog)
	one := article.GetByIdString(c.MongoSession, IdBlog)
  fmt.Println("blog: ", one)
  if one == nil {
    return c.NotFound("Not found: " + IdBlog)
  } else {
	  return c.Render(one)
  }
}

func (c Blog) Add() revel.Result {
	if c.User != nil {
		article := models.Article{}
		ObjectId := bson.ObjectId.Hex(article.Id)
		action := "/Blog/Create"
		return c.Render(action, ObjectId, article)
	}
	return c.Forbidden("You must be logged in to create articles.")
}

func (c Blog) Create(article *models.Article) revel.Result {
	if c.User != nil {
		article.Validate(c.Validation)
		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			c.Flash.Error("Please correct the errors below.")
			return c.Redirect(Blog.Add)
		}

		// Set calculated fields
		article.Author_id = c.User.Id
		article.Published = true
		article.Posted = time.Now()
		article.Id = bson.NewObjectId()
		article.Save(c.MongoSession)
	}
	return c.Redirect(Application.Index)
}
