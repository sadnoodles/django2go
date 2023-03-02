# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import models

class Model1(models.Model):
    field1 = models.CharField(max_length=200)
    field_with_under_score = models.IntegerField()
    fieldWithUpperCase = models.IntegerField()
    fieldWith_Case = models.IntegerField()


class Model2(models.Model):
    field1 = models.CharField(max_length=200)
    field_with_under_score = models.IntegerField()
    fieldWithUpperCase = models.IntegerField()
    fieldWith_Case = models.IntegerField()


class Model3(models.Model):
    field1 = models.CharField(max_length=200)

class ModelCase2(models.Model):
    field1 = models.CharField(max_length=200)
    field2 = models.IntegerField()
    field3 = models.BigIntegerField()
    field4 = models.BinaryField()
    field5 = models.FloatField()
    field6 = models.TextField()
    field7 = models.TimeField()
    field8 = models.DateTimeField()
    field9 = models.ForeignKey(Model1, on_delete=models.CASCADE)
    field10 = models.OneToOneField(Model2, on_delete=models.CASCADE)
    field11 = models.ManyToManyField(Model3)
    

class ModelCase3(models.Model):
    field1 = models.CharField(null=True,max_length=200)
    field2 = models.IntegerField(null=True)
    field3 = models.BigIntegerField(null=True)
    field4 = models.BinaryField(null=True)
    field5 = models.FloatField(null=True)
    field6 = models.TextField(null=True)
    field7 = models.TimeField(null=True)
    field8 = models.DateTimeField(null=True)
    field9 = models.ForeignKey(Model1,null=True, on_delete=models.SET_NULL)
    field10 = models.OneToOneField(Model2,null=True, on_delete=models.SET_NULL)
    field11 = models.ManyToManyField(Model3)
    