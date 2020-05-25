/**
 * @author  BakaErii & dem0r
 * @date    2020-1-14 22:06
 */

package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Config template
var configTemplate string = `{
    "secKey": "MUST CHANGE THIS!",
    "port": 80,
	"enableFront": true,
	"enableBuiltInFront": true,
	"enableDatabase": false,
	"enableCache": true,
	"cacheLife": 3,
	"enableHTTPS": false,
	"certFile": "Your cert file path(You should have permission)",
	"keyFile": "Your key file path(You should have permission)"
}
`

// Type : configs of server.
type ServerConfig struct {
	SecKey             string   `json:"secKey"`
	Port               string   `json:"port"`
	EnableFront        bool     `json:"enableFront"`
	EnableBuiltInFront bool     `json:"enableBuiltInFront"`
	FrontPort          string   `json:"frontPort"`
	FrontAddr          []string `json:"frontAddr"`
	EnableDatabase     bool     `json:"enableDatabase"`
	DBConnectURI       string   `json:"dbConnectURI"`
	DBMaxConnection    int      `json:"dbMaxConnection"`
	EnableCache        bool     `json:"enableCache"`
	CacheLife          int      `json:"cacheLife"`
	EnableHTTPS        bool     `json:"enableHTTPS"`
	CertFile           string   `json:"certFile"`
	KeyFile            string   `json:"keyFile"`
}

var serverConfig ServerConfig

