# A Web Crawler App

# 快速啟動

Client

```bash
cd web/app/ && yarn && yarn dev
```

Server

```bash
cd cmd/web_crawler && go run main.go
```

# Client

前端由 TypeScript/React 所撰寫，使用者可以在畫面上輸入 URL 已建立一個建立爬蟲，系統會返回一個 ID，使用者隨時可以使用此 ID 查詢任務的狀態

![Screenshot 2023-01-04 at 8 32 41 PM](https://user-images.githubusercontent.com/58166555/210576103-5d05fc47-ce2f-48d5-b9a8-21776eb46941.png)

當使用者輸入完 URL 且按下「建立爬蟲任務」後，前端會將 URL 送給後端處理

# Server

後端由 Go 所撰寫，包含以下幾個服務

- HTTP Service — 負責處理 HTTP 請求與回應
- Job Producer — 負責爬蟲任務的前處理
- Job Queue — 負責緩存爬蟲任務，提供系統能處理大量請求的彈性
- Job Comsumer — 負責執行緩存中的爬蟲任務，與資料後處理
- Crawler — 負責實際執行爬蟲

建立爬蟲任務的執行流程

![Screenshot 2023-01-04 at 10 26 06 PM](https://user-images.githubusercontent.com/58166555/210578579-c9ca743b-99ce-4e5a-af02-cce75ee77256.png)

1. 前端在畫面中輸入 URL，按下送出後，向後端發出一個請求。
2. 後端由 HTTP Service 接收到請求後，解析請求中的 URL，接著交由 Job Producer 處理。
3. Producer 會產生出一個 ID 作為此任務的唯一 ID。接著會將此 ID 作為 key 存入資料庫後，隨即將 ID 返回給 HTTP Service。
4. HTTP Serveice 收到 ID 後，回應使用者「新增任務成功」。
5. 接著，Producer 會把 ID 與 URL 組成一個任務，把任務推到 Job Queue 內。
6. Job Queue 為一個 queue，只要長度不為 0，Job Queue 就會把 queue 中最前面的任務交給 Job Consumer 執行。
7. Consumer 會非同步的調用 Crawler 執行爬蟲任務，爬蟲結束後，會把結果存到資料庫中，使用同一個 ID 複寫掉 Result 的欄位。

查詢爬蟲結果的流程

![Screenshot 2023-01-04 at 10 28 19 PM](https://user-images.githubusercontent.com/58166555/210578606-2135d605-5354-4ee7-9780-5ee279d0856d.png)

1. 前端在畫面中輸入任務 ID，按下送出後，向後端發出一個請求。
2. 後端使用此 ID 在資料庫中查找任務，若找到結果，回傳給前端。

## Services

### HTTP Service

- 使用 Gin 框架建立
- 能夠處理兩個 POST 請求，一個是建立爬蟲任務的請求、另一個是查詢爬蟲任務的請求
- 需要與資料庫連線，用於查詢任務時

### Job Producer

- 需要與資料庫連線，用於將新的任務存入資料庫
- 能夠建立一個完整個任務，需要包含至少以下資訊：任務 ID、URL
- 能夠呼叫 Job Queue 的 Equeue API

### Job Queue

- 使用 Go channel 實作
- 內部包含一個 queue，儲存完整的任務
- 當 queue 長度不為 0 時，要能夠把任務交由 Job Consumer 處理
- API：
  - Enqueue — 外部將任務加入 queue 的方法

### Job Consumer

- 能夠監聽 Job Queue，將 queue 內的任務 dequeue 出來執行
- 能夠解析任務內容，並且能併發執行，使用 Goroutine
- 能夠呼叫 Crawler 的 Parse API 以執行爬蟲任務
- 需要與資料庫連線，用於儲存任務時

### Crawler

- 使用 Colly 套件
- 返回爬蟲後的結果
- API
  - Parse — 外部調用爬蟲功能

# 容錯機制

- 若 Server 壞了，能夠將 queue 內的任務回復嗎？
  - 可以將 queue 內的任務存在資料庫中，以確保伺服器 crash 後，還能夠還原任務進度
  - 可以寫一個排程執行的服務，定期檢查 queue 內任務，並將其備份等等
  - 可以換成更可靠的 Message Queue 服務，例如 Kafka 或 RabbitMQ
- 若任務執行失敗，queue 內需要怎麼處理任務？
  - 可能需要在任務內定義好 fallback function，當 Consumer 執行失敗時執行
  - 可以再設計一個 queue，專門放失敗的任務，在特定時間點啟動，例如主要的 queue 全部執行完後
- 若 Consumer 執行很久，要如何判斷是否需要 dequeue？
  - 可以在 Consumer 中設計一個 API，用以檢查 Consumer 是否還存活，若還存活，可以等到 Consumer 執行完成
- 若任務會執行很久，可以怎麼優化？
  - 可以使用多個 goroutine 來處理任務
  - 可以把任務拆成很多個小任務
- 若 Consumer 執行到一半壞了怎麼辦？
  - 設計重試機制，當 Consumer 執行到一半壞了，直接重新開始執行
  - 可以採用事件驅動的架構來處理，例如當某個 consumer 執行失敗時，發送事件給系統，請系統來決定該任務如何繼續

# 擴展性

- 使用分散式 Message Queue 服務，來 scale message queue 的處理與容錯
- 建立一個 Goroutine pool，設定 Goroutine 數量的上下限，能夠動態調配執行任務的 Goroutine

# 部署

前端與後端服務都已經寫好 Dockerfile，夠確保順利 build 出能夠執行的 image，且我也寫了一份 docker-compose.yaml，所以部署時只需要再跑一次 docker build ，建立好 image 後在目標機器上使用 docker-compose up 就可以啟用服務了。
