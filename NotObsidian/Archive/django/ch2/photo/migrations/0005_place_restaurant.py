# Generated by Django 4.2.6 on 2023-11-09 06:23

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('photo', '0004_publication'),
    ]

    operations = [
        migrations.CreateModel(
            name='Place',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=50)),
                ('address', models.CharField(max_length=80)),
            ],
        ),
        migrations.CreateModel(
            name='Restaurant',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='DefRestName', max_length=50)),
                ('serves_pizza', models.BooleanField(default=False)),
                ('place', models.OneToOneField(on_delete=django.db.models.deletion.CASCADE, to='photo.place')),
            ],
        ),
    ]