package main

import "fmt"

// userStatus represents user profile info
type userStatus struct {
	isMember    bool
	isGold      bool
	isFirstTime bool
}

// getDiscount returns a discount based on user type
func getDiscount(user userStatus) int {
	if !user.isMember {
		return 0
	}

	if user.isGold {
		return 50
	}

	if user.isFirstTime {
		return 30
	}

	return 10
}

// getDiscount uses nested if-else instead of guard clauses
// func getDiscount(user userStatus) int {
// 	discount := 0

// 	if user.isMember {
// 		if user.isGold {
// 			discount = 50
// 		} else {
// 			if user.isFirstTime {
// 				discount = 30
// 			} else {
// 				discount = 10
// 			}
// 		}
// 	} else {
// 		discount = 0
// 	}

// 	return discount
// }

func main() {
	u1 := userStatus{false, false, false} // Not a member
	u2 := userStatus{true, true, false}   // Gold member
	u3 := userStatus{true, false, true}   // First time member
	u4 := userStatus{true, false, false}  // Regular member

	fmt.Println("u1 Discount:", getDiscount(u1)) // 0
	fmt.Println("u2 Discount:", getDiscount(u2)) // 50
	fmt.Println("u3 Discount:", getDiscount(u3)) // 30
	fmt.Println("u4 Discount:", getDiscount(u4)) // 10
}
