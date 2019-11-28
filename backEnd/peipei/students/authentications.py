from django.core.cache import cache
from rest_framework.authentication import BaseAuthentication

from .models import Student


class StudentAuthentication(BaseAuthentication):
    def authenticate(self, request):
        # if request.Method =="GET":
        token = request.query_params.get("token")
        try:
            id = cache.get(token)
            user = Student.objects.get(id =id)
            return user, token
        except:
            return
