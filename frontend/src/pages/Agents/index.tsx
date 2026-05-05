import React from 'react';
import { Card, Table, Button, Tag, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';

const Agents: React.FC = () => {
  const data = [
    { key: '1', name: '客服助手Agent', framework: 'LangChain', model: 'gpt-4o', calls: 15420, status: 'running' },
    { key: '2', name: '代码审查Agent', framework: 'AutoGen', model: 'claude-3.5', calls: 3200, status: 'running' },
    { key: '3', name: '数据分析Agent', framework: 'CrewAI', model: 'qwen-max', calls: 890, status: 'stopped' },
  ];

  return (
    <Card title="Agent管理" extra={<Button type="primary" icon={<PlusOutlined />}>注册Agent</Button>}>
      <Table dataSource={data} pagination={false} columns={[
        { title: '名称', dataIndex: 'name' },
        { title: '框架', dataIndex: 'framework', render: (f) => <Tag>{f}</Tag> },
        { title: '模型', dataIndex: 'model' },
        { title: '调用次数', dataIndex: 'calls' },
        { title: '状态', dataIndex: 'status', render: (s) => <Tag color={s === 'running' ? 'green' : 'red'}>{s}</Tag> },
        { title: '操作', render: () => <Space><a>详情</a><a>版本</a><a>评测</a></Space> },
      ]} />
    </Card>
  );
};

export default Agents;