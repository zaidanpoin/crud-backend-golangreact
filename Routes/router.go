package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zaidanpoin/crud-golang-react/controllers"
)

func ServeApps() {

	r := gin.Default()
	// 8 MiB

	r.Use(cors.Default())

	r.Static("/uploads", "./uploads")

	authRoutes := r.Group("/api/v1")
	{

		MemberRoutes(authRoutes)

	}

	r.Run(":8080")
}

func MemberRoutes(router *gin.RouterGroup) {
	router.GET("/members", controllers.GetMembers)
	router.GET("/members/:id", controllers.GetMemberByID)
	router.POST("/members", controllers.CreateMembers)
	router.DELETE("/members/:id", controllers.DeleteMembers)
	router.PATCH("/members/:id", controllers.UpdateMembers)

}
