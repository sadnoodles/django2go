package main

import (
	"encoding/json"
	"fmt"
	"strings"

	{%if null%}
	"github.com/guregu/null"{%endif%}
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const TABLE_PREFIX = "{{app_name}}_"

{% for model_name, table_name, use_prefix, fields in structs %}

type {{model_name}} struct {{% for field, field_type, tag in fields %}
	{{field | ljust:15}} {{field_type| ljust:10}}  {%if tag%}`{{tag|safe }}`{%endif%}{%endfor%}
}

// TableName 使用指定的数据库表名
func ({{model_name}}) TableName() string {
	{% if use_prefix %}
	return TABLE_PREFIX + "{{table_name}}"
	{% else %}
	return "{{table_name}}"
	{% endif %}
}

{% endfor %}

// MigrateTabels 迁移数据库表
func MigrateTabels(db *gorm.DB) {
	{% for model_name, _, _, _ in structs %}
	db.AutoMigrate(&{{model_name}}{}){% endfor %}

}

func Unmarshal(name string, data []byte) (interface{}, error) {
	name = strings.ToLower(name)
	switch name {{% for model_name, _, _, _ in structs %}
	case "{{model_name|lower}}":
		var source {{model_name}}
		err := json.Unmarshal(data, &source)
		return &source, err
	{% endfor %}
	}
	return nil, fmt.Errorf("Not a validated source type.")

}

func main() {
	db, err := gorm.Open("sqlite3", "{{app_name}}.db")
	// db, err := gorm.Open("mysql", "root:root@/db?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Migrate the schema
	MigrateTabels(db)

	fmt.Println("Done.")
}
