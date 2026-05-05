import React from 'react';
import { Card, Row, Col, Statistic, Table, Tag } from 'antd';
import { SafetyOutlined, AuditOutlined, LockOutlined } from '@ant-design/icons';

const Security: React.FC = () => {
  const auditLogs = [
    { key: '1', event: 'Prompt修改', resource: '客服问答Prompt', actor: '张三', time: '5分钟前', hash: '0xa3f2...' },
    { key: '2', event: '版本发布', resource: 'RAG配置v2.1', actor: '李四', time: '1小时前', hash: '0xb7c1...' },
    { key: '3', event: '权限变更', resource: '开发团队角色', actor: '管理员', time: '3小时前', hash: '0xc9d8...' },
  ];

  return (
    <div>
      <Row gutter={[16, 16]}>
        <Col span={8}>
          <Card><Statistic title="安全事件(24h)" value={2} prefix={<SafetyOutlined />} valueStyle={{ color: '#faad14' }} /></Card>
        </Col>
        <Col span={8}>
          <Card><Statistic title="审计条目" value={1580} prefix={<AuditOutlined />} /></Card>
        </Col>
        <Col span={8}>
          <Card><Statistic title="数据脱敏规则" value={105} prefix={<LockOutlined />} /></Card>
        </Col>
      </Row>
      <Card title="审计日志" style={{ marginTop: 16 }}>
        <Table
          dataSource={auditLogs}
          pagination={false}
          columns={[
            { title: '事件', dataIndex: 'event', render: (e) => <Tag>{e}</Tag> },
            { title: '资源', dataIndex: 'resource' },
            { title: '操作人', dataIndex: 'actor' },
            { title: '时间', dataIndex: 'time' },
            { title: '哈希', dataIndex: 'hash' },
          ]}
        />
      </Card>
    </div>
  );
};

export default Security;