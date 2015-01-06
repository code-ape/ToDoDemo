package main

import (
	log "github.com/Sirupsen/logrus"
	gin "github.com/gin-gonic/gin"
)

func init() {
	ConfigLogging()
}

func main() {
	log.Info("Configuring ToDo server.")
	r := gin.Default()
	//r.Use(gin.Logger())
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/static/html/app.html")
	})

	api := r.Group("/api")
	{
		api.POST("/login", HandleLogin)
		api.POST("/logout", HandleLogout)
		api.POST("/get-to-dos", HandleGetToDos)
		api.POST("/add-to-dos", HandlePostToDos)
		api.POST("/delete-to-dos", HandleDeleteToDos)
	}

	log.Info("Starting ToDo server.")
	r.Run(":8080")
}

func HandleLogin(c *gin.Context) {
	req := new(AuthReq)
	c.Bind(req)

	token, success := AuthUser(req)

	if success {
		c.JSON(200, gin.H{"status": "success", "token": token})
	} else {
		c.JSON(401, gin.H{"status": "failed", "token": ""})
	}
}

func HandleLogout(c *gin.Context) {
	req := new(LogoutReq)
	c.Bind(req)
	success := VerifyUserToken(req.User, req.Token)
	if !success {
		c.JSON(401, gin.H{"status": "bad token"})
		return
	}

	success = LogoutUser(req)
	if success {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		c.JSON(401, gin.H{"status": "failed"})
	}
}

func HandleGetToDos(c *gin.Context) {
	req := new(GetToDosReq)
	c.Bind(req)
	success := VerifyUserToken(req.User, req.Token)
	if !success {
		c.JSON(401, gin.H{"status": "bad token"})
		return
	}
	to_dos := GetToDos()
	users := GetUsers()
	c.JSON(200, gin.H{"status": "success", "to_dos": to_dos, "users": users})
}

func HandlePostToDos(c *gin.Context) {
	req := new(AddToDoReq)
	c.Bind(req)
	success := VerifyUserToken(req.User, req.Token)
	if !success {
		c.JSON(401, gin.H{"status": "bad token"})
		return
	}
	success = AddToDos(req.ToDos)
	success_str := "failed"
	if success {
		success_str = "success"
	}
	c.JSON(200, gin.H{"status": success_str})
}

func HandleDeleteToDos(c *gin.Context) {
	req := new(DeleteToDoReq)
	c.Bind(req)
	success := VerifyUserToken(req.User, req.Token)
	if !success {
		c.JSON(401, gin.H{"status": "bad token"})
		return
	}
	success = DeleteToDos(req.IDs)
	c.JSON(200, gin.H{"status": success})
}
