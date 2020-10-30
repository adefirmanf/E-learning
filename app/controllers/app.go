package controllers

import (
	"e-learning/app/models"
	"e-learning/app/routes"
	"e-learning/app/storage/material"
	"e-learning/app/storage/material/inmem"
	"e-learning/app/storage/user"
	inmemuser "e-learning/app/storage/user/inmem"
	"fmt"

	"github.com/revel/revel"
)

// App .
type App struct {
	*revel.Controller
}

// Index .
func (c App) Index() revel.Result {
	m := material.NewMaterial(inmem.NewMaterialInMem(inmem.SeederMaterial()))
	materials, err := m.List()
	if err != nil {
		panic(err)
	}
	return c.Render(materials)
}

// Login .
func (c App) Login() revel.Result {
	return c.Render()
}

// Verify .
func (c App) Verify(email, password string) revel.Result {
	c.Validation.Required(email).Message("Please input your email")
	c.Validation.Required(password).Message("Please input your password")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	return c.Redirect(App.Auth, email, password)
}

// Auth .
func (c App) Auth(email, password string) revel.Result {
	user := user.NewUser(inmemuser.NewUserInMem(inmemuser.SeederUser()))
	if !user.IsValid(email, password) {
		c.Flash.Error("Invalid e-mail / password")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	c.Session.Set("isLoggedIn", true)
	getID, err := c.Session.Get("recentMaterial")
	if err == nil && getID != nil {
		fmt.Println("Exist!")
		convID := getID.(float64)
		return c.Redirect(routes.App.Material(int(convID)))
	}
	return c.Redirect(App.Index)
}

// Dashboard .
func (c App) Dashboard() revel.Result {
	return c.Render()
}

// Material .
func (c App) Material(id int) revel.Result {
	_, err := c.Session.Get("isLoggedIn")
	if err != nil {
		c.Session.Set("recentMaterial", id)
		c.Flash.Error("You need login before continue")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	m := material.NewMaterial(inmem.NewMaterialInMem(inmem.SeederMaterial()))
	material, err := m.Get(id)
	if err != nil {
		panic(err)
	}
	return c.Render(material)
}

func (c App) getUser(email string) *models.User {
	var user models.User
	_, err := c.Session.GetInto("email", user, false)
	if err == nil && user.Email == email {
		return &user
	}
	return nil
}
