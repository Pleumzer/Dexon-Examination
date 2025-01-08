package main


import (

"github.com/gin-gonic/gin"

"github.com/Pleumzer/Dexon-Examination/controller"

"github.com/Pleumzer/Dexon-Examination/entity"

)


func main() {

entity.SetupDatabase()

r := gin.Default()

r.Use(CORSMiddleware())

// Info Routes

r.GET("/infos", controller.ListInfo)

r.GET("/info/:id", controller.GetInfo)

r.POST("/infos", controller.CreateInfo)

r.PATCH("/infos", controller.UpdateInfo)

r.DELETE("/infos/:id", controller.DeleteInfo)

// CML Routes

r.GET("/cmls", controller.ListCml)

r.GET("/cml/:id", controller.GetCml)

r.POST("/cmls", controller.CreateCml)

r.PATCH("/cmls", controller.UpdateCml)

r.DELETE("/cmls/:id", controller.DeleteCml)

// TestPoint Routes

r.GET("/test_points", controller.ListTestPoint)

r.GET("/test_point/:id", controller.GetTestPoint)

r.POST("/tesst_points", controller.CreateTestPoint)

r.PATCH("/test_points", controller.UpdateTestPoint)

r.DELETE("/test_points/:id", controller.DeleteTestPoint)

// Thickness Routes

r.GET("/thicknesses", controller.ListThickness)

r.GET("/thickness/:id", controller.GetThickness)

r.POST("/thicknesses", controller.CreateThickness)

r.PATCH("/thicknesses", controller.UpdateThickness)

r.DELETE("/thicknesses/:id", controller.DeleteThickness)

// Run the server

r.Run()

}


func CORSMiddleware() gin.HandlerFunc {

return func(c *gin.Context) {

c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")


if c.Request.Method == "OPTIONS" {

c.AbortWithStatus(204)

return

}


c.Next()

}

}