from __future__ import unicode_literals # python 2.x 지원용

from django.db import models
from django.contrib.auth.models import User

class Bookmark(models.Model):
    title = models.CharField(max_length = 100, blank = True, null = True)
    url = models.URLField('url', unique = True)
    owner = models.ForeignKey(User, on_delete=models.CASCADE, null = True)

    def __str__(self):
        return self.title
    

