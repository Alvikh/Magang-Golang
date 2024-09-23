package provinsi_v010

import (
	"net/http"

	"rsudlampung/middlewares/mid_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProvinsiServer interface {
	Init()
}

type provinsiServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewProvinsiServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) ProvinsiServer {

	return &provinsiServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *provinsiServer) Init() {

	provinsiControl := NewProvinsiController(s.database)

	s.apiRoutes.GET("/"+s.version+"/provinsi/all", func(ctx *gin.Context) {
		ctx.JSON(200, provinsiControl.FindAll())
	})

	s.apiRoutes.POST("/"+s.version+"/provinsi", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := provinsiControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/provinsi", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := provinsiControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}

	})

	s.apiRoutes.DELETE("/"+s.version+"/provinsi/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := provinsiControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})
}
