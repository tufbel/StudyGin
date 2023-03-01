package middleware

import (
	"StudyGin/src/tools/myLog"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"reflect"
)

func ExceptionCaptureMiddleware(ctx *gin.Context) {
	// 异常捕获中间件
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		switch errType := err.(type) {
		case validator.ValidationErrors:
			ctx.JSON(http.StatusBadRequest, gin.H{"detail": errType.Error()})
		case *mysql.MySQLError:
			switch errType.Number {
			case 1062:
				ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Duplicate Entry"})
			default:
				myLog.Logger.Warn(fmt.Sprintf("Mysql error msg: %v", err))
				ctx.JSON(http.StatusInternalServerError, gin.H{"detail": "Server Error"})
			}
		case error:
			switch errType.(error) {
			case gorm.ErrRecordNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"detail": "Not Found"})
			default:
				myLog.Logger.Warn(fmt.Sprintf("Error msg [%s]: %v ", reflect.TypeOf(err), err))
				ctx.JSON(http.StatusInternalServerError, gin.H{"detail": "Server Error"})
			}
		default:
			myLog.Logger.Error(fmt.Sprintf("Unknown error [%s]: %v", reflect.TypeOf(err), err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"detail": "Server Error"})
		}
	}()
	ctx.Next()
}
