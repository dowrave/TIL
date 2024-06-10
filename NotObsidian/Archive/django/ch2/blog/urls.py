from django.urls import include, re_path # url 대신 re_path 사용 O
from blog.views import *

urlpatterns = [
    
    # Example: /
    re_path(r'^$', PostLV.as_view(), name = 'index'),

    # Example: /post/ (same as /)
    re_path(r'^post/$', PostLV.as_view(), name = 'post_list'),

    # Example: /post/django-example/
    re_path(r'^post/(?P<slug>[-\w]+)/$', PostDV.as_view(), name = 'post_detail'),

    # Example: /archive/
    re_path(r'archive/$', PostAV.as_view(), name = 'post_archive'),

    # Example: /2012/
    re_path(r'^(?P<year>\d{4})/$', PostYAV.as_view(), name = 'post_year_archive'),

    # Example : /2012/nov/
    re_path(r'^(?P<year>\d{4})/(?P<month>[a-z]{3})/$', PostMAV.as_view(), name = 'post_month_archive'),

    # Example : /2012/nov/10/
    re_path(r'^(?P<year>\d{4})/(?P<month>[a-z]{3})/(?P<day>\d{1,2})/$', PostDAV.as_view(), name = 'post_day_archive'),

    # Example: /today/
    re_path(r'^today/$', PostTAV.as_view(), name = 'post_today_archive'),

    # Example : /tag/
    re_path(r'^tag/$', TagTV.as_view(), name = 'tag_cloud'),
    
    # Example : /tag/tagname/
    re_path(r'^tag/(?P<tag>[^/]+(?u))/$', PostTOL.as_view(), name = 'tagged_object_list'),

    # Example : /search/
    re_path(r'^search/$', SearchFormView.as_view(), name = 'search'),

    # Example : /add/
    re_path(r'^add/$', PostCreateView.as_view(), name='add'),

    # Example : /change/
    re_path(r'^change/$', PostChangeLV.as_view(), name = 'change'),

    # Example : /99/update/
    re_path(r'^(?P<pk>[0-9]+)/update/$', PostUpdateView.as_view(), name = 'update'),

    # Example : /99/delete/
    re_path(r'^(?P<pk>[0-9]+)/delete/$', PostDeleteView.as_view(), name = 'delete'),

    # Example: /test/word/
    re_path(r'^test/(?P<word>[\w]+)/$', TestPostLV.as_view(), name='post_test'),

]