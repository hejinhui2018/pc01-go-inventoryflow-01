# InventoryFlow 现场说明

InventoryFlow 是一个自建的库存事件与订单预留服务。仓库接收入库、出库和调整事件，系统将事件写入本地 JSON 快照，按 SKU/仓库维护库存状态，再输出对账报表和浏览器台账页面。

## 业务链

HTTP/CLI 输入 -> `workflow.Service` 校验并登记事件 -> `inventory.Ledger` 保存事件、`ReservationBook` 管理预留 -> `reconcile.Reconciler` 按版本推进状态 -> `report` 生成汇总 -> `/v1/report` 和 `web/dist` 输出。

## 本地运行

`go run ./cmd/inventoryflow -data ./data -port 8080`

健康检查：`curl http://localhost:8080/health`；报表：`curl http://localhost:8080/v1/report`。前端页面位于 `http://localhost:8080/`。

## 验证

`go test ./... -count=1`、`go vet ./...`、`go build ./...`。前端在 `web` 目录执行 `npm install` 后运行 `npm run build`。
