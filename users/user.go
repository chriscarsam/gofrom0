package users

import (
	"fmt"
	"time"

	"github.com/chriscarsam/gofrom0/models"
)

func UserRegistration() {
	u := new(models.User)
	u.AddUser(10, "Charly", time.Now(), true)
	fmt.Println(u)
}
