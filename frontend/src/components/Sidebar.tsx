// src/components/Sidebar.tsx
import React, { useState } from 'react';
import './Sidebar.css';
import RedCrossIcon from '../assets/red-cross.svg';

interface File {
  name: string;
  path: string;
  content: string;
}

interface SidebarProps {
  files: File[];
  onFileSelect: (file: File) => void;
  onCloseFile: (file: File) => void;
  onRunCode: () => void;
  onSaveFile: (file: File) => Promise<void>;
  activeFile: File | null;
}

const Sidebar: React.FC<SidebarProps> = ({ files, onFileSelect, onCloseFile, onRunCode, onSaveFile, activeFile }) => {
  const [showPopup, setShowPopup] = useState(false);
  const [fileToClose, setFileToClose] = useState<File | null>(null);

  const handleCloseClick = (file: File, event: React.MouseEvent) => {
    event.stopPropagation(); // Evita que el evento se propague al `li`
    
    // Aquí puedes añadir tu lógica de "archivo modificado".
    // Por ahora, para la demostración, siempre abrimos el popup.
    const isModified = true; // Simula que siempre está modificado para probar el popup
    
    if (isModified) {
      setFileToClose(file);
      setShowPopup(true);
    } else {
      onCloseFile(file);
    }
  };

  const handleConfirmClose = () => {
    if (fileToClose) {
      onCloseFile(fileToClose);
      setShowPopup(false);
      setFileToClose(null);
    }
  };

  const handleSaveAndClose = async () => {
    if (fileToClose) {
      try {
        await onSaveFile(fileToClose);
        onCloseFile(fileToClose);
      } catch (err) {
        alert(`Error al guardar: ${err instanceof Error ? err.message : String(err)}`);
      }
      setShowPopup(false);
      setFileToClose(null);
    }
  };


  const handleCancelClose = () => {
    setShowPopup(false);
    setFileToClose(null);
  };

  return (
    <div className="sidebar">
      <div className="sidebar-header">EXPLORER</div>
      
      <div className="sidebar-actions">
        <button onClick={onRunCode} className="run-button-sidebar">
          Run Code
        </button>
      </div>

      <ul className="file-list">
        {files.map((file) => (
          <li
            key={file.path}
            className={`file-item ${activeFile?.path === file.path ? 'active' : ''}`}
            onClick={() => onFileSelect(file)}
          >
            <span className="file-name">📄 {file.name}</span>
            <button 
              className="close-file-button" 
              onClick={(e) => handleCloseClick(file, e)}
              title="Close File"
            >
              x
            </button>
          </li>
        ))}
      </ul>

      {showPopup && (
        <div className="popup-overlay">
          <div className="popup-content">
            <div className="popup-header">
              <img src={RedCrossIcon} alt="Cerrar" className="popup-icon" />
              <h3 className="popup-title">CONFIRMACIÓN DE CIERRE</h3>
              <button className="popup-close-button" onClick={handleCancelClose}>✕</button>
            </div>
            <p className="popup-message">
              Tienes cambios sin guardar. ¿Quieres cerrar el archivo sin guardar?
            </p>
            <div className="popup-buttons">
              <button className="popup-confirm-button" onClick={handleSaveAndClose}>Guardar y cerrar</button>
              <button className="popup-confirm-button" onClick={handleConfirmClose}>Cerrar sin guardar</button>
              <button className="popup-cancel-button" onClick={handleCancelClose}>Cancelar</button>
            </div>
            <div className="popup-source">
              <span>Source: Editor de Código</span>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default Sidebar;