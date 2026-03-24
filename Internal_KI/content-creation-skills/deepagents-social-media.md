# deepagents-social-media — AI 驱动社媒帖子 + 自动配图

- **来源**: `langchain-ai/deepagents@social-media`
- **安装**: `npx skills add langchain-ai/deepagents@social-media`
- **GitHub Stars**: 17.1K | **周安装量**: 473
- **安全**: 全部通过 (Gen Agent Trust Hub / Socket / Snyk)

---

## 核心工作流

```
调研（强制）→ 撰写帖子 → 生成配图 → 输出到目录
```

### 1. 调研前置（必须）

写任何社媒内容前 **必须** 先完成话题调研。

**deepagents 框架内调用:**
```python
task(
    subagent_type="researcher",
    description="Research [话题]. Save findings to research/[slug].md"
)
```

**独立运行替代方案:**
- 使用 WebSearch / WebFetch 搜集话题资料
- 或由用户直接提供素材/briefing
- 将调研结果保存到 `research/[slug].md`，后续步骤读取

调研完成后，**先读取 findings 文件**，再开始写作。

### 2. 输出结构（必须）

**每篇帖子必须包含文案 + 配图，缺一不可。**

```
linkedin/<slug>/
├── post.md       # 帖子文案
└── image.png     # 必须: 生成的配图

tweets/<slug>/
├── thread.md     # Thread 文案
└── image.png     # 必须: 生成的配图
```

### 3. 图片生成

**deepagents 框架内调用:**
```python
generate_social_image(
    prompt="详细描述...",
    platform="linkedin",
    slug="your-post-slug"
)
```

**独立运行替代方案:**
- 使用任意图片生成工具（DALL-E / Midjourney / 本地 SD）
- 或标注 `[需配图: 描述]` 占位，由人工补充
- 配图保存到 `<platform>/<slug>/image.png`

---

## 平台指南

### LinkedIn

**格式:**
- 1,300 字符限制（~210 字符后折叠）
- 首行至关重要 — 必须是 Hook
- 用空行增加可读性
- 末尾 3-5 个 hashtag

**结构模板:**
```
[Hook - 1 行引人注目]

[背景 - 为什么这很重要]

[主要洞察 - 2-3 短段]

[CTA 或提问]

#hashtag1 #hashtag2 #hashtag3
```

**语气:** 专业但有个人温度，分享洞察和经验，用问题驱动互动

### X/Twitter

**格式:**
- 每条 280 字符
- Thread 格式: `1/🧵`
- 每条不超过 2 个 hashtag

**Thread 结构:**
```
1/🧵 [Hook - 核心洞察]
2/ [支撑论点 1]
3/ [支撑论点 2]
4/ [示例或证据]
5/ [结论 + CTA]
```

---

## 配图最佳实践

社媒配图需要在小尺寸、拥挤的信息流中脱颖而出:

- **大胆简洁** — 一个清晰焦点
- **高对比** — 滚动时能抓住眼球
- **不要嵌入文字** — 太小看不清，平台会自加
- **方形或 4:5** — 跨平台通用

### 配图提示词模板

**洞察/技巧帖:**
```
Single glowing lightbulb floating against a deep purple gradient background,
lightbulb made of interconnected golden geometric lines, rays of soft light
emanating outward. Minimal, striking, high contrast. Square composition.
```

**公告/新闻帖:**
```
Abstract rocket ship made of colorful geometric shapes launching upward with
a trail of particles. Bright coral and teal color scheme against clean white
background. Energetic, celebratory mood. Bold flat illustration style.
```

**思考型帖子:**
```
Two overlapping translucent circles, one blue one orange, creating a glowing
intersection in the center. Dark charcoal background, soft ethereal glow.
Minimalist and contemplative.
```

---

## 内容类型

| 类型 | 要点 |
|------|------|
| **公告** | 新闻先行 → 影响说明 → 链接/下一步 |
| **洞察** | 一个具体发现 → 简要背景 → 可操作建议 |
| **提问** | 真诚提问 → 先给出你的观点 → 聚焦一个话题 |

---

## 质量清单

- [ ] 帖子已保存到正确路径
- [ ] 配图已生成并放在帖子旁
- [ ] 首行能抓住注意力
- [ ] 内容符合平台字数限制
- [ ] 语气匹配平台规范
- [ ] 有明确 CTA 或提问
- [ ] Hashtag 相关（非通用）

---

## 融合 creative-copywriting: LinkedIn/X Hook 强化

### LinkedIn 首行 Hook 公式

```
[情绪触发词] + [具体成果] + [时间框架]
"The hidden technique that 10x'd my reach in 7 days"
```

**高效 Hook 类型（LinkedIn 最佳）:**
- 个人故事 Hook（8.7% 停留率）: "I failed at..." 脆弱性开头
- 数据冲击（8.1%）: 惊人统计数字开场
- 好奇缺口（9.8%）: 不完整信息引发继续阅读

### X/Twitter Thread Hook 公式

```
[Hook] + [承诺] + [具体性]
"I analyzed 1,000 viral tweets. Here's the pattern nobody talks about: 🧵"
```

**Read More 触发技巧:**
- 数字序列: "7 things I learned..."（大脑想补全）
- 不完整列表: "Here's what changed: 1."
- 故事开头: "Last week something happened..."
- 反常识: "Everyone says X. They're wrong."

### 情绪触发词快速参考

| 场景 | 推荐触发词 |
|------|-----------|
| 新功能公告 | Discover, Finally, Unlock, Transform |
| 经验分享 | Proven, Real, Authentic, Tested |
| 紧急通知 | Now, Limited, Warning, Before |
| 思考型帖子 | Truth, Hidden, Secret, Reveal |
