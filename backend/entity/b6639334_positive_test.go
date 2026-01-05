package entity
import(
	"testing"
	."github.com/onsi/gomega"	
)
func TestCustomer_Positive(t *testing.T){
	g:=NewWithT(t)
	customer := Customer{
		Name:"Som",
		Email:"Som@pim.com",
		CustomerID: "L1234567",
	}

	
	err:=customer.Validate()
	g.Expect(err).To(BeNil())
}