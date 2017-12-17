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