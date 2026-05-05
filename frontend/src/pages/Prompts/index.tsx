import React from 'react';
import { Card, Table, Button, Tag, Space } from 'antd';
import { PlusOutlined, RollbackOutlined } from '@ant-design/icons';

const Prompts: React.FC = () => {
  const data = [
    { key: '1', name: '客服问答Prompt', version: 'v3.2', branch: 'main', author: '张三', updated: '2026-05-05', status: 'production' },
    { key: '2', name: '代码生成Prompt', version: 'v2.1', branch: 'feature/optimize', author: '李四', updated: '2026-05-04', status: 'staging' },
    { key: '3', name: '文档摘要Prompt', version: 'v1.5', branch: 'main', author: '王五', updated: '2026-05-03', status: 'development' },
  ];

  return (
    <Card title="Prompt版本管理" extra={<Button type="primary" icon={<PlusOutlined />}>新建Prompt</Button>}>
      <Table dataSource={data} pagination={false} columns={[
        { title: '名称', dataIndex: 'name' },
        { title: '版本', dataIndex: 'version' },
        { title: '分支', dataIndex: 'branch' },
        { title: '作者', dataIndex: 'author' },
        { title: '更新时间', dataIndex: 'updated' },
        { title: '状态', dataIndex: 'status', render: (s) => <Tag color={s === 'production' ? 'green' : s === 'staging' ? 'blue' : 'default'}>{s}</Tag> },
        { title: '操作', render: () => <Space><a>编辑</a><a>回滚</a></Space> },
      ]} />
    </Card>
  );
};

export default Prompts;