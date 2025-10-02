package storage

import (
	"time"
)

type Pornhub struct {
	Id        int64     `xorm:"pk autoincr notnull comment('主键id') INT(11)"`
	Key       string    `xorm:"varchar(255) comment(下载视频的唯一标识)"`
	Name      string    `xorm:"varchar(255) comment(视频名称)"`
	From      string    `xorm:"varchar(255) comment(视频来源)"`
	CreatedAt time.Time `xorm:"created comment(创建时间)"`
	UpdatedAt time.Time `xorm:"updated comment(修改时间)"`
	DeletedAt time.Time `xorm:"deleted comment(删除时间)"`
}

func (p *Pornhub) InsertOne() (int64, error) {
	return GetMysql().InsertOne(p)
}

func (p *Pornhub) FindByKey() (bool, error) {
	return GetMysql().Where("`key` = ?", p.Key).Get(p)
}
