package service

import (
	"amin/core/sdk/pkg/logger"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logger.Logger
}
