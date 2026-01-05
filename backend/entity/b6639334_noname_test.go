package entity

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCustomer_Negative_EmptyName(t *testing.T) {
	g := NewWithT(t)

	customer := Customer{
		Name:       "", //  ไม่กรอกชื่อ
		Email:      "som@example.com",
		CustomerID: "M1234567",
	}

	err := customer.Validate()
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("name must not be empty"))
}

