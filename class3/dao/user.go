package dao

var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

var messageBin = make([]string, 100)

var passwordQuestion = make(map[string]string, 100)

var pwdQuetionAnswer = make(map[string]string, 100)

var likes = make(map[string][]string, 100)

func AddUser(username, password string) {
	database[username] = password
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

// 实现添加密保问题与答案
func AddPwdQuestion(username string, question string) {
	passwordQuestion[username] = question
}

func AddPwdAnswer(username string, answer string) {
	question := passwordQuestion[username] //得到用户名对应的密保问题
	pwdQuetionAnswer[question] = answer
}

// 找回密码
func FindPassword(username string, newpassword string) {
	database[username] = newpassword
}

// 检索密保问题与答案
func SelectQuestion(username string) bool {
	if passwordQuestion[username] == "" {
		return false
	}
	return true
}

func SelectAnswerFromQuestion(username string) string {
	question := passwordQuestion[username]
	return pwdQuetionAnswer[question]
}

// 留言
func LeaveMessage(message string) {
	messageBin = append(messageBin, message)
}

// 点赞
func SelectLike(friendname string) []string {
	var ret []string
	//为这个用户点赞过的用户名存在一个数组里
	ret = likes[friendname]
	return ret
}

func Like(username string, friendname string) {
	s1 := make([]string, 100)
	s1 = append(s1, username)
	likes[friendname] = s1
}
