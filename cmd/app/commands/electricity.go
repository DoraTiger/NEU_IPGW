package commands

import (
	"fmt"
	"strings"

	"github.com/DoraTiger/NEU_IPGW/config"
	"github.com/DoraTiger/NEU_IPGW/internal/handler"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/spf13/cobra"
)

var (
	powerUsername string
	powerPassword string
	powerSaveCred bool
	powerAccount  string
	factoryCode   string
	roomNumber    string
	infoOnly      bool
	allInfo       bool
)

var PowerCmd = &cobra.Command{
	Use:   "power",
	Short: "query electricity room info",
	Run: func(cmd *cobra.Command, args []string) {
		lang := currentLanguage()
		credentialUsername := strings.TrimSpace(powerUsername)
		credentialPassword := powerPassword

		store, err := newCredentialStore()
		if err != nil {
			fmt.Println(err)
			return
		}

		if credentialUsername == "" && credentialPassword == "" {
			cred, err := loadCredentialForLogin(store, strings.TrimSpace(powerAccount))
			if err != nil {
				file, line := utils.GetErrorLocation()
				logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
				fmt.Println(err)
				fmt.Println("Please login first: NEU_IPGW login -u <username> -p <password> --save")
				return
			}
			credentialUsername = cred.Username
			credentialPassword = cred.Password
		} else if credentialUsername == "" || credentialPassword == "" {
			fmt.Println("both username and password are required when using inline credential")
			return
		}

		gwHandler := handler.NewGWHandler()
		gwHandler.SetLogger(logger)
		err = gwHandler.Login(credentialUsername, credentialPassword)
		if err != nil {
			fmt.Println(err)
			return
		}

		if powerSaveCred {
			if err := store.Save(credentialUsername, credentialPassword); err != nil {
				file, line := utils.GetErrorLocation()
				logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
				fmt.Println(err)
				return
			}
		} else {
			if err := store.MarkLast(credentialUsername); err != nil {
				file, line := utils.GetErrorLocation()
				logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			}
		}

		electricityHandler := handler.NewElectricityHandler()
		electricityHandler.SetLogger(logger)
		electricityHandler.SetClient(gwHandler.GetClient())

		if err := electricityHandler.Login(); err != nil {
			fmt.Printf("Login error: %v\n", err)
			return
		}

		defer func() {
			electricityHandler.Logout()
		}()

		var elecRoomNo string
		var dormitoryName string

		if roomNumber != "" {
			elecRoomNo = strings.TrimSpace(roomNumber)
			dormitoryName = config.TranslateDormitory(elecRoomNo)
		} else {
			roomInfo, err := electricityHandler.QueryDefaultRoomInfo(strings.TrimSpace(factoryCode))
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			elecRoomNo = roomInfo.ElecRoomNo
			dormitoryName = roomInfo.DormitoryName
		}

		if infoOnly {
			roomNumberPart := ""
			if len(elecRoomNo) > 3 {
				roomNumberPart = elecRoomNo[3:]
			}
			printRoomOnlyInfo(lang, elecRoomNo, dormitoryName, roomNumberPart)
			return
		}

		elecInfo, err := electricityHandler.QueryRoomInfo(elecRoomNo, strings.TrimSpace(factoryCode))
		if err != nil {
			fmt.Printf("Query electricity info error: %v\n", err)
			return
		}

		printElectricityInfo(lang, elecRoomNo, dormitoryName, elecInfo)
	},
}

func init() {
	registerFlagsPowerCmd(PowerCmd)
}

func registerFlagsPowerCmd(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&powerUsername, "username", "u", "", "the username for authentication")
	cmd.Flags().StringVarP(&powerPassword, "password", "p", "", "the password for authentication")
	cmd.Flags().BoolVar(&powerSaveCred, "save", false, "save credential locally and overwrite same username")
	cmd.Flags().StringVar(&powerAccount, "account", "", "saved account username to load")
	cmd.Flags().StringVarP(&factoryCode, "factory", "f", config.DefaultFactoryCode, "factory code for electricity query")
	cmd.Flags().StringVarP(&roomNumber, "room", "r", "", "manually specify room number to query")
	cmd.Flags().BoolVarP(&infoOnly, "info", "i", false, "only query room info (no electricity)")
	cmd.Flags().BoolVarP(&allInfo, "all", "a", false, "output full info (including description, last read time, etc.)")
}

func printRoomOnlyInfo(lang string, elecRoomNo, dormitoryName, roomNumber string) {
	if lang == "en" {
		fmt.Printf("Room No: %s\n", elecRoomNo)
		fmt.Printf("Dormitory: %s\n", dormitoryName)
		fmt.Printf("Room Number: %s\n", roomNumber)
	} else {
		fmt.Printf("宿舍编号: %s\n", elecRoomNo)
		fmt.Printf("宿舍名称: %s\n", dormitoryName)
		fmt.Printf("寝室号: %s\n", roomNumber)
	}
}

func printElectricityInfo(lang string, elecRoomNo, dormitoryName string, elecInfo *handler.ElectricityInfo) {
	roomNumber := ""
	if len(elecRoomNo) > 3 {
		roomNumber = elecRoomNo[3:]
	}

	if lang == "en" {
		fmt.Printf("Room No: %s\n", elecRoomNo)
		fmt.Printf("Dormitory: %s\n", dormitoryName)
		fmt.Printf("Room Number: %s\n", roomNumber)
		fmt.Printf("Electricity Remaining: %.2f kWh\n", elecInfo.RemainDegree)
		fmt.Printf("Electricity Balance: %.2f CNY\n", elecInfo.RemainYuan)
		if allInfo {
			fmt.Printf("Room Description: %s\n", elecInfo.ElecRoomInfo)
			fmt.Printf("Last Read Time: %s\n", elecInfo.LastReadTime)
		}
	} else {
		fmt.Printf("宿舍编号: %s\n", elecRoomNo)
		fmt.Printf("宿舍名称: %s\n", dormitoryName)
		fmt.Printf("寝室号: %s\n", roomNumber)
		fmt.Printf("剩余度数: %.2f 度\n", elecInfo.RemainDegree)
		fmt.Printf("电费余额: %.2f 元\n", elecInfo.RemainYuan)
		if allInfo {
			fmt.Printf("房间描述: %s\n", elecInfo.ElecRoomInfo)
			fmt.Printf("最后读表时间: %s\n", elecInfo.LastReadTime)
		}
	}
}
