# social-media-posts — 多平台帖子生成器

- **来源**: `jezweb/claude-skills@social-media-posts`
- **安装**: `npx skills add jezweb/claude-skills@social-media-posts`
- **GitHub Stars**: 643 | **周安装量**: 110
- **安全**: Gen Agent Trust Hub (Pass) / Socket (Pass) / Snyk (Warn)

---

## 三种运行模式

| 模式 | 输入 | 用途 |
|------|------|------|
| **From Scratch** | 话题 + 关键点 | 从零创建各平台帖子 |
| **From Content** | 现有内容（博客/通讯/公告） | 改写为各平台帖子 |
| **Campaign** | 话题/产品 | 协调多平台帖子 + 排程时间线 |

### 输入参数
- 话题/源内容（必须）
- 目标平台（必须）
- 语气（可选）
- CTA（可选）
- 链接（可选）
- 是否有配图（可选）

---

## 平台指南

### LinkedIn
- **受众**: 专业人士 / B2B
- **最佳长度**: 1,200–1,500 字符
- **Hook**: 前 2 行必须有效（~210 字符后折叠）
- **结构**: Hook → 正文段落（空行分隔）→ CTA 提问 → 3-5 hashtag
- **图片**: 1200×627px

### Facebook
- **受众**: 混合（朋友/家人/社区）
- **最佳长度**: 80–150 字符（互动最高）；故事型可到 500
- **链接帖**: Facebook 自动生成预览卡 — 不要重复标题
- **图片**: 1200×630px

### Instagram
- **受众**: 视觉优先、年轻群体
- **关键限制**: 前 125 字符必须传递核心信息
- **Hashtag**: 20-30 个，放首评
- **图片**: 1080×1080 (feed) / 1080×1350 (竖图) / 1080×1920 (Story)

### Reddit
- **受众**: 社区特定，重视真实性
- **核心**: **绝不能有推销感** — 价值先行
- **标题就是一切** — Reddit 成败看标题
- **语气**: 真诚、对话式、绝不商业化
- **无 hashtag、无过度 emoji**

---

## 质量标准（7 条铁律）

1. **绝不** 跨平台复制粘贴相同文字
2. **避免** 企业腔调和 jargon
3. **前置** 价值在开头行
4. **匹配** 平台文化
5. **包含** 互动引导
6. **遵守** 字数限制
7. **绝不** 使用虚假互动技巧

---

## 输出格式

```
═══ LINKEDIN ═══
[帖子内容]

═══ FACEBOOK ═══
[帖子内容]

═══ INSTAGRAM ═══
[帖子内容]

═══ REDDIT ═══
[帖子内容]
```

可选保存到 `.jez/artifacts/social-posts-[topic].md`

---

## Campaign 排程模板

| 时间 | 平台 | 内容类型 |
|------|------|---------|
| Day -1 | LinkedIn | 预告 |
| Day 0 (上午) | 全平台 | 正式公告 |
| Day 0 (下午) | Instagram | Story |
| Day +1 | Reddit | 讨论帖 |
| Day +3 | LinkedIn | 跟进 |
| Day +7 | Facebook | 用户评价 |

---

## 融合 creative-copywriting: 各平台 Hook 强化

### LinkedIn Hook 公式

**Hook-Story-Action 框架:**
```
HOOK (0-3 sec) → STORY (3-30 sec) → ACTION (最后)
    ↓               ↓                  ↓
停住滚动        构建情绪           驱动行为
```

**LinkedIn 最佳 Hook 模板:**
```
[情绪触发词] + [具体成果] + [时间框架]
例: "The hidden technique that 10x'd my reach in 7 days"
```

### Instagram 轮播叙事架构 (7 Slide)

```
Slide 1: HOOK — 停住滚动（大胆视觉 + 好奇）
Slide 2: SETUP — 建立问题/背景
Slide 3: TENSION — 激化（为什么这很重要）
Slide 4: BRIDGE — 暗示解决方案
Slide 5: VALUE — 核心洞察/转变
Slide 6: PROOF — 证据、示例或故事
Slide 7: CTA — 明确下一步
```

**幻灯片间滑动触发:**
- 句子中断: "But here's where it gets..."
- 视觉延续: 箭头、编号序列
- 模式打断: 变色、引入角色
- 悬念: "The result surprised everyone..."

**轮播类型 × 滑动率:**

| 类型 | 结构 | 滑动率 | 最佳用途 |
|------|------|--------|---------|
| 教育 | 问题→步骤→结果 | 65% | 权威建立 |
| 故事 | 设定→冲突→解决 | 72% | 情感连接 |
| 排行/列表 | 引言→条目→总结 | 58% | 收藏/分享 |
| 转变 | 前→过程→后 | **78%** | 证明/激励 |
| 破迷思 | 迷思→真相→证据 | 61% | 互动/分享 |

### Reddit 标题 Hook

Reddit 不适用营销式 Hook，但好标题仍需:
- **具体性**: 精确数字、确切时间
- **好奇缺口**: 不完整但不 clickbait
- **真实性**: 第一人称、承认局限

**好标题示例:**
```
"I built a tool that does X in 20 lines. Here's the approach and where it breaks."
```

### 各平台情绪触发词速查

| 平台 | 适合触发词 | 避免触发词 |
|------|-----------|-----------|
| LinkedIn | Proven, Transform, Discover, Real | Warning, Sick of, Fed up |
| Instagram | Ultimate, Dream, Secret, Hidden | Risk, Mistake, Don't |
| Facebook | Finally, Truth, Authentic | Limited, Ending |
| Reddit | (不用触发词) | 全部营销触发词 |
