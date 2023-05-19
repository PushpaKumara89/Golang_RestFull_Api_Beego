package controllers

import (
	// "encoding/json"
	// "fmt"
	// "net/http"

	"ApiBeeGo/db"
	"ApiBeeGo/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "Invalid request body "+err.Error())
		return
	}

	password, bcErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if bcErr != nil {
		c.CustomAbort(http.StatusNotAcceptable, "Invalid request body "+bcErr.Error())
		return
	}

	user.Password = string(password)
	id, err := db.CreateUser(&user)
	user.Id = id
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "Check data again..! "+err.Error())
		return
	}
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = user
	c.ServeJSON()
	return
}

func (c *UserController) DeleteUser() {
	userID := c.GetString(":ID")
	err := db.Delete(userID)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "User Not Found...!")
		return
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	c.Data["json"] = "Delete Success ...!"
	c.ServeJSON()
	return
}

func (c *UserController) UpdateUser() {
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "Invalid request body")
		return
	}
	password, bcErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if bcErr != nil {
		c.CustomAbort(http.StatusNotAcceptable, "Invalid request body "+bcErr.Error())
		return
	}
	user.Password = string(password)
	err = db.UpdateUser(&user)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "check data again..!")
		return
	}
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = user
	c.ServeJSON()
	return
}

func (c *UserController) GetAllUsers() {
	users, err := db.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "Something error in db or server...!")
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = users
	c.ServeJSON()
	return
}

func (c *UserController) GetUser() {
	email := c.GetString(":email")
	user, err := db.GetUser(email)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "check data again..!")
		return
	}
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = user
	c.ServeJSON()
	return
}

// ---------------------AUTH SESSION----------------------

func (c *UserController) Login() {
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var userUc models.User
	err := decoder.Decode(&userUc)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "Invalid request body")
		return
	}
	user, err := db.GetUser(userUc.Email)
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "User Not Found "+err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userUc.Password))
	if err != nil {
		log.Println(err)
		c.CustomAbort(http.StatusBadRequest, "Check Your Password "+err.Error())
		return
	}
	c.Ctx.SetCookie("username", user.Email)
	c.Data["json"] = "loggin Successfuly"
	c.ServeJSON()
}

func (c *UserController) Logout() {

	fmt.Println(c.Ctx.GetCookie("username"))
	c.DestroySession()
	c.ServeJSON()
}
