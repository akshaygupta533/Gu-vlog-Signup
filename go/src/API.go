package main

import (
   "fmt"
   "github.com/gin-contrib/cors"                      // Why do we need this package?
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/mattn/go-sqlite3"
   _ "github.com/jinzhu/gorm/dialects/sqlite"           // If you want to use mysql or any other db, replace this line
)

var db *gorm.DB                                         // declaring the db globally
var err error

//Defining the schema for users table
type User struct {
	Name string `json:"name"`
	Username	string `json:"username" gorm:"primary_key"`
	Password string `json:"password"`
	Admin bool `json:"admin"`
	LoggedIn bool `json:"logged_in"`
}

type Msg struct{
	Message string `json:"message"`
}


func main(){
	db,err = gorm.Open("sqlite3" , "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	r := gin.Default()
	r.POST("/newuser",CreateUser)
	r.POST("/user",GetUser)
	r.GET("/users/",GetUsers)
	r.DELETE("/users/:username",DeleteUser)
	r.Use((cors.Default()))
	r.Run(":8080")    
}

func DeleteUser(c *gin.Context){
	//delete a user
	var user User
	username := c.Params.ByName("username")
	 
	   d := db.Table("users").Where("username = ?", username).Delete(&user)
	   fmt.Println(d)
	   
   	c.Header("access-control-allow-origin", "*")//enabling cors
   	c.JSON(200, gin.H{"user #" + username: "deleted"})


}

func GetUsers(c *gin.Context){
	//get all users
	var user []User
	if err := db.Table("users").Find(&user).Error; err!=nil{
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.Header("access-control-allow-origin","*")//enabling cors
		c.JSON(200,user)
	}
}

func GetUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	var jmsg Msg
	var userr User
	//check if the user exists in database
	if db.Table("users").Where("username = ?", user.Username).First(&userr).RecordNotFound()==true {
	jmsg.Message = "User Doesn't exist"
	c.Header("access-control-allow-origin", "*") //enabling cors
	c.JSON(404, jmsg)
	} else {
		if userr.Password==user.Password{
		jmsg.Message = "User logged in"
		userr.LoggedIn=!(userr.LoggedIn)
		db.Table("users").Save(&userr)
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, jmsg)
		} else{
		jmsg.Message = "Wrong username or Password"
		c.Header("access-control-allow-origin", "*")
		c.JSON(401, jmsg)

	}
}

}

func CreateUser(c *gin.Context) {
	if db.HasTable(&User{})!=true{
		db.CreateTable(&User{})
	}
	var jmsg Msg
	var user User
	var userr User
	c.BindJSON(&user)
	//check if the user already exists
	if db.Table("users").Where("username = ?", user.Username).First(&userr).RecordNotFound()==true {
		user.LoggedIn=false
		user.Admin=false
		db.Table("users").Create(&user)
	jmsg.Message = "User successfully Created"
	c.Header("access-control-allow-origin", "*") // enabling cors
	c.JSON(201, jmsg)
	} else {
		jmsg.Message = "User already exists"
		c.Header("access-control-allow-origin", "*") 
		c.JSON(404, jmsg)
	}
}

