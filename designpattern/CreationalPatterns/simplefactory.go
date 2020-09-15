package CreationalPatterns



type Greet interface {
	Say(content string) string
}

func simpleFactory(i int) string {
	if i == 1 {
		c := &chineseGreet{}
		return c.Say("你好")
	} else if i == 2 {
		e := &englishGreet{}
		return e.Say("hello")
	}
	return ""
}

type chineseGreet struct{}

func (chinese *chineseGreet) Say(content string) string {
	return content
}

type englishGreet struct{}

func (english *englishGreet) Say(content string) string {
	return content
}
