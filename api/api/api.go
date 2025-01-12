package api

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/GlennLucy1/learn-ai222/ctrl"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type DetectArg struct {
	Pixel    [][]uint8 `json:"pixel"`
	RightKey string    `json:"right_key"`
}

func Detect(c *gin.Context) {
	arg := &DetectArg{}
	if err := c.ShouldBindBodyWithJSON(arg); err != nil {
		c.String(500, "param parse error")
		return
	}
	apiKey := c.GetHeader("ApiKey")
	if apiKey != "123e4567-e89b-12d3-a456-426614174000" {
		c.String(400, "auth failed")
		return
	}

	fmt.Println("new request, right key is ", arg.RightKey)

	dir, err := os.MkdirTemp("", "ocr-*")
	if err != nil {
		c.String(500, "mkdir temp")
		return
	}
	defer os.RemoveAll(dir)

	pngOutPath := filepath.Join(dir, "out.png")
	if err := ctrl.DrawFromPixelArray(arg.Pixel, pngOutPath); err != nil {
		c.String(500, "draw pic")
		return
	}
	fmt.Println("png draw done, saved at ", pngOutPath)

	fileDat, err := os.ReadFile(pngOutPath)
	if err != nil {
		c.String(500, "read file error")
		return
	}

	b64Img := base64.StdEncoding.EncodeToString(fileDat)
	result, err := ctrl.Detect(context.Background(), b64Img)
	if err != nil {
		c.String(500, "detect error")
		return
	}
	fmt.Println("ocr result ", result)
	result = ctrl.CleanResult(result)

	rightKey := strings.Split(arg.RightKey, "#")
	rightNum := ctrl.CalcRate(result, rightKey)

	// 识别率 <= 2个文字，添加到待处理表
	if rightNum <= 2 {
		fmt.Println("detect num is 2")
		c.String(200, "1234")
		return
	}

	// 识别率 == 3个文字，开始结果推测
	if rightNum == 3 {
		fmt.Println("detect num is 3")
		result = ctrl.Speculate(result, rightKey)
		fmt.Println("speculated result is ", result)
	} else {
		fmt.Println("ocr detected all")
	}

	c.String(200, ctrl.GetOrder(result, rightKey))
}

func StartServer(addr string) {
	r := gin.Default()
	r.POST("/", Detect)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	fmt.Println("listening on ", addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
