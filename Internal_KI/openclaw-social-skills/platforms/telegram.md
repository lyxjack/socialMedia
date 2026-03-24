# Telegram Skills

## publora-telegram — Telegram 频道发布与排程

- **Slug**: `sergebulaev/publora-telegram`
- **安装**: `clawhub install sergebulaev/publora-telegram`
- **安全**: Benign(VT) / Suspicious(OC) — 建议审查源码
- **版本**: 1.2.0 | **下载**: 356

### 功能
- 发布内容到 Telegram 频道和群组
- 定时发布（ISO 格式时间戳）
- 自动限流队列管理（Telegram 30 msg/sec 全局 + 20 msg/min/群）
- 图片上传（最大 10MB）
- 视频上传（最大 50MB）
- 文本内容（最大 4096 字符）

### 配置
- Publora API Key（`x-publora-key` header）
- Telegram Bot Token（@BotFather 获取）
- Publora 账户 + 在 Dashboard 连接 Telegram 频道
- Bot 需有目标频道/群的管理员或发送消息权限

### 使用场景
- 产品发布公告排程
- 社区图片更新发布
- 自动化 Release Notes 推送
- 每日小贴士定时发送
- 私有群组视频预览分发
