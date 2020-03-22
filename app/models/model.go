/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/10/21 17:07
*/
package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"leeblog.com/config"
	"leeblog.com/pkg/utils"
	"log"
)

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id" form:"id" uri:"id"`
	CreatedAt utils.JSONTime `json:"createdAt"`
	UpdatedAt utils.JSONTime `json:"updatedAt"`
	DeletedAt utils.JSONTime `json:"deletedAt"`
}

var db *gorm.DB

func init() {
	setup()
}

// 获得MySQL的资源链接
func DB() *gorm.DB {
	return db
}

// Setup initializes the database instance
func setup() {
	var err error
	mysql := config.MySQLConfig()
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.Username,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.DbName))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	if server := config.ServerConfig(); server.Debug {
		db.LogMode(true)
	}

	db.AutoMigrate(Admin{}, Permission{}, User{}, Role{}, News{})

	db.SingularTable(false)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
