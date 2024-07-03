package demo

import (
	"dagger/lib/sys"
	"errors"
	"fmt"
	"net/http"

	"dagger/lib/logger"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, sys.SystemErrorResp())
	c.JSON(http.StatusOK, sys.ParamErrorResp())
	c.JSON(http.StatusOK, sys.ForbiddenErrorResp())
	c.JSON(http.StatusOK, sys.NotfoundErrorResp())
	c.JSON(http.StatusOK, sys.ErrorResp(errors.New("123")))
	c.JSON(http.StatusOK, sys.ErrorResp(sys.NewDaggerError(10001, errors.New("123"))))
	c.JSON(http.StatusOK, sys.SuccResp(123))
	fmt.Printf("%+v\n", c.Request.Context())
	logger.Debug(c.Request.Context(), "IndexHandler", "succ")
	logger.Info(c.Request.Context(), "IndexHandler", "succ")
	logger.Warn(c.Request.Context(), "IndexHandler", "succ")
	logger.Error(c.Request.Context(), "IndexHandler", errors.New("123"))
	logger.Error(c.Request.Context(), "IndexHandler", sys.NewDaggerError(10001, errors.New("123")))
	logger.ErrorWithMsg(c.Request.Context(), "IndexHandler", "succ")
	// logger.Panic(c.Request.Context(), "IndexHandler", "succ")
	// logger.Fatal(c.Request.Context(), "IndexHandler", "succ")
}
