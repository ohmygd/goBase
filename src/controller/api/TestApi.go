package api

import (
	"context"
	"controller"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/ohmygd/mgo/merror"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"io/ioutil"
	"log"
	"model/mysql"
	"net/http"
	"pconst"
	"runtime"
	"sync"
	"time"
)

type Logger struct {
	logPath string
	logName string
}

var (
	b *int
	ee chan int
	aa int

	lock sync.Mutex
	syncMap sync.Map
)

var db *gorm.DB

func init() {
	c := 5
	b = &c

	ee = make(chan int, 100)

	var err error
	db, err = gorm.Open("mysql", "root:123456@/mall?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(6)
	db.DB().SetMaxIdleConns(4)
	db.DB().SetConnMaxLifetime(time.Minute)

	db.LogMode(true)

	pool = &redis.Pool{     //实例化一个连接池
		MaxIdle:2,    //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout:40,    //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn ,error){     //要连接的redis数据库
			return redis.Dial("tcp","localhost:6379")
		},
	}
}

func NewB() *int {
	return b
}

var aaa = map[string]interface{}{
	"mc":1,
	"dj":2,
}

var pool *redis.Pool

func Test25(c *gin.Context) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//操作etcd
	_, err = cli.Put(ctx, "/name123", "mc12")
	fmt.Println(err, "---------")
	//操作完毕，取消etcd
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	//取值，设置超时为1秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "name2211")
	fmt.Println(resp, err, "===========")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func Test24(c *gin.Context) {
	err := merror.New(pconst.ErrorOk)
	controller.Render(c, err, nil)
	//controller.Render(c, nil, nil)
	return
}

func Test23(c *gin.Context) {
	log.Println("eeeeee")
}

/**
config 读取
 */
func Test22(c *gin.Context) {
	viper.SetConfigName("etcd")
	viper.AddConfigPath("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("22222")
	}

	a := viper.Get("name")
	b := viper.Get("age")
	d := viper.Get("arr")
	e := viper.Get("info.add")

	fmt.Println(a,b,d,e)
}

func Test21(c *gin.Context) {
	cc := pool.Get()
	defer cc.Close()

	_, err := cc.Do("set", "mc", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success")
}


func Test20(c *gin.Context) {
	fmt.Println("222")
	mc()
	fmt.Println(444)
}

func mc() {
	fmt.Println("123")
	defer func(){
		fmt.Println("456")
	}()
	fmt.Println(789)
}

func Test19(c *gin.Context) {


	var user = &mysql.User{}
	db.First(&user)

	user = &mysql.User{
		Name:"ok",
		Mobile:"123",
		Openid:"222",
	}

	fmt.Println(user)
}

func Test18(c *gin.Context) {
	ioutil.ReadFile("config/etcd.yaml")
	config, err := yaml.ReadFile("config/etcd.yaml")
	if err != nil {
		fmt.Println(err)
		panic("config file not exists")
	}

	fmt.Println(config.Get("path"))
	fmt.Println(config.Get("info"))
	fmt.Println(config.Get("arr"))
}

func Test17(c *gin.Context) {
	//log.Debug("mc test log")
	//log.Info("mc test log")
	//log.Warn("mc test log")
	//log.Error("mc test log")
}

func Testmc1() {
	Testmc()
}

func Testmc() {
	for i:=0;i<5;i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			file = "???"
			line = 0
		}

		fmt.Println(file, line, "-------")
	}
}

func Test16(c *gin.Context) {
	for i:=0;i<10;i++ {
		go func(){
			aaa["dj"] = i
		}()
	}

	fmt.Println(aaa)
}

func Test15(c *gin.Context) {
	var once sync.Once
	res := make(chan bool)

	ss := func() {
		fmt.Println("once test")
	}

	for i := 0;i<10;i++ {
		go func() {
			once.Do(ss)

			res <- true
		}()
	}

	for j:=0;j<10;j++ {
		<- res
	}
}

func Test14(c *gin.Context) {
	timer := time.NewTimer(time.Second)
	go func() {
		for{
			<-timer.C
			fmt.Println("Timer 2 expired")
		}

	}()
}

func Test13(c *gin.Context) {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		//for t := range ticker.C {
		//	fmt.Println("Tick at", t)
		//}
		for {
			fmt.Println("----")
			time.Sleep(time.Second)

			fmt.Println(len(ticker.C), <-ticker.C)
		}
	}()
}

func Test12(c *gin.Context) {
	c1 := make(chan string, 1)
	go func(){
		time.Sleep(time.Second * 5)

		c1 <- "result 1"
	}()

	select {
	case a := <-c1:
		fmt.Println(a, "=====")
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	default:
		fmt.Println("default--------")
	}
}

func Test11(c *gin.Context) {
	go func(s chan int) {
		for i:=0;i<5;i++ {
			s <- i
		}

	}(ee)

	go func(a chan int) {
		for i:=5;i<10;i++ {
			a <- i
		}

		close(a)
	}(ee)


	fmt.Println(cap(ee), "-----")
}

func Test9(c *gin.Context) {
	d := NewB()
	e := NewB()

	fmt.Println(d, e, "======")
}

func Test10(c *gin.Context) {
	flag.Parse()
	defer glog.Flush()

	glog.Error("mc merror")
}

func Test(c *gin.Context) {
	fmt.Println("test func")
	c.String(http.StatusOK, "test")
}

func Test1(c *gin.Context) {
	fmt.Println("test1 func")
	c.String(http.StatusOK, "test1")
}

func Test2(c *gin.Context) {
	fmt.Println(123)
}

func Test3(c *gin.Context) {
	var param controller.Test3

	if c.ShouldBindQuery(&param) == nil {
		fmt.Println(param.Name, param.Age, "_________")
	}

	c.JSON(http.StatusOK, gin.H{
		"code":1001,
		"msg":"success",
		"data":nil,
	})

	fmt.Println("-------- 123")
}

func Test4(c *gin.Context) {
	var param controller.Test3

	// shouldBind post
	// shouldBindQuery get
	if err := c.ShouldBind(&param); err == nil {
		fmt.Println(param.Name, param.Age, "_________")
	} else {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":1001,
		"msg":"success",
		"data":nil,
	})

	fmt.Println("-------- 123")
}

func Test5(c *gin.Context) {
	 a := map[string]interface{}{}
	a["name"] = "mc"
	fmt.Println(a)
}

var a = 1

func Test6(c *gin.Context) {
	for i:=0;i<10;i++ {
		go func(){
			a++
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(a, "_____")
}

func Test7(c *gin.Context) {
	var param controller.Test7

	var err error

	if err = c.ShouldBindQuery(&param); err == nil {
		fmt.Println(param, "=======")
	}
	fmt.Println(err, "err-------")

	fmt.Println(param, "over")
}

func Test8(c *gin.Context) {
	//log.Println("hello")
}