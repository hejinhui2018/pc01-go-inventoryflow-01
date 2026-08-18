# Bug 复现记录

## 现象

同一个库存事件因消息重试再次到达时，事件版本没有前进，但对账器仍会再次累加数量。订单看见的可用库存会少于真实库存，重复重试还可能触发库存不足告警。

## 基线复现

在 Bug baseline 分支运行：

`go test ./internal/reconcile -run '^TestReplaySameVersionIsIgnored$' -count=1`

预期基线失败：测试收到的 `OnHand` 为 20 而不是 10。

## 正确行为

版本小于或等于已应用版本的事件都应视为重复/过期事件，保持状态不变并返回 `Duplicate=true`。版本严格前进时才可以应用事件数量。测试还覆盖过期版本、正常入库、报表和持久化行为。
