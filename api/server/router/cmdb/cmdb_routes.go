package cmdb

import (
	"context"
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/api/server/httputils"
)

func (cm *cmdbRouter) createCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) getCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listCMDBs(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) resourceCount(c *gin.Context) {
	r := httputils.NewResponse()

	resources, err := cm.control.CMDB().GetResource(context.TODO())
	if err != nil {
		httputils.SetFailed(c, r, err)
	}

	r.Result = resources
	httputils.SetSuccess(c, r)
}
