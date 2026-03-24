# epicenter-social-media — 技术向社媒文案

- **来源**: `epicenterhq/epicenter@social-media`
- **安装**: `npx skills add epicenterhq/epicenter@social-media`
- **GitHub Stars**: 4.3K | **周安装量**: 65
- **安全**: 全部通过 (Gen Agent Trust Hub / Socket / Snyk)

---

## 核心原则

> 技术细节 > 营销话术。写开发者真正关心的内容。

### 适用场景
- LinkedIn/X/Reddit 上发布产品或技术更新
- 发布功能公告（无 hashtag、无炒作用语）
- 分享实现细节（开发者真正在意的）
- 把 AI 生成的文案改写成简洁、平台原生的写法
- 写短 Thread，每条帖子都有独立的实质内容

---

## 平台简洁度标准

**LinkedIn**: 3-5 行。说明功能 + 附链接，完毕。

**X/Twitter**: 每条推文 **一个想法**。不要过度解释。

**Reddit 技术子版块**: 聚焦实现细节，不讲好处。

---

## 必须删除的内容

- 所有 hashtag（除非平台文化确实需要）
- 帖子中的 section 标题（"## Implementation"）
- 功能/好处的列表
- 营销用语: "game-changing" / "seamless" / "powerful"
- CTA 套话: "See it in action!" / "Try it today!"
- 冗余形容词: "excellent" / "really"

## 必须添加的内容

- 开发者关心的 **具体技术细节**
- 实际的 **实现挑战和解决方案**
- 相关 **库/API 的链接**
- 一个独特功能细节（"with your model of choice"）
- 推荐工具时的 **免责声明**（"Not affiliated, it just..."）
- 个人标准/观点（"by my standards"）

---

## 示例

### X/Twitter — 好的功能公告
```
Whispering now does file uploads. Drag audio/video files in,
get transcription out. Works with any OpenAI-compatible API.
Open source: github.com/EpicenterHQ/epicenter
```

### X/Twitter — 好的技术洞察
```
TIL: Tauri's onDragDropEvent works where web drag-and-drop doesn't.
If you're building a desktop app with web tech, native drag-and-drop
is ~20 lines and handles file paths the OS way.
```

### LinkedIn — 好的公告
```
Whispering now supports direct file uploads! Simply drag and drop
(or click to browse) your audio files for instant transcription,
with your model of choice. Free open-source app:
https://github.com/EpicenterHQ/epicenter
```

### Reddit — 好的实现分享
```
包含: 具体组件名、代码示例、实际技术挑战、
所用库链接、对代码质量的个人评价。
推荐工具时说明非关联关系。
```

### Thread 规则 (2-4 条推文最多)
每条推文必须 **独立可理解**，包含具体事实而非 Hook 或铺垫。

---

## 融合 creative-copywriting: 技术内容的 Hook 优化

虽然 epicenter 风格反对营销式 Hook，但好的技术 Hook 仍然能提升停留率:

### 技术内容适用的 Hook 类型

| Hook 类型 | 技术适配 | 示例 |
|-----------|---------|------|
| 数据冲击 | 性能数据、基准测试 | "20 lines of code replaced our entire drag-and-drop stack" |
| 好奇缺口 | 意外发现 | "TIL: web drag-and-drop is broken in Tauri. Here's what works." |
| 反常识 | 技术迷思打破 | "Everyone uses feature flags for this. You don't need to." |

### 技术内容不适用的 Hook

- 个人故事 Hook（"I failed at..."）— 在技术 Reddit 会被踩
- 紧迫感触发词（"Limited" / "Now"）— 开源没有稀缺性
- 过度情绪化触发词 — 开发者社区偏好理性

### 技术帖子的心理学原则

| 原则 | 技术场景用法 |
|------|-------------|
| 具体性 | 精确数字（"~20 lines"、"3x faster"） |
| 权威 | 展示你真的写了代码（链接 repo） |
| 互惠 | 开源、免费工具、完整代码示例 |
| 好奇缺口 | TIL 格式，意外的技术发现 |
