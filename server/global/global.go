package global

import (
	"bjback/config"
	"github.com/jinzhu/gorm"
)

var (
	DB     *gorm.DB      //  数据库的句柄
	CONFIG config.Server // 配置文件
)
