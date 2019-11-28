package constants

const (
	HTTP_ACTION_LOGIN    = "login"
	HTTP_ACTION_REGISTER = "register"
)

var (
	SPORTS_CHOICES = map[int]string{
		1: "跑步",
		2: "羽毛球",
		3: "篮球",
		4: "足球",
		5: "自行车",
		6: "游泳",
		7: "撸铁",
		8: "兵乓球",
		9: "瑜伽舞蹈",
	}
	NOVEL_CHOICES = map[int]string{
		1:  "技术知识类",
		2:  "文学类小说",
		3:  "漫画",
		4:  "言情小说",
		5:  "玄幻小说",
		6:  "科幻小说",
		7:  "武侠小说",
		8:  "推理小说",
		9:  "恐怖小说",
		10: "网游小说",
		11: "二次元小说",
	}
	GAMES_CHOICES = map[int]string{
		1:  "moba",
		2:  "角色扮演",
		3:  "射击类",
		4:  "棋牌类",
		5:  "即时战略",
		6:  "沙盒类",
		7:  "竞速游戏",
		8:  "王者荣耀",
		9:  "和平精英",
		10: "第五人格",
		11: "捉精灵",
		12: "阴阳师",
		13: "明日方舟",
		14: "炉石",
		15: "崩坏3",
		16: "FGO",
	}
	VIDEOS_CHOICES = map[int]string{
		1:  "泰剧",
		2:  "日韩",
		3:  "欧美",
		4:  "短视频",
		5:  "游戏类",
		6:  "音乐类",
		7:  "其他类",
		8:  "体育",
		9:  "歌唱",
		10: "美食",
		11: "情感",
		12: "说唱",
		13: "真人秀",
	}

	GENDER_CHOICES = map[bool]string{
		true: "男",
		false: "女",
	}

	SOUTHNORTH_CHOICES = map[bool]string{
		true: "南方",
		false: "北方",
	}

	SLEEP_TIME = map[int]string{
		1: "21:00之前",
		2: "21:01-22:00",
		3: "22:01-23:00",
		4: "23:01-23:59",
		5: "24:00之后",
	}

	WAKE_TIME = map[int]string{
		1: "7:00之前",
		2: "7:01-8:00",
		3: "8:01-9:00",
		4: "9:01-10:00",
		5: "10:01之后",
	}

	SMOKE_CHOICES = map[bool]string{
		false: "不吸烟",
		true:  "吸烟",
	}
)
