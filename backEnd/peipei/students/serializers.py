from .models import Student
from rest_framework.serializers import ModelSerializer

class StudentSerializer(ModelSerializer):

    class Meta:
        model = Student
        fields= (
            "id",
            "name",
            "school",

            "is_super",

            "gender",
            "major",

            "province",
            "south_north",
            "south_north_weight",

            "sports_type",
            "sports_type_weight",

            "novels_type",
            "novels_type_weight",

            "videos_type",
            "videos_type_weight",

            "games_type",
            "games_type_weight",

            "winter_temperature",
            "summer_temperature",
            "temperature_weight",

            "sleep_time",
            "wake_time",
            "sleep_weight",

            "smoke",
            "smoke_weight",
        )