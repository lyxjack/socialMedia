# Instagram / Threads Skills

## instagram-api — Instagram & Threads 官方 API 发帖

- **Slug**: `lifeissea/instagram-api`
- **安装**: `clawhub install lifeissea/instagram-api`
- **安全**: Suspicious(VT) / Suspicious(OC) — **高风险，务必审查源码后使用**
- **版本**: 最新

### 功能
- 通过 Meta Graph API 直接发帖
- 支持内容类型: Feed 帖子、Story、Reels、Carousel（轮播图）
- Threads 跨平台发帖
- Imgur 集成（本地媒体转公开 URL）

### 配置
```bash
INSTAGRAM_ACCESS_TOKEN=           # Meta Graph API 长期令牌
INSTAGRAM_BUSINESS_ACCOUNT_ID=    # 商业账号 ID
IMGUR_CLIENT_ID=                  # 匿名使用即可
# Threads 额外需要:
THREADS_ACCESS_TOKEN=
THREADS_USER_ID=
```

### 前置条件
- Meta for Developers 账号 + Business App
- `instagram_content_publish` 权限已启用

### 工作流示例
1. 上传图片到 Imgur 获取公开 URL
2. 通过 Meta API 创建 media container
3. 使用账号凭证发布
4. 通过 API 响应验证成功

### 使用场景
- 产品图发帖
- 短视频 Reels
- 多图轮播
- Story 图片发布
- Threads 文字分享
