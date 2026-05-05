import React from 'react';
import { Card, Row, Col, Statistic, Table, Tag } from 'antd';
import { ApiOutlined, ClockCircleOutlined, WarningOutlined } from '@ant-design/icons';

const Observability: React.FC = () => {
  const traces = [
    { key: '1', traceId: 'trace-a1b2c3', duration: '1.2s', tokens: 850, status: 'ok', time: '10秒前' },
    { key: '2', traceId: 'trace-d4e5f6', duration: '3.5s', tokens: 2100, status: 'slow', time: '30秒前' },
    { key: '3', traceId: 'trace-g7h8i9', duration: '0.8s', tokens: 420, status: 'ok', time: '1分钟前' },
  ];

  return (
    <div>
      <Row gutter={[16, 16]}>
        <Col span={6}>
          <Card><Statistic title="请求量(24h)" value={45280} prefix={<ApiOutlined />} /></Card>
        </Col>
        <Col span={6}>
          <Card><Statistic title="平均延迟" value={1.2} suffix="s" prefix={<ClockCircleOutlined />} /></Card>
        </Col>
        <Col span={6}>
          <Card><Statistic title="错误率" value={0.3} suffix="%" valueStyle={{ color: '#cf1322' }} prefix={<WarningOutlined />} /></Card>
        </Col>
        <Col span={6}>
          <Card><Statistic title="Token消耗(24h)" value={1280000} /></Card>
        </Col>
      </Row>
      <Card title="链路追踪" style={{ marginTop: 16 }}>
        <Table
          dataSource={traces}
          pagination={false}
          columns={[
            { title: 'Trace ID', dataIndex: 'traceId' },
            { title: '耗时', dataIndex: 'duration' },
            { title: 'Token数', dataIndex: 'tokens' },
            { title: '状态', dataIndex: 'status', render: (s) => <Tag color={s === 'ok' ? 'green' : 'orange'}>{s}</Tag> },
            { title: '时间', dataIndex: 'time' },
          ]}
        />
      </Card>
    </div>
  );
};

export default Observability;