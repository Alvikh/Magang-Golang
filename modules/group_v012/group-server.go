package group_v012

import (
	"net/http"

	"rsudlampung/middlewares/mid_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupServer interface {
	Init()
}

type groupServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewGroupServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) GroupServer {

	return &groupServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *groupServer) Init() {

	groupControl := NewGroupController(s.database)

	s.apiRoutes.GET("/"+s.version+"/group", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		ctx.JSON(200, groupControl.FindAll())
	})

	s.apiRoutes.GET("/"+s.version+"/group/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := groupControl.FindById(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.POST("/"+s.version+"/group", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := groupControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/group", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := groupControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}

	})

	s.apiRoutes.DELETE("/"+s.version+"/group/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := groupControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})
}
