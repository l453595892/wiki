package main

import "fmt"

type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct {
}

type KoreanGirlFriendFactory struct {
}

func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

func getGirlFriend(typeGf string) GirlFriend {
	var gffact AbstractFactory
	switch typeGf {
	case "Indian":
		gffact = IndianGirlFriendFactory{}
		return gffact.CreateMyLove()
	case "Korean":
		gffact = KoreanGirlFriendFactory{}
		return gffact.CreateMyLove()
	}
	return GirlFriend{}
}

func main() {

	a := getGirlFriend("Indian")

	fmt.Println(a.eyesColor)
}
