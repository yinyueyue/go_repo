package main

import "fmt"

/**
  使用面向对象的方式，编写的收支软件
*/
func main() {
	account := &FamilyAccount{
		balance: 10000,
		details: "",
	}
	account.showMenu()
}

type FamilyAccount struct {
	balance float64

	details string // 历史所有收支明细，所有的使用字符串；拼接
}

func (familyAccount *FamilyAccount) showMenu() {
	for {
		fmt.Println("--------------欢迎进入家庭收支记账软件--------------")
		fmt.Println("              1.收支明细")
		fmt.Println("              2.登记收入")
		fmt.Println("              3.登记支出")
		fmt.Println("              4.查询余额")
		fmt.Println("              5.退出软件")
		key := ""
		fmt.Scanln(&key)
		switch key {
		case "1":
			familyAccount.showDetails()
		case "2":
			familyAccount.income()
		case "3":
			familyAccount.expend()
		case "4":
			familyAccount.showBalance()
		case "5":
			{
				exit := familyAccount.exit()
				if exit {
					return
				}
			}

		default:
			fmt.Println("未知的操作")
		}

	}
}

/**
  查询余额
*/
func (familyAccount *FamilyAccount) showBalance() {
	fmt.Printf("当前账户余额明细:%v\n", familyAccount.balance)
}

func (familyAccount *FamilyAccount) showDetails() {
	fmt.Println("-----------当前收支明细------------")
	if familyAccount.details == "" {
		fmt.Println("-----------没有任何收支明细------------")
	} else {
		fmt.Println(familyAccount.details)
	}
}

/**
收入
*/
func (familyAccount *FamilyAccount) income() {

	fmt.Println("请输入收入金额")
	var amount float64
	_, _ = fmt.Scanln(&amount)
	familyAccount.balance += amount
	notes := fmt.Sprintf("收到一笔钱款,金额为:%v\n", amount)
	fmt.Println(notes)
	familyAccount.details += notes
}

/**
支出
*/
func (familyAccount *FamilyAccount) expend() {
	fmt.Println("请输入支出金额")
	var amount float64
	_, _ = fmt.Scan(&amount)
	if familyAccount.balance < amount {
		fmt.Println("当前账户余额不足")
		return
	} else {
		familyAccount.balance -= amount
	}
	notes := fmt.Sprintf("支出一笔钱款,金额为:%v\n", amount)
	fmt.Println(notes)
	familyAccount.details += notes
}

/**
退出
*/
func (familyAccount *FamilyAccount) exit() bool {
	fmt.Println("确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			fmt.Println("您已经成功退出")
			return true
		}
		if choice == "n" || choice == "N" {
			fmt.Println("不退出")
			return false
		}
		fmt.Println("输入错误，请重新输入")
	}
}
