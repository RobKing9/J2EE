package controler

import (
	"Experiment1/common"
	"Experiment1/dao"
	"Experiment1/models"
	"Experiment1/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginSuccess(c *gin.Context) {
	//数据库
	db := dao.GetDB()
	//获取参数
	tel := c.PostForm("tel")
	psw := c.PostForm("psw")
	//验证数据
	//手机号码必须为11位
	if len(tel) != 11 {
		//response.Response(c, http.StatusUnprocessableEntity, 422, nil,"密码必须超过六位！" )
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "手机号码必须为11位！",
		})
		return
	}
	//密码必须超过六位
	if len(psw) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "密码必须超过六位！",
		})
		return
	}
	//判断手机号是否存在
	var u models.User
	db.Where("Tel=?", tel).First(&u)
	if u.Name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "该手机号未注册！"})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(u.Psw), []byte(psw)); err != nil {
		//response.Response(c, http.StatusUnprocessableEntity, 422, nil,"密码必须超过六位！" )
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确！"})
		return
	}
	//返回token
	token, err := common.ReleaseToken(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "系统异常！"})
		log.Println("系统异常 err=", err.Error())
		return
	}
	response.Success(c, gin.H{"token": token}, "登录成功!!!")
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegSuccess(c *gin.Context) {
	////数据库
	db := dao.GetDB()
	//获取参数
	username := c.PostForm("username")
	psw := c.PostForm("password")
	email := c.PostForm("email")
	name := c.PostForm("rename")
	tel := c.PostForm("telphone")
	sex := c.PostForm("gender")
	borndate := c.PostForm("birthday")
	//验证数据
	//手机号码必须为11位
	if len(tel) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "手机号码必须为11位！",
		})
		return
	}
	//密码必须超过六位
	if len(psw) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "密码必须超过六位！",
		})
		return
	}
	//用户名为空时 默认为robking
	if len(name) == 0 {
		name = "robking"
		return
	}
	//手机号码必须为11位
	if len(tel) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "手机号码必须为11位！",
		})
		return
	}

	//判断手机号码是否存在
	if IsTelExist(db, tel) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码已存在！")
		return
	}

	//密码加密
	hashedpsw, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 422, nil, "加密错误！")
		return
	}
	user := &models.User{
		Username: username,
		Psw:      string(hashedpsw),
		Email:    email,
		Name:     name,
		Tel:      tel,
		Sex:      sex,
		Borndate: borndate,
	}
	c.BindJSON(&user)
	//存入数据库
	err = db.Create(&user).Error
	if err != nil {
		response.Fail(c, nil, "注册失败！")
	} else {
		response.Success(c, nil, "注册成功！！！")
		//c.HTML(http.StatusOK, "login.html", nil)
		log.Println(name, tel, psw)
	}
}

// IsTelExist 判断手机号码是否存在
func IsTelExist(db *gorm.DB, tel string) bool {
	var u models.User
	db.Where("Tel=?", tel).First(&u)
	if u.Name != "" {
		return true
	}
	return false
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Resume(c *gin.Context) {
	c.HTML(http.StatusOK, "resume.html", nil)
}

func Mydream(c *gin.Context) {
	c.HTML(http.StatusOK, "mydream.html", nil)
}
func Mylife(c *gin.Context) {
	c.HTML(http.StatusOK, "mylife.html", nil)
}
func Learn(c *gin.Context) {
	c.HTML(http.StatusOK, "learn.html", nil)
}
