package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/lanyeit/wg-gen-web/api/v1/auth"
	"github.com/lanyeit/wg-gen-web/api/v1/client"
	"github.com/lanyeit/wg-gen-web/api/v1/server"
	"github.com/lanyeit/wg-gen-web/api/v1/status"
)

// ApplyRoutes apply routes to gin router
func ApplyRoutes(r *gin.RouterGroup, private bool) {
	v1 := r.Group("/v1.0")
	{
		if private {
			client.ApplyRoutes(v1)
			server.ApplyRoutes(v1)
			status.ApplyRoutes(v1)
		} else {
			auth.ApplyRoutes(v1)
		}
	}
}
