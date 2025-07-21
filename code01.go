package main

import "fmt"

type Product struct {
	id       string
	name     string
	price    float64
	quantity int
}

//打印菜单
func showMenu() {
	fmt.Println("--- 库存管理系统 ---")
	fmt.Println("1. 添加商品")
	fmt.Println("2. 查看所有商品")
	fmt.Println("3. 更新商品库存")
	fmt.Println("4. 删除商品")
	fmt.Println("5. 查找指定商品")
	fmt.Println("q. 退出")
	fmt.Printf("请输入你的选择:")
}

//添加商品
func addProduct(inv map[string]Product) {
	var ad Product
	for {
		fmt.Print("请输入你要插入的商品编号：")
		fmt.Scanln(&ad.id)
		if ad.id == "" {
			fmt.Println("错误!商品编号不能为空，请重新输入。") // 优化：提示重新输入
			continue
		}
		_, exist := inv[ad.id]
		if exist {
			fmt.Println("商品ID已存在，请重新输入。") // 优化：提示重新输入
			ad.id = ""
			continue
		}
		break
	}
	for {
		fmt.Print("请输入商品名称：")
		fmt.Scanln(&ad.name)
		if ad.name == "" {
			fmt.Println("错误!商品名称不能为空，请重新输入。") // 优化：提示重新输入
			continue
		}
		break
	}
	for {
		fmt.Print("请输入商品价格：")
		n, err := fmt.Scanln(&ad.price)
		if n != 1 || err != nil || ad.price < 0 {
			fmt.Println("错误!请输入有效的非负价格。") // 优化：价格校验
			// 清空输入缓冲区
			var dump string
			fmt.Scanln(&dump)
			continue
		}
		break
	}
	for {
		fmt.Print("请输入商品库存：")
		n, err := fmt.Scanln(&ad.quantity)
		if n != 1 || err != nil || ad.quantity < 0 {
			fmt.Println("错误!请输入有效的非负库存数量。") // 优化：库存校验
			var dump string
			fmt.Scanln(&dump)
			continue
		}
		break
	}
	inv[ad.id] = ad
	fmt.Println("添加成功！") // 优化：更友好提示
}

//打印商品信息
func listProducts(inv map[string]Product) {

	if len(inv) == 0 {
		fmt.Println("库存为空")
	} else {
		fmt.Println("ID\t名称\t价格\t数量")
		fmt.Println("---------------------------------")
		for _, value := range inv {
			fmt.Printf("%s\t%s\t%.2f\t%d\n", value.id, value.name, value.price, value.quantity)
		}
	}
}

//更新库存
func updateStock(inv map[string]Product) {
	if inv == nil {
		fmt.Println("库存未初始化")
		return
	}
	var up Product
	for {
		fmt.Print("请输入要更新的产品ID：")
		fmt.Scanln(&up.id)
		if up.id == "" {
			fmt.Println("商品ID不能为空，请重新输入。") // 优化：ID校验
			continue
		}
		_, exist := inv[up.id]
		if !exist {
			fmt.Println("商品不存在，请重新输入。") // 优化：ID校验
			continue
		}
		break
	}
	var i int
	for {
		fmt.Print("请输入商品库存变动数量(正数为增加，负数为减少)：")
		n, err := fmt.Scanln(&i)
		if n != 1 || err != nil {
			fmt.Println("请输入有效的整数。") // 优化：数量校验
			var dump string
			fmt.Scanln(&dump)
			continue
		}
		if inv[up.id].quantity+i < 0 {
			fmt.Println("操作失败，库存不足，请重新输入。") // 优化：库存校验
			continue
		}
		break
	}
	p := inv[up.id]
	p.quantity = p.quantity + i
	inv[up.id] = p
	fmt.Println("库存更新成功！") // 优化：更友好提示
}

//删除商品
func deleteProduct(inv map[string]Product) {
	if inv == nil {
		fmt.Println("库存未初始化")
		return
	}
	var i string
	for {
		fmt.Print("请输入你要删除的商品ID:")
		fmt.Scanln(&i)
		if i == "" {
			fmt.Println("商品ID不能为空，请重新输入。") // 优化：ID校验
			continue
		}
		_, exist := inv[i]
		if !exist {
			fmt.Println("商品不存在，请重新输入。") // 优化：ID校验
			continue
		}
		break
	}
	delete(inv, i)
	fmt.Println("商品已删除！") // 优化：更友好提示
}

//查找商品
func findProduct(inv map[string]Product) {
	if inv == nil {
		fmt.Println("库存未初始化")
		return
	}
	var i string
	for {
		fmt.Print("请输入你要查找的商品ID:")
		fmt.Scanln(&i)
		if i == "" {
			fmt.Println("商品ID不能为空，请重新输入。") // 优化：ID校验
			continue
		}
		_, exist := inv[i]
		if !exist {
			fmt.Println("商品不存在，请重新输入。") // 优化：ID校验
			continue
		}
		break
	}
	fmt.Println("ID\t名称\t价格\t数量")
	fmt.Println("---------------------------------")
	fmt.Printf("%s\t%s\t%.2f\t%d\n", inv[i].id, inv[i].name, inv[i].price, inv[i].quantity)
	fmt.Println("查找完成！") // 优化：更友好提示

}

func main() {
	//初始化仓库数据结构
	inventory := make(map[string]Product)
	//接受用户选择

	for {
		showMenu()
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			addProduct(inventory)
		case "2":
			listProducts(inventory)
		case "3":
			updateStock(inventory)
		case "4":
			deleteProduct(inventory)
		case "5":
			findProduct(inventory)
		case "q", "Q":
			fmt.Println("感谢您的使用，再见！")
			return
		default:
			fmt.Println("非法输入，请重新选择菜单项。") // 优化：更友好提示
		}
		fmt.Println("------------------------------") // 优化：分隔符
	}
}
