package kabkot_v010

import (
	"net/http"

	"rsudlampung/middlewares/mid_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KabkotServer interface {
	Init()
}

type kabkotServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewKabkotServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) KabkotServer {

	return &kabkotServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *kabkotServer) Init() {

	kabkotControl := NewKabkotController(s.database)

	s.apiRoutes.GET("/"+s.version+"/kabkot/all", func(ctx *gin.Context) {
		ctx.JSON(200, kabkotControl.FindAll())
	})

	s.apiRoutes.GET("/"+s.version+"/kabkot/byprovinsi/:provinsi_id", func(ctx *gin.Context) {
		ctx.JSON(200, kabkotControl.FindByProvinsi(ctx))
	})

	s.apiRoutes.POST("/"+s.version+"/kabkot", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := kabkotControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/kabkot", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := kabkotControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}

	})

	s.apiRoutes.DELETE("/"+s.version+"/kabkot/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := kabkotControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/kabkot/ubahstatus/:id/:status", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := kabkotControl.ChangeStatus(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}

	})
}
