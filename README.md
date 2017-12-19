# Django Model 2 Go Struct

for gorm or other ORM operation.

Depends in go code:

https://github.com/jinzhu/gorm

https://github.com/guregu/null

## Install

```bash
git clone /django2go.git
cd django2go
python setup.py install .
```

## Usage:

settings.py

```python
INSTALLED_APPS = [
    ...
    'django2go',
    ...
]

```

run CMD:

`python manage.py model2go app_name`

Example generated go code: [testapp.go](./test_project/testapp.go)

Django model:

```python

class Model1(models.Model):
    field1 = models.CharField(max_length=200)
    field_with_under_score = models.IntegerField()
    fieldWithUpperCase = models.IntegerField()
    fieldWith_Case = models.IntegerField()


```

Generated go struct:

```go

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

```