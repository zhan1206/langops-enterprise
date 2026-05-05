import React from 'react';
import { Card, Table, Button, Tag, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';

const Release: React.FC = () => {
  const data = [
    { key: '1', name: '客服Prompt v3.2上线', from: 'staging', to: 'production', canary: '20%', status: 'canary' },
    { key: '2', name: 'RAG配置 v2.1上线', from: 'development', to: 'staging', canary: '-', status: 'approved' },
    { key: '3', name: 'Agent工作流 v1.0', from: 'preproduction', to: 'production', canary: '100%', status: 'completed' },
  ];

  return (
    <Card title="发布管控" extra={<Button type="primary" icon={<PlusOutlined />}>创建发布</Button>}>
      <Table dataSource={data} pagination={false} columns={[
        { title: '发布名称', dataIndex: 'name' },
        { title: '源环境', dataIndex: 'from', render: (e) => <Tag>{e}</Tag> },
        { title: '目标环境', dataIndex: 'to', render: (e) => <Tag color="blue">{e}</Tag> },
        { title: '灰度比例', dataIndex: 'canary' },
        { title: '状态', dataIndex: 'status', render: (s) => {
          const colors: any = { canary: 'orange', approved: 'green', completed: 'blue' };
          return <Tag color={colors[s] || 'default'}>{s}</Tag>;
        }},
        { title: '操作', render: () => <Space><a>审批</a><a>灰度</a><a>回滚</a></Space> },
      ]} />
    </Card>
  );
};

export default Release;