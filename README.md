# LangOps Enterprise

> 全球首个生产级、全生态兼容的开源LLM应用全生命周期Ops平台

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://go.dev/)
[![Python](https://img.shields.io/badge/Python-3.12+-3776AB?logo=python)](https://python.org/)
[![React](https://img.shields.io/badge/React-18-61DAFB?logo=react)](https://react.dev/)

## 🎯 项目定位

LangOps Enterprise 是闭源工具 LangSmith 的开源生产级替代品，覆盖企业级 Prompt/RAG/Agent 应用的**统一开发协作平台、效果评测中枢、发布管控底座、可观测与持续优化引擎**。

## 🏗️ 核心架构

`
┌─────────────────────────────────────────────────────────┐
│              前端交互与协作管控层                          │
│  可视化编排 · 版本管理 · 协作评审 · 模板市场 · 管控控制台  │
├─────────────────────────────────────────────────────────┤
│              安全合规与全链路审计层                        │
│  权限管控 · 不可篡改审计 · 内容安全 · 敏感数据防护          │
├─────────────────────────────────────────────────────────┤
│              全生命周期发布管控层                          │
│  环境隔离 · 灰度发布 · A/B测试 · 自动回滚 · 流量调度       │
├─────────────────────────────────────────────────────────┤
│            评测与持续优化核心层（项目心脏）                 │
│  自动化评测 · 回归测试 · 效果退化检测 · 根因分析与优化      │
├─────────────────────────────────────────────────────────┤
│              全生态兼容与可观测层                          │
│  多框架适配器 · 模型适配 · 全链路追踪 · 监控告警 · 成本管控 │
└─────────────────────────────────────────────────────────┘
`

## ✨ 核心能力

| 能力域 | 核心特性 |
|--------|---------|
| **开发协作** | 低代码可视化编排、全链路版本管理、团队协作评审、环境隔离 |
| **效果评测** | 10+维度自动化评测、回归测试、效果退化检测、根因分析 |
| **发布管控** | 灰度发布、A/B测试、自动回滚、智能流量调度 |
| **可观测** | OpenTelemetry全链路追踪、监控告警、成本精细化管控 |
| **安全合规** | 不可篡改审计、内容安全护栏、敏感数据防护、合规规则库 |
| **全生态兼容** | LangChain/LlamaIndex/AutoGen/CrewAI/OpenHands 全兼容 |

## 🚀 快速开始

### 前置条件

- Go 1.22+
- Python 3.12+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+

### Docker 部署

``bash
cd deploy/docker
docker-compose up -d
``

### 从源码构建

`ash
# 后端
cd backend && go build -o langops-server ./cmd/server

# 评测引擎
cd eval && pip install -r requirements.txt

# 前端
cd frontend && npm install && npm run build
``

## 📁 项目结构

`
langops-enterprise/
├── backend/               # Go 后端核心引擎
│   ├── cmd/server/        # 服务入口
│   └── internal/          # 核心模块
│       ├── core/          # 评测/回归/退化/根因/发布/AB/流量/协作
│       ├── adapter/       # 框架适配器 (LangChain/LlamaIndex/AutoGen/CrewAI/OpenHands)
│       ├── security/      # 安全模块 (认证/审计/护栏/脱敏/合规)
│       └── observability/ # 可观测 (追踪/指标/成本/告警)
├── eval/                  # Python 评测引擎
│   ├── engine/            # 评测核心引擎
│   ├── metrics/           # 评测指标库
│   └── report/            # 报告生成器
├── frontend/              # React + Ant Design Pro 前端
│   └── src/pages/         # 9大功能页面
├── configs/               # 配置文件
└── deploy/                # 部署配置 (Docker/K8s)
`

## 📜 开源协议

Apache License 2.0 — 永久开源免费、商用友好、无厂商锁定