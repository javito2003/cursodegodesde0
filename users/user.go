package users

import (
	"cursodegodesde0/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Javico", time.Now(), true)
	fmt.Println(u)
}
