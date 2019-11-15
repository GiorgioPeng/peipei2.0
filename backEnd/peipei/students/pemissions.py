from rest_framework.permissions import BasePermission
from .models import Student
from .constants import HTTP_ACTION_LOGIN
from .constants import HTTP_ACTION_REGISTER

class ManagerPermission(BasePermission):
    def has_permission(self, request, view):
        if request.method == "POST":
            if request.query_params.get("action") == HTTP_ACTION_LOGIN:
                return True
            elif request.query_params.get("action") == HTTP_ACTION_REGISTER:
                if isinstance(request.user,Student):
                    return request.user.is_super
            return False

        elif request.method == 'GET':
            if isinstance(request.user, Student):
                return request.user.is_super
            return False
