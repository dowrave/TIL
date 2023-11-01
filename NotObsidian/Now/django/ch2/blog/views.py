from django.shortcuts import render
from django.views.generic import ListView, DetailView
from django.views.generic.dates import ArchiveIndexView, YearArchiveView, MonthArchiveView, DayArchiveView, TodayArchiveView

from blog.models import Post

# Create your views here.

#--- ListView
class PostLV(ListView):
    model = Post # PostLV 클래스의 대상 테이블
    template_name = 'blog/post_all.html' # 지정 X 시 blog/post_list.html
    context_object_name = 'posts' # 템플릿에 넘겨주는 컨텍스트 변수명. 디폴트인 object_list도 여전히 사용 가능
    paginate_by = 2 # 한 페이지에 보여주는 객체 리스트의 숫자. 페이징 기능 활성화 & 이동 버튼 자동 활성화.

#--- DetailView
class PostDV(DetailView):
    model = Post

#--- ArchiveView
class PostAV(ArchiveIndexView):
    model = Post
    date_field = 'modify_date'

class PostYAV(YearArchiveView):
    model = Post
    date_field = 'modify_date'
    make_object_list = True # 해당 연도의 객체 리스트를 템플릿에 넘겨준다 : 템플릿에선 object_list로 사용 가능

class PostMAV(MonthArchiveView):
    model = Post
    date_field = 'modify_date'

class PostDAV(DayArchiveView):
    model = Post
    date_field = 'modify_date'

class PostTAV(TodayArchiveView):
    model = Post
    date_field = 'modify_date'
