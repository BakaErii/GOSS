/**
 * @author  BakaErii & dem0r
 * @date    2020-1-14 22:06
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Config template
var configTemplate string = `
{
    "secKey": "This is your server's secret key",
    "name": "Client's name (UNIQUE)",
	"description": "Anything....",
    "reportServer": "http://example.com/report",
	"reportInterval": 5,
    "reportService": [
		"nginx",
        "process name",
		"Windows process should add .exe like this.exe"
    ]
}
`

// Type : report info about client.
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

var reportData ReportData

// Type : configs of client.
type ClientConfig struct {
	SecKey         string   `json:"secKey"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ReportServer   string   `json:"reportServer"`
	ReportInterval uint64   `json:"reportInterval"`
	ReportService  []string `json:"reportService"`
}

var clientConfig ClientConfig

// Load config file.
func loadConfigFile() error {
	// Check if config not exist
	_, err := os.Stat("goss-config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[Warning] Is the first time run goss client?\n")
			fmt.Printf("[Warning] Creating config template goss-config.json.\n")
			fmt.Printf("[Warning] Update it and run goss client again\n")
			configFile, err := os.Create("goss-config.json")
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
	viper.SetConfigName("goss-config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	return err
}

// Load config.
func loadConfig() {
	// Load client's secret key
	clientConfig.SecKey = viper.GetString("SecKey")
	fmt.Printf("[Info] Client secret key : %s .\n", clientConfig.SecKey)
	reportData.SecKey = clientConfig.SecKey
	// Load client's name
	clientConfig.Name = viper.GetString("name")
	fmt.Printf("[Info] Client name : %s .\n", clientConfig.Name)
	reportData.Name = clientConfig.Name
	// Load client's description
	clientConfig.Description = viper.GetString("description")
	fmt.Printf("[Info] Client description : %s .\n", clientConfig.Description)
	reportData.Description = clientConfig.Description
	// Load client's report interval
	clientConfig.ReportInterval = viper.GetUint64("reportInterval")
	fmt.Printf("[Info] Client report interval : %d second.\n", clientConfig.ReportInterval)
	// Load client's report server address
	clientConfig.ReportServer = viper.GetString("reportServer")
	fmt.Printf("[Info] Client report server address : %v .\n", clientConfig.ReportServer)
	// Load client's report status of service
	clientConfig.ReportService = viper.GetStringSlice("reportService")
	fmt.Printf("[Info] Client report status of service : %v .\n", clientConfig.ReportService)
	return
}

// Update client's status
func updateStatus() {
	// Get CPU usage
	cpuPer, _ := cpu.Percent(time.Duration(clientConfig.ReportInterval)*time.Second, false)
	reportData.CpuUsage = int(cpuPer[0])
	// Get Mem & MaxMem
	memInfo, _ := mem.VirtualMemory()
	reportData.MemUsage = int(memInfo.Used / 1024 / 1024)
	reportData.MaxMemLimit = int(memInfo.Total / 1024 / 1024)
	// Get disk usage
	diskInfo, _ := disk.Usage("/")
	reportData.DiskUsage = int(diskInfo.Used / 1024 / 1024)
	reportData.MaxDiskLimit = int(diskInfo.Total / 1024 / 1024)
	// Get network I/O
	netInfo, _ := net.IOCounters(false)
	reportData.InBound = (float64(netInfo[0].BytesRecv)/1024/1024 - reportData.InBoundTotalUsage) / float64(clientConfig.ReportInterval)
	reportData.InBoundTotalUsage = float64(netInfo[0].BytesRecv) / 1024 / 1024
	reportData.OutBound = (float64(netInfo[0].BytesSent)/1024/1024 - reportData.OutBoundTotalUsage) / float64(clientConfig.ReportInterval)
	reportData.OutBoundTotalUsage = float64(netInfo[0].BytesSent) / 1024 / 1024
	// Get service status
	for i := range clientConfig.ReportService {
		reportData.ServiceStatus[clientConfig.ReportService[i]] = false
	}
	pids, err := process.Pids()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range pids {
		p := process.Process{Pid: pids[i]}
		name, _ := p.Name()
		if _, ok := reportData.ServiceStatus[name]; ok {
			reportData.ServiceStatus[name] = true
		}
	}
	return
}

// Report status to server usd POST method
func reportStatus() {
	postData := new(bytes.Buffer)
	reqClient := &http.Client{}
	json.NewEncoder(postData).Encode(reportData)
	req, err := http.NewRequest("POST", clientConfig.ReportServer, postData)
	if err != nil {
		fmt.Errorf("[ERROR] %s\n", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := reqClient.Do(req)
	if err != nil {
		fmt.Errorf("[ERROR] %s\n", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
	return
}

// Main Entry
func main() {
	// Load config
	configErr := loadConfigFile()
	if configErr != nil {
		panic(fmt.Errorf("[ERROR] Load config error. %s \n", configErr))
		return
	}
	loadConfig()
	reportData.ServiceStatus = make(map[string]bool, len(clientConfig.ReportService))
	// Update client status and report it to server
	for true {
		updateStatus()
		reportStatus()
		time.Sleep(time.Duration(clientConfig.ReportInterval) * time.Second)
	}
	return
}
