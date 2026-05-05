import React from 'react';
import { Card, Form, Input, Select, Button, Tabs, Switch, Divider, message } from 'antd';

const Settings: React.FC = () => {
  const handleSave = () => {
    message.success('设置已保存');
  };

  return (
    <Card title="系统设置">
      <Tabs items={[
        {
          key: 'general',
          label: '基本配置',
          children: (
            <Form layout="vertical" style={{ maxWidth: 600 }}>
              <Form.Item label="平台名称" initialValue="LangOps Enterprise">
                <Input />
              </Form.Item>
              <Form.Item label="默认评测模型" initialValue="gpt-4o">
                <Select options={[
                  { value: 'gpt-4o', label: 'GPT-4o' },
                  { value: 'claude-3.5', label: 'Claude 3.5' },
                  { value: 'qwen-max', label: 'Qwen Max' },
                ]} />
              </Form.Item>
              <Form.Item label="启用效果退化检测" valuePropName="checked" initialValue={true}>
                <Switch />
              </Form.Item>
              <Divider />
              <Button type="primary" onClick={handleSave}>保存</Button>
            </Form>
          ),
        },
        {
          key: 'notification',
          label: '告警通知',
          children: (
            <Form layout="vertical" style={{ maxWidth: 600 }}>
              <Form.Item label="告警渠道">
                <Select mode="multiple" options={[
                  { value: 'email', label: '邮件' },
                  { value: 'wechat', label: '企业微信' },
                  { value: 'dingtalk', label: '钉钉' },
                  { value: 'slack', label: 'Slack' },
                ]} />
              </Form.Item>
              <Form.Item label="效果退化阈值(%)" initialValue={10}>
                <Input type="number" />
              </Form.Item>
              <Divider />
              <Button type="primary" onClick={handleSave}>保存</Button>
            </Form>
          ),
        },
        {
          key: 'compliance',
          label: '合规设置',
          children: (
            <Form layout="vertical" style={{ maxWidth: 600 }}>
              <Form.Item label="合规标准">
                <Select mode="multiple" options={[
                  { value: 'djl_2_0', label: '等保2.0' },
                  { value: 'gdpr', label: 'GDPR' },
                  { value: 'sox', label: 'SOX' },
                  { value: 'data_security_law', label: '数据安全法' },
                ]} />
              </Form.Item>
              <Form.Item label="启用内容安全护栏" valuePropName="checked" initialValue={true}>
                <Switch />
              </Form.Item>
              <Divider />
              <Button type="primary" onClick={handleSave}>保存</Button>
            </Form>
          ),
        },
      ]} />
    </Card>
  );
};

export default Settings;