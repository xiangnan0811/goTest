package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required,min=3"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpParam struct {
    Age        uint8  `json:"age" binding:"gte=1,lte=130"`
    Name       string `json:"name" binding:"required,min=5,max=10"`
    Email      string `json:"email" binding:"required,email"`
    Password   string `json:"password" binding:"required"`
    RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

var trans ut.Translator

func removeTopStruct(fields map[string]string) map[string]string {
    r := make(map[string]string, len(fields))
    for field, val := range fields {
        r[field[strings.Index(field, ".")+1:]] = val
    }
    return r
}

func InitTrans(locale string) (err error) {
    // 修改 gin 的 Validator 引擎属性，实现自定义翻译器
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        // 注册一个获取 json tag 的自定义方法
        v.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
            if name == "-" {
                return ""
            }
            return name
        })
        zhT := zh.New() // 中文翻译器
        enT := en.New() // 英文翻译器
        // 第一个参数是备用（fallback）的语言环境, 后面参数是应该支持的语言环境（支持多个）
        uni := ut.New(enT, enT, zhT)
        trans, ok = uni.GetTranslator(locale)
        if !ok {
            return fmt.Errorf("uni.GetTranslator(%s)", locale)
        }
        switch locale {
        case "en":
            en_translations.RegisterDefaultTranslations(v, trans)
        case "zh":
            zh_translations.RegisterDefaultTranslations(v, trans)
        default:
            en_translations.RegisterDefaultTranslations(v, trans)
        }
        return
    }
    return
}

func TimeConsumingMiddleware() gin.HandlerFunc {
    return func (c *gin.Context) {
        s := time.Now()
        c.Next()
        fmt.Printf("time consuming: %v\n", time.Since(s))
        status := c.Writer.Status()
        fmt.Println("status: ", status)
    }
}

func main() {
    if err := InitTrans("zh"); err != nil {
        fmt.Println("获取翻译器出错")
        return
    }
	r := gin.Default()
    r.Use(TimeConsumingMiddleware())
    r.POST("/loginJSON", loginJSON)
    r.POST("/signup", signUp)
    r.Run(":8080")
}

func signUp(c *gin.Context) {
    var data SignUpParam
    if err := c.ShouldBindJSON(&data); err != nil {
        errs, ok := err.(validator.ValidationErrors)
        if !ok {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{
            "error": removeTopStruct(errs.Translate(trans)),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status": "注册成功",
    })
}

func loginJSON(c *gin.Context) {
    var data Login
    if err := c.ShouldBindJSON(&data); err != nil {
        errs, ok := err.(validator.ValidationErrors)
        if !ok {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{
            "error": removeTopStruct(errs.Translate(trans)),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status": "you are logged in",
    })
}