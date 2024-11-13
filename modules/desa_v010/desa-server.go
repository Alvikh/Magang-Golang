package desa_v010

import (
	"net/http"
	"rsudlampung/middlewares/mid_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DesaServer interface {
	Init()
}

type desaServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewDesaServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) DesaServer {
	return &desaServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *desaServer) Init() {
	desaControl := NewDesaController(s.database)

	s.apiRoutes.GET("/"+s.version+"/desa/all", func(ctx *gin.Context) {
		desas := desaControl.FindAll(ctx)
		ctx.JSON(http.StatusOK, gin.H{"data": desas})
	})

	s.apiRoutes.GET("/"+s.version+"/desa/bykecamatan/:kecamatan_id", func(ctx *gin.Context) {
		desas := desaControl.FindByKecamatan(ctx)
		ctx.JSON(http.StatusOK, gin.H{"data": desas})
	})

	s.apiRoutes.POST("/"+s.version+"/desa", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := desaControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"data": result})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/desa", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := desaControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Desa berhasil diperbarui"})
		}
	})

	s.apiRoutes.DELETE("/"+s.version+"/desa/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := desaControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Desa berhasil dihapus"})
		}
	})
}
