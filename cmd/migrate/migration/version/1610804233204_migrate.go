package version

import (
	"amin/app/admin/models/system"

	//"amin/app/admin/models"
	"gorm.io/gorm"
	"runtime"

	"amin/cmd/migrate/migration"
	common "amin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1610804233204Test)
}

func _1610804233204Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// TODO: 这里开始写入要变更的内容
		post := system.SysPost{}
		err := tx.Model(&post).Where("status = ?", 0).Update("status", 2).Error
		if err != nil {
			return err
		}

		role := system.SysRole{}
		err = tx.Model(&role).Where("status = ?", 0).Update("status", 2).Error
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
