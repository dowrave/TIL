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

    # Example: /album/add/
    re_path(r'^album/add/$', AlbumPhotoCV.as_view(), name = 'album_add'),

    # Example: /album/change/
    re_path(r'^album/change/$', AlbumChangeLV.as_view(), name='album_change'),

    # Example: /album/99/update/
    re_path(r'^album/(?P<pk>[0-9]+)/update/$', AlbumPhotoUV.as_view(), name = 'album_update'),

    # Example: /album/99/delete/
    re_path(r'^album/(?P<pk>[0-9]+)/delete/$', AlbumDeleteView.as_view(), name = 'album_delete'),

    # Example: /photo/add/
    re_path(r'^photo/add/$', PhotoCreateView.as_view(), name = 'photo_add'),

    # Example: /photo/change/
    re_path(r'^photo/change/$', PhotoChangeLV.as_view(), name='photo_change'),

    # Example: /photo/99/update/
    re_path(r'^photo/(?P<pk>[0-9]+)/update/$', PhotoUpdateView.as_view(), name = 'photo_update'),

    # Example: /photo/99/delete/
    re_path(r'^photo/(?P<pk>[0-9]+)/delete/$', PhotoDeleteView.as_view(), name = 'photo_delete'),

]