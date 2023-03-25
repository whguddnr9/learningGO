package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VipLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s Age: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP User: %s, Id: %s, Age: %d VIP Level :%d VIP Price: %dman won\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VipLevel,
		vip.Price,
	)

}
