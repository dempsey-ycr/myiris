package tool

import (
	"myiris/cache/session"
	"myiris/library/logger"
)

type User struct {
	Name string `bson:"name"`
	Age  int32  `bson:"age"`
	Addr string `bson:"addr"`
}

type MysqlUserInfo struct {
	CountryCode int32 `xorm:"countryCode"`
	PhoneNo     int64 `xorm:"phoneNo"`
	UserId      int64 `xorm:"userId"` // primary key
	CreatedAt   int64 `xorm:"createdAt"`
}

// 该方法的返回值即为MysqlUserInfo 结构对应的表名
func (*MysqlUserInfo) TableName() string {
	return "users"
}

func Stores_test() {
	record := &User{
		Name: "xoxifu",
		Age:  15,
		Addr: "guizhou",
	}
	cr := session.Mongo.Database(session.Conf.MongoDB).Collection("questions")
	res, err := cr.InsertOne(nil, record)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(res)

	// redis test
	if _, err = session.Redis.Sadds("zxu", "redis", "test", "aaa", 1,2,3); err != nil {
		logger.Error(err)
		return
	}

	// mysql test

	//_ = session.Mysql.DropTables("users")
	//
	//createSql := `CREATE TABLE users (   userId bigint(13) PRIMARY KEY,  countryCode  int   NOT NULL COMMENT "国际区号",    phoneNo     bigint(13)  NOT NULL unique COMMENT "手机号",  createdAt  bigint(13)    NOT NULL COMMENT "注册时间");`
	//_, err = session.Mysql.Exec(createSql)
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}

	row := &MysqlUserInfo{
		CountryCode: 86,
		PhoneNo:     18795394022,
		UserId:      100007,
		CreatedAt:   1525909008,
	}
	n, err := session.Mysql.InsertOne(row)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("res count:", n)
}