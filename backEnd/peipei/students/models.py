from django.db import models


class  School(models.Model):
    '''
    学校信息
    '''
    school = models.CharField("学校名称", unique=True, max_length=20, null=False, blank=False)
    province = models.CharField("省份", max_length=20, null=False, blank=False)
    city = models.CharField("城市",max_length=20, null = False, blank=False)
    remark = models.TextField("备注",max_length=100 ,null=True, blank=True)
    add_time = models.DateTimeField("添加时间", null=True, blank=True,auto_now_add=True)

    def __str__(self):
        return self.school

    class Meta:
        verbose_name = '学校信息表'
        verbose_name_plural = '学校信息表'


# 通过views函数进行修改，每增加一个或删除一个时，获取已有数据，修改后专业的性别人数，再更新
class Major(models.Model):
    '''
    学校专业信息
    '''
    school = models.ManyToManyField(School, verbose_name="学校")
    major = models.CharField("专业", max_length=30, null=False, blank=False)
    # maleNumber = models.IntegerField("男生人数", default=0)
    # femaleNumber = models.IntegerField("女生人数", default=0)
    add_time = models.DateTimeField("添加时间", auto_now= True)

    def __str__(self):
        return self.major

    class Meta:
        verbose_name = '学校专业信息表'
        verbose_name_plural = '学校专业信息表'


class Sports(models.Model):
    '''
    学生运动信息
    '''
    SPORTS_CHOICES = (
        (1, "跑步"),
        (2, "羽毛球"),
        (3, "篮球"),
        (4, "足球"),
        (5, "自行车"),
        (6, "游泳"),
        (7, "撸铁"),
        (8, "兵乓球"),
        (9, "瑜伽舞蹈"),
    )
    sport_type = models.IntegerField("运动", primary_key=True, choices=SPORTS_CHOICES)
    add_time = models.DateTimeField("运动信息添加时间", auto_now_add=True)

    def __str__(self):
        return dict(self.SPORTS_CHOICES)[self.sport_type]

    class Meta:
        verbose_name = ' 学生运动信息表'
        verbose_name_plural = ' 学生运动信息表'


class Novels(models.Model):
    '''
    学生阅读小说类型信息
    '''

    NOVEL_CHOICES = (
        (1, "技术知识类"),
        (2, "文学类小说"),
        (3, "漫画"),
        (4, "言情小说"),
        (5, "玄幻小说"),
        (6, "科幻小说"),
        (7, "武侠小说"),
        (8, "推理小说"),
        (9, "恐怖小说"),
        (10, "网游小说"),
        (11, "二次元小说"),
    )
    # 多选问题
    novel_type = models.IntegerField("小说类型",  primary_key=True, choices=NOVEL_CHOICES)
    add_time = models.DateTimeField("小说类型信息添加时间", auto_now_add= True)

    def __str__(self):
        return dict(self.NOVEL_CHOICES)[self.novel_type]

    class Meta:
        verbose_name = ' 学生阅读小说信息表'
        verbose_name_plural = ' 学生阅读小说信息表'


class Games(models.Model):
    '''
    学生游戏类型信息
    '''

    GAMES_CHOICES = (
        (1, "moba"),
        (2, "角色扮演"),
        (3, "射击类"),
        (4, "棋牌类"),
        (5, "即时战略"),
        (6, "沙盒类"),
        (7, "竞速游戏"),
        (8, "王者荣耀"),
        (9, "和平精英"),
        (10, "第五人格"),
        (11, "捉精灵"),
        (12, "阴阳师"),
        (13, "明日方舟"),
        (14, "炉石"),
        (15, "崩坏3"),
        (16, "FGO"),
    )
    # 多选问题
    game_type = models.IntegerField("游戏类型", primary_key=True, choices=GAMES_CHOICES)
    add_time = models.DateTimeField("游戏类型信息添加时间", auto_now_add= True)

    def __str__(self):
        return dict(self.GAMES_CHOICES)[self.game_type]

    class Meta:
        verbose_name = ' 学生游戏类型信息表'
        verbose_name_plural = ' 学生游戏类型信息表'


