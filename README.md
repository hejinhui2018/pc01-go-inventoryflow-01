# InventoryFlow

InventoryFlow 是一个面向仓库运营的库存事件、订单预留与对账服务。系统接收入库、出库和人工调整事件，将事件按 SKU 和仓库形成台账，维护预留状态，并通过 HTTP API、CSV 和浏览器页面输出可用库存。

项目完全使用 Go 标准库，持久化层采用受控目录中的 JSON 快照和 JSONL 日志，便于本地部署和审计。核心链路为：输入校验、事件登记、状态对账、库存策略、持久化、报表输出。

## 开发命令

```text
go test ./... -count=1
go vet ./...
go build ./...
```

前端构建：进入 `web`，运行 `npm install` 和 `npm run build`。服务启动后访问 `http://localhost:8080/`。
