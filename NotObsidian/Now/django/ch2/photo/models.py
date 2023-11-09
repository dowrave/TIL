from django.db import models
from django.urls import reverse

from photo.fields import ThumbnailImageField

from django.contrib.auth.models import User

# Create your models here.
class Album(models.Model):
    name = models.CharField(max_length = 50, unique= True)
    description = models.CharField('One Line Description', max_length = 100, blank = True)
    owner = models.ForeignKey(User, on_delete = models.SET_NULL, null = True)

    class Meta:
        ordering = ['name']

    def __str__(self):
        return self.name
    
    def get_absolute_url(self):
        return reverse('photo:album_detail', args = (self.id, ))
    
# N:N 구축을 위한 모델
class Publication(models.Model):
    title = models.CharField(max_length = 30)
    albums = models.ManyToManyField(Album)

    def __str__(self):
        return self.title


class Photo(models.Model):
    album = models.ForeignKey(Album, on_delete= models.CASCADE)
    title = models.CharField(max_length = 50)
    image = ThumbnailImageField(upload_to= 'photo/%Y/%m')
    description = models.TextField('Photo Description', blank = True)
    upload_date = models.DateTimeField('Upload Date', auto_now_add = True)
    owner = models.ForeignKey(User, on_delete = models.SET_NULL, null = True)


    class Meta:
        ordering = ['title']

    def __str__(self):
        return self.title
    
    def get_absolute_url(self):
        return reverse('photo:photo_detail', args = (self.id, ))
    

# 1:1 구축
class Place(models.Model):
    name = models.CharField(max_length = 50)
    address = models.CharField(max_length = 80)

    def __str__(self):
        return '%s the place' % self.name
    
class Restaurant(models.Model):
    place = models.OneToOneField(Place, on_delete = models.CASCADE)
    name = models.CharField(max_length= 50,
                            default = 'DefRestName')
    serves_pizza = models.BooleanField(default = False)

    def __str__(self):
        return '%s the restaurant' % self.name