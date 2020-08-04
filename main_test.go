package main

import (
	"testing"
	"reflect"
	"encoding/json"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	DB.First(&result, user.ID)
	
	userJson, _ := json.Marshal(user)
	resultJson, _ := json.Marshal(result)
	
	if string(resultJson) != string(userJson) {
		t.Errorf("jsonEqual Failed, resultJson: %v, userJson: %v", string(resultJson), string(userJson))
	}
	if !result.CreatedAt.Equal(user.CreatedAt) {
		t.Errorf("time.Equal Failed, result.CreatedAt: %v, user.CreatedAt: %v", result.CreatedAt, user.CreatedAt)
	}
	if !reflect.DeepEqual(result.CreatedAt, user.CreatedAt) {
		t.Errorf("reflect.DeepEqual Failed, result.CreatedAt: %v, user.CreatedAt: %v", result.CreatedAt, user.CreatedAt)
	}
}
