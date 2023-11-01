from django.urls import re_path 
from bookmark.views import BookmarkLV, BookmarkDV

urlpatterns = [
    # class based views
    re_path(r'^$', BookmarkLV.as_view(), name = 'index'),
    re_path(r'^(?P<PK>\d+)$', BookmarkDV.as_view(), name = 'detail')
]