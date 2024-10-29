package user

type Pagination struct {
	Page int32 `json:"page" form:"page" query:"page"`
	Size int32 `json:"size" form:"size" query:"size"`
}

type GetUsersReq struct {
	Pagination
	Account
}

type Account struct {
	Name  string `json:"name"  form:"name" query:"name"`
	Phone string `json:"phone"  form:"phone" query:"phone"`
}

type PhoneChangeReq struct {
	UserPk string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	Phone  string `json:"phone"  form:"phone" query:"phone"`
}

type _UserPk struct {
	UserPk string `json:"user_pk"  form:"user_pk" query:"user_pk"`
}

type VerifyUserReq struct {
	Name   string `json:"name"  form:"name" query:"name"`
	IdCard string `json:"id_card"  form:"id_card" query:"id_card"`
	Face   string `json:"face"  form:"face" query:"face"`
	Url    string `json:"url"  form:"url" query:"url"`
}

type CreateUserCredentialsReq struct {
	Name          string `json:"name"  form:"name" query:"name"`
	Serial        string `json:"serial"  form:"serial" query:"serial"`
	Type          int32  `json:"type"  form:"type" query:"type"`
	FrontFileName string `json:"front_file_name"  form:"front_file_name" query:"front_file_name"`
	BackFileName  string `json:"back_file_name"  form:"back_file_name" query:"back_file_name"`
}

type DeleteUserCredentialsReq struct {
	Pk string `json:"pk"  form:"pk" query:"pk"`
}

type SetUserBaseInfoReq struct {
	Pk       string `json:"pk"  form:"pk" query:"pk"`
	Name     string `json:"name"  form:"name" query:"name"`
	IdCard   string `json:"id_card"  form:"id_card" query:"id_card"`
	Birthday string `json:"birthday"  form:"birthday" query:"birthday"`
	Gender   int32  `json:"gender"  form:"gender" query:"gender"`
	Height   string `json:"height"  form:"height" query:"height"`
	Weight   string `json:"weight"  form:"weight" query:"weight"`
}
