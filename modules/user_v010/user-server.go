package user_v010

import (
	"net/http"

	"rsudlampung/middlewares/mid_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServer interface {
	Init()
}

type userServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewUserServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) UserServer {

	return &userServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *userServer) Init() {

	userControl := NewUserController(s.database)

	s.apiRoutes.GET("/"+s.version+"/user", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		ctx.JSON(200, userControl.FindAll())
	})

	s.apiRoutes.GET("/"+s.version+"/user/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := userControl.FindByNik(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.POST("/"+s.version+"/user", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := userControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/user", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := userControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}

	})

	s.apiRoutes.DELETE("/"+s.version+"/user/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := userControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})
}