class Videos(models.Model):
    '''
    学生看剧类型信息
    '''

    VIDEOS_CHOICES = (
        (1, "泰剧"),
        (2, "日韩"),
        (3, "欧美"),
        (4, "短视频"),
        (5, "游戏类"),
        (6, "音乐类"),
        (7, "其他类"),
        (8, "体育"),
        (9, "歌唱"),
        (10, "美食"),
        (11, "情感"),
        (12, "说唱"),
        (13, "真人秀"),

    )
    # 多选问题
    video_type = models.IntegerField("看剧类型",primary_key=True, choices=VIDEOS_CHOICES)
    addTime = models.DateTimeField("看剧类型信息添加时间", auto_now_add=True)

    def __str__(self):
        return dict(self.VIDEOS_CHOICES)[self.video_type]

    class Meta:
        verbose_name = ' 学生看剧类型信息表'
        verbose_name_plural = ' 学生看剧类型信息表'


class Dormitory(models.Model):
    '''
    分配宿舍成员信息
    '''
    DORMITORY_CHOICE = (
        (False, "女生宿舍"),
        (True, "男生宿舍"),
    )
    school = models.ForeignKey(School, verbose_name="学校", on_delete=models.CASCADE)
    gender = models.BooleanField(verbose_name="宿舍类别", choices=DORMITORY_CHOICE)
    number = models.IntegerField(verbose_name="宿舍号")
    remark = models.CharField(verbose_name='备注', null=True, blank=True, max_length=100)
    addTime = models.DateTimeField("信息更新时间", null=True, blank=True, auto_now=True)

    @property
    def members(self):
        mates = self.student_set
        return  [each.name for each in mates]

    def __str__(self):
        return self.school, ":", dict(self.DORMITORY_CHOICE)[self.gender], ":", self.number

    class Meta:
        verbose_name = ' 分配宿舍成员信息表'
        verbose_name_plural = ' 分配宿舍成员信息表'


class Student(models.Model):
    '''
    学生基本信息
    '''
    #  基本映射关系表
    GENDER_CHOICES = ((True, '男'), (False, '女'))
    SOUTHNORTH_CHOICES = ((True, '南方'), (False, '北方'))
    SLEEP_TIME = ((1,'21:00之前'),(2,'21:01-22:00'),(3,'22:01-23:00'),(4,'23:01-23:59'),(5,'24:00之后'))
    WAKE_TIME =((1,'7:00之前'),(2,'7:01-8:00'),(3,'8:01-9:00'),(4,'9:01-10:00'),(5,'10:01之后'))
    SMOKE_CHOICES = ((False, '不吸烟'), (True, '吸烟'))
    _attr_list=["id","name","is_super","school","gender","major","province","south_north","south_north_weight",
                "sports","sports_weight","novels_type","novels_type_weight","videos_type","videos_type_weight",
                "games_type","games_type_weight","winter_temperature","summer_temperature","temperature_weight",
                "sleep_time","wake_time","sleep_weight","smoke","smoke_weight","add_time"]

    # 用于验证身份的两个信息
    id = models.IntegerField(verbose_name="学生学号", primary_key=True)
    name = models.CharField(verbose_name="学生姓名", max_length=30, null= False, blank= False)

    # 用于获得权限的信息
    is_super = models.BooleanField(verbose_name="超级管理员",default=False, null=True,blank=True)

    # 绑定的学校和宿舍
    dormitory = models.ForeignKey(Dormitory,verbose_name="宿舍号",  null=True, blank=False, on_delete=models.CASCADE)
    school = models.ForeignKey(School,verbose_name="学校名称",on_delete=models.CASCADE)

    gender = models.BooleanField(verbose_name="性别", null=True, blank=False, choices=GENDER_CHOICES)
    major = models.ForeignKey(Major, verbose_name="专业", on_delete=models.CASCADE, null=True, blank=True)

    province = models.CharField(verbose_name="学生省份", max_length=10, null=True, blank = True)
    south_north = models.BooleanField(verbose_name="南北方", null=True, blank=True, choices=SOUTHNORTH_CHOICES)
    south_north_weight = models.IntegerField(verbose_name="南北方权重", null=True, blank=True)

    sports_type = models.ManyToManyField(Sports, verbose_name="运动",blank=True)
    sports_type_weight = models.IntegerField(verbose_name="运动权重", null=True, blank=True)

    novels_type = models.ManyToManyField(Novels, verbose_name="小说类型",blank=True)
    novels_type_weight = models.IntegerField(verbose_name="小说类型权重", null=True, blank=True)

    videos_type = models.ManyToManyField(Videos,verbose_name="看剧类型",blank=True)
    videos_type_weight = models.IntegerField(verbose_name="看剧类型权重", null= True, blank = True)

    games_type = models.ManyToManyField(Games,verbose_name="游戏类型",blank=True)
    games_type_weight = models.IntegerField(verbose_name="游戏类型权重", null= True, blank = True)

    winter_temperature = models.IntegerField(verbose_name="冬天空调合适温度", null=True, blank=True)
    summer_temperature = models.IntegerField(verbose_name="夏天空调合适温度", null=True, blank=True)
    temperature_weight = models.IntegerField(verbose_name="空调温度权重", null=True, blank=True)

    sleep_time = models.IntegerField(verbose_name="睡觉时间", null=True, blank=True, choices=SLEEP_TIME)
    wake_time = models.IntegerField(verbose_name="起床时间", null=True, blank=True, choices=WAKE_TIME)
    sleep_weight = models.IntegerField(verbose_name="睡眠重要性权重", null=True, blank=True)

    smoke = models.BooleanField(verbose_name="吸烟", null=True, blank=False, choices=SMOKE_CHOICES)
    smoke_weight = models.IntegerField(verbose_name="吸烟权重", null=True, blank=True)

    add_time = models.DateTimeField(verbose_name="学生信息更新时间", null=True, blank=True, auto_now=True)

    @property
    def integrity(self):
        for x in self._attr_list:
            if getattr(self,x) is None:
                return False
        return True

    def __str__(self):
        return self.school.school,":",self.id,":",self.name

    class Meta:
        verbose_name = ' 学生基本信息表'
        verbose_name_plural = ' 学生基本信息表'




