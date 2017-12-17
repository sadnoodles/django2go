# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.template.loader import render_to_string
from django2go.models import serializer_field_mapping, null_map
from django.apps import apps
# Create your views here.


def trans_name(name):
    return name.replace('_', " ").title().replace(' ', '')



def map_field(field, null=True):
    # field, field_type, tag in fields
    name, _type, tag = trans_name(field.column), "", []
    _type = serializer_field_mapping.get(field.get_internal_type(), "unknown")
    if _type == "unknown":
        print("Field %s unknown"%field.get_internal_type())

    if null and field.null:
        _type = null_map[_type]
    
    tag += ['json:"%s"'%field.column]
    if field.primary_key:
        tag += ['gorm:"primary_key"']


    return name, _type, ' '.join(tag)

def get_app_structs(app_name, models=[], null=True):
    _models = apps.all_models[app_name]
    structs = []
    # model_name, table_name, use_prefix, fields in structs
    for mname, model in _models.iteritems():
        if models:
            if mname not in models:
                continue
        go_fields = [map_field(f, null) for f in model._meta.fields]
        n = model.__name__
        tablename = model._meta.db_table
        use_prefix = tablename.startswith(app_name+"_")
        if use_prefix:
            tablename = tablename[len(app_name)+1:]
            
        structs.append([
            n, 
            tablename, 
            use_prefix, 
            go_fields
        ])
        
    return structs


def model2go(app_name, models=[], null=True):
    env = dict(
        null=null,
        app_name=app_name,
        structs=get_app_structs(app_name, models, null),
    )
    return render_to_string("model.go", env)

