package service

import (
	"github.com/JathamJ/fisgo/log"
	"github.com/JathamJ/fisgo/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type Receiver struct {
	Port 	string
	Engine 	*gin.Engine
}

// init a receiver instance.
func InitReceiver(port string) *Receiver {
	return &Receiver{
		Port:	port,
	}
}

// ServerStart start a http server handle upload.
func (r *Receiver) ServerStart() {
	r.Engine = gin.Default()
	r.Engine.POST("/upload", HandleUpload)
	err := r.Engine.Run(":" + r.Port)
	if err != nil {
		log.Fatalf("ServerStart failed, err: %s", err.Error())
	} else {
		log.Infof("server start port :%s", r.Port)
	}
}

// Handle upload.
func HandleUpload(ctx *gin.Context) {
	op := ctx.PostForm("op")
	to := ctx.PostForm("to")
	if op == "" || to == "" {
		ctx.String(http.StatusOK, "0")
		return
	}

	_, fileHeaders, _ := ctx.Request.FormFile("file")
	f, err := fileHeaders.Open()
	if err != nil {
		log.Warnf("HandleUpload fileOpen failed, err: %s", err.Error())
		ctx.String(http.StatusOK, "0")
		return
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Warnf("HandleUpload IOReadAll failed, err: %s", err.Error())
		ctx.String(http.StatusOK, "0")
		return
	}

	// check and create filepath
	err = utils.CheckAndCreateFileDir(to)
	if err != nil {
		log.Warnf("HandleUpload CheckAndCreateFileDir failed, err: %s", err.Error())
		ctx.String(http.StatusOK, "0")
		return
	}

	fileObj, err := os.OpenFile(to, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Warnf("HandleUpload OpenFile failed, err: %s", err.Error())
		ctx.String(http.StatusOK, "0")
		return
	}
	defer fileObj.Close()
	_, _ = fileObj.Write(b)
	ctx.String(http.StatusOK, "1")
}
