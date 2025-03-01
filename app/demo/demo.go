package demo

import (
	"dagger/app/demo/dto"
	"dagger/lib/logger"
	"dagger/lib/sys"
	"dagger/model/redis"
	"dagger/service"
	"errors"
	"fmt"
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
//	@Success		200	{object}	sys.ApiResult{data=mysql.Demo}	"succ"
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
// @Success		200	{object}	sys.ApiResult{data=mysql.Demo}	"succ"
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
// @Success		200	{object}	sys.ApiResult{data=[]mysql.Demo}	"succ"
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

// AddHandler godoc
//
// @Summary		AddHandler
// @Description	这是一个添加数据示例接口
// @Tags		demo
// @Accept		json
// @method		post
// @Param		reqAddDemo	body	dto.ReqAddDemo	true	"ReqAddDemo"
// @Produce		json
// @Success		200	{object}	sys.ApiResult	"succ"
// @Failure		200	{object}	sys.ApiResult
// @Router		/demo/v1/add [post]
func AddHandler(c *gin.Context) {
	var req = dto.ReqAddDemo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	if req.Email == "" || req.Name == "" {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	id, err := service.NewDemoService().AddDemo(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{"id": id}))
}

// UpdateHandler godoc
//
// @Summary		UpdateHandler
// @Description	这是一个更新数据示例接口
// @Tags		demo
// @Accept		json
// @method		post
// @Param		reqUpdateDemo	body	dto.ReqUpdateDemo	true	"ReqUpdateDemo"
// @Produce		json
// @Success		200	{object}	sys.ApiResult	"succ"
// @Failure		200	{object}	sys.ApiResult
// @Router		/demo/v1/update [post]
func UpdateHandler(c *gin.Context) {
	var req = dto.ReqUpdateDemo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	if req.Id < 1 || req.Email == "" || req.Name == "" {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	rows, err := service.NewDemoService().UpdateDemo(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{"rows": rows}))
}

// DeleteHandler godoc
//
// @Summary		DeleteHandler
// @Description	这是一个删除数据示例接口
// @Tags		demo
// @Accept		json
// @method		post
// @Param		reqDeleteDemo	body	dto.ReqDeleteDemo	true	"ReqDeleteDemo"
// @Produce		json
// @Success		200	{object}	sys.ApiResult	"succ"
// @Failure		200	{object}	sys.ApiResult
// @Router		/demo/v1/delete [post]
func DeleteHandler(c *gin.Context) {
	var req = dto.ReqDeleteDemo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	if req.Id < 1 {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	if err := service.NewDemoService().DeleteDemo(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}

// redis
func RedisHandler(c *gin.Context) {
	var res string
	var err error
	res, err = redis.TestGet(c.Request.Context(), "test")
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
	err = redis.TestSet(c.Request.Context(), "test", "123")
	fmt.Printf("%+v\n", err)
	res, err = redis.TestGet(c.Request.Context(), "test")
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
}
