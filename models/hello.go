package models

type Hello struct{
	Msg 	string `json:"hello"`
	Lang 	string `json:"language"`
}

var hellos []Hello

func init() {
	hellos = []Hello{
		Hello {
			Msg: "Hallo",
			Lang: "de",
		},
		Hello {
			Msg: "Hola",
			Lang: "es",
		},
	}
}

// get all Hello
func GetHellos() []Hello {
	return hellos
}

// get the message from the Hello corresponding to the specified language
func GetHello(lang string) string {
	for _, item := range hellos {
		if item.Lang == lang {
			return item.Msg
		}
	}
	return ""
}

// create a new Hello
func CreateNewHello(newHello Hello) Hello {
	// do not add if the Hello is already in the slice
	if GetHello(newHello.Lang) == "" {
		hellos = append(hellos, newHello)
	}
	return newHello
}

// delete from the slice a hello corresponding to the specified language
func DeleteHello(lang string) Hello {
	ret := Hello{Msg:"", Lang:""}
	for i, item := range hellos {
		if item.Lang == lang {
			ret = item
			hellos = append(hellos[:i], hellos[i+1:]...)
			break
		}
	}
	return ret
}
