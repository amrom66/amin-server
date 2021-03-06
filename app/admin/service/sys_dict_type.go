package service

import (
	"errors"

	"gorm.io/gorm"

	"amin/app/admin/models/system"
	"amin/app/admin/service/dto"
	cDto "amin/common/dto"
	"amin/common/service"
)

type SysDictType struct {
	service.Service
}

// GetPage 获取列表
func (e *SysDictType) GetPage(c *dto.SysDictTypeSearch, list *[]system.SysDictType, count *int64) error {
	var err error
	var data system.SysDictType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取对象
func (e *SysDictType) Get(d *dto.SysDictTypeById, model *system.SysDictType) error {
	var err error
	var data system.SysDictType

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建对象
func (e *SysDictType) Insert(model *system.SysDictType) error {
	var err error
	var data system.SysDictType

	err = e.Orm.Model(&data).
		Create(model).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改对象
func (e *SysDictType) Update(c *system.SysDictType) error {
	var err error
	var data system.SysDictType

	db := e.Orm.Model(&data).
		Where(c.GetId()).Updates(c)
	if db.Error != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// Remove 删除
func (e *SysDictType) Remove(d *dto.SysDictTypeById, c *system.SysDictType) error {
	var err error
	var data system.SysDictType

	db := e.Orm.Model(&data).
		Where(d.GetId()).Delete(c)
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetAll 获取所有
func (e *SysDictType) GetAll(c *dto.SysDictTypeSearch, list *[]system.SysDictType) error {
	var err error
	var data system.SysDictType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
