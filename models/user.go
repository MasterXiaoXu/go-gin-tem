
package models


type User struct {

	Id             int64     `gorm:"column:id;primary_key"`
	Email          string    `gorm:"column:email;not null;unique"`
	Password       string    `gorm:"column:password;not null"`
	IsGoogle       bool      `gorm:"column:isGoogle;not null"`
	GoogleSeceret  string    `gorm:"column:googleSeceret"` // 谷歌私钥
	UserPrivateKey string    `gorm:"column:userPrivateKey;not null;unique"`
	// WalletAddress  string    `gorm:"column:walletAddress;not null;unique"` 
	AccountState   bool      `gorm:"column:accountState;"` // 账户是否可登陆
	
	// 邀请相关
	InviteCode     string    `gorm:"column:inviteCode;"`  //邀请码
    InviteUrl 	   string    `gorm:"column:inviteUrl;"`   //邀请链接
    InviteUserCode string    `gorm:"column:inviteUserCode;"`  //邀请人邀请码


}