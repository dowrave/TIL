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
from django.urls import include, re_path # url 대신 re_path 사용 O
from django.contrib import admin
from django.conf.urls.static import static
from django.conf import settings
import os 

from mysite.views import HomeView
from mysite.views import UserCreateView, UserCreateDoneTV


urlpatterns = [
    re_path(r'^$', HomeView.as_view(), name = 'home'),

    re_path(r'^accounts/', include('django.contrib.auth.urls')),
    re_path(r'^accounts/register/$', UserCreateView.as_view(), name = 'register'),
    re_path(r'^accounts/register/done/$', UserCreateDoneTV.as_view(), name = 'register_done'),

    re_path(r'^admin/', admin.site.urls),
    re_path(r'^bookmark/', include(('bookmark.urls', 'bookmark'), namespace = 'bookmark')),
    re_path(r'^blog/', include(('blog.urls', 'blog'), namespace = 'blog')),
    re_path(r'^photo/', include(('photo.urls', 'photo'), namespace = 'photo')), 

] + static(settings.MEDIA_URL, document_root = settings.MEDIA_ROOT)
