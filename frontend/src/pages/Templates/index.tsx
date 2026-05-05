import React from 'react';
import { Card, Row, Col, Button, Tag } from 'antd';
import { DownloadOutlined } from '@ant-design/icons';

const templates = [
  { id: '1', name: '客服问答模板', category: '客服', desc: '含Prompt/RAG/评测完整配置', downloads: 1280 },
  { id: '2', name: '代码辅助模板', category: '研发', desc: '含代码生成/审查Prompt', downloads: 890 },
  { id: '3', name: '合同审核模板', category: '法务', desc: '含合规检测/脱敏规则', downloads: 560 },
  { id: '4', name: '知识库问答模板', category: '通用', desc: '含RAG检索/事实性评测', downloads: 2100 },
  { id: '5', name: '内容创作模板', category: '营销', desc: '含合规/创意性评测', downloads: 430 },
  { id: '6', name: '数据分析模板', category: '数据', desc: '含SQL生成/结果校验', downloads: 720 },
];

const Templates: React.FC = () => {
  return (
    <Card title="模板市场">
      <Row gutter={[16, 16]}>
        {templates.map((t) => (
          <Col span={8} key={t.id}>
            <Card size="small" title={t.name} extra={<Tag>{t.category}</Tag>}>
              <p>{t.desc}</p>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <span>下载: {t.downloads}</span>
                <Button size="small" type="primary" icon={<DownloadOutlined />}>使用</Button>
              </div>
            </Card>
          </Col>
        ))}
      </Row>
    </Card>
  );
};

export default Templates;