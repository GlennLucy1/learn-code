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
	"strconv"
)

type DetectArg struct {
	Pixel    [][]uint8 `json:"pixel"`
	RightKey string    `json:"right_key"`
}

type SingleText struct {
    pixels [][]uint8
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
	// defer os.RemoveAll(dir)

	textList := make([]*SingleText, 0)
	for i := 0; i < 4; i++ {
		singleText := &SingleText{
			pixels: make([][]uint8, 0),
		}
		for j := 0; j < 36; j++ {
			singleText.pixels = append(singleText.pixels, arg.Pixel[j][i*48:(i+1)*48])
		}
		textList = append(textList, singleText)
	}

	finalResult := make([]string, 0)
	for i, v := range textList {
		pngOutPath := filepath.Join(dir, fmt.Sprintf("single%d.png", i+1))
		if err := ctrl.DrawFromPixelArray(v.pixels, pngOutPath); err != nil {
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

		if len(result) <= 0 {
			finalResult = append(finalResult, strconv.Itoa(i+1))
		} else {
			finalResult = append(finalResult, result[0])
		}
	}
	fmt.Println("ocr result ", finalResult)

	rightKey := strings.SplitN(arg.RightKey, "", 4)
	rightNum := ctrl.CalcRate(finalResult, rightKey)

	// 识别率 <= 2个文字
	if rightNum <= 2 {
		fmt.Println("detect num is ", rightNum)
		// 结果随机推测
		finalResult = ctrl.RandomGenerate(finalResult, rightKey)
		fmt.Println("random result is ", finalResult)
	}

	// 识别率 == 3个文字
	if rightNum == 3 {
		fmt.Println("detect num is 3")
		// 结果准确推测
		finalResult = ctrl.Speculate(finalResult, rightKey)
		fmt.Println("speculated result is ", finalResult)
	} else {
		fmt.Println("ocr detected all")
	}

	c.String(200, ctrl.GetOrder(finalResult, rightKey))
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
