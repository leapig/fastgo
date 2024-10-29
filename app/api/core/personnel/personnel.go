package personnel

import (
	"github.com/gin-gonic/gin"
	utils "github.com/leapig/fastgo/app/library/util"
)

func Load(r *gin.RouterGroup) {
	//管理平台:部门
	r.GET("department", utils.TenantAccessTokenJWTMiddleware(), SelectAllDepartment)
	r.POST("department", utils.TenantAccessTokenJWTMiddleware(), CreateDepartment)
	r.PUT("department", utils.TenantAccessTokenJWTMiddleware(), UpdateDepartment)
	r.DELETE("department", utils.TenantAccessTokenJWTMiddleware(), DeleteDepartment)
	r.GET("department/member", utils.TenantAccessTokenJWTMiddleware(), GetDepartmentMember)
	//人员管理
	r.GET("enterprise/user", utils.TenantAccessTokenJWTMiddleware(), SelectEnterpriseUser)
	r.POST("enterprise/user", utils.TenantAccessTokenJWTMiddleware(), CreateEnterpriseUser)
	r.PUT("enterprise/user", utils.TenantAccessTokenJWTMiddleware(), UpdateEnterpriseUser)
	r.PUT("enterprise/user/dimission", utils.TenantAccessTokenJWTMiddleware(), DimissionEnterpriseUser)
	r.DELETE("enterprise/user", utils.TenantAccessTokenJWTMiddleware(), DeleteEnterpriseUser)
	r.GET("enterprise/user/detail", utils.TenantAccessTokenJWTMiddleware(), SelectEnterpriseUserDetail)
	r.GET("enterprise/user/inservice", utils.TenantAccessTokenJWTMiddleware(), SelectInServiceEnterpriseUser)
}

type Pagination struct {
	Page int32 `json:"page" form:"page" query:"page"`
	Size int32 `json:"size" form:"size" query:"size"`
}

type DeleteReq struct {
	Pk string `json:"pk"  form:"pk" query:"pk"`
}

type GetEnterpriseUserReq struct {
	Pagination
	Name      string `json:"name"  form:"name" query:"name"`
	Phone     string `json:"phone"  form:"phone" query:"phone"`
	Gender    int32  `json:"gender"  form:"gender" query:"gender"`
	JobTitle  string `json:"job_title"  form:"job_title" query:"job_title"`
	JobNumber string `json:"job_number"  form:"job_number" query:"job_number"`
	Status    int32  `json:"status"  form:"status" query:"status"`
}

type PostEnterpriseUserReq struct {
	Phone                 string   `json:"phone"  form:"phone" query:"phone"`
	Name                  string   `json:"name"  form:"name" query:"name"`
	DepartmentPk          []string `json:"department_pk"  form:"department_pk" query:"department_pk"`
	JobNumber             string   `json:"job_number"  form:"job_number" query:"job_number"`
	JobTitle              string   `json:"job_title"  form:"job_title" query:"job_title"`
	Birthday              string   `json:"birthday"  form:"birthday" query:"birthday"`
	Gender                int32    `json:"gender"  form:"gender" query:"gender"`
	Height                string   `json:"height"  form:"height" query:"height"`
	Weight                string   `json:"weight"  form:"weight" query:"weight"`
	Nation                string   `json:"nation"  form:"nation" query:"nation"`
	Education             string   `json:"education"  form:"education" query:"education"`
	EmergencyContact      string   `json:"emergency_contact"  form:"emergency_contact" query:"emergency_contact"`
	EmergencyContactPhone string   `json:"emergency_contact_phone"  form:"emergency_contact_phone" query:"emergency_contact_phone"`
	TermOfContract        string   `json:"term_of_contract"  form:"term_of_contract" query:"term_of_contract"`
	Address               string   `json:"address"  form:"address" query:"address"`
}

type EnterpriseUserAttachment struct {
	Pk           string `json:"pk"  form:"pk" query:"pk"`
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	UserPk       string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	FileName     string `json:"file_name"  form:"file_name" query:"file_name"`
	FileType     string `json:"file_type"  form:"file_type" query:"file_type"`
}

type PutEnterpriseUserReq struct {
	Pk                    string                     `json:"pk"  form:"pk" query:"pk"`
	Phone                 string                     `json:"phone"  form:"phone" query:"phone"`
	Name                  string                     `json:"name"  form:"name" query:"name"`
	JobNumber             string                     `json:"job_number"  form:"job_number" query:"job_number"`
	JobTitle              string                     `json:"job_title"  form:"job_title" query:"job_title"`
	Birthday              string                     `json:"birthday"  form:"birthday" query:"birthday"`
	Gender                int32                      `json:"gender"  form:"gender" query:"gender"`
	Height                string                     `json:"height"  form:"height" query:"height"`
	Weight                string                     `json:"weight"  form:"weight" query:"weight"`
	DepartmentPk          []string                   `json:"department_pk"  form:"department_pk" query:"department_pk"`
	Nation                string                     `json:"nation"  form:"nation" query:"nation"`
	Education             string                     `json:"education"  form:"education" query:"education"`
	EmergencyContact      string                     `json:"emergency_contact"  form:"emergency_contact" query:"emergency_contact"`
	EmergencyContactPhone string                     `json:"emergency_contact_phone"  form:"emergency_contact_phone" query:"emergency_contact_phone"`
	EntryTime             string                     `json:"entry_time"  form:"entry_time" query:"entry_time"`
	TermOfContract        string                     `json:"term_of_contract"  form:"term_of_contract" query:"term_of_contract"`
	Address               string                     `json:"address"  form:"address" query:"address"`
	ContractFileName      []EnterpriseUserAttachment `json:"contract_file_name"  form:"contract_file_name" query:"contract_file_name"`
	BodyReportFileNames   []EnterpriseUserAttachment `json:"body_report_file_names"  form:"body_report_file_names" query:"body_report_file_names"`
	ReviewFileNames       []EnterpriseUserAttachment `json:"review_file_names"  form:"review_file_names" query:"review_file_names"`
}

type GetDepartmentMemberReq struct {
	Pagination
	DepartmentPk string `json:"department_pk"  form:"department_pk" query:"department_pk"`
	JobTitle     string `json:"job_title"  form:"job_title" query:"job_title"`
	JobNumber    string `json:"job_number"  form:"job_number" query:"job_number"`
	IsLeader     int32  `json:"is_leader"  form:"is_leader" query:"is_leader"`
	Name         string `json:"name" form:"name" query:"name"`
	Phone        string `json:"phone" form:"phone" query:"phone"`
}

type PostDepartmentReq struct {
	Name     string `json:"name"  form:"name" query:"name"`
	ParentPk string `json:"parent_pk"  form:"parent_pk" query:"parent_pk"`
}

type PutDepartmentReq struct {
	Pk       string `json:"pk"  form:"pk" query:"pk"`
	Name     string `json:"name"  form:"name" query:"name"`
	Leader   string `json:"leader" form:"leader" query:"leader"`
	ParentPk string `json:"parent_pk"  form:"parent_pk" query:"parent_pk"`
}

type SelectApplicationRecordReq struct {
	Pagination
	PositionPk      string `json:"position_pk"  form:"position_pk" query:"position_pk"`
	InterviewResult string `json:"interview_result"  form:"interview_result" query:"interview_result"`
	SourceType      string `json:"source_type"  form:"source_type" query:"source_type"`
	EnterprisePk    string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
	StartTime       string `json:"start_time"  form:"start_time" query:"start_time"`
	EndTime         string `json:"end_time"  form:"end_time" query:"end_time"`
}

type SelectUserEntryReq struct {
	Pagination
	Pk        string `json:"pk"  form:"pk" query:"pk"`
	Status    string `json:"status"  form:"status" query:"status"`
	Name      string `json:"name"  form:"name" query:"name"`
	Phone     string `json:"phone"  form:"phone" query:"phone"`
	StartTime string `json:"start_time"  form:"start_time" query:"start_time"`
	EndTime   string `json:"end_time"  form:"end_time" query:"end_time"`
}

type SelectEnterpriseUserDetailReq struct {
	UserPk       string `json:"user_pk"  form:"user_pk" query:"user_pk"`
	EnterprisePk string `json:"enterprise_pk"  form:"enterprise_pk" query:"enterprise_pk"`
}

type DimissionEnterpriseUserReq struct {
	Pk string `json:"pk"  form:"pk" query:"pk"`
}
