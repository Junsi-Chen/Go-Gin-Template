package server

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"template/conf"
	"template/models"
	"template/router"
	"template/tool/log"
	"template/tool/mysql"
	"template/tool/util"
)

type Server struct {
	gin *gin.Engine // 路由服务
	db  *gorm.DB    // db数据库服务
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(conf *conf.Config) error {
	var err error
	log.Init(conf.App.ServerName)

	// 初始化mysql
	s.db, err = mysql.InitEngine(conf.Mysql)
	if err != nil {
		return err
	}

	// 初始化redis
	//err = redis.InitClient(conf.Redis)
	//if err != nil {
	//	return err
	//}
	// 赋值到models包的db
	models.InitDb(s.db)
	// 表同步
	models.AutoMigrate()
	// jwt
	util.InitJwtSecret()
	// gin
	gin.SetMode(conf.Http.Mode)
	s.gin = gin.Default()
	router.InitRouter(s.gin)
	// validator // 入参校验翻译器
	err = util.InitTrans()
	if err != nil {
		return err
	}
	log.Logger.Info("server init success")
	return nil
}
func (s *Server) Run() error {
	// 读取CA证书
	caCert, err := os.ReadFile("conf/ca/ca.crt")
	if err != nil {
		return err
	}
	// 创建一个CA证书池并添加CA证书
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// 配置 TLS
	tlsConfig := &tls.Config{
		ClientCAs: certPool,
		// 不需要客户端的证书，只需要服务器验证（常用）
		ClientAuth: tls.NoClientCert,
	}

	// 启动HTTPS 并配置客户端证书验证
	srv := &http.Server{
		Addr:      conf.Conf.Http.Port,
		Handler:   s.gin,
		TLSConfig: tlsConfig,
	}

	return srv.ListenAndServeTLS("conf/ca/server.crt", "conf/ca/server.key")
}
