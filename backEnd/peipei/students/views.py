import uuid
from django.core.cache import cache
from rest_framework import status
from rest_framework import exceptions
from rest_framework.renderers import JSONRenderer
from rest_framework.response import Response
from .serializers import StudentSerializer
from rest_framework.views import APIView
from .models import Student, School, Major
from django.http import Http404
from .constants import HTTP_ACTION_LOGIN
from .constants import HTTP_ACTION_REGISTER
from .authentications import StudentAuthentication
from .pemissions import ManagerPermission
from rest_framework.request import Request


class StudentList(APIView):
    authentication_classes = (StudentAuthentication,)
    permission_classes = (ManagerPermission,)

    def _clean(self, data):
        if not data["smoke"]:
            data["smoke"] = False
        if not data["gender"]:
            data["gender"] = False
        if not data["south_north"]:
            data["south_north"] = False
        if not data["is_super"]:
            data["is_super"] = False
        return data

    def get(self, request, format=None):
        if isinstance(request.user, Student):
            school = request.user.school
            students = Student.objects.filter(school=school)
            serializer = StudentSerializer(students, many=True)
            result = []
            for data in serializer.data:
                data = dict(data)
                data = self._clean(data)
                school_id = data["school"]
                major_id = data["major"]
                if school_id:
                    school = School.objects.get(pk=school_id).school
                    data["school"] =school
                if major_id:
                    major = Major.objects.get(pk= major_id).major
                    data["major"] = major
                result.append(data)
            return Response(result)
        else:
            raise exceptions.NotAuthenticated

    def post(self, request, format=None):
        action = request.query_params.get('action')
        if action == HTTP_ACTION_LOGIN:
            id = int(request.data.get("id"))
            name = request.data.get("name")
            school = request.data.get("school")

            try:
                student = Student.objects.get(id=id)

                if student.name == name and student.school.__str__() == school:
                    token = uuid.uuid4().hex
                    cache.set(token, student.id)
                    data = {
                        'msg': 'login success',
                        'status': '200',
                        'token': token
                    }
                    return Response(data)
                else:
                    raise exceptions.AuthenticationFailed
            except Student.DoesNotExist:
                raise exceptions.NotFound
        if action == HTTP_ACTION_REGISTER:
            school_name = request.data.get("school")
            try:
                school = School.objects.get(school=school_name)
            except:
                raise exceptions.NotFound
            request.data["school"] = school.pk
            request.data["is_super"] = False
            serializer = StudentSerializer(data=request.data)
            if serializer.is_valid():
                serializer.save()
                return Response({'msg': 'register success', 'status': 201}, status=status.HTTP_201_CREATED)
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class StudentDetail(APIView):
    authentication_classes = (StudentAuthentication,)

    def _clean(self, data):
        if not data["smoke"]:
            data["smoke"] = False
        if not data["gender"]:
            data["gender"]  =False
        if not data["south_north"]:
            data["south_north"] = False
        if not data["is_super"]:
            data["is_super"] = False
        return data

    def get_object(self, pk):
        try:
            return Student.objects.get(pk=pk)
        except Student.DoesNotExist:
            raise Http404

    def get(self, request, pk, format=None):
        if request.user.id != pk:
            raise exceptions.PermissionDenied
        serializer = StudentSerializer(request.user)
        data = dict(serializer.data)

        school_id = serializer.data.get("school")
        if school_id:
            school = School.objects.get(id= school_id).school
            data["school"] = school
        major_id  =serializer.data.get("major")
        if major_id:
            major = Major.objects.get(id= major_id).major
            data["major"] = major
        data = self._clean(data)
        return Response(data)

    def put(self, request, pk, format=None):
        if request.user.id != pk:
            raise exceptions.PermissionDenied
        print(request.user.id)
        name = request.data.get("name")
        id = request.data.get("id")
        school = request.data.get("school")
        if name != request.user.name or int(id) != request.user.id or school != request.user.school.__str__():
            raise exceptions.AuthenticationFailed
        try:
            major = Major.objects.get(major=request.data.get("major"))
            school = School.objects.get(school=school)
        except:
            raise exceptions.NotFound
        request.data["school"] = school.pk
        request.data["major"] = major.pk
        request.data["is_super"] = False
        serializer = StudentSerializer(request.user, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response({'msg': 'put success', 'status': 202}, status=status.HTTP_202_ACCEPTED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

