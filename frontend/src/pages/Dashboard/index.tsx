import React from 'react';
import { Row, Col, Card, Statistic, Table, Tag } from 'antd';
import { CheckCircleOutlined, RocketOutlined, AlertOutlined, DollarOutlined } from '@ant-design/icons';

const Dashboard: React.FC = () => {
  const recentAlerts = [
    { key: '1', resource: '客服问答Prompt', type: '效果退化', level: 'warning', time: '5分钟前' },
    { key: '2', resource: 'RAG检索配置', type: '版本发布', level: 'success', time: '1小时前' },
    { key: '3', resource: '代码辅助Agent', type: '评测完成', level: 'info', time: '2小时前' },
  ];

  return (
    <div>
      <Row gutter={[16, 16]}>
        <Col span={6}>
          <Card><Statistic title="活跃应用" value={128} prefix={<CheckCircleOutlined />} /></Card>
        </Col>
        <Col span={6}>
          <Card><Statistic title="今日评测" value={45} prefix={<RocketOutlined />} /></Card>
        </Col>
        <Col span={6}>
          <Card><Statistic title="告警数" value={3} prefix={<AlertOutlined />} valueStyle={{ color: '#cf1322' }} /></Card>
        </Col>
        <Col span={6}>
          <Card><Statistic title="本月成本" value={2840.5} prefix={<DollarOutlined />} suffix="元" /></Card>
        </Col>
      </Row>
      <Card title="最近告警" style={{ marginTop: 16 }}>
        <Table
          dataSource={recentAlerts}
          pagination={false}
          columns={[
            { title: '资源', dataIndex: 'resource' },
            { title: '类型', dataIndex: 'type' },
            { title: '级别', dataIndex: 'level', render: (v) => <Tag color={v === 'warning' ? 'orange' : v === 'success' ? 'green' : 'blue'}>{v}</Tag> },
            { title: '时间', dataIndex: 'time' },
          ]}
        />
      </Card>
    </div>
  );
};

export default Dashboard;