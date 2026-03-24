# 多平台社媒管理 Skills

## adaptlypost — 多平台统一发布 (SaaS)

- **Slug**: `tarasshyn/adaptlypost`
- **安装**: `clawhub install tarasshyn/adaptlypost`
- **安全**: Benign / Benign
- **版本**: 1.0.3 | **下载**: 606

### 覆盖平台
Instagram, X(Twitter), Bluesky, TikTok, Threads, LinkedIn, Facebook, Pinterest, YouTube — **共 9 个平台**

### 功能
- 单次 API 调用发布/排程到最多 9 个平台
- 支持内容类型: 文本、图片、视频、轮播图
- 安全三步媒体上传流程
- 逐平台文本自定义（同一帖子，不同平台不同文案）
- 草稿创建、排程发布、即时发布

### 配置
```bash
ADAPTLYPOST_API_KEY=    # adaptlypost.com 注册获取
```

### 关键端点
- `GET /social-accounts` — 获取已连接的平台账号 ID
- `POST /social-posts` — 发布/排程/草稿

### 工作流
1. 获取账号连接信息
2. 确认发布偏好
3. 发送帖子请求（含平台指定配置）
4. 报告各平台成功/失败状态

---

## mixpost — 自托管社媒管理 (Self-hosted)

- **Slug**: `lao9s/mixpost`
- **安装**: `clawhub install lao9s/mixpost`
- **安全**: Benign / Benign
- **版本**: 1.0.0 | **下载**: 2,752 | **安装量**: 10

### 覆盖平台
Facebook, X, Instagram, LinkedIn, Pinterest, TikTok, YouTube, Threads — **共 8 个平台**

### 功能
- 跨平台同时排程发布
- 批量上传和标记媒体文件
- 查询已排程/失败帖子状态
- 自动化帖子审批工作流
- 一键删除已发布内容（App + 社交平台同步）
- REST API 完整的账户/帖子/媒体管理

### 配置
```bash
MIXPOST_URL=             # 自托管 Mixpost 实例地址
MIXPOST_ACCESS_TOKEN=    # Dashboard 生成
MIXPOST_WORKSPACE_UUID=  # Social Accounts 页面获取
```

### 特点
- **自托管**: 数据完全自控，适合对隐私要求高的场景
- **审批工作流**: 团队协作，帖子需审批后发布

---

## metricool — 社媒排程管理 (Metricool API)

- **Slug**: `willscott-v2/metricool`
- **安装**: `clawhub install willscott-v2/metricool`
- **安全**: Benign(VT) / Suspicious(OC) — 建议审查源码
- **版本**: 1.0.0 | **下载**: 2,331

### 覆盖平台
LinkedIn, X, Bluesky, Threads, Instagram, Facebook — **共 6 个平台**

### 功能
- 单次 API 调用多平台排程
- 逐平台文本自定义
- 指定日期/时间和时区排程
- 查询近期排程帖子列表
- 获取各平台最佳发帖时间
- 附加图片到排程帖子

### 配置
```bash
METRICOOL_USER_TOKEN=    # Metricool 账户 token
METRICOOL_USER_ID=       # Metricool 用户 ID
```

### 使用场景
- 跨网络产品公告排程
- 查看未来一周排程帖子
- 查询特定平台最佳发帖时间
- 单次操作发布平台定制消息

---

## 三者对比

| 特性 | adaptlypost | mixpost | metricool |
|------|-------------|---------|-----------|
| 部署方式 | SaaS | 自托管 | SaaS |
| 平台数 | 9 | 8 | 6 |
| 数据控制 | 第三方 | 完全自控 | 第三方 |
| 审批工作流 | - | 有 | - |
| 最佳发帖时间 | - | - | 有 |
| 安全评级 | Benign | Benign | Suspicious(OC) |
| 适合场景 | 快速多平台发布 | 隐私优先/团队协作 | 数据驱动排程 |
