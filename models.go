package global

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Client struct {
	Model
	UUID          string         `gorm:"index:idx_client_uuid,unique" json:"uuid"`
	Name          string         `json:"name"`
	AddressLine1  string         `json:"addressline1"`
	AddressLine2  string         `json:"addressline2"`
	AddressLine3  string         `json:"addressline3"`
	Phone         string         `json:"phone"`
	Email         string         `gorm:"index:idx_client_email,unique" json:"email" binding:"required"`
	Password      string         `gorm:"-" json:"password"`
	NewPwd        string         `gorm:"-" json:"new_pwd"`
	BcryptHash    string         `json:"-"`
	Organizations []Organization `gorm:"foreignKey:ClientID" json:"organizations"`
	Applications  []Application  `gorm:"foreignKey:ClientID" json:"applications"`
	Gorm          *gorm.DB       `gorm:"-" json:"-"`
	Sess          *gin.Context   `gorm:"-" json:"-"`
}

type Organization struct {
	Model
	ClientUUID string
	ClientID   uint   `gorm:"index:idx_organizations_name" json:"-"`
	UUID       string `gorm:"index:idx_organizations_uuid,unique" json:"uuid" binding:"required"`
	Name       string `gorm:"index:idx_organizations_name" json:"name" binding:"required"`
	Category   string `json:"category"`
	Comments   string `json:"comments"`
	JWTSecret  string `json:"-"`
	//Applications []Application `gorm:"foreignKey:OrganizationID" json:"applications"`
	//Users []User       `gorm:"foreignKey:OrganizationID" json:"users"`
	Gorm *gorm.DB     `gorm:"-" json:"-"`
	Sess *gin.Context `gorm:"-" json:"-"`
}

type Application struct {
	Model
	ClientUUID    string
	ClientID      uint   `gorm:"index:idx_applications_name,unique" json:"-"`
	UUID          string `gorm:"index:idx_applications_uuid,unique" json:"uuid"`
	Name          string `gorm:"index:idx_applications_name,unique" json:"name" binding:"required"`
	Category      string `json:"category"`
	ApiKey        string `json:"apikey"`
	RedirectRoute string `json:"redirect"`
	Description   string `json:"description"`
	// TODO:  create an AllowedRoles struct for the application to establish what is supported
	AllowedRoles []Role       `gorm:"foreignKey:ApplicationID" json:"roles"`
	Users        []User       `gorm:"many2many:user_application" json:"users"`
	Preload      bool         `gorm:"-" json:"preload"`
	Gorm         *gorm.DB     `gorm:"-" json:"-"`
	Sess         *gin.Context `gorm:"-" json:"-"`
}

type LoginRedirect struct {
	Model
	//UserID        uint
	ClientUUID      string
	ApplicationUUID uint         `gorm:"-"`
	RedirectRoute   string       `gorm:"-"`
	Gorm            *gorm.DB     `gorm:"-" json:"-"`
	Sess            *gin.Context `gorm:"-" json:"-"`
}

type User struct {
	Model
	OrganizationID uint `json:"-"`
	ApplicationID  uint
	UUID           string        `gorm:"index:idx_users_uuid,unique" json:"uuid"`
	FirstName      string        `json:"first"`
	MiddleName     string        `json:"middle"`
	LastName       string        `json:"last"`
	PreferredName  string        `json:"preferred"`
	Email          string        `gorm:"index:idx_users_unique,unique" json:"email" binding:"required"`
	Email2         string        `json:"email2"`
	Phone1         string        `json:"phone1"`
	Phone2         string        `json:"phone2"`
	AddressLine1   string        `json:"addressline1"`
	AddressLine2   string        `json:"addressline2"`
	AddressLine3   string        `json:"addressline3"`
	Location       string        `json:"location"`
	Username       string        `gorm:"-" json:"username"`
	Password       string        `gorm:"-" json:"password"`
	LoggedIn       bool          `gorm:"default:false" json:"-"`
	Roles          []Role        `gorm:"foreignKey:UserID" json:"roles"`
	Credentials    []Credential  `gorm:"foreignKey:UserID" json:"credentials"`
	Applications   []Application `gorm:"many2many:user_application" json:"applications"`
	//Attributes     []Attrs       `gorm:"foreignKey:UserID" json:"attributes"`
	Gorm *gorm.DB     `gorm:"-" json:"-"`
	Sess *gin.Context `gorm:"-" json:"-"`
}

type Role struct {
	Model
	ClientUUID    string
	UserID        uint         `gorm:"index:idx_roles_userid_applicationid,unique" json:"-"`
	ApplicationID uint         `gorm:"index:idx_roles_userid_applicationid,unique" json:"-"`
	UUID          string       `gorm:"index:idx_roles_uuid,unique" json:"uuid" binding:"required"`
	Name          string       `gorm:"index:idx_roles_userid_applicationid,unique" json:"name" binding:"required"`
	Gorm          *gorm.DB     `gorm:"-" json:"-"`
	Sess          *gin.Context `gorm:"-" json:"-"`
}

type ClientCredential struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Credential struct {
	Model
	UserID        uint   `gorm:"index:idx_creds_app_id_user_id_name,unique" json:"-"`
	ApplicationID uint   `gorm:"index:idx_creds_app_id_user_id_name,unique" json:"-"`
	UserName      string `gorm:"index:idx_creds_app_id_user_id_name,unique" json:"username"`
	Password      string `gorm:"-" json:"password"`
	NewPwd        string `gorm:"-" json:"new_pwd"`
	BcryptHash    string `json:"-"`
}

type LoginToken struct {
	Client Client
	Token  string
}

// type Item struct {
// 	Model
// 	UserID uint
// 	Attrs  Attrs
// }

// type Attrs map[string]interface{}
