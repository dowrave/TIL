"""
URL configuration for mysite project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/4.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
# from django.conf.urls import url # 아예 없어짐(4버전부터)
from django.urls import re_path # url 대신 re_path 사용 O

from django.contrib import admin
import os 
from bookmark.views import BookmarkLV, BookmarkDV

urlpatterns = [
    re_path(r'^admin/', admin.site.urls),
    # Class-based views for Bookmark app
    re_path(r'^bookmark/$', BookmarkLV.as_view(), name = 'index'),
    re_path(r'^bookmark/(?P<pk>\d+)/$', BookmarkDV.as_view(), name = 'detail'),
]

