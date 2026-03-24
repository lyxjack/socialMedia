# 社媒文案创作 Skills 索引

> 来源: skills.sh 开放生态
> 抓取日期: 2026-03-24
> 安装方式: `npx skills add <owner/repo@skill>`
> 注: creative-copywriting 的 Hook 公式、情绪触发词、轮播叙事、心理学技巧已融入各 skill

## 工作流调度映射

> **给调度器用**: 每个 skill 可独立运行。输入统一为话题/素材，输出为平台帖子文案。

| Skill ID | 切片文件 | 输入 | 输出 | 覆盖平台 | 外部依赖 |
|----------|---------|------|------|---------|---------|
| `content-creation` | [content-creation.md](content-creation.md) | 话题 + 目标渠道 | 结构化文案（博客/社媒/邮件/落地页） | LinkedIn, X, IG, FB | 无 |
| `deepagents` | [deepagents-social-media.md](deepagents-social-media.md) | 话题 | 帖子文案 + 配图 (png) | LinkedIn, X | `task(subagent_type="researcher")`, `generate_social_image()` — deepagents 框架专属，独立运行时替换为手动调研 + 图片工具 |
| `epicenter` | [epicenter-social-media.md](epicenter-social-media.md) | 话题/技术更新 + 链接 | 平台原生简洁文案 | LinkedIn, X, Reddit | 无 |
| `social-media-posts` | [social-media-posts.md](social-media-posts.md) | 话题/源内容 + 目标平台 | copy-paste-ready 帖子（按平台分块） | LinkedIn, FB, IG, Reddit | 无 |

### 独立运行说明

每个 skill 文件 **自包含全部指令**，无需读取本索引或其他 skill 文件即可运行。工作流中:
1. 调度器读取本索引，按需求选择 skill 或全部并行
2. 将话题/素材 + 对应 skill 文件内容注入 agent
3. agent 按 skill 文件中的模板/规则生成文案
4. 调度器收集各 agent 输出，进行对比评审

### deepagents 独立运行适配

`deepagents` 原始设计依赖其框架的 `researcher` 子 agent 和 `generate_social_image` 工具。独立运行时:
- **调研**: 替换为 WebSearch/WebFetch 或手动提供素材
- **配图**: 替换为任意图片生成工具，或标注"需配图"由人工补充
- 其余文案模板、平台指南、质量清单均可直接使用

## 快速查找

| 需求场景 | 推荐 Skill | 切片文件 |
|----------|-----------|---------|
| 完整营销内容体系（博客/邮件/社媒/落地页/SEO） | content-creation | [content-creation.md](content-creation.md) |
| AI 驱动多平台帖子 + 自动配图 | deepagents | [deepagents-social-media.md](deepagents-social-media.md) |
| 技术产品社媒宣传（反营销、重细节） | epicenter | [epicenter-social-media.md](epicenter-social-media.md) |
| 多平台 copy-paste-ready 帖子（含 Campaign） | social-media-posts | [social-media-posts.md](social-media-posts.md) |

## 概览对比

| 维度 | content-creation | deepagents | epicenter | social-media-posts |
|------|-----------------|------------|-----------|-------------------|
| GitHub Stars | 10.3K | 17.1K | 4.3K | 643 |
| 周安装量 | 950 | 473 | 65 | 110 |
| 覆盖平台 | LinkedIn/X/IG/FB | LinkedIn/X | LinkedIn/X/Reddit | LinkedIn/FB/IG/Reddit |
| 自动配图 | 无 | 有 | 无 | 无 |
| SEO 集成 | 强 | 无 | 无 | 无 |
| 博客/邮件/落地页 | 有 | 无 | 无 | 无 |
| Campaign 模式 | 无 | 无 | 无 | 有 |
| 技术向 | 一般 | 一般 | 强 | 一般 |
| 调研前置 | 无 | 强制 | 无 | 可选 |
| 外部依赖 | 无 | researcher + image_gen | 无 | 无 |

## 组合建议

- **全栈内容营销**: content-creation（主） + social-media-posts（Campaign 排程）
- **AI 自动化发帖**: deepagents（自动调研 + 配图）
- **开发者产品推广**: epicenter（LinkedIn/X/Reddit 技术向）
- **快速多平台分发**: social-media-posts（从现有内容改写）
- **对比评审**: 4 个 skill 并行生成 → 人工选最佳 → 通过 openclaw-social-skills 发布

## 目录结构

```
content-creation-skills/
├── index.md                       # 本文件（调度映射 + 总索引）
├── content-creation.md            # Anthropic 营销内容框架（自包含）
├── deepagents-social-media.md     # LangChain AI 社媒 + 配图（自包含）
├── epicenter-social-media.md      # 技术向社媒文案（自包含）
└── social-media-posts.md          # 多平台帖子生成器（自包含）
```
