from django.core.management.base import BaseCommand, CommandError
from django2go.views  import model2go
class Command(BaseCommand):
    help = 'Closes the specified poll for voting'

    def add_arguments(self, parser):
        parser.add_argument('app_name', nargs='+', type=str)
        # parser.add_argument('models', nargs='+', type=str)

    def handle(self, *args, **options):
        for app_name in options['app_name']:
            try:
                open("%s.go"%app_name, 'w').write(model2go(app_name).encode("utf8"))
            except:
                raise
                raise CommandError('App "%s" does not exist' % app_name)

            self.stdout.write(self.style.SUCCESS('Successfully generated model "%s"' % app_name))