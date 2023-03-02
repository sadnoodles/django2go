from django.core.management.base import BaseCommand, CommandError
from django2go.views  import model2go

import six 
class Command(BaseCommand):
    help = 'Generate structs from models.py from given app.'

    def add_arguments(self, parser):
        parser.add_argument('app_name', nargs='+', type=str)
        parser.add_argument('-c', '--use_column_name', help='Use DB column name or model field name(default).', action='store_true')

    def handle(self, *args, **options):
        for app_name in options['app_name']:
            code = model2go(app_name, use_column_name=options['use_column_name'])
            if six.PY2:
                open("%s.go"%app_name, 'w').write(code.encode('utf-8'))
            else:
                open("%s.go"%app_name, 'w', encoding='utf-8').write(code)

            self.stdout.write(self.style.SUCCESS('Successfully generated model "%s"' % app_name))