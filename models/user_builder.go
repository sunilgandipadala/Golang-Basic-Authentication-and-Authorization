package models

import "strconv"

// these register users can only see the employees data and can add employees

// now I have to convert this Struct to BUILDER DESIGN PATTREN
//Product
type User struct {
	Name            string `json:"username" form:"username"  gorm:"type:text"`
	Email           string `json:"email" form:"email" gorm:"unique"`
	Age             int    `json:"age" form:"age" `
	Phone           string `json:"phone" form:"phone"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"cpassword" form:"cpassword"`
}

//Builder - an Abstract Interface defines the methods for building the parts of the Email.
//here actually, by constructors for each parameter, it will define each step..
type UserBuilder interface {
	SetAttributes(attributes map[string]interface{}) UserBuilder
	//this method will be
	Build() User
}

//Concrete Builders: we will make a dynamic Concrete Builder Class to accept all the types of same object

// ConcreteBuilder: DynamicUserBuilder implements the UserBuilder interface to build a dynamic user.
//Actually we can define different types of builders for different attributes group, but Making it dynamic is more finer right..,!
type DynamicUserBuilder struct {
	user User
}

//Builder Methods Implementation

//Here we can directly use the user data,but to use this build process step-by-step
func (b *DynamicUserBuilder) SetAttributes(attributes map[string]interface{}) UserBuilder {
	if username, ok := attributes["username"].(string); ok {
		b.user.Name = username
	}
	if email, ok := attributes["email"].(string); ok {
		b.user.Email = email
	}
	if age, ok := attributes["age"].(string); ok {
		var err error
		b.user.Age, err = strconv.Atoi(age)

		if err != nil {
			b.user.Age = 0
		}
	}
	if phone, ok := attributes["phone"].(string); ok {
		b.user.Phone = phone
	}
	if password, ok := attributes["password"].(string); ok {
		b.user.Password = password
	}
	if cppassword, ok := attributes["cpassword"].(string); ok {
		b.user.ConfirmPassword = cppassword
	}

	// Handle one attributes if needed...
	return b
}

func (b *DynamicUserBuilder) Build() User {
	var db = ConnectDB()
	err := db.Create(&b.user).Error
	if err != nil {
		return b.user
	}
	return b.user
}

// Director: UserRegister manages the user registration process using a builder.
type RegisterUser struct {
	Builder UserBuilder
}

func (c *RegisterUser) UserRegistration(attributes map[string]interface{}) User {
	return c.Builder.SetAttributes(attributes).Build()
	//here first SetAttributes will be called then Build method will called..
	// The set Attrinutes is where the step by step processes gets completed
	//this director will call the concrete builders
}
