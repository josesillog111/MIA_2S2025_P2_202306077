import React from 'react';
// IMPORTA AQUÍ el componente completo del explorador de archivos
import Toolbar from '../components/file_explorer/Toolbar';
import FileContent from '../components/file_explorer/FileContent';

const FileExplorerPage: React.FC = () => {
  return (
    // El contenedor de la página no necesita un div con relleno (p-8)
    // porque ExploradorArchivos ya está configurado para tomar toda la pantalla.
   <div className="flex flex-col h-screen w-full">
      <Toolbar />
      <FileContent />
    </div>
  );
};

export default FileExplorerPage;