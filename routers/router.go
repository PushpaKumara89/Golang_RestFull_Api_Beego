package routers

import (
	"ApiBeeGo/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	// Enable CORS middleware
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:Hello")

	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/deleteUser/:ID", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/updateUser", &controllers.UserController{}, "put:UpdateUser")
	beego.Router("/getAllUsers", &controllers.UserController{}, "get:GetAllUsers")
	beego.Router("/getUser/:email", &controllers.UserController{}, "get:GetUser")

	//--------------Authentication Routers---------------------------------------
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")

	//------------------image upload-----------------------------------------
	beego.Router("/upload/image", &controllers.ImageController{}, "post:UploadImage")
	beego.Router("/images/:id", &controllers.ImageController{}, "get:GetImage")

}
