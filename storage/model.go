package storage

import (
	"time"
)

type Video struct {
	Id        int64     `xorm:"pk autoincr notnull comment('主键id') INT(11)"`
	Keyword   string    `xorm:"varchar(255) comment(下载视频的唯一标识)"`
	Name      string    `xorm:"TEXT(255) comment(视频名称)"`
	From      string    `xorm:"varchar(255) comment(视频来源)"`
	Url       string `xorm:"varchar(255) comment(视频链接)"`
	CreatedAt time.Time `xorm:"created comment(创建时间)"`
	UpdatedAt time.Time `xorm:"updated comment(修改时间)"`
	DeletedAt time.Time `xorm:"deleted comment(删除时间)"`
}

func (v *Video) InsertOne() (int64, error) {
	return GetMysql().InsertOne(v)
}

func (v *Video) FindByKeyword() (bool, error) {
	return GetMysql().Where("keyword = ?", v.Keyword).Get(v)
}
func (v *Video) FindByUrl() (bool, error) {
	return GetMysql().Where("uri = ?", v.Url).Get(v)
}
