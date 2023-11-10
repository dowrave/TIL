from photo.models import Album, Photo
from django.forms.models import inlineformset_factory 
from django import forms
from django.forms.models import modelform_factory

PhotoInlineFormSet = inlineformset_factory(Album, Photo, 
                                           fields = ['image', 'title', 'description'],
                                           extra = 2)


# 17장 - 일반 폼 
class PhotoForm(forms.Form):
    album = forms.ModelChoiceField(queryset = Album.objects.all())
    title = forms.CharField(max_length = 50)
    image = forms.ImageField()
    description = forms.CharField(label = 'Photo Description', 
                                  widget = forms.Textarea,
                                  requried = False)
    upload_date = forms.DateTimeField(label = 'Upload Date')

# 17장 - ModelForm
class PhotoForm(forms.ModelForm):
    class Meta:
        model = Photo
        fields = ['title', 'image', 'description']

# 17장 - modelform_factory
PhotoForm = modelform_factory(Photo, fields = '__all__')