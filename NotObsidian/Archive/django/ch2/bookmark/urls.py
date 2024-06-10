from django.urls import re_path 
# from bookmark.views import BookmarkLV, BookmarkDV
from bookmark.views import *

urlpatterns = [
    # class based views
    re_path(r'^$', BookmarkLV.as_view(), name = 'index'),
    re_path(r'^(?P<pk>\d+)/$', BookmarkDV.as_view(), name = 'detail'),

    # example : /add/
    re_path(r'^add/$', BookmarkCreateView.as_view(), name = 'add'),
    
    # example : /change/
    re_path(r'^change/$', BookmarkChangeLV.as_view(), name ='change'),

    # example : /99/update/
    re_path(r'^(?P<pk>[0-9]+)/update/$', BookmarkUpdateView.as_view(), name = 'update'),

    # example : /99/delete/
    re_path(r'^(?P<pk>[0-9]+)/delete/$', BookmarkDeleteView.as_view(), name='delete')
]