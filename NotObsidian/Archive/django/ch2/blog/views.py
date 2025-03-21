from django.shortcuts import render
from django.views.generic import ListView, DetailView, TemplateView
from django.views.generic.dates import ArchiveIndexView, YearArchiveView, MonthArchiveView, DayArchiveView, TodayArchiveView
from tagging.models import Tag, TaggedItem
from tagging.views import TaggedObjectList

from blog.models import Post

from django.views.generic.edit import FormView
from blog.forms import PostSearchForm
from django.db.models import Q
from django.shortcuts import render

from django.views.generic.edit import CreateView, UpdateView, DeleteView
from django.urls import reverse_lazy
from mysite.views import LoginRequiredMixin

# Create your views here.

#--- TemplateView
class TagTV(TemplateView):
    template_name = 'tagging/tagging_cloud.html'

#--- ListView
class PostLV(ListView):
    model = Post # PostLV 클래스의 대상 테이블
    template_name = 'blog/post_all.html' # 지정 X 시 blog/post_list.html
    context_object_name = 'posts' # 템플릿에 넘겨주는 컨텍스트 변수명. 디폴트인 object_list도 여전히 사용 가능
    paginate_by = 2 # 한 페이지에 보여주는 객체 리스트의 숫자. 페이징 기능 활성화 & 이동 버튼 자동 활성화.

class TestPostLV(ListView):
    # model = Post
    queryset = Post.objects.all()[:5]
    # template_name = 'blog/post_all.html'
    template_name = 'blog/post_test.html'
    context_object_name = 'posts'
    paginate_by = 2

    def get_queryset(self):
        return Post.objects.filter(Q(content__icontains = self.kwargs['word'])).distinct()
    
    def get_context_data(self, **kwargs):
        context = super(TestPostLV, self).get_context_data(**kwargs)
        context['SearchWord'] = self.kwargs['word']
        return context

class PostTOL(TaggedObjectList):
    model = Post
    template_name = 'tagging/tagging_post_list.html'

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

class SearchFormView(FormView):
    form_class = PostSearchForm
    template_name = 'blog/post_search.html'

    def form_valid(self, form):
        schWord = '%s' % self.request.POST['search_word']
        post_list = Post.objects.filter(Q(title__icontains = schWord) | Q(description__icontains = schWord) | Q(content__icontains = schWord)).distinct()

        context = {}
        context['form'] = form
        context['search_term'] = schWord
        context['object_list'] = post_list

        return render(self.request, self.template_name, context)
    
# CRUD 관련
class PostCreateView(LoginRequiredMixin, CreateView):
    model = Post
    fields = ['title', 'slug', 'description', 'content', 'tag']
    initial = {'slug': 'auto-filling-do-not-input'}
    success_url = reverse_lazy('blog:index')

    def form_valid(self, form):
        form.instance.owner = self.request.user
        return super(PostCreateView, self).form_valid(form)
    
class PostChangeLV(LoginRequiredMixin, ListView):
    template_name = 'blog/post_change_list.html'

    def get_queryset(self):
        return Post.objects.filter(owner = self.request.user)
    
class PostUpdateView(LoginRequiredMixin, UpdateView):
    model = Post
    fields = ['title', 'slug', 'description', 'content', 'tag']
    success_url = reverse_lazy('blog:index')

class PostDeleteView(LoginRequiredMixin, DeleteView):
    model = Post
    success_url = reverse_lazy('blog:index')