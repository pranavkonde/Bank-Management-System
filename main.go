package main

import (
	"fmt"
	"time"
)

type customer struct {
	email_id     string
	phone_no     string
	account_ID   int
	password     int
	balance      int
	trans_amount int
	trans_type   string
}
type hist struct {
	sr_no              int
	date               string
	transaction_type   string
	transaction_amount int
	total_balance      int
}

var histo = map[int][]hist{}
var map1 = make(map[int]customer)
var bankacc_id = 10000
var password1 = 123456

func Accountant(acc customer) {
	fmt.Println("Hello Customer kindly tell your email-id and phone-no")
	var email_id1, phone_no string
	fmt.Println("Enter the email-id:")
	fmt.Scanln(&email_id1)
	fmt.Println("Enter the phone:")
	fmt.Scanln(&phone_no)
	bankacc_id = bankacc_id + 1
	password1 = password1 + 1
	fmt.Println("Customer your Account_id and Password is: ", bankacc_id, password1)
	acc.account_ID = bankacc_id
	acc.password = password1
	acc.balance = 0
	acc.phone_no = phone_no
	acc.email_id = email_id1
	map1[acc.account_ID] = acc
}
func transhistory(id int) {
	var make ([]hist) = histo[id]
	for i := 0; i < len(make); i++ {
		var typ hist = make[i]
		fmt.Println("Transaction History")
		fmt.Println("Sr no. ", typ.sr_no)
		fmt.Println("Date: ", typ.date)
		fmt.Println("Type of Transaction", typ.transaction_type)
		fmt.Println("Transaction amount: ", typ.transaction_amount)
		fmt.Println("Total Balance: ", typ.total_balance)
	}
}
func customer1() {
	fmt.Println("Hello Customer")
	var a, id, pass int
	fmt.Println("Enter the account_no and password")
	fmt.Scanln(&id)
	fmt.Scanln(&pass)
	temp := map1[id]
	if id == temp.account_ID && pass == temp.password {
		for {
			fmt.Println("What you want to do 1. Deposit 2. Withdraw 3. View Balance 4. Delete Account 5. Transaction History 6. Exit")
			fmt.Scanln(&a)
			flag := false
			switch a {
			case 1:
				var dep int
				fmt.Println("Enter the amount to deposit")
				fmt.Scanln(&dep)
				temp1 := map1[id]
				temp1.balance += dep
				map1[id] = temp1
				_, ok := histo[temp1.account_ID]
				if ok {
					var flag int = 0
					for i := 0; i <= len(histo); i++ {
						flag++
					}
					var a hist
					t := time.Now()
					a.date = t.Format("01-02-2006")
					a.sr_no = flag
					a.total_balance = temp1.balance
					a.transaction_amount = dep
					a.transaction_type = "Deposit"
					temp3 := histo[id]
					temp3 = append(temp3, a)
					histo[id] = temp3
				} else {
					var a hist
					t := time.Now()
					a.date = t.Format("01-02-2006")
					a.sr_no = 1
					a.total_balance = temp1.balance
					a.transaction_amount = dep
					a.transaction_type = "Deposit"
					temp3 := histo[id]
					temp3 = append(temp3, a)
					histo[id] = temp3
				}
			case 2:
				var withd int
				if map1[id].balance > 0 {
					fmt.Println("Enter the amount to withdraw")
					fmt.Scanln(&withd)
					temp2 := map1[id]
					temp2.balance -= withd
					map1[id] = temp2
					_, ok := histo[temp2.account_ID]
					if ok {
						var flag int = 0
						for i := 0; i <= len(histo); i++ {
							flag++
						}
						var a hist
						t := time.Now()
						a.date = t.Format("01-02-2006")
						a.sr_no = flag
						a.total_balance = temp2.balance
						a.transaction_amount = withd
						a.transaction_type = "Withdraw"
						temp3 := histo[id]
						temp3 = append(temp3, a)
						histo[id] = temp3
					} else {
						var a hist
						t := time.Now()
						a.date = t.Format("01-02-2006")
						a.sr_no = 1
						a.total_balance = temp2.balance
						a.transaction_amount = withd
						a.transaction_type = "Withdraw"
						temp3 := histo[id]
						temp3 = append(temp3, a)
						histo[id] = temp3
					}
				} else {
					fmt.Println("Insufficient Balance")
				}
			case 3:
				fmt.Println("Your balance is: ", map1[id].balance)
			case 4:
				delete(map1, temp.account_ID)
				fmt.Println("Account Successfully Deleted")
			case 5:
				transhistory(id)
			case 6:
				flag = true
			}
			if flag {
				break
			}
		}
	}
}
func main() {
	fmt.Println("Welcome")
	for {
		fmt.Println("To whom you have to connect with 1. Accountant 2. Customer 3. Exit")
		var a int
		fmt.Scanln(&a)
		if a == 1 {
			A_email_id := "account@bank.com"
			A_password := "josh@123"
			var EA_email_id, EA_password string
			fmt.Println("Enter the email-id:")
			fmt.Scanln(&EA_email_id)
			fmt.Println("Enter the password:")
			fmt.Scanln(&EA_password)
			if A_email_id == EA_email_id && A_password == EA_password {
				var acc customer
				Accountant(acc)
			} else {
				fmt.Println("Invalid Credentials")
			}
		} else if a == 2 {
			customer1()
		} else {
			break
		}
	}
}
