package user

import "golang.org/x/crypto/bcrypt"

// Service ini mewakili bisnis logic / mewakili feature
// atau kata kerja di aplikasi kita
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	// RegisterUser(input RegisterUserInput)
	// RegisterUser => nama method
	// (input, RegisterUserInput) => parameter dengan tipe data RegisterUserInput
	// (User, error) => balikannya bisa berupa object struct User atau error
}

type service struct {
	repository Repository
	// si service punya ketergantungan ke siapa si
	// pada akhirnya akan memanggil repository
	// karena dia bisa mensave data
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// mapping struct input ke struct User
// simpan struct User melalui repository
