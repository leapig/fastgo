package enterprise

type Pagination struct {
	Page int32 `json:"page" form:"page" query:"page"`
	Size int32 `json:"size" form:"size" query:"size"`
}

type EnterpriseReq struct {
	Pagination
	Name        string `json:"name"  form:"name" query:"name"`
	Type        int32  `json:"type" form:"type" query:"type"`
	Cover       string `json:"cover" form:"cover" query:"cover"`
	Serial      string `json:"serial" form:"serial"  query:"serial"`
	License     string `json:"license" form:"license"  query:"license"`
	Country     string `json:"country" form:"country"  query:"country"`
	Province    string `json:"province" form:"province"  query:"province"`
	City        string `json:"city" form:"city"  query:"city"`
	District    string `json:"district" form:"district"  query:"district"`
	County      string `json:"county" form:"county"  query:"county"`
	Site        string `json:"site" form:"site"  query:"site"`
	CorporatePk string `json:"corporate_pk" form:"corporate_pk" query:"corporate_pk"`
	Longitude   string `json:"longitude" db:"longitude"  query:"longitude"`
	Latitude    string `json:"latitude" db:"latitude"  query:"latitude"`
	AddressCode string `json:"address_code" db:"address_code"  query:"address_code"`
}

type PostSuperviseReq struct {
	Name        string `json:"name"`
	Country     string `json:"country" `
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	County      string `json:"county" `
	Site        string `json:"site"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AddressCode string `json:"address_code"`
}

type PostSecurityReq struct {
	Name        string `json:"name"`
	StaffSize   string `json:"staff_size"`
	Country     string `json:"country" `
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	County      string `json:"county" `
	Site        string `json:"site"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AddressCode string `json:"address_code"`
}

type PostEmploymentReq struct {
	Name        string `json:"name"`
	Country     string `json:"country" `
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	County      string `json:"county" `
	Site        string `json:"site"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AddressCode string `json:"address_code"`
}

type PutSecurityReq struct {
	Pk          string `json:"pk"`
	Name        string `json:"name"`
	StaffSize   string `json:"staff_size"`
	Country     string `json:"country" `
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	County      string `json:"county" `
	Site        string `json:"site"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AddressCode string `json:"address_code"`
}

type PutEmploymentReq struct {
	Pk          string `json:"pk"`
	Name        string `json:"name"`
	Country     string `json:"country" `
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	County      string `json:"county" `
	Site        string `json:"site"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AddressCode string `json:"address_code"`
}

type PutCorporateReq struct {
	Pk          string `json:"pk"`
	CorporatePk string `json:"corporatePk"`
}

type PostNewReq struct {
	Phone          string `json:"phone"  form:"phone" query:"phone"`
	UserName       string `json:"user_name"  form:"user_name" query:"user_name"`
	Country        string `json:"country"  form:"country" query:"country"`
	Province       string `json:"province"  form:"province" query:"province"`
	City           string `json:"city"  form:"city" query:"city"`
	District       string `json:"district"  form:"district" query:"district"`
	County         string `json:"county"  form:"county" query:"county"`
	Site           string `json:"site"  form:"site" query:"site"`
	Longitude      string `json:"longitude"  form:"longitude" query:"longitude"`
	Latitude       string `json:"latitude"  form:"latitude" query:"latitude"`
	EnterpriseName string `json:"enterprise_name"  form:"enterprise_name" query:"enterprise_name"`
	Code           string `json:"code"  form:"code" query:"code"`
	StaffSize      string `json:"staff_size" form:"staff_size" query:"staff_size"`
	AddressCode    string `json:"address_code" form:"address_code" query:"address_code"`
}

type SelectEnterpriseAreaPermissionReq struct {
	Pagination
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
}

type CreateEnterpriseAreaPermissionReq struct {
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	Province     string `json:"province"  form:"province" query:"province"`
	City         string `json:"city"  form:"city" query:"city"`
	District     string `json:"district"  form:"district" query:"district"`
	County       string `json:"county"  form:"county" query:"county"`
	AddressCode  string `json:"address_code" form:"address_code" query:"address_code"`
}

type DeleteEnterpriseAreaPermissionReq struct {
	Pk string `json:"pk"  form:"pk" query:"pk"`
}
