# X/Twitter Skills

## bird — X/Twitter CLI

- **Slug**: `steipete/bird`
- **安装**: `clawhub install steipete/bird`
- **安全**: Benign / Benign
- **版本**: 1.0.0

### 功能
- 搜索 X 上的推文（按关键词）
- 阅读完整推文串
- 直接从终端发推和回复
- 查看认证状态 (`bird check`)
- 查看自己的 X 资料 (`bird whoami`)
- 无需 OAuth 注册

### 认证方式
1. **浏览器 Cookie** — 从 Firefox/Chrome 提取已有登录态
2. **Sweetistics API** — 通过 API Key（需设 `SWEETISTICS_API_KEY`）

### 用法示例
```bash
bird check                        # 验证认证来源
bird search "Claude MCP" -n 5    # 搜索最近 5 条推文
```

---

## readx — X/Twitter 数据分析工具包

- **Slug**: `wxtsky/readx`
- **安装**: `clawhub install wxtsky/readx`
- **安全**: Suspicious(VT) / Benign(OC) — 建议审查源码
- **版本**: 1.1.4

### 功能
- 分析 X 账号和推文表现
- 计算互动率、病毒传播比、争议评分、粉丝质量指标
- 自动计算衍生指标（非原始数据）
- 支持分页覆盖完整数据
- 可通过 MCP 工具或 curl API 调用

### 配置
- **环境变量**: `READX_API_KEY`
- **账户**: readx.cc 活跃账户

### 典型工作流
1. 解析用户名到 user ID
2. 获取并分页推文时间线
3. 逐推计算指标
4. 按互动排名
5. 输出 Markdown 表格 + 关键发现

### 使用场景
- 竞品账号互动率对标
- 按语言/互动阈值筛选爆款推文
- 粉丝权威性和质量对比
- 通过回复/点赞比检测负面反馈
- 关键词社区监控

---

## bird-dms — X/Twitter 私信读取

- **Slug**: `tolibear/bird-dms`
- **安装**: `clawhub install tolibear/bird-dms`
- **安全**: Benign / Benign
- **版本**: 1.0.0

### 功能
- 列出 DM 对话及最新消息预览
- 按对话 ID 获取完整消息串
- 复用 bird 的浏览器 cookie，无需额外 X API 凭证
- 导出 DM 为 JSON 存档
- 监控群聊新活动

### 前置条件
- 已安装并认证 `bird` skill

### 用法
```bash
bird-dm inbox                   # 列出最近对话
bird-dm read <conv-id>          # 读取特定对话
```

---

## kiro-x-publisher — X 热点发推自动化

- **Slug**: `vmining/kiro-x-publisher`
- **安装**: `clawhub install vmining/kiro-x-publisher`
- **安全**: Suspicious(VT) / Benign(OC) — 建议审查源码
- **版本**: 1.1.0

### 功能
- 拉取 X 热门推文
- 通过 FxTwitter 富化候选推文（完整文本 + 文章数据）
- 按互动和时效评分排名
- 生成单条推文草稿
- 可选自动发布到 X（OAuth 1.0a）

### 配置（5 个环境变量）
```bash
X_BEARER_TOKEN=          # 搜索用
X_API_KEY=               # 发帖用
X_API_SECRET=            # 发帖用
X_ACCESS_TOKEN=          # 发帖用
X_ACCESS_TOKEN_SECRET=   # 发帖用
```
**要求**: X Developer 账号，Elevated 或更高 API 权限

### 使用场景
- 实时 AI 话题推文起草
- 定时趋势感知自动发帖
- 每日 X 信号简报
- 批量竞品/关键词监控
- 内容管线自动化
