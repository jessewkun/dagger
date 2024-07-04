package demo

import (
	"dagger/lib/sys"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, sys.SystemErrorResp(c))
	// c.JSON(http.StatusOK, sys.ParamErrorResp(c))
	// c.JSON(http.StatusOK, sys.ForbiddenErrorResp(c))
	// c.JSON(http.StatusOK, sys.NotfoundErrorResp(c))
	// c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("123")))
	// c.JSON(http.StatusOK, sys.ErrorResp(c, sys.NewDaggerError(10001, errors.New("123"))))
	// c.JSON(http.StatusOK, sys.SuccResp(c, 123))
	// logger.Debug(c.Request.Context(), "IndexHandler", "succ")
	// logger.Info(c.Request.Context(), "IndexHandler", "succ")
	// logger.Warn(c.Request.Context(), "IndexHandler", "succ")
	// logger.Error(c.Request.Context(), "IndexHandler", errors.New("123"))
	// logger.Error(c.Request.Context(), "IndexHandler", sys.NewDaggerError(10001, errors.New("123")))
	// logger.ErrorWithMsg(c.Request.Context(), "IndexHandler", "succ")
	// logger.Panic(c.Request.Context(), "IndexHandler", "succ")
	// logger.Fatal(c.Request.Context(), "IndexHandler", "succ")
}
