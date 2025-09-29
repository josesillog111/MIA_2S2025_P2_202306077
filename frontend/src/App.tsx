// src/App.tsx
import React, { useState, useRef } from 'react';
import Split from 'react-split'; 
import Sidebar from './components/Sidebar';
import Editor from './components/Editor';
import StatusBar from './components/StatusBar';
import Ribbon from './components/Ribbon';
import Console from './components/Console';
import './App.css';

import { saveFile, openFolder } from './services/fileService';
import type { EditorFile } from './services/fileService';

const initialFiles: EditorFile[] = [
  { name: 'index.ts', path: 'src/index.ts', content: 'console.log("Hello from a TypeScript file!");' },
  { name: 'App.tsx', path: 'src/App.tsx', content: '/* JSX code here */' },
  { name: 'README.md', path: 'README.md', content: '# My Code Editor' },
];

const App: React.FC = () => {
  const [files, setFiles] = useState<EditorFile[]>(initialFiles);
  const [activeFile, setActiveFile] = useState<EditorFile | null>(initialFiles[0]);
  const [consoleOutput, setConsoleOutput] = useState('');

  const editorRef = useRef<any>(null);

  const handleFileSelect = (file: EditorFile) => {
    setActiveFile(file);
    setConsoleOutput('');
  };

  const handleContentChange = (newContent: string) => {
    if (activeFile) {
      const updatedFile = { ...activeFile, content: newContent };
      setActiveFile(updatedFile);
      setFiles(files.map(f => f.path === activeFile.path ? updatedFile : f));
    }
  };

  const handleRunCode = async () => {
    if (!activeFile) return;

    setConsoleOutput('Executing code...');

    try {
      const API_URL = 'http://localhost:8080/execute';

      const response = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'text/plain' },
        body: activeFile.content, // aquí va el texto directo
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();

      if (result.success) {
        setConsoleOutput(result.output);
      } else {
        setConsoleOutput(result.error);
      }
    } catch (error) {
      setConsoleOutput(
        `Error de conexión: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  };


  const handleSaveFile = async (file: EditorFile) => {
    try {
      const updatedFile = await saveFile(file);
      setActiveFile(updatedFile);
      setFiles(files.map(f => f.path === file.path ? updatedFile : f));
      setConsoleOutput(`Archivo "${updatedFile.name}" guardado exitosamente.`);
    } catch (err) {
      setConsoleOutput(String(err));
    }
  };

  const handleOpenFolder = async () => {
    try {
      const newFiles = await openFolder();
      setFiles([...files, ...newFiles]);
      setConsoleOutput(`Carpeta cargada con ${newFiles.length} archivo(s).`);
    } catch (err) {
      setConsoleOutput(String(err));
    }
  };

  const handleCloseFile = (fileToClose: EditorFile) => {
    const updatedFiles = files.filter(file => file.path !== fileToClose.path);
    setFiles(updatedFiles);

    if (activeFile?.path === fileToClose.path) {
      setActiveFile(updatedFiles.length > 0 ? updatedFiles[0] : null);
    }
  };

  const handleUndo = () => editorRef.current?.undo();
  const handleRedo = () => editorRef.current?.redo();
  const handleCut = () => editorRef.current?.cut();
  const handleCopy = () => editorRef.current?.copy();
  const handlePaste = () => editorRef.current?.paste();

  return (
    <div className="vscode-container">
      <Ribbon 
        onNewFile={() => {
          const newFile: EditorFile = {
            name: `untitled-${files.length + 1}.txt`,
            path: `untitled-${files.length + 1}.txt`,
            content: ''
          };
          setFiles([...files, newFile]);
          setActiveFile(newFile);
        }}
        onSaveFile={() => activeFile && handleSaveFile(activeFile)} 
        onOpenFolder={handleOpenFolder}
        onUndo={handleUndo}
        onRedo={handleRedo}
        onCut={handleCut}
        onCopy={handleCopy}
        onPaste={handlePaste}
      />
      <div className="main-content">
        <Split
          className="split-horizontal"
          sizes={[20, 80]}
          minSize={[100, 300]}
          direction="horizontal"
        >
          <Sidebar
            files={files}
            onFileSelect={handleFileSelect}
            onCloseFile={handleCloseFile}
            activeFile={activeFile}
            onSaveFile={handleSaveFile}
            onRunCode={handleRunCode}
          />
          <Split
            className="split-vertical"
            sizes={[70, 25]}
            minSize={[250, 150]}
            direction="vertical"
          >
            <Editor file={activeFile} onContentChange={handleContentChange} ref={editorRef} />
            <Console output={consoleOutput} />
          </Split>
        </Split>
      </div>
      <StatusBar />
    </div>
  );
};

export default App;
