package service

import (
	"alta-dashboard-be/features/users"
	"alta-dashboard-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		userId           uint
		Email            string
		Password         string
		loggedInUserId   uint
		loggedInUserRole string
		userInput        users.UserEntity
		limit            int
		offset           int
		queryParams      map[string][]string
	}
	Output struct {
		IsError   bool
		Token     string
		Result    interface{}
		errResult error
	}
}

func LoginTestTable() []TestTable {
	tname := "test login"
	return []TestTable{
		{
			Name: tname + " blank email",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				Password: "qwerty",
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Token:   "",
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " blank password",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				Email: "Joko@gmail.com",
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Token:   "",
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect failed",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				Email:    "joko@gmail.com",
				Password: "qwerty",
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Token:   "Token",
				Result: users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				Email:    "joko@gmail.com",
				Password: "qwerty",
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Token:   "Token",
				Result: users.UserEntity{
					Id:       1,
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Team:     consts.E_USER_Mentor,
					Role:     consts.E_USER_Admin,
					Status:   consts.E_USER_Active,
				},
				errResult: nil,
			},
		},
	}
}

func CreateTestTable() []TestTable {
	tname := "test register"
	return []TestTable{
		{
			Name: tname + " forbidden request",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " blank password",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					FullName: "Joko",
					Email:    "Joko@gmail.com",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " blank email",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					FullName: "Joko",
					Password: "qwerty",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " blank full name",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					Email:    "joko@gmail.com",
					Password: "qwerty",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect failed",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result: users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result: users.UserEntity{
					Id:       1,
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
					Team:     consts.E_USER_Mentor,
					Role:     consts.E_USER_User,
					Status:   consts.E_USER_Active,
				},
				errResult: nil,
			},
		},
	}
}

func ModifyDataTestTable() []TestTable {
	tname := "test update user data"
	return []TestTable{
		{
			Name: tname + " forbidden request",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				userId:           1,
				loggedInUserId:   2,
				loggedInUserRole: consts.E_USER_User,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " blank email",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				userId:           2,
				loggedInUserId:   1,
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					FullName: "Joko",
					Password: "qwerty",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " blank full name",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				userId:           2,
				loggedInUserId:   1,
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					Email:    "joko@gmail.com",
					Password: "qwerty",
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect failed",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				userId:           1,
				loggedInUserId:   1,
				loggedInUserRole: consts.E_USER_User,
				userInput: users.UserEntity{
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
					Team:     consts.E_USER_Mentor,
					Role:     consts.E_USER_User,
					Status:   consts.E_USER_Active,
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result: users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success (admin update another user data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				userId:           2,
				loggedInUserId:   1,
				loggedInUserRole: consts.E_USER_Admin,
				userInput: users.UserEntity{
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
					Team:     consts.E_USER_Mentor,
					Role:     consts.E_USER_User,
					Status:   consts.E_USER_Active,
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result: users.UserEntity{},
				errResult: nil,
			},
		},
		{
			Name: tname + " expect success (user update self data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				userId:           1,
				loggedInUserId:   1,
				loggedInUserRole: consts.E_USER_User,
				userInput: users.UserEntity{
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
					Team:     consts.E_USER_Mentor,
					Role:     consts.E_USER_User,
					Status:   consts.E_USER_Active,
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result: users.UserEntity{
					Id:       1,
					FullName: "Joko",
					Email:    "joko@gmail.com",
					Password: "qwerty",
					Team:     consts.E_USER_Mentor,
					Role:     consts.E_USER_User,
					Status:   consts.E_USER_Active,
				},
				errResult: nil,
			},
		},
	}
}

func GetAllTestTable() []TestTable {
	tname := "test get all user data"
	return []TestTable{
		{
			Name: tname + " expect failed",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				limit:  2,
				offset: 2,
				queryParams: map[string][]string{
					"role": []string{"Admin"},
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result: map[string]any{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				limit:  2,
				offset: 2,
				queryParams: map[string][]string{
					"role": []string{"Admin"},
				},
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result: map[string]any{
					"data": map[string]any{
						"total_page": 2,
						"page":       2,
						"data": users.UserEntity{
							Id:       1,
							FullName: "Joko",
							Email:    "joko@gmail.com",
							Password: "qwerty",
							Team:     consts.E_USER_Mentor,
							Role:     consts.E_USER_User,
							Status:   consts.E_USER_Active,
						},
					},
				},
				errResult: nil,
			},
		},
	}
}

func GetDataTestTable() []TestTable {
	tname := "test get user data"
	return []TestTable{
		{
			Name: tname + " forbidden request (user get another user data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
				loggedInUserId:   1,
				userId:           2,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
			},
		},
		{
			Name: tname + " expect failed",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
				loggedInUserId:   1,
				userId:           1,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				Result:  users.UserEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success (user get self data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
				loggedInUserId:   1,
				userId:           1,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result:  users.UserEntity{
					Id: 1,
					FullName: "Joko",
					Email: "Joko@gmail.com",
					Team: consts.E_USER_Mentor,
					Role: consts.E_USER_User,
					Status: consts.E_USER_Active,
				},
				errResult: nil,
			},
		},
		{
			Name: tname + "  expect success (admin get another user data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				loggedInUserId:   1,
				userId:           2,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result:  users.UserEntity{
					Id: 2,
					FullName: "Joko",
					Email: "Joko@gmail.com",
					Team: consts.E_USER_Mentor,
					Role: consts.E_USER_User,
					Status: consts.E_USER_Active,
				},
				errResult: nil,
			},
		},
		{
			Name: tname + " expect success (admin get self data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				loggedInUserId:   1,
				userId:           1,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				Result:  users.UserEntity{
					Id: 1,
					FullName: "Joko",
					Email: "Joko@gmail.com",
					Team: consts.E_USER_Mentor,
					Role: consts.E_USER_Admin,
					Status: consts.E_USER_Active,
				},
				errResult: nil,
			},
		},
	}
}

func RemoveTestTable() []TestTable {
	tname := "test remove user"
	return []TestTable{
		{
			Name: tname + " forbidden request (user get another user data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
				loggedInUserId:   1,
				userId:           2,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect failed",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
				loggedInUserId:   1,
				userId:           1,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: true,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success (user get self data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_User,
				loggedInUserId:   1,
				userId:           1,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				errResult: nil,
			},
		},
		{
			Name: tname + "  expect succes (admin get another user data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				loggedInUserId:   1,
				userId:           2,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
			},
		},
		{
			Name: tname + " expect succes (admin get self data)",
			Input: struct {
				userId           uint
				Email            string
				Password         string
				loggedInUserId   uint
				loggedInUserRole string
				userInput        users.UserEntity
				limit            int
				offset           int
				queryParams      map[string][]string
			}{
				loggedInUserRole: consts.E_USER_Admin,
				loggedInUserId:   1,
				userId:           1,
			},
			Output: struct {
				IsError   bool
				Token     string
				Result    interface{}
				errResult error
			}{
				IsError: false,
				errResult: nil,
			},
		},
	}
}
