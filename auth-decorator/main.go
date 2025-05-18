package main 

import(
	"fmt"
)

func withAuth(f func(user string), user string) func() {
	return func() {
		if user != "admin" {
			fmt.Println("Unauthorized access")
			return 
		}
		f(user)
	}
}

func showDashboard(user string)  {
	fmt.Println("Welcome: ", user)
}

func main() {
	adminView := withAuth(showDashboard, "admin")
	guest := withAuth(showDashboard, "guest")

	adminView()
	guest()
}