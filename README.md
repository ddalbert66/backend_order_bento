

# Golang Web專案

- 採前後分離
- 對應前端Vue : <https://github.com/ddalbert66/vue_order_bento>

# 前端:
	使用前端框架: vue.js(v2.5.2)
-	套件版本:
	>axios: v0.21.1,<br>
	element-ui: v2.14.1,<br>
    	vue-router: v3.0.1,<br>
    	vuex: v3.6.0,<br>
	net: v1.0.2,
	
# 後端:
	使用語言:golang
-	依賴:
	>gin-gonic/gin v1.6.3  // 網頁框架<br>
	go-sql-driver/mysql v1.5.0 // SQL DRIVER<br>
	jinzhu/gorm v1.9.16 // ORM套件<br>
	google/uuid v1.1.4 // uuid套件<br>
	uber-go/zap v1.16.0 // 日誌記錄套件<br>
	github.com/spf13/viper v1.7.1 // 設定檔存讀套件<br>
	
# 其他技術應用:
-
	>ngnix : http跳轉使用/IP限制使用<br>
	redis server : 登入緩存使用<br>
	mysql server : 系統資料保存<br>

## 依賴包說明
### 1. gin-gonic

- 一種go的http web framework
- 使用gin來快速配置路由以及中間件(達成類似JAVA AOP)
- 在benchmark中，效能高於其他frame

### 2. uber-go/zap

- uber團隊的zap包用於日誌紀錄
- 高性能的日誌紀錄工具，效能高於其餘日誌紀錄
- 具有報錯顯示完整路徑並不會強制中斷程序的error錯誤日誌類型

### 3. viper

- viper為設定檔管理工具
- 支持JSON，TOML，YAML，HCL，envfile，Properties等等檔案
- 支持從環境變數中存取
- 支持熱讀取

### 4. gorm 

- ORM工具，方便從資料庫中存寫數據。
- 可快速的使用預設CRUD，使用方法輕鬆易懂
- 可自定義Logger，且默認Logger已有足夠的資料
- 支持事務回滾

### 5. go-redis

- 緩存存儲工具
- 使用於登入緩存、在線會員緩存、資料鎖

### 6. google/uuid

- 快速產生不會重複的唯一值
- 目前作用存取用戶識別(產生session id給予用戶)

### 7. gorilla/websocket

- 用於快速部屬長連線通訊協定
- 用於即時回傳訊息給用戶
- 用於即時給前端新增好友通知、新群組通知等等提示訊息


## 功能簡介

### 1. 註冊

![](https://github.com/ddalbert66/backend_order_bento/raw/master/resource/image/02.註冊頁面.jpg) 
> /src/controller/userController.go Register

```
	// 取得前端請求
	var data userReq
	err := ctx.Bind(&data)	
	if err != nil {
		zapLog.ErrorW("register error!:", err)
		return
	}
	resp := make(gin.H)
	user := userService.QueryUserByName(data.Name)
	zapLog.WriteLogInfo("user register", zap.String("name", user.Name))

	// 判斷帳號是否可註冊，並回傳json訊息給予前端
	if user.ID != 0 {
		resp["msg"] = "已註冊的帳號"
		resp["code"] = "error"
	} else {
		user.Name = data.Name
		user.Pwd = data.Password
		userService.Insert(user)
		resp["msg"] = "註冊成功"
	}
	ctx.JSON(http.StatusOK, resp)
```

### 2. 登入

![](https://github.com/ddalbert66/backend_order_bento/raw/master/resource/image/01.登入頁面.jpg) 

- 登入主要為驗證使用者輸入資料正確，以及存入REDIS和用戶端cookie
> /src/controller/userController.go Login
```
...
user := userService.QueryUserByName(data.Name)  // 由user.userService查詢DB
if user.ID != 0 { 								// 若不存在則登入失敗
	if user.Pwd != data.Password { /* 密碼檢查 start */
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "密碼錯誤",
			"code": "error",
		})
		return
	}

	user.SessionId = uuid.New().String()		//產生新的UUID
	... /* 解析部分資料 */
	redisdb.Set(constant.LoginKey+user.SessionId, userJson, time.Hour*3) 				//資料放入redis緩存 存活時間三小時
	redisdb.HSet(constant.LoginOnlineHash, constant.LoginKey+user.SessionId, userJson)	//資料放入在線會員清單
	ctx.SetCookie("sessionId", user.SessionId, int(time.Hour*3), "/", "", false, true)	//資料放入用戶
...
```

### 3. 中間件驗證登入狀態

> - 設定於所有登入後API，由/src/middleware/loginCheck.go驗證是否登入狀態

```
	//非登入時的輸出狀態
	var out gin.H = gin.H{
		"code": "notLogin",
		"msg":  "尚未登入",
	}

...
	data, err := ctx.Cookie("sessionId") 			//讀取用戶cookie

	if err != nil {
		zapLog.ErrorW("login check error!:", err)
		ctx.JSON(http.StatusOK, out)
		ctx.Abort()									//Abort: 不繼續執行其餘handler
		return
	}

	// 存取redsi 若已經有資料且可轉models.User則Pass
	redisdb := utils.GetRedisDb()
	cmd := redisdb.Get(constant.LoginKey + data)
	if cmd.Err() != nil || cmd.Val() == "" {		//無登入資訊判定為非登入
		fmt.Printf("err: %v , value %v\n", cmd.Err(), cmd.Val())
		ctx.JSON(http.StatusOK, out)
		ctx.Abort()
		return
	} else {
		redisdb.Expire(constant.LoginKey, time.Hour*3) 	//若驗證通過則延長登入時效三小時
		var user models.User
		err := json.Unmarshal([]byte(cmd.Val()), &user)	//從redis中取得的user資料解析為物件
		if err != nil {
			zapLog.ErrorW("login check err!", err)
			ctx.Abort()
			return
		}
		ctx.Set("user", user)  // 存入gin.context資料 此流程後續任何handler皆可取用使用者登入資訊
	}
	ctx.Next() //繼續執行其餘handler
```

- 若驗證不通過 前端接收回傳json判斷code為notLogin時 則出現彈窗並跳轉到登入頁面

![](https://github.com/ddalbert66/backend_order_bento/raw/master/resource/image/04.未登入提示.jpg)