package setups

import (
	"fmt"
	"log"

	"github.com/RyukiKuwahara/Bio-Map/repositories"
	_ "github.com/lib/pq"
)

func Initialization() {
	fmt.Println("Initialization code in setup.go")

	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}

	tableName := "users"
	exists, err := ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreateUsers()
	}

	tableName = "session"
	exists, err = ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreateSession()
	}

	tableName = "posts"
	exists, err = ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreatePosts()
	}

	tableName = "genres"
	exists, err = ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreateGenres()
	}

	tableName = "species"
	exists, err = ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreateSpecies()
	}

	tableName = "badges"
	exists, err = ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreateBadges()
	}

	tableName = "user_badge_history"
	exists, err = ur.TableExits(tableName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("テーブル %s は存在します。\n", tableName)
	} else {
		fmt.Printf("テーブル %s は存在しません。\n", tableName)
		ur.CreateUserBadgeHistory()
	}
}
