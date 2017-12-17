# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import models

# Create your models here.

IntegerField = "int64"
BooleanField = "bool"
CharField = "string"
DateField = "string"
DateTimeField = "string"
EmailField = "string"
DecimalField = "float"
ModelField = "unknown"
FileField = "string"
ImageField = "string"
FloatField = "float64"
NullBooleanField = "int"
URLField = "string"
TimeField = "string"
IPAddressField = "string"
FilePathField = "string"

# null_map = {
#     "string":"sql.NullString",
#     "unknown":"unknown",
#     "int":"sql.NullInt64",
#     "int64":"sql.NullInt64",
#     "bool":"sql.NullBool",
#     "float64":"sql.NullFloat64"
# }

null_map = {
    "string":"null.String",
    "unknown":"unknown",
    "int":"null.Int",
    "int64":"null.Int",
    "bool":"null.Bool",
    "float64":"null.Float"
}

serializer_field_mapping = {
    'AutoField': IntegerField,
    'ForeignKey': IntegerField,
    'OneToOneField': IntegerField,
    'ManyToManyField': IntegerField,
    'BigIntegerField': IntegerField,
    'BooleanField': BooleanField,
    'CharField': CharField,
    'CommaSeparatedIntegerField': CharField,
    'DateField': DateField,
    'DateTimeField': DateTimeField,
    'DecimalField': DecimalField,
    'EmailField': EmailField,
    'Field': ModelField,
    'FileField': FileField,
    'FloatField': FloatField,
    'ImageField': ImageField,
    'IntegerField': IntegerField,
    'NullBooleanField': NullBooleanField,
    'PositiveIntegerField': IntegerField,
    'PositiveSmallIntegerField': IntegerField,
    'SmallIntegerField': IntegerField,
    'TextField': CharField,
    'TimeField': TimeField,
    'URLField': URLField,
    'GenericIPAddressField': IPAddressField,
    'FilePathField': FilePathField,
}

extras = {
    'BinaryField': CharField,
}

serializer_field_mapping.update(extras)