# class SleepTime(models.Model):
#     '''
#     学生睡觉时间信息
#     '''
#     sl_id = models.AutoField(primary_key=True)
#     wakeTime = models.CharField("起床时间", null=False, blank = False)
#     sleepTime = models.CharField("睡觉时间", null= False, blank = False)
#     addTime = models.DateTimeField("睡觉信息添加时间", null= True, blank = False)


# class SouthNorth(models.Model):
#     '''
#     学生所在地南北方信息
#     '''
#     sn_id=models.AutoField(primary_key=True)
#     southNorth = models.BooleanField("南北方", null= True, blank = False)
#     addTime = models.DateTimeField("南北方信息添加时间", null= True, blank = False)


#保留
# class TVShow(models.Model):
#     '''
#     学生看剧类型信息
#     '''
#     TVSHOW_CHOICES=(
#         (1, "游戏类直播"),
#         (2, "唱歌类直播"),
#         (3, "短视频"),
#         (4, "日韩"),
#         (5, "欧美"),
#         (6, "综艺"),
#     )
#     no_id=models.AutoField(primary_key=True)
#     tvShowType = models.IntegerField("看剧类型", null= False, choices = TVSHOW_CHOICES,blank = False)#0-
#     addTime = models.DateTimeField("看剧类型信息添加时间", null= True, blank = False)


# class AirconditionTemperature(models.Model):
#     '''
#     冬夏空调温度信息
#     '''
#     temp_id = models.AutoField(primary_key=True)
#     wintertemperature = models.IntegerField("冬天温度", null=False, blank = False)
#     summertemperature = models.IntegerField("夏天空调温度", null=False, blank=False)
#     addTime = models.DateTimeField("夏天空调温度添加时间", null=True, blank=False)


# class Smoking(models.Model):
#     '''
#     学生吸烟信息
#     '''
#     sm_id = models.AutoField(primary_key=True)
#     smoke = models.BooleanField("吸烟", null= False, blank=False)#0不吸烟 1吸烟
#     addTime = models.DateTimeField("添加时间", null= True, blank=False)

