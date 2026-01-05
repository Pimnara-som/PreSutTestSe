package entity

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCustomer_Negative_InvalidCustomerID(t *testing.T) {
	g := NewWithT(t)

	customer := Customer{
		Name:       "Som",
		Email:      "som@example.com",
		CustomerID: "A1234567", // ❌ ไม่ขึ้นต้นด้วย L/M/H
	}

	err := customer.Validate()
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("customer_id format invalid"))
}
