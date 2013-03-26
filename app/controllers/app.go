package controllers

import (
	"github.com/jgraham909/bloggo/app/models"
	m "github.com/jgraham909/revmgo/app/controllers"
	"github.com/robfig/revel"
)

type Application struct {
	m.MgoController
	User *models.User
}

func init() {
	revel.InterceptMethod((*Application).Setup, revel.BEFORE)
	revel.TemplateFuncs["nil"] = func(a interface{}) bool {
		return a == nil
	}
}

// Responsible for doing any necessary setup for each web request.
func (c *Application) Setup() revel.Result {
	// If there is an active user session load the User data for this user.
	if email, ok := c.Session["user"]; ok {
		c.User = c.User.GetByEmail(c.MSession, email)
	}
	return nil
}

func (c Application) Index() revel.Result {
	if c.User != nil {
		user := c.User
		return c.Render(user)
	}
	return c.Render()
}

func (c Application) UserAuthenticated() bool {
	if _, ok := c.Session["user"]; ok {
		return true
	}
	return false
}
