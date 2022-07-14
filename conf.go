package conf

var (
	CLogLevel            int   = 3           //控製臺日誌等級
	FLogLevel            int   = 1           //檔案日誌等級
	LogLevel             int   = 1           //日誌等級 範例=>目前0:全印 設置1:不顯示debug訊息 ( 0:調試信息(全部列印) 1:輸出日誌信息 2:警告 3:錯誤 4:嚴重錯誤 )
	ConfigDirName              = "config"    //配置文件目錄名
	EVTCHANSIZE                = 20000       //事件通道大小
	MSGPACKERBUF         int32 = 1024 * 1000 //消息緩沖區大小
	READBUFLEN           int   = 1024 * 1000 //接收消息緩沖區大小
	DefaultEventThrottle int   = 10000       //預設的事件並發數
	ClearLog             bool  = false
)

var (
	TCPSERVERNAME_HALL_CLIENT     = "hallTCP_Client"   //大廳麵嚮客戶端TCP服務器名稱
	TCPSERVERNAME_HALL_GAMESERVER = "hallTCP_GSMANAGE" //大廳遊戲服務器管理TCP服務器名稱
	TCPSERVERNAME_ROOM_CLIENT     = "roomTcp_Client"   //房間麵嚮客戶端TCP服務器名稱

	TCPSERVERNAME_DBSERVER = "dbserverTCP" //dbserver tcp服務器名稱

	TCPSERVERCONN_HALL_CLIENT     = 100000 //大廳客戶端連接最大連接數
	TCPSERVERCONN_HALL_GAMESERVER = 20000  //大廳遊戲服務器連接最大連接數
	TCPSERVERCONN_ROOM_CLIENT     = 20000  //遊戲客戶端連接最大連接數

	TCPSERVERCONN_DBSERVER = 20000 //dbserverTCP連接數

	TCPSERVERCONN_NOTIFYSERVER = 20000             //notifyserver tcp連接數
	TCPSERVERNAME_NOTIFYSERVER = "notifyserverTCP" //notifyserver tcp服務器名稱
)

//服務器類型
const (
	ServerType_GameServer = iota //gameserver
	ServerType_HallServer        //hallserver
)

//服務器狀態
const (
	ServerStatus_Runing  int32 = 1 //運行中
	ServerStatus_Stoping int32 = 0 //停服中
	ServerStatus_Close   int32 = 2 //已關閉
	ServerStatus_Recover int32 = 3 //維護中
)
