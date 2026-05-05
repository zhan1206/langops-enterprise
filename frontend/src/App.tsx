import React from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { Layout, Menu } from 'antd';
import {
  DashboardOutlined,
  CodeOutlined,
  CloudOutlined,
  RobotOutlined,
  CheckCircleOutlined,
  RocketOutlined,
  EyeOutlined,
  SafetyOutlined,
  AppstoreOutlined,
  SettingOutlined,
} from '@ant-design/icons';
import Dashboard from './pages/Dashboard';
import Prompts from './pages/Prompts';
import RAG from './pages/RAG';
import Agents from './pages/Agents';
import Eval from './pages/Eval';
import Release from './pages/Release';
import Observability from './pages/Observability';
import Security from './pages/Security';
import Templates from './pages/Templates';
import Settings from './pages/Settings';
import './App.css';

const { Sider, Content, Header } = Layout;

const menuItems = [
  { key: '/dashboard', icon: <DashboardOutlined />, label: '运维大盘' },
  { key: '/prompts', icon: <CodeOutlined />, label: 'Prompt管理' },
  { key: '/rag', icon: <CloudOutlined />, label: 'RAG配置' },
  { key: '/agents', icon: <RobotOutlined />, label: 'Agent管理' },
  { key: '/eval', icon: <CheckCircleOutlined />, label: '评测中心' },
  { key: '/release', icon: <RocketOutlined />, label: '发布管控' },
  { key: '/observability', icon: <EyeOutlined />, label: '可观测' },
  { key: '/security', icon: <SafetyOutlined />, label: '安全合规' },
  { key: '/templates', icon: <AppstoreOutlined />, label: '模板市场' },
  { key: '/settings', icon: <SettingOutlined />, label: '系统设置' },
];

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Layout className="app-layout">
        <Sider className="app-sider" width={220}>
          <div className="logo">LangOps Enterprise</div>
          <Menu
            theme="dark"
            mode="inline"
            items={menuItems}
            defaultSelectedKeys={['/dashboard']}
          />
        </Sider>
        <Layout>
          <Header style={{ background: '#fff', padding: '0 24px', borderBottom: '1px solid #f0f0f0' }}>
            <h2 style={{ margin: 0, lineHeight: '64px' }}>LangOps Enterprise - LLM应用全生命周期管控平台</h2>
          </Header>
          <Content className="app-content">
            <Routes>
              <Route path="/dashboard" element={<Dashboard />} />
              <Route path="/prompts" element={<Prompts />} />
              <Route path="/rag" element={<RAG />} />
              <Route path="/agents" element={<Agents />} />
              <Route path="/eval" element={<Eval />} />
              <Route path="/release" element={<Release />} />
              <Route path="/observability" element={<Observability />} />
              <Route path="/security" element={<Security />} />
              <Route path="/templates" element={<Templates />} />
              <Route path="/settings" element={<Settings />} />
              <Route path="/" element={<Navigate to="/dashboard" replace />} />
            </Routes>
          </Content>
        </Layout>
      </Layout>
    </BrowserRouter>
  );
};

export default App;