package sqlboilerorm

import (
	"fmt"
	"github.com/kazune-br/go-orm-challange/internal/infrastructure/sqlboiler/dao"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"math/rand"
	"time"
)

type BoilerRepository struct {
	connector RDBConnector
}

type IBoilerRepository interface {
	Create()
	SelectOne()
	SelectAll()
}

func NewBoilerRepository(connector RDBConnector) IBoilerRepository {
	return &BoilerRepository{
		connector: connector,
	}
}

func (br *BoilerRepository) Create() {
	fmt.Println("Executing BoilerRepository.Create")
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1000)
	u := dao.User{
		UID: int64(i),
	}
	err := u.Insert(br.connector.GetContext(), br.connector.GetDB(), boil.Infer())
	if err != nil {
		panic(err)
	}
}

func (br *BoilerRepository) SelectOne() {
	fmt.Println("Executing BoilerRepository.SelectOne")
	one, err := dao.Users().One(br.connector.GetContext(), br.connector.GetDB())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", one)
}

func (br *BoilerRepository) SelectAll() {
	fmt.Println("Executing BoilerRepository.SelectAll")
	all, err := dao.Users().All(br.connector.GetContext(), br.connector.GetDB())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", all)
}
