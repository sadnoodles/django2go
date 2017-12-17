# Django Model 2 Go Struct

for gorm or other ORM operation.

depends:
https://github.com/jinzhu/gorm
https://github.com/guregu/null

## install

```bash
git clone /django2go.git
cd django2go
python setup.py install .
```

## 使用

在settings.py里增加
```python
INSTALLED_APPS = [
    ...
    'django2go',
    ...
]

```

使用

`python manage.py model2go app_name`