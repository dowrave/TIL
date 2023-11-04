from django.urls import re_path 
from photo.views import *

urlpatterns = [

    # Example : /
    re_path(r'^$', AlbumLV.as_view(), name = 'index'),

    # Example : /album/ (/와 동일)
    re_path(r'^album/$', AlbumLV.as_view(), name = 'album_list'),

    # Example: /album/99/
    re_path(r'^album/(?P<pk>\d+)/$', AlbumDV.as_view(), name = 'album_detail'),

    # Example: /photo/99/
    re_path(r'^photo/(?P<pk>\d+)/$', PhotoDV.as_view(), name = 'photo_detail'),

]