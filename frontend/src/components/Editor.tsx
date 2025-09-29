// src/components/Editor.tsx
import { useRef, useImperativeHandle, forwardRef } from 'react';
import EditorComponent from '@monaco-editor/react';

interface File {
  name: string;
  path: string;
  content: string;
}

interface EditorProps {
  file: File | null;
  onContentChange: (newContent: string) => void;
}

// 1. Usamos `forwardRef` para que el componente padre pueda obtener una referencia.
const Editor = forwardRef<any, EditorProps>(({ file, onContentChange }, ref) => {
  // Usa useRef para guardar la instancia del editor.
  const editorRef = useRef<any>(null);

  // 2. Con `useImperativeHandle`, exponemos los métodos del editor al componente padre.
  useImperativeHandle(ref, () => ({
    // Aquí definimos los métodos que el componente padre puede llamar.
    undo: () => {
      editorRef.current?.trigger('ribbon-action', 'undo', null);
    },
    redo: () => {
      editorRef.current?.trigger('ribbon-action', 'redo', null);
    },
    cut: () => {
      editorRef.current?.trigger('ribbon-action', 'editor.action.clipboardCutAction', null);
    },
    copy: () => {
      editorRef.current?.trigger('ribbon-action', 'editor.action.clipboardCopyAction', null);
    },
    paste: () => {
      editorRef.current?.trigger('ribbon-action', 'editor.action.clipboardPasteAction', null);
    },
  }));

  const handleEditorDidMount = (editor: any) => {
    editorRef.current = editor;
  };

  if (!file) {
    return (
      <div className="editor-container">
        <div className="no-file-open">Selecciona un archivo para empezar a editar.</div>
      </div>
    );
  }

  const getFileLanguage = (fileName: string) => {
    const extension = fileName.split('.').pop();
    switch (extension) {
      case 'ts':
      case 'tsx':
        return 'typescript';
      case 'js':
      case 'jsx':
        return 'javascript';
      case 'json':
        return 'json';
      case 'html':
        return 'html';
      case 'css':
        return 'css';
      case 'md':
        return 'markdown';
      default:
        return 'plaintext';
    }
  };

  return (
    <div className="editor-container">
      <div className="tab-bar">
        <div className="tab-item active">{file.name}</div>
      </div>
      <EditorComponent
        height="89%"
        width="100%"
        theme="vs-dark"
        language={getFileLanguage(file.name)}
        value={file.content}
        onChange={(value) => onContentChange(value || '')}
        onMount={handleEditorDidMount}
        options={{
          minimap: { enabled: false },
          fontSize: 14,
          quickSuggestions: false,
          automaticLayout: true,
          glyphMargin: false,
          folding: false,
          renderValidationDecorations: 'off',
          scrollbar: {
            vertical: 'auto',
            horizontal: 'auto',
          },
          lineNumbersMinChars: 3,
          suggest: {
            preview: false
          },
        }}
      />
    </div>
  );
});

export default Editor;