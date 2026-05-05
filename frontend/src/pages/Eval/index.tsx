import React from 'react';
import { Card, Table, Button, Tag, Space, Progress } from 'antd';
import { PlusOutlined, PlayCircleOutlined } from '@ant-design/icons';

const Eval: React.FC = () => {
  const data = [
    { key: '1', name: '客服问答准确率评测', dimensions: 6, score: 0.87, passRate: 0.83, status: 'completed' },
    { key: '2', name: 'RAG检索相关性评测', dimensions: 5, score: 0.92, passRate: 0.90, status: 'completed' },
    { key: '3', name: '代码生成安全性评测', dimensions: 4, score: 0.78, passRate: 0.65, status: 'failed' },
  ];

  return (
    <Card title="评测中心" extra={<Button type="primary" icon={<PlusOutlined />}>新建评测</Button>}>
      <Table dataSource={data} pagination={false} columns={[
        { title: '评测名称', dataIndex: 'name' },
        { title: '评测维度', dataIndex: 'dimensions' },
        { title: '综合评分', dataIndex: 'score', render: (s) => <Progress percent={Math.round(s * 100)} size="small" /> },
        { title: '通过率', dataIndex: 'passRate', render: (r) => ${(r * 100).toFixed(0)}% },
        { title: '状态', dataIndex: 'status', render: (s) => <Tag color={s === 'completed' ? 'green' : s === 'failed' ? 'red' : 'blue'}>{s}</Tag> },
        { title: '操作', render: () => <Space><a>报告</a><a>回归测试</a><PlayCircleOutlined style={{ cursor: 'pointer' }} /></Space> },
      ]} />
    </Card>
  );
};

export default Eval;