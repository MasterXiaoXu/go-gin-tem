package app

import (
	
   "bv/models"
   "bv/util"
   "github.com/gin-gonic/gin"
   "net/http"
   "log"
//    "fmt"
)


func Register(c *gin.Context)  {

	// 定义一个结构体用来接受数据
	type userStruct struct{
		Email            string 
		EmailCode        int
		Password         string
		InvitationCode   string
	}

	var userData userStruct
	// 数据绑定
	c.BindJSON(&userData)
	if userData.Email == ""{

		log.Print("func Register error: the email parameter cannot be empty")
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"the email parameter cannot be empty",
		 })
		 return
	}
	// TODO 验证 email格式

	if userData.Password == ""{

		log.Print("func Register error: the password parameter cannot be empty")
		
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"the password parameter cannot be empty",
		 })
		 return
	}

	if userData.EmailCode == 0{

		log.Print("func Register error: the emailCode parameter cannot be empty")
		
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"the password emailCode cannot be empty",
		 })
		 return

	}

	// TODO验证emailCode
	if userData.InvitationCode == ""{

		userData.InvitationCode = "BVBVBV"
	}

	var users []models.User
	
	// 验证是否已经注册
	_ = util.MysqlDb.Where("email = ?",userData.Email).Find(&users).Error
	if len(users) > 0{

		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"The user has registered",
		 })

		 return
	}
	
	userPrivateKey := util.RandStringBytesMaskImpr(32)
	inviteCode := util.RandStringBytesMaskImpr(6)
	passwordHash,_ := util.PasswordHash(userData.Password)
	// 拼接数据
	user := models.User{

		Email:           userData.Email,
		Password:        passwordHash,
		IsGoogle:        false,
		GoogleSeceret:   "",
		UserPrivateKey:  userPrivateKey,
		AccountState:    true,
		
		// 邀请相关
		InviteCode:      inviteCode,
		InviteUrl:       "https://www.binance.com/user/register?inviteCode=" + inviteCode,
		InviteUserCode:  userData.InvitationCode,

	}
	
	err := util.MysqlDb.Create(&user).Error
	if err!= nil{
		
		log.Printf("func CreateUser error:",err)
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"user registration failure",
		 })

		 return 
	}
	
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"success",
	 })
 }

 
 func Login(c *gin.Context)  {

	// 定义一个结构体用来接受数据
	type userStruct struct{
		Email            string 
		Password         string
	}

	var userData userStruct
	// 数据绑定
	c.BindJSON(&userData)

	// 验证数据	
	if userData.Email == ""{

		log.Print("func Login error: the email parameter cannot be empty")
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"the email parameter cannot be empty",
		 })
		 return
	}


	if userData.Password == ""{

		log.Print("func Login error: the password parameter cannot be empty")
		
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"the password parameter cannot be empty",
		 })
		 return
	}

	var user models.User
	
	// 验证是否已经注册
	var err error
	err = util.MysqlDb.Where("email = ?",userData.Email).First(&user).Error

	if err != nil{
		if user.Id == 0{
			
			c.JSON(http.StatusOK,gin.H{
				"code":-2,
				"msg":"user not registered",
			 })

			return
	
		}
		// 记录log
		log.Print("func Login error: data query exception")
		c.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"Service exception",
		 })

		 return

	}
	
	if util.PasswordVerify(userData.Password,user.Password){

		c.JSON(http.StatusOK,gin.H{
			"code":-3,
			"msg":"wrong password",
		 })

		 return
		
	}

	
	// 拼接返回数据
	
	

	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"success",
	 })
	 return

	
 }
