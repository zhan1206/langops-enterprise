import React from 'react';
import { Card, Table, Button, Tag, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';

const RAG: React.FC = () => {
  const data = [
    { key: '1', name: '产品知识库RAG', vectorDB: 'Milvus', chunks: 12580, version: 'v2.3', status: 'active' },
    { key: '2', name: '合同审查RAG', vectorDB: 'Pinecone', chunks: 3420, version: 'v1.1', status: 'active' },
    { key: '3', name: '技术文档RAG', vectorDB: 'Weaviate', chunks: 8900, version: 'v3.0', status: 'draft' },
  ];

  return (
    <Card title="RAG配置管理" extra={<Button type="primary" icon={<PlusOutlined />}>新建RAG配置</Button>}>
      <Table dataSource={data} pagination={false} columns={[
        { title: '名称', dataIndex: 'name' },
        { title: '向量库', dataIndex: 'vectorDB' },
        { title: '文档块数', dataIndex: 'chunks' },
        { title: '版本', dataIndex: 'version' },
        { title: '状态', dataIndex: 'status', render: (s) => <Tag color={s === 'active' ? 'green' : 'default'}>{s}</Tag> },
        { title: '操作', render: () => <Space><a>配置</a><a>评测</a><a>历史</a></Space> },
      ]} />
    </Card>
  );
};

export default RAG;