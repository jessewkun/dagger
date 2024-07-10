package demo

import (
	"dagger/lib/logger"
	"dagger/lib/sys"
	"dagger/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// IndexHandler godoc
//
//	@Summary		IndexHandler
//	@Description	这是一个示例接口
//	@Tags			demo
//	@Accept			x-www-form-urlencoded
//	@method			get
//	@Param			id	query	int	true	"ID"
//	@Produce		json
//	@Success		200	{object}	sys.ApiResult{data=mysql.Demo}	"成功"
//	@Failure		200	{object}	sys.ApiResult
//	@Router			/demo/v1/index [get]
func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, sys.ParamErrorResp(c))
	c.JSON(http.StatusOK, sys.ForbiddenErrorResp(c))
	c.JSON(http.StatusOK, sys.NotfoundErrorResp(c))
	c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("123")))
	c.JSON(http.StatusOK, sys.ErrorResp(c, sys.NewDaggerError(10001, errors.New("123"))))
	c.JSON(http.StatusOK, sys.SuccResp(c, 123))
	logger.Debug(c.Request.Context(), "IndexHandler", "succ")
	logger.Info(c.Request.Context(), "IndexHandler", "succ")
	logger.Warn(c.Request.Context(), "IndexHandler", "succ")
	logger.Error(c.Request.Context(), "IndexHandler", errors.New("123"))
	logger.Error(c.Request.Context(), "IndexHandler", sys.NewDaggerError(10001, errors.New("123")))
	logger.ErrorWithMsg(c.Request.Context(), "IndexHandler", "succ")
	logger.Panic(c.Request.Context(), "IndexHandler", "succ")
	logger.Fatal(c.Request.Context(), "IndexHandler", "succ")
}

// OneHandler godoc
//
// @Summary		OneHandler
// @Description	这是一个添加数据示例接口
// @Tags		demo
// @Accept		x-www-form-urlencoded
// @method		get
// @Param		id	query	int	true	"ID"
// @Produce		json
// @Success		200	{object}	sys.ApiResult{data=mysql.Demo}	"成功"
// @Failure		200	{object}	sys.ApiResult
// @Router		/demo/v1/one [get]
func OneHandler(c *gin.Context) {
	ids, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	id := cast.ToInt(ids)
	if id <= 0 {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	data, err := service.NewDemoService().GetDemoById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{"demo": data}))
}

// ListHandler godoc
//
// @Summary		ListHandler
// @Description	这是一个获取列表示例接口
// @Tags		demo
// @Accept		x-www-form-urlencoded
// @method		get
// @Param		id	query	int	true	"ID"
// @Param		pagesize	query	int	false	"分页大小，默认20"
// @Produce		json
// @Success		200	{object}	sys.ApiResult{data=[]mysql.Demo}	"成功"
// @Failure		200	{object}	sys.ApiResult
// @Router		/demo/v1/list [get]
func ListHandler(c *gin.Context) {
	ids, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	id := cast.ToInt(ids)
	if id < 0 {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	pagesize := 20
	if pagesizes, ok := c.GetQuery("pagesize"); ok {
		pagesize = cast.ToInt(pagesizes)
	}
	if pagesize > 100 {
		pagesize = 20
	}

	data, err := service.NewDemoService().GetDemoList(c.Request.Context(), id, pagesize)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{
		"list":     data,
		"pagesize": pagesize,
		"id":       id,
	}))
}

func AddHandler(c *gin.Context) {

}

func UpdateHandler(c *gin.Context) {}

func DeleteHandler(c *gin.Context) {}
