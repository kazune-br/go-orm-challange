package sqlboilerorm

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kazune-br/go-orm-challange/internal/infrastructure/sqlboiler/dao"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"math/rand"
	"time"
)

func RunSqlboiler() {
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?parseTime=true",
		"root",
		"password",
		"localhost",
		"4402",
		"sqlboilerdb",
	)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	fmt.Println("Executing Sqlboiler Example")

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1000)
	u := dao.User{
		UID: int64(i),
	}
	err = u.Insert(ctx, db, boil.Infer())
	if err != nil {
		panic(err)
	}

	one, err := dao.Users().One(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", one)

	all, err := dao.Users().All(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", all)
}

//type User struct {
//	ID  int64 `boil:"users.id" json:"id"`
//	UID int64 `boil:"users.uid" json:"uid"`
//}
