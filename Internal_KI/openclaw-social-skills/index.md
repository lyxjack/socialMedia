# OpenClaw 社媒 Skills 索引

> 来源: [VoltAgent/awesome-openclaw-skills](https://github.com/VoltAgent/awesome-openclaw-skills)
> 抓取日期: 2026-03-24
> 安装方式: `clawhub install <slug>`

## 工作流调度映射

> **给调度器用**: 本目录的 skill 负责"发布/执行"环节，与 `content-creation-skills/` 的"文案生成"环节配合。

| 环节 | 能力 | 切片文件 | 输入 | 输出 |
|------|------|---------|------|------|
| 发帖到 X/Twitter | 搜索/发帖/DM/热点 | [platforms/x-twitter.md](platforms/x-twitter.md) | 文案文本 | 已发布帖子 |
| 发帖到 Bluesky | 发帖/互动/管理 | [platforms/bluesky.md](platforms/bluesky.md) | 文案文本 | 已发布帖子 |
| 发帖到 Instagram | Feed/Story/Reels/Carousel | [platforms/instagram.md](platforms/instagram.md) | 文案 + 图片URL | 已发布帖子 |
| YouTube 分析 | 转录/元数据/情绪 | [platforms/youtube.md](platforms/youtube.md) | 视频 URL | 分析报告 |
| 发帖到 Telegram | 频道发布/排程 | [platforms/telegram.md](platforms/telegram.md) | 文案 + 媒体 | 已发布/排程消息 |
| Reddit 调研 | 子版块匹配/规则 | [platforms/reddit.md](platforms/reddit.md) | 产品/话题 | subreddit 推荐 |
| 多平台排程 | 统一发布/管理 | [management/multi-platform.md](management/multi-platform.md) | 文案 + 平台列表 | 跨平台排程 |
| 内容日历 | 周/月规划 | [management/content-planning.md](management/content-planning.md) | 品牌/话题 | 日历 MD + CSV |
| 竞品情报 | 跨平台数据 | [management/analytics.md](management/analytics.md) | 竞品名/URL | 分析报告 |

### 与 content-creation-skills 的协作

```
content-creation-skills (文案生成)
    ↓ 输出: 各平台文案
openclaw-social-skills (发布执行)
    ↓ 输出: 已发布帖子 / 排程任务
```

每个切片文件 **自包含**，可独立加载使用，无需读取本索引或其他切片。

## 快速查找

| 需求场景 | 推荐 Skill | 切片文件 |
|----------|-----------|---------|
| X/Twitter 发帖/搜索 | bird | [platforms/x-twitter.md](platforms/x-twitter.md) |
| X/Twitter 数据分析 | readx | [platforms/x-twitter.md](platforms/x-twitter.md) |
| X/Twitter DM | bird-dms | [platforms/x-twitter.md](platforms/x-twitter.md) |
| X/Twitter 热点发推 | kiro-x-publisher | [platforms/x-twitter.md](platforms/x-twitter.md) |
| Bluesky 全功能 | bluesky | [platforms/bluesky.md](platforms/bluesky.md) |
| Instagram/Threads 发帖 | instagram-api | [platforms/instagram.md](platforms/instagram.md) |
| YouTube 分析/转录 | youtube-pro | [platforms/youtube.md](platforms/youtube.md) |
| Telegram 发布/排程 | publora-telegram | [platforms/telegram.md](platforms/telegram.md) |
| 多平台统一排程 | adaptlypost | [management/multi-platform.md](management/multi-platform.md) |
| 自托管社媒管理 | mixpost | [management/multi-platform.md](management/multi-platform.md) |
| 社媒排程 (Metricool) | metricool | [management/multi-platform.md](management/multi-platform.md) |
| 内容日历生成 | social-media-content-calendar | [management/content-planning.md](management/content-planning.md) |
| 周内容规划 | weekly-content-planner | [management/content-planning.md](management/content-planning.md) |
| 竞品社媒情报 | apify-competitor-intelligence | [management/analytics.md](management/analytics.md) |
| Reddit 子版块侦察 | subreddit-scout | [platforms/reddit.md](platforms/reddit.md) |

## 安全概览

| Skill | VirusTotal | OpenClaw | 建议 |
|-------|-----------|---------|------|
| bird | Benign | Benign | 可用 |
| readx | Suspicious | Benign | 谨慎使用，审查源码 |
| bird-dms | Benign | Benign | 可用 |
| kiro-x-publisher | Suspicious | Benign | 谨慎使用 |
| bluesky | Benign | Benign | 可用 |
| instagram-api | Suspicious | Suspicious | 高风险，务必审查源码 |
| youtube-pro | Benign | Benign | 可用 |
| publora-telegram | Benign | Suspicious | 谨慎使用 |
| adaptlypost | Benign | Benign | 可用 |
| mixpost | Benign | Benign | 可用 |
| metricool | Benign | Suspicious | 谨慎使用 |
| social-media-content-calendar | Benign | Benign | 可用 |
| weekly-content-planner | Benign | Benign | 可用 |
| apify-competitor-intelligence | Benign | Benign | 可用 |
| subreddit-scout | - | - | 需进一步确认 |

## 目录结构

```
openclaw-social-skills/
├── index.md                          # 本文件 - 总索引
├── platforms/                        # 按平台切片
│   ├── x-twitter.md                  # X/Twitter 相关 (4 skills)
│   ├── bluesky.md                    # Bluesky (1 skill)
│   ├── instagram.md                  # Instagram/Threads (1 skill)
│   ├── youtube.md                    # YouTube (1 skill)
│   ├── telegram.md                   # Telegram (1 skill)
│   └── reddit.md                     # Reddit (1 skill)
└── management/                       # 跨平台管理工具
    ├── multi-platform.md             # 多平台统一管理 (3 skills)
    ├── content-planning.md           # 内容规划与日历 (2 skills)
    └── analytics.md                  # 竞品分析与情报 (1 skill)
```
