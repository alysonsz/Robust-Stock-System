package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Product string
	Acquired uint
	Available uint
	Safety uint
	Price float32
	Overbooking bool
}

func main() {
	dsn := "host=localhost user=dojo password=dojopw dbname=dojodb port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Stock{})

	if err != nil {
		log.Panicln(err)
	}

	service := NewService(db)
	id, _ := service.Create("Notebook", 30, 10, 30.21)
  service.UpdateAvailable(id, 25)
}
 
type Service struct{
	db *gorm.DB
}

func NewService (db *gorm.DB) *Service{
	return &Service{db}
}

func (s *Service) Create(product string, acquired uint, safety uint, price float32)(uint, error)  {
	stock := Stock{
		Product: product,
		Acquired: acquired,
		Available: acquired,
		Safety: safety,
		Price: price,
		Overbooking: false,
	}

	result := s.db.Create(&stock)
	if result.Error != nil{
		log.Println("Erro ao adicionar produto no estoque")
		log.Panicln(result.Error)
	}

	return stock.ID , nil
}

func (s *Service) UpdateAvailable(id uint, available uint) error {
	var stock *Stock
	result := s.db.First(&stock, id)
	if result.Error != nil {
		return result.Error
	}

	stock.Available = available

	result = s.db.Save(stock)
	if result.Error != nil {
		log.Println("Erro ao atualizar o produto no estoque")
		return result.Error
	}

	return nil
}

func (s *Service) Find(id uint) (*Stock, error) {
	var stock *Stock
	result := s.db.First(&stock, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return stock, nil
}