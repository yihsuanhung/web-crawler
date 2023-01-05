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
- Job Queue — 負責暫存爬蟲任務，提供系統能處理大量請求的彈性
- Job Comsumer — 負責執行暫存中的爬蟲任務，與資料後處理
- Crawler — 負責實際執行爬蟲

建立爬蟲任務的執行流程

![Page 2 (2)](https://user-images.githubusercontent.com/58166555/210790475-aedfa27a-a04f-4c1b-a506-52058d2e7b0e.png)

1. 前端在畫面中輸入 URL，按下送出後，向後端發出一個請求。
2. 後端由 HTTP Service 接收到請求後，解析請求中的 URL，接著交由 Job Producer 處理。
3. Producer 會產生出一個 ID 作為此任務的唯一 ID。接著會將此 ID 作為 key 存入資料庫後，隨即將 ID 返回給 HTTP Service。
4. HTTP Serveice 收到 ID 後，回應使用者「新增任務成功」。
5. 接著，Producer 會把 ID 與 URL 等資訊組成一個「完整的任務」，把任務推到 Job Queue 內。
6. Job Queue 為一個 queue，能夠讓 Job Consumer 從 queue 中把取出並執行。
7. Consumer 會非同步的調用 Crawler 執行爬蟲任務，爬蟲完成後，會把結果存到資料庫中，使用同一個 ID 複寫掉 Result 的欄位。

查詢爬蟲結果的流程

![Page 2 (3)](https://user-images.githubusercontent.com/58166555/210790716-33de3236-e0e0-463f-96cf-a851ee066e04.png)

1. 前端在畫面中輸入任務 ID，按下送出後，向後端發出一個請求。
2. 後端使用此 ID 在資料庫中查找任務，若找到結果，回傳給前端。

## Services

### HTTP Service

- 使用 Gin 框架建立
- 能夠處理至少兩個請求：建立爬蟲任務的請求與查詢爬蟲任務的請求
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
- 能夠解析任務內容，
- 能併發執行任務，使用 Goroutine 設計
- 能夠呼叫 Crawler 的 Parse API 以執行爬蟲任務
- 需要與資料庫連線，用於儲存任務狀態

### Crawler

- 使用 Colly 套件
- 返回爬蟲後的結果
- API
  - Parse — 外部調用爬蟲功能

# 容錯機制與擴充性

為了增強系統的容錯機制與擴充性，可以進行以下調整

- Job Queue 能夠跟所有 Consumers 溝通，並且自己維護一些狀態，包含：
  - Consumer 的執行狀態，例如成功、失敗、執行中。
  - 任務與 Consumer 的對應關係，例如 A 任務由 ID 為 123 的 Consumer 執行。
- Job Queue 設定一個 time out，當 Consmer 執行超時，派定為失敗。
- Job Queue 設計每隔固定時間，例如 5 秒，發請求詢問執行中的 Consumer 狀態的功能。類似 ELK 的 Heartbeat，如果 Consumer 回覆執行中，就繼續等待任務處理完成，如果 Consumer 回傳失敗，或是 time out，即可判定任務失敗，該任務重新執行或重新排隊。
- Job Queue 只會在 Consumer 回傳任務成功時，才會真的把任務從 queue 內 pop 掉。
- Job Queue 只會在確定任務成功，才更新資料庫內的任務狀態。

能夠解決或應付以下問題：

- 若 Consumer 壞了，執行到一半的任務怎麼辦？
  - 由於 Job Queue 只會在 Consumer 回傳成功時才把任務從 queue 中 pop 掉，所以當 Consumer 壞掉時，任務還在 queue 內等待，不會寫入 db 也不會影響結果正確性。
  - Job Queue 會有固定排程與所有 Consumer 溝通，當 Job Queue 檢查到某個 Consumer 壞掉時，只要把該任務重新排隊等待執行即可。
  - 另外也可以採用事件驅動的軟體設計，當 Consumer 失敗時，發送事件給系統，交由系統處理失敗任務的後續行為。
- 若 Queue 壞了，排隊中的任務如何回復？
  - Queue 壞了重啟之後，走訪一遍所有的 Consumer ，檢查 Consumer 狀態，把 Consumer 正在執行中的任務重新 enqueue 即可。
  - 若很擔心 Queue 的備份，可以寫一個排程服務，定期備份 queue 的狀態。
- 若任務執行失敗，Queue 內需要怎麼處理任務？
  - 當 Consumer 回傳給 Queue 任務執行失敗時，可以把任務重新排列到 queue 最後重新等待執行。
- 若 Consumer 需要執行很久，要如何判斷結果？
  - Heartbeat 的設計能夠確保 Consumer 的執行正確性，若 Job Queue 定時傳送 heartbeat 給 consumer，且 consumer 都有回傳正確，就讓 Consumer 繼續執行任務，直到失敗或 Time out。

# 部署

前端與後端服務都已經寫了 Dockerfile，方便部署前建立 image，其中後端採用 multi-stage build 以最小化 image size。可以使用 Gitlab 或 Github action 將 image 部署至目標機器上。
