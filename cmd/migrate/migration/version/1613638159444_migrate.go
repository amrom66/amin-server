package version

import (
	"amin/app/admin/models/system"

	"runtime"

	"gorm.io/gorm"

	"amin/cmd/migrate/migration"
	common "amin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1613638159444Test)
}

func _1613638159444Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// TODO: 这里开始写入要变更的内容

		// TODO: 例如 修改表字段 使用过程中请删除此段代码
		err := tx.Migrator().AlterColumn(&system.SysRole{}, "create_by")
		if err != nil {
			return err
		}
		err = tx.Migrator().AlterColumn(&system.SysRole{}, "update_by")
		if err != nil {
			return err
		}

		err = tx.Migrator().AlterColumn(&system.SysMenu{}, "update_by")
		if err != nil {
			return err
		}
		err = tx.Migrator().AlterColumn(&system.SysMenu{}, "create_by")
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
