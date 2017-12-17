package main

import (
	"encoding/json"
	"fmt"
	"strings"

	
	"github.com/guregu/null"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const TABLE_PREFIX = "testapp_"



type Model1 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Field1          string      `json:"field1"`
	FieldWithUnderScore int64       `json:"field_with_under_score"`
	Fieldwithuppercase int64       `json:"fieldWithUpperCase"`
	FieldwithCase   int64       `json:"fieldWith_Case"`
}

// TableName 使用指定的数据库表名
func (Model1) TableName() string {
	
	return TABLE_PREFIX + "model1"
	
}



type Model2 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Field1          string      `json:"field1"`
	FieldWithUnderScore int64       `json:"field_with_under_score"`
	Fieldwithuppercase int64       `json:"fieldWithUpperCase"`
	FieldwithCase   int64       `json:"fieldWith_Case"`
}

// TableName 使用指定的数据库表名
func (Model2) TableName() string {
	
	return TABLE_PREFIX + "model2"
	
}



type Model3 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Field1          string      `json:"field1"`
}

// TableName 使用指定的数据库表名
func (Model3) TableName() string {
	
	return TABLE_PREFIX + "model3"
	
}



type ModelCase2_field11 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Modelcase2Id    int64       `json:"modelcase2_id"`
	Model3Id        int64       `json:"model3_id"`
}

// TableName 使用指定的数据库表名
func (ModelCase2_field11) TableName() string {
	
	return TABLE_PREFIX + "modelcase2_field11"
	
}



type ModelCase2 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Field1          string      `json:"field1"`
	Field2          int64       `json:"field2"`
	Field3          int64       `json:"field3"`
	Field4          string      `json:"field4"`
	Field5          float64     `json:"field5"`
	Field6          string      `json:"field6"`
	Field7          string      `json:"field7"`
	Field8          string      `json:"field8"`
	Field9Id        int64       `json:"field9_id"`
	Field10Id       int64       `json:"field10_id"`
}

// TableName 使用指定的数据库表名
func (ModelCase2) TableName() string {
	
	return TABLE_PREFIX + "modelcase2"
	
}



type ModelCase3_field11 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Modelcase3Id    int64       `json:"modelcase3_id"`
	Model3Id        int64       `json:"model3_id"`
}

// TableName 使用指定的数据库表名
func (ModelCase3_field11) TableName() string {
	
	return TABLE_PREFIX + "modelcase3_field11"
	
}



type ModelCase3 struct {
	Id              int64       `json:"id" gorm:"primary_key"`
	Field1          null.String  `json:"field1"`
	Field2          null.Int    `json:"field2"`
	Field3          null.Int    `json:"field3"`
	Field4          null.String  `json:"field4"`
	Field5          null.Float  `json:"field5"`
	Field6          null.String  `json:"field6"`
	Field7          null.String  `json:"field7"`
	Field8          null.String  `json:"field8"`
	Field9Id        null.Int    `json:"field9_id"`
	Field10Id       null.Int    `json:"field10_id"`
}

// TableName 使用指定的数据库表名
func (ModelCase3) TableName() string {
	
	return TABLE_PREFIX + "modelcase3"
	
}



// MigrateTabels 迁移数据库表
func MigrateTabels(db *gorm.DB) {
	
	db.AutoMigrate(&Model1{})
	db.AutoMigrate(&Model2{})
	db.AutoMigrate(&Model3{})
	db.AutoMigrate(&ModelCase2_field11{})
	db.AutoMigrate(&ModelCase2{})
	db.AutoMigrate(&ModelCase3_field11{})
	db.AutoMigrate(&ModelCase3{})

}

func Unmarshal(name string, data []byte) (interface{}, error) {
	name = strings.ToLower(name)
	switch name {
	case "model1":
		var source Model1
		err := json.Unmarshal(data, &source)
		return &source, err
	
	case "model2":
		var source Model2
		err := json.Unmarshal(data, &source)
		return &source, err
	
	case "model3":
		var source Model3
		err := json.Unmarshal(data, &source)
		return &source, err
	
	case "modelcase2_field11":
		var source ModelCase2_field11
		err := json.Unmarshal(data, &source)
		return &source, err
	
	case "modelcase2":
		var source ModelCase2
		err := json.Unmarshal(data, &source)
		return &source, err
	
	case "modelcase3_field11":
		var source ModelCase3_field11
		err := json.Unmarshal(data, &source)
		return &source, err
	
	case "modelcase3":
		var source ModelCase3
		err := json.Unmarshal(data, &source)
		return &source, err
	
	}
	return nil, fmt.Errorf("Not a validated source type.")

}

func main() {
	db, err := gorm.Open("sqlite3", "testapp.db")
	// db, err := gorm.Open("mysql", "root:root@/db?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Migrate the schema
	MigrateTabels(db)

	fmt.Println("Done.")
}
