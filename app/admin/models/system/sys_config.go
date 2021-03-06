package system

import (
	"amin/common/models"
)

type SysConfig struct {
	models.Model
	ConfigName  string `json:"configName" gorm:"type:varchar(128);comment:ConfigName"`   //
	ConfigKey   string `json:"configKey" gorm:"type:varchar(128);comment:ConfigKey"`     //
	ConfigValue string `json:"configValue" gorm:"type:varchar(255);comment:ConfigValue"` //
	ConfigType  string `json:"configType" gorm:"type:varchar(64);comment:ConfigType"`    //
	Remark      string `json:"remark" gorm:"type:varchar(128);comment:Remark"`           //
	models.ControlBy
	models.ModelTime
}

func (SysConfig) TableName() string {
	return "sys_config"
}

func (e *SysConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysConfig) GetId() interface{} {
	return e.Id
}
