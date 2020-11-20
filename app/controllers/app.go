package controllers

import (
	"database/sql"
	"e-learning/app/storage"
	"e-learning/app/storage/user/db"
	"e-learning/app/models"
	"e-learning/app/routes"
	"e-learning/app/storage/material"
	mdb "e-learning/app/storage/material/db"
	"e-learning/app/storage/user"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

// App .
type App struct {
	*revel.Controller
	db *sql.DB
}

// Index .
func (c App) Index() revel.Result {
	connection, err := storage.CreateConnectionPostgres("localhost",5232,"billfazz", "billfazz", "e_learning", "disable")
	if err != nil {
		panic(err)
	}
	m := material.NewMaterial(mdb.NewMaterialDB(connection))
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

// Register .
func (c App) Register() revel.Result {
	
	return c.Render()
}

// Auth . 
func (c App) Auth(email, password string) revel.Result {
	c.Validation.Required(email).Message("Please input your email")
	c.Validation.Required(password).Message("Please input your password")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	connection, err := storage.CreateConnectionPostgres("localhost",5232,"billfazz", "billfazz", "e_learning", "disable")
	if err != nil {
		panic(err)
	}
	user := user.NewUser(db.NewUserDB(connection))
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

// AddUser .
func (c App) AddUser(email, password, password2 string) revel.Result {
	connection, err := storage.CreateConnectionPostgres("localhost",5232,"billfazz", "billfazz", "e_learning", "disable")
	if err != nil {
		panic(err)
	}
	user := user.NewUser(db.NewUserDB(connection))

	c.Validation.Required(email).Message("Please input your email")
	c.Validation.Required(password).Message("Please input your password")
	c.Validation.Required(password2).Message("Please input your repeat password")
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Register)
	}
	_, err = user.Create(email, password, email)
	if err != nil {
		c.Flash.Error("Something went wrong")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Register)
	}
	c.Flash.Success("Congratulations, your account was created. You can login with your account")
	return c.Redirect(App.Login)
}

// Dashboard .
func (c App) Dashboard() revel.Result {
	return c.Render()
}

// Material .
func (c App) Material(id int) revel.Result {
	connection, err := storage.CreateConnectionPostgres("localhost",5232,"billfazz", "billfazz", "e_learning", "disable")
	if err != nil {
		panic(err)
	}
	_, err = c.Session.Get("isLoggedIn")
	if err != nil {
		c.Session.Set("recentMaterial", id)
		c.Flash.Error("You need login before continue")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	m := material.NewMaterial(mdb.NewMaterialDB(connection))
	material, err := m.Get(id)
	if err != nil {
		panic(err)
	}
	return c.Render(material)
}

// ManageMaterials .
func (c App) ManageMaterials() revel.Result {
	connection, err := storage.CreateConnectionPostgres("localhost",5232,"billfazz", "billfazz", "e_learning", "disable")
	if err != nil {
		panic(err)
	}
	m := material.NewMaterial(mdb.NewMaterialDB(connection))
	materials, err := m.List()
	if err != nil {
		panic(err)
	}
	return c.Render(materials)
}

// Upload .
func (c App) Upload(title, description, author, category, resource string) revel.Result {
	fmt.Println(title, description, author, category, resource)
	return c.Render(title, description, author, category, resource)
}

func (c App) getUser(email string) *models.User {
	var user models.User
	_, err := c.Session.GetInto("email", user, false)
	if err == nil && user.Email == email {
		return &user
	}
	return nil
}