// Type : report data from client.
type ReportData struct {
	SecKey             string          `json:"secKey"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	CpuUsage           int             `json:"cpuUsage"`
	MemUsage           int             `json:"memUsage"`
	MaxMemLimit        int             `json:"maxMemLimit"`
	DiskUsage          int             `json:"diskUsage"`
	MaxDiskLimit       int             `json:"maxDiskLimit"`
	InBound            float64         `json:"inBound"`
	InBoundTotalUsage  float64         `json:"inBoundTotalUsage"`
	OutBound           float64         `json:"outBound"`
	OutBoundTotalUsage float64         `json:"outBoundTotalUsage"`
	ServiceStatus      map[string]bool `json:"serviceStatus"`
}

// Type : report info about client.
type ClientReportInfo struct {
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	CpuUsage           int             `json:"cpuUsage"`
	MemUsage           int             `json:"memUsage"`
	MaxMemLimit        int             `json:"maxMemLimit"`
	DiskUsage          int             `json:"diskUsage"`
	MaxDiskLimit       int             `json:"maxDiskLimit"`
	InBound            float64         `json:"inBound"`
	InBoundTotalUsage  float64         `json:"inBoundTotalUsage"`
	OutBound           float64         `json:"outBound"`
	OutBoundTotalUsage float64         `json:"outBoundTotalUsage"`
	ServiceStatus      map[string]bool `json:"serviceStatus"`
	LastReportTime     int64           `json:"lastReportTime"`
}

var clientReportInfos map[string]ClientReportInfo

// Load config file.
func loadConfigFile() error {
	// Check if config not exist
	_, err := os.Stat("goss-server-config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[Warning] Is the first time run goss server?\n")
			fmt.Printf("[Warning] Creating config template goss-server-config.json.\n")
			fmt.Printf("[Warning] Update it and run goss server again\n")
			configFile, err := os.Create("goss-server-config.json")
			if err != nil {
				fmt.Errorf("[ERROR] Create config error. %s\n", err.Error())
				return err
			}
			defer configFile.Close()
			_, err = configFile.Write([]byte(configTemplate))
			if err != nil {
				fmt.Errorf("[ERROR] Write config error. %s\n", err.Error())
				return err
			}
			fmt.Printf("[Info] Already create config template.\n")
		}
		return err
	}
	// Try load config file
	viper.SetConfigName("goss-server-config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	return err
}

// Load config.
func loadConfig() {
	// Load server's secret key
	serverConfig.SecKey = viper.GetString("SecKey")
	fmt.Printf("[Info] Server secret key : %s .\n", serverConfig.SecKey)
	// Load server's port
	serverConfig.Port = viper.GetString("port")
	fmt.Printf("[Info] Server port : %s .\n", serverConfig.Port)
	// Load if server enable front
	serverConfig.EnableFront = viper.GetBool("enableFront")
	fmt.Printf("[Info] Server enable front : %v .\n", serverConfig.EnableFront)
	if serverConfig.EnableFront {
		// Load if server enable built-in front
		serverConfig.EnableBuiltInFront = viper.GetBool("enableBuiltInFront")
		fmt.Printf("[Info] Server enable built-in front : %v .\n", serverConfig.EnableBuiltInFront)
		if serverConfig.EnableBuiltInFront {
			// Load server's built-in front port
			serverConfig.FrontPort = viper.GetString("frontPort")
			fmt.Printf("[Info] Server built-in front address : %s .\n", serverConfig.FrontPort)
		} else {
			// Load server's front address
			serverConfig.FrontAddr = viper.GetStringSlice("frontAddr")
			fmt.Printf("[Info] Server front address : %s .\n", serverConfig.FrontAddr)
		}
	}
	// Load if server enable database
	serverConfig.EnableDatabase = viper.GetBool("enableDatabase")
	fmt.Printf("[Info] Server enable database : %v .\n", serverConfig.EnableDatabase)
	if serverConfig.EnableDatabase {
		// Load server's database(MongoDB) connection URI
		serverConfig.DBConnectURI = viper.GetString("dbConnectURI")
		fmt.Printf("[Info] Server database connection URI : %s second.\n", serverConfig.DBConnectURI)
		// Load server's max database connection limit
		serverConfig.DBMaxConnection = viper.GetInt("dbMaxConnection")
		fmt.Printf("[Info] Server max database connection limit : %d .\n", serverConfig.DBMaxConnection)
	}
	// Load if server enable cache
	serverConfig.EnableCache = viper.GetBool("enableCache")
	fmt.Printf("[Info] Server enable cache : %v .\n", serverConfig.EnableCache)
	if serverConfig.EnableCache {
		// Load cache life time
		serverConfig.CacheLife = viper.GetInt("cacheLife")
		fmt.Printf("[Info] Server cache life time is : %d second(s).\n", serverConfig.CacheLife)
	}
	// Load if server enable HTTPS
	serverConfig.EnableHTTPS = viper.GetBool("enableHTTPS")
	fmt.Printf("[Info] Server enable HTTPS : %v .\n", serverConfig.EnableHTTPS)
	if serverConfig.EnableHTTPS {
		// Load CRT path
		serverConfig.CertFile = viper.GetString("certFile")
		fmt.Printf("[Info] Cert file path is : %s .\n", serverConfig.CertFile)
		// Load KEY path
		serverConfig.KeyFile = viper.GetString("keyFile")
		fmt.Printf("[Info] Key file path is : %s .\n", serverConfig.KeyFile)
	}
	return
}

// Handle built-in front
func handlerBuiltInFront(ctx iris.Context) {
	return
}

// Handle client's report
func handlerClientReport(ctx iris.Context) {
	var reportData ReportData
	var reportInfo ClientReportInfo
	err := ctx.ReadJSON(&reportData)
	if err != nil {
		ctx.HTML(err.Error())
		return
	}
	if reportData.SecKey != serverConfig.SecKey {
		ctx.HTML("SecKey do not match")
		return
	}
	reportInfo.Name = reportData.Name
	reportInfo.Description = reportData.Description
	reportInfo.CpuUsage = reportData.CpuUsage
	reportInfo.MemUsage = reportData.MemUsage
	reportInfo.MaxMemLimit = reportData.MaxMemLimit
	reportInfo.DiskUsage = reportData.DiskUsage
	reportInfo.MaxDiskLimit = reportData.MaxDiskLimit
	reportInfo.InBound = reportData.InBound
	reportInfo.InBoundTotalUsage = reportData.InBoundTotalUsage
	reportInfo.OutBound = reportData.OutBound
	reportInfo.OutBoundTotalUsage = reportData.OutBoundTotalUsage
	reportInfo.ServiceStatus = reportData.ServiceStatus
	reportInfo.LastReportTime = time.Now().Unix()
	clientReportInfos[reportInfo.Name] = reportInfo
	return
}

// Handle get client's status
func handlerStatus(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"status": clientReportInfos,
	})
	return
}

// Main Entry
func main() {
	app := iris.New()
	app.Options("{root:path}", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
	})
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://127.0.0.1:8080"},
	//	AllowCredentials: true,
	//})
	//apiParty := app.Party("/", crs).AllowMethods(iris.MethodOptions)
	apiParty := app.Party("/").AllowMethods(iris.MethodOptions)
	clientReportInfos = make(map[string]ClientReportInfo, 8)
	// Load config
	configErr := loadConfigFile()
	if configErr != nil {
		panic(fmt.Errorf("[ERROR] Load config error. %s \n", configErr))
		return
	}
	loadConfig()
	// Reg router
	apiParty.Post("/report", handlerClientReport)
	apiParty.Post("/status", iris.Cache(time.Duration(serverConfig.CacheLife)*time.Second), handlerStatus)
	// Check if enable built-in front
	if serverConfig.EnableBuiltInFront {
		app.RegisterView(iris.HTML("./front", ".html"))
		apiParty.Get("/", func(ctx iris.Context) {
			ctx.View("index.html")
		})
		app.HandleDir("/","./front")
	}
	// Check if enable HTTPS
	if serverConfig.EnableHTTPS {
		app.Run(iris.TLS(":"+serverConfig.Port, serverConfig.CertFile, serverConfig.KeyFile))
	} else {
		app.Run(iris.Addr(":" + serverConfig.Port))
	}
	return
}
