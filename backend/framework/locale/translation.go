package locale

type Entry struct {
	tag, key string
	msg      interface{}
}

// func Load(entries []Entry) {
// 	for _, e := range entries {
// 		tag := language.MustParse(e.tag)
// 		switch msg := e.msg.(type) {
// 		case string:
// 			message.SetString(tag, e.key, msg)
// 		case catalog.Message:
// 			message.Set(tag, e.key, msg)
// 		case []catalog.Message:
// 			message.Set(tag, e.key, msg...)
// 		}
// 	}
// }
