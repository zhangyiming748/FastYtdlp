package storage

import (
	"time"
)

type Video struct {
	Id        int64     `xorm:"pk autoincr notnull comment('主键id') INT(11)"`
	Name      string    `xorm:"Varchar(255) comment(视频名称)"`
	Url       string    `xorm:"Varchar(255) comment(视频链接)"`
	CreatedAt time.Time `xorm:"created comment(创建时间)"`
	UpdatedAt time.Time `xorm:"updated comment(修改时间)"`
	DeletedAt time.Time `xorm:"deleted comment(删除时间)"`
}

func (v *Video) InsertOne() (int64, error) {
	return GetMysql().InsertOne(v)
}

func (v *Video) FindByUrl() (bool, error) {
	return GetMysql().Where("uri = ?", v.Url).Get(v)
}
func (v *Video) FindByName() (bool, error) {
	return GetMysql().Where("name = ?", v.Name).Get(v)
}
