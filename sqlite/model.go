package sqlite

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type YtdlpHistory struct {
	Id        int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Name      string         `gorm:"comment:视频名称"`
	Url       string         `gorm:"comment:视频链接"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (yt *YtdlpHistory) Sync() {
	log.Printf("开始同步表结构")

	// 使用 GORM 自动迁移功能创建表
	if err := GetSqlite().AutoMigrate(yt); err != nil {
		log.Printf("同步表结构失败: %v", err)
	}
	log.Printf("同步表结构完成")
}
func (yt *YtdlpHistory) InsertOne() (int64, error) {
	result := GetSqlite().Create(&yt)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, result.Error
}

/*
根据原始url判断是否下载过
*/
func (yt *YtdlpHistory) FindByOriginURL() (bool, error) {
	result := GetSqlite().Where("url = ?", yt.Url).First(&yt)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

/*
根据自定义文件名判断是否下载过
*/
func (yt *YtdlpHistory) FindByFilename() (bool, error) {
	result := GetSqlite().Where("name = ?", yt.Name).First(&yt)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

/*
Check if a record exists with the same name
*/
func (yt *YtdlpHistory) ExistsByName() (bool, error) {
	var count int64
	err := GetSqlite().Model(&YtdlpHistory{}).Where("name = ?", yt.Name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

/*
Check if a record exists with the same URL
*/
func (yt *YtdlpHistory) ExistsByUrl() (bool, error) {
	var count int64
	err := GetSqlite().Model(&YtdlpHistory{}).Where("url = ?", yt.Url).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
