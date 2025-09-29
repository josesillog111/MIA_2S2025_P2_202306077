// src/components/StatusBar.tsx
import React from 'react';

const StatusBar: React.FC = () => {
  return (
    <div className="status-bar">
      <div className="status-item">
        <span className="icon"></span>
        <span>Go Disk</span>
      </div>
      <div className="status-item">
        <span className="icon"></span>
        <span>USAC</span>
      </div>
      <div className="status-item">
        <span className="icon"></span>
        <span>UTF-8</span>
      </div>
    </div>
  );
};

export default StatusBar;