package models

import (
	"fmt"
	mdw "meetup/middleware"
	"meetup/utils/logging"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Error("error in bcrypt")
	}
	user.Password = string(hashedPassword)
	if err = db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginToken(mobile, password string) (string, error) {
	usr_query := &User{}
	db.Where("mobile_number = ?", mobile).First(&usr_query)
	err_pass_check := bcrypt.CompareHashAndPassword([]byte(usr_query.Password), []byte(password))
	if err_pass_check != nil {
		err_string := fmt.Sprintf("error: %v", err_pass_check)
		logging.Error(err_string)
		return "", err_pass_check
	}
	jwt_token := mdw.GenerateJWTToken(mobile, usr_query.FirstName)
	return jwt_token, nil
}

func GetUserProfile(mobile string) (*UserGenInfo, error) {
	usr := &UserGenInfo{}
	if err := db.Model(&User{}).Where("mobile_number = ?", mobile).First(usr).Error; err != nil {
		return nil, err
	}
	return usr, nil
}

func UpdateUserProfile(mobile string, user *User) error {
	usr, validErr := validateUpdateParams(user, mobile)
	if validErr != nil {
		return validErr
	}
	if err := db.Model(&User{}).Where("mobile_number = ?", mobile).Updates(usr).Error; err != nil {
		return err
	}
	return nil
}

func age(birthdate, today time.Time) uint8 {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return uint8(age)
}

func validateUpdateParams(user *User, mobile string) (interface{}, error) {
	if len(user.UserAuth.MobileNumber) > 0 && mobile != user.UserAuth.MobileNumber {
		return user, fmt.Errorf("you can not update mobile number")
	}
	usr := map[string]interface{}{}
	if len(user.UserProfile.BirthDate) > 0 {
		dob_string, err := time.Parse("02/01/2006", user.UserProfile.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("please provide the data of birth in DD/MM/YYYY format")
		}
		usr["birth_date"] = user.UserProfile.BirthDate
		usr["age"] = age(dob_string, time.Now())
	}
	if len(user.UserAvatar.FirstName) > 0 {
		usr["first_name"] = user.UserAvatar.FirstName
	}
	if len(user.UserAvatar.LastName) > 0 {
		usr["last_name"] = user.UserAvatar.LastName
	}
	// if len(user.UserProfile.UID) > 0 {
	// 	usr["u_id"] = user.UserProfile.UID
	// }
	if len(user.UserProfile.Language) > 0 {
		usr["language"] = user.UserProfile.Language
	}
	if len(user.UserProfile.Bio) > 0 {
		usr["bio"] = user.UserProfile.Bio
	}
	if len(user.UserProfile.Gender) > 0 {
		usr["gender"] = user.UserProfile.Gender
	}
	return usr, nil
}

func ValidateFullName(username string) bool {
	//Remove the extra space
	space := regexp.MustCompile(`\s+`)
	name := space.ReplaceAllString(username, " ")

	//Remove trailing spaces
	name = strings.TrimSpace(name)

	//To support all possible languages
	matched, _ := regexp.Match(`^[^±!@£$%^&*_+§¡€#¢§¶•ªº«\\/<>?:;'"|=.,0123456789]{3,20}$`, []byte(name))
	return matched
}
