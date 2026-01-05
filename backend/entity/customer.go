package entity
import(
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

)

func init(){
	govalidator.SetFieldsRequiredByDefault(false)//required แค่ fields ที่มี tag required เท่านั้น

}
type Customer struct {
	gorm.Model
	Name string `valid:"required~name must not be empty"`
	Email string `valid:"email~email is invalid"`
	CustomerID string `valid:"required~customer_id must not be empty,matches(^[LMH][0-9]{7}$)~customer_id format invalid"`
}
	
func (c *Customer) Validate() error {
	_,err := govalidator.ValidateStruct(c)
	return err
}