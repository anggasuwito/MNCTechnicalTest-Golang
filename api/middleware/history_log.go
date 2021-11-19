package middleware

import (
	"TechnicalTest-Golang/api/model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func CreateHistoryLog(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var historyLog model.HistoryLog
		historyLog.AccessedPath = context.FullPath()
		historyLog.IpAddress = context.ClientIP()
		historyLog.CreatedDate = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		historyLog.UpdatedDate = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
		err := db.Debug().Table("history_log").Create(&historyLog).Error
		if err != nil {
			log.Println(err)
		}
		context.Next()
	}
}
