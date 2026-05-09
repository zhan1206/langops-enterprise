# LangOps Enterprise：企业级LLM应用全生命周期管控与持续优化平台

[English](./README_en.md) | 中文

## 项目核心信息总览

| 项目属性 | 详情说明 |
|----------|----------|
| 项目全称 | LangOps Enterprise |
| 核心定位 | 全球首个生产级、全生态兼容的开源LLM应用全生命周期Ops平台，闭源工具LangSmith的开源生产级替代品 |
| 开源协议 | Apache 2.0（永久开源免费、商用友好、无厂商锁定） |
| 兼容生态 | LangChain、LlamaIndex、AutoGen、CrewAI、OpenHands等所有主流框架 |

---

## 一、行业背景与核心痛点

2026年企业级AI应用已从"demo验证期"全面进入"规模化生产落地期"，90%以上的企业AI应用通过Prompt工程、RAG检索增强、Agent编排实现业务闭环。现有开源生态的核心空白：**闭源商业工具LangSmith几乎垄断市场，但价格高昂、强绑定生态、不支持私有化部署**。

### 核心痛点与解决方案

| 痛点维度 | 解决方案 |
|----------|----------|
| 开发协作混乱 | 全链路版本管理、团队协作、在线评审、环境隔离 |
| 效果评测缺失 | 多维度自动化评测、回归测试、效果退化检测、根因分析 |
| 发布管控空白 | 灰度发布、A/B测试、自动回滚、流量调度 |
| 全链路黑盒 | OpenTelemetry全链路追踪、监控告警、成本管控 |
| 安全合规不达标 | 链式审计、内容安全护栏、敏感数据脱敏、合规规则库 |
| 厂商锁定严重 | 全生态兼容、零侵入接入、私有化部署 |

---

## 二、项目架构

```
┌─────────────────────────────────────────┐
│ 前端交互与协作管控层 │
│ 可视化编排、版本管理、协作评审、模板市场 │
├─────────────────────────────────────────┤
│ 安全合规与全链路审计层 │
│ 权限管控、审计链、内容安全、数据防护 │
├─────────────────────────────────────────┤
│ 全生命周期发布管控层 │
│ 环境隔离、灰度发布、A/B测试、自动回滚 │
├─────────────────────────────────────────┤
│ 评测与持续优化核心层 │
│ 自动化评测、回归测试、退化检测、根因分析 │
├─────────────────────────────────────────┤
│ 全生态兼容与可观测层 │
│ 多框架适配器、模型适配、链路追踪、监控 │
└─────────────────────────────────────────┘
```

---

## 二、核心技术栈

| 模块 | 技术选型 |
|------|---------|
| 后端核心 | Go + Gin |
| AI适配层 | Python |
| 前端 | React + TypeScript + Ant Design |
| 数据存储 | MySQL + Redis |
| 日志与追踪 | Elasticsearch + OpenTelemetry + Jaeger |
| 消息队列 | Kafka |
| 部署 | Docker + Kubernetes |

---

## 四、快速开始

### 前置要求

- Go 1.21+
- Python 3.10+
- MySQL 8.0+
- Redis 7.0+

### 本地开发

```bash
# 克隆项目
git clone https://github.com/zhan1206/langops-enterprise.git
cd langops-enterprise

# 启动后端
cd backend
cp config.example.yaml config.yaml
# 编辑 config.yaml 配置数据库等
go run cmd/server/main.go

# 启动前端
cd ../frontend
npm install
npm run dev
```

### Docker部署

```bash
docker-compose up -d
```

访问 http://localhost:3000

---

## 五、核心功能模块

### 1. 版本管理系统
- Prompt/RAG配置/Agent工作流全版本可追溯
- 分支管理、版本对比、一键回滚
- 变更记录完整追溯

### 2. 效果评测引擎
- 10+维度自动化量化评测
- 自动化回归测试
- 7×24效果退化检测
- 智能根因分析

### 3. 发布管控
- 开发/测试/预发/生产环境隔离
- 灰度发布、A/B测试
- 多条件自动回滚
- 智能流量调度

### 4. 可观测平台
- 全链路OpenTelemetry追踪
- 自定义监控告警
- 多维度成本统计

### 5. 安全合规
- 链式审计、不可篡改
- 内容安全护栏
- 敏感数据脱敏
- 合规规则库（等保2.0、GDPR等）

---

## 六、路线图

| 版本 | 周期 | 核心目标 |
|------|------|----------|
| v0.1 | 3个月 | MVP，核心框架适配、基础评测、监控 |
| v0.5 | 6个月 | Beta，完整团队协作、灰度发布、合规 |
| v1.0 | 12个月 | 正式版，企业级高可用、私有化部署 |
| v2.0 | 24个月 | 生态扩展、多租户、自学习能力 |

---

## 七、配套项目

本项目是**企业级AI应用落地技术体系**的核心组件之一，配套项目包括：

- [AIChat-Router](https://github.com/zhan1206/aigateway-universal) - 智能路由与全局调度引擎
- [AgentShield](https://github.com/zhan1206/agentshield-enterprise) - Agent安全沙箱与权限管控

---

## 八、协议与联系

- 开源协议: Apache 2.0
- 项目主页: https://github.com/zhan1206/langops-enterprise
- 问题反馈: GitHub Issues
- 贡献指南: CONTRIBUTING.md

---

## 九、Star历史

[![Star History Chart](https://api.star-history.com/svg?repos=zhan1206/langops-enterprise&type=Date)](https://star-history.com/#zhan1206/langops-enterprise&type=Date)