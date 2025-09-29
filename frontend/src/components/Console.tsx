// src/components/Console.tsx
import React from 'react';
import './Console.css';

interface ConsoleProps {
  output: string;
}

const Console: React.FC<ConsoleProps> = ({ output }) => {
  return (
    <div className="console-container">
      <div className="console-header">OUTPUT</div>
      <pre className="console-output">{output}</pre>
    </div>
  );
};

export default Console;