# 内容规划与日历 Skills

## social-media-content-calendar — 结构化社媒内容日历

- **Slug**: `seanwyngaard/social-media-content-calendar`
- **安装**: `clawhub install seanwyngaard/social-media-content-calendar`
- **安全**: Benign / Benign
- **版本**: 1.0.0 | **下载**: 1,100+

### 覆盖平台
Instagram, TikTok, LinkedIn, X(Twitter), Facebook, YouTube, Threads

### 功能
- 品牌调性和受众分析
- 平台特定格式化和最佳实践
- 按周生成帖子（含 CTA）
- 标签库（按流量分级: 高量/中量/利基）
- CSV 导出（兼容 Buffer / Hootsuite / Later）
- 内容类型配比优化
- 无障碍功能集成

### 输入参数
- 品牌描述
- 目标平台列表
- 时间范围（如 "4 weeks"）

### 输出结构
```
output/
├── week1.md          # 周计划
├── week2.md
├── ...
├── calendar.csv      # 导入社媒管理工具
├── hashtag-library.md
└── image-prompts.md  # 汇总图片生成提示词
```

### 使用场景
- 新品牌社媒冷启动
- 月度内容规划
- 团队内容协作（CSV 导入现有工具）

---

## weekly-content-planner — 一周社媒内容快速生成

- **Slug**: `claudiodrusus/weekly-content-planner`
- **安装**: `clawhub install claudiodrusus/weekly-content-planner`
- **安全**: Benign / Benign
- **版本**: 1.0.0 | **下载**: 324

### 覆盖平台
X(Twitter), LinkedIn, Instagram

### 功能
- 一次生成 7 天 × 3 平台 = **21 条帖子草稿**
- 平台适配的标签策略
- 推荐发帖时间
- 每日主题分配（如周一激励、周二教程）
- 输出完整内容日历 Markdown

### 使用场景
- 产品发布内容日历
- 通讯主题排程
- 自由职业客户内容创作
- 周主题活动帖子
- 个人品牌社媒文案

---

## 两者对比

| 特性 | content-calendar | weekly-planner |
|------|-----------------|----------------|
| 平台数 | 7 | 3 |
| 时间跨度 | 可定义（多周） | 固定 1 周 |
| 标签策略 | 分级库 | 平台适配 |
| 输出格式 | MD + CSV + 图片提示 | MD |
| CSV 导出 | 有 | 无 |
| 适合场景 | 完整内容策略 | 快速周计划 |
