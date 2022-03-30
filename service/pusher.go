package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/JathamJ/fisgo/conf"
	"github.com/JathamJ/fisgo/log"
	"github.com/radovskyb/watcher"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Pusher config
type PusherConfig struct {
	// Root dir path
	RootPath 		string 		`yaml:"root"`
	// Paths not push
	IgnorePath 		[]string	`yaml:"ignore"`
	// Receiver api address, only support http/https
	ReceiverApi	 	string		`yaml:"receiver"`
	// Remote root path
	ServerRootPath	string 		`yaml:"to"`
	// Push break when error effect
	ErrorBreak 		bool 		`yaml:"errorBreak"`
	// Thread num limit when push
	ThreadNum 		int 		`yaml:"threadNum"`
}

type Pusher struct {
	Config 		*PusherConfig
	Watcher 	*watcher.Watcher
}

// init a pusher instance by config file
func InitPusherByFile(configPath string) *Pusher {
	var config *PusherConfig

	conf.LoadConf(configPath, &config)
	return InitPusher(config)
}

// Init a pusher instance
func InitPusher(conf *PusherConfig) *Pusher {
	pusher := &Pusher{
		Config: 	conf,
		Watcher:	watcher.New(),
	}
	return pusher
}

// WatcherFiles init all files which need watch
func (p *Pusher) WatcherFiles() {
	r := regexp.MustCompile("[^~]$")
	p.Watcher.AddFilterHook(watcher.RegexFilterHook(r, false))
	p.Watcher.IgnoreHiddenFiles(true)
	if err := p.Watcher.AddRecursive(p.Config.RootPath); err != nil {
		log.Fatalf("Pusher WatcherFiles AddRecursive failed, err: %s", err.Error())
	}
	if err := p.Watcher.Ignore(p.Config.IgnorePath...); err != nil {
		log.Fatalf("Pusher WatcherFiles Ignore failed, err: %s", err.Error())
	}
	// Print a list of all of the files and folders currently
	i := 0
	for path, _ := range p.Watcher.WatchedFiles() {
		i++
		log.Infof("watch path: %s\n", path)
	}
	log.Infof("all watched files, total: %d", i)
}

// Watch start a time interval watch file changes.
func (p *Pusher) Watch() {
	go func() {
		for {
			select {
			case event := <-p.Watcher.Event:
				log.Debugf("event: %v", event)
				if !event.IsDir() {	// push file if not dir
					go func() {
						err := p.PushFile(event.Op.String(), event.Path)
						if err != nil {
							log.Warnf("push file: %s ...failed", event.Path)
						} else {
							log.Infof("push file: %s ...ok", event.Path)
						}
					}()
				}
			case err := <-p.Watcher.Error:
				log.Fatalf("Pusher Watch error: %s", err.Error())
			case <-p.Watcher.Closed:
				log.Fatal("Pusher watcher closed")
			}
		}
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := p.Watcher.Start(time.Millisecond * 100); err != nil {
		log.Fatalf("Pusher Watch Monitor failed, err: %s", err.Error())
	}
}

func (p *Pusher) getToPath(localPath string) string {
	return strings.Replace(localPath, p.Config.RootPath, p.Config.ServerRootPath, 1)
}

func (p *Pusher) PushFile(op, localPath string) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(localPath)
	if err != nil {
		log.Debugf("PushFile Open localPath failed, err: %s", err.Error())
		return err
	}
	defer file.Close()
	part1, err := writer.CreateFormFile("file", filepath.Base(localPath))
	_, err = io.Copy(part1, file)
	if err != nil {
		log.Debugf("PushFile Io Copy failed, err: %s", err.Error())
		return err
	}
	_ = writer.WriteField("op", op)
	_ = writer.WriteField("to", p.getToPath(localPath))

	if err = writer.Close(); err != nil {
		log.Debugf("PushFile Close Writer failed, err: %s", err.Error())
		return err
	}

	client := &http.Client {}
	req, err := http.NewRequest("POST", p.Config.ReceiverApi, payload)
	if err != nil {
		log.Debugf("PushFile NewRequest failed, err: %s", err.Error())
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Debugf("PushFile DoRequest failed, err: %s", err.Error())
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debugf("PushFile Ioutil ReadAll failed, err: %s", err.Error())
		return err
	}
	ret := string(body)
	if ret != "1" {
		err = errors.New(fmt.Sprintf("PushFile response: %s, not '1'", ret))
		log.Debugf("PushFile Ioutil ReadAll failed, err: %s", err.Error())
		return err
	}
	return nil
}





























