package version

import (
	"amin/app/admin/models"
	"runtime"

	"gorm.io/gorm"

	"amin/cmd/migrate/migration"
	common "amin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1603465014697Test)
}

func _1603465014697Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Model(&models.Menu{}).Where("path = ?", "/api/v1/syscategoryList").Update("path", "/api/v1/syscategory").Error
		if err != nil {
			return err
		}
		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
