from django.urls import path,include
from rest_framework.urlpatterns import  format_suffix_patterns
from . import  views

urlpatterns = [
    path(r'',views.StudentList.as_view()),
    path(r'<int:pk>',views.StudentDetail.as_view()),
]

urlpatterns = format_suffix_patterns(urlpatterns)