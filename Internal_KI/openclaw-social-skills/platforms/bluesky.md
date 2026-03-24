# Bluesky Skills

## bluesky — Bluesky 全功能 CLI

- **Slug**: `jeffaf/bluesky`
- **安装**: `clawhub install jeffaf/bluesky`
- **安全**: Benign / Benign
- **版本**: 1.6.0 | **下载**: 5,000+ | **安装量**: 14

### 功能
- 发帖、创建帖子串
- 回复、点赞、转发
- 关注/取关、屏蔽/静音
- 搜索内容
- 图片上传
- 个人资料查看
- 通知监控
- JSON 输出（便于脚本化）
- 认证使用 App Password（不在磁盘存储 session token）

### 配置
- Bluesky 活跃账号
- App Password（bsky.app → Settings → Privacy and Security 生成）

### 关键命令
```bash
bsky login                    # 认证
bsky whoami                   # 检查登录状态
bsky create-thread            # 创建帖子串
bsky post                     # 发帖
# 标准互动命令: like, repost, follow, block, mute
```

### 使用场景
- 终端自动化发帖
- 批量账号管理（通过脚本）
- 帖子串排程与创建
- 无浏览器话题搜索
- 通知监控
