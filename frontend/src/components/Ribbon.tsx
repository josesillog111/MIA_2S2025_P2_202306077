// src/components/Ribbon.tsx
import React from 'react';
import RibbonItem from './RibbonItem';
import './Ribbon.css';

interface RibbonProps {
  onNewFile: () => void;
  onSaveFile: () => void;
  onOpenFolder: () => void;
  onUndo: () => void;
  onRedo: () => void;
  onCut: () => void;
  onCopy: () => void;
  onPaste: () => void;
  onFileExplorer: () => void;
}

const Ribbon: React.FC<RibbonProps> = ({ 
  onNewFile, 
  onSaveFile, 
  onOpenFolder, 
  onUndo, 
  onRedo, 
  onCut, 
  onCopy, 
  onPaste,
  onFileExplorer
}) => {
  return (
    <nav className="ribbon">
      <ul className="ribbon-menu">
        <RibbonItem label="File">
          <li className="ribbon-menu-item" onClick={onNewFile}>New File</li>
          <li className="ribbon-menu-item" onClick={onOpenFolder}>Open Folder...</li>
          <li className="ribbon-menu-item" onClick={onSaveFile}>Save</li>
          <li className="ribbon-menu-item separator"></li>
          <li className="ribbon-menu-item" onClick={() => {}}>Exit</li>
        </RibbonItem>

        <RibbonItem label="Edit">
          <li className="ribbon-menu-item" onClick={onUndo}>Undo</li>
          <li className="ribbon-menu-item" onClick={onRedo}>Redo</li>
          <li className="ribbon-menu-item separator"></li>
          <li className="ribbon-menu-item" onClick={onCut}>Cut</li>
          <li className="ribbon-menu-item" onClick={onCopy}>Copy</li>
          <li className="ribbon-menu-item" onClick={onPaste}>Paste</li>
        </RibbonItem>

        <RibbonItem label="Disk">
          <li className="ribbon-menu-item" onClick={onFileExplorer}>Explorador de Archivos</li>
        </RibbonItem>
      </ul>
    </nav>
  );
};

export default Ribbon;