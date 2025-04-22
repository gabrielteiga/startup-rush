package database

import (
	"log"

	"gorm.io/gorm"
)

var events []Events = []Events{
	{Name: "Pitch convincente", Score: 6},
	{Name: "Produto com bugs", Score: -4},
	{Name: "Boa tração de usuários", Score: 3},
	{Name: "Investidor irritado", Score: -6},
	{Name: "Fake news no pitch", Score: -8},
}

func SeedEvents(db *gorm.DB) error {
	for _, event := range events {
		if err := db.Where("Name = ?", event.Name).First(&event).Error; err != nil {
			if err := db.Create(&event).Error; err != nil {
				log.Fatal("Error seeding events: ", err)
				return err
			}
			log.Println("Event seeded: ", event.Name)
		}
	}
	log.Println("Events seeded successfully")
	return nil
}
