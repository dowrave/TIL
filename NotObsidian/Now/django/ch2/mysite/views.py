from django.views.generic.base import TemplateView

# auth 관련
from django.views.generic.edit import CreateView
from django.contrib.auth.forms import UserCreationForm
from django.urls import reverse_lazy

from django.contrib.auth.decorators import login_required

# 홈페이지
class HomeView(TemplateView):
    template_name = 'home.html'

# auth 관련
class UserCreateView(CreateView):
    template_name = 'registration/register.html'
    form_class = UserCreationForm
    success_url = reverse_lazy('register_done')

class UserCreateDoneTV(TemplateView):
    template_name = 'registration/register_done.html'

# login_required() 데코레이터 제공 클래스
class LoginRequiredMixin(object):
    @classmethod
    def as_view(cls, **init_kwargs):
        view = super(LoginRequiredMixin, cls).as_view(**init_kwargs)
        return login_required(view)