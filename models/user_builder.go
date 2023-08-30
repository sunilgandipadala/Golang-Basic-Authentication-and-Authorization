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
//here actually, by constructors for each attribute, it will define each step..
//like here  we can declare like instead of the below
/*
type UserBuilder interface {
	SetUserName(username string) UserBuilder
	SetEmail(email string) UserBuilder
	SetAge(age int) UserBuilder //optional attribute
	SetPhone(phone string) UserBuilder //optional attribute
	SetPassword(password string) UserBuilder
	//and in the concrete builder which implements this interface will call all these methods .. and different concrete builders can use different optional attributes

	Build() User
}
*/
type UserBuilder interface {
	SetAttributes(attributes map[string]interface{}) UserBuilder
	Build() User
}

// ConcreteBuilder: DynamicUserBuilder implements the UserBuilder interface to build a dynamic user.
//As here we are doing only registration of user type only definded like this..
type DynamicUserBuilder struct {
	user User
}

//Builder Methods Implementation

//Here we can directly use the user data,but to use this build process step-by-step
func (b *DynamicUserBuilder) SetAttributes(attributes map[string]interface{}) UserBuilder {

	//HERE ACCORDING TO -- Builder Pattern we can divide these all If conditions into Individual methods instead of
	//SetAttributes(),we can make SetUsername(),SetEmail() like that

	//-- We can even Create a different method --- like for making registration for admin, and there we can take another attribute
	//like Admin_ID also..
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
	//Here we are directly calling the SetAttributes..
	//If we have different types of registeration like users, admins and some more... Then we have to check the type first
	//and based on that type the SetUserName()... and other will be called and Build...
	return c.Builder.SetAttributes(attributes).Build()
	//here first SetAttributes will be called then Build method will called..
	// The set Attrinutes is where the step by step processes gets completed
	//this director will call the concrete builders
}
