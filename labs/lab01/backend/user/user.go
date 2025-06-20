package user

import (
	"errors"
	"strings"
	"strconv"
)

// Predefined errors
var (
	ErrInvalidName  = errors.New("invalid name: must be between 1 and 30 characters")
	ErrInvalidAge   = errors.New("invalid age: must be between 0 and 150")
	ErrInvalidEmail = errors.New("invalid email format")
)

// User represents a user in the system
type User struct {
	Name  string
	Age   int
	Email string
}

<<<<<<< HEAD
// Validate checks if the user data is valid, returns an error for each invalid field
func (u *User) Validate() error {
	if !IsValidName(u.Name) {
		return ErrInvalidName
	}

	if !IsValidAge(u.Age) {
		return ErrInvalidAge
	}

=======
// NewUser creates a new user with validation
func NewUser(name string, age int, email string) (*User, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrEmptyName
	}

	// Проверяем возраст
	if age < 0 || age > 150 {
		return nil, ErrInvalidAge
	}

	// Проверяем email
	if !IsValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	return &User{
		Name:  name,
		Age:   age,
		Email: strings.TrimSpace(email),
	}, nil
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	// Проверяем имя
	if strings.TrimSpace(u.Name) == "" {
		return ErrEmptyName
	}

	// Проверяем возраст
	if u.Age < 0 || u.Age > 150 {
		return ErrInvalidAge
	}

	// Проверяем email
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	return nil
}

// String returns a string representation of the user, formatted as "Name: <name>, Age: <age>, Email: <email>"
func (u *User) String() string {
<<<<<<< HEAD
	// TODO: Implement this function
	return ""
=======
    return "User{Name: " + u.Name + ", Age: " + strconv.Itoa(u.Age) + ", Email: " + u.Email + "}"
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// NewUser creates a new user with validation, returns an error if the user is not valid
func NewUser(name string, age int, email string) (*User, error) {
	// TODO: Implement this function
	return nil, nil
}

// IsValidEmail checks if the email format is valid
// You can use regexp.MustCompile to compile the email regex
func IsValidEmail(email string) bool {
<<<<<<< HEAD
	// TODO: Implement this function
	return false
}

// IsValidName checks if the name is valid, returns false if the name is empty or longer than 30 characters
func IsValidName(name string) bool {
	// TODO: Implement this function
	return false
}

// IsValidAge checks if the age is valid, returns false if the age is not between 0 and 150
func IsValidAge(age int) bool {
	// TODO: Implement this function
	return false
=======
    email = strings.TrimSpace(email)
    if email == "" {
        return false
    }

    // Проверяем наличие пробелов в email
    if strings.ContainsAny(email, " \t\n\r") {
        return false
    }

    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return false
    }

    if parts[0] == "" || parts[1] == "" {
        return false
    }

    // Проверяем наличие точки в доменной части
    if !strings.Contains(parts[1], ".") {
        return false
    }

    return true
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}