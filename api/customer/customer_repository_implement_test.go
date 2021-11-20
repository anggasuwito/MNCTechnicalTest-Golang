package customer

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerRepository_GetCustomerByEmail(t *testing.T) {
	t.Run("test customer repository", func(t *testing.T) {
		t.Run("test get customer by email", func(t *testing.T) {
			var result model.Customer
			db := config.MockDatabase()
			repository := NewCustomerRepository(db)
			repository.GetCustomerByEmail("admin@gmail.com", &result)
			assert.Equal(t, result.Email, "admin@gmail.com", "customer repository found customer by email")

			var result2 model.Customer
			repository.GetCustomerByEmail("adminnotexist@gmail.com", &result2)
			assert.Equal(t, result2.Email, "", "customer repository not found customer by email")
		})
	})
}
