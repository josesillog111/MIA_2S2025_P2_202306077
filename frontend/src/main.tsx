import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
// 1. Importar componentes necesarios de react-router-dom
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LoginPage from './pages/LoginPage.tsx';
import './index.css';

// 2. Importar los componentes de las páginas
import App from './App.tsx'; // Tu editor de código principal
import FileExplorerPage from './pages/FileExplorerPage'; // El nuevo explorador de archivos

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    {/* 3. El BrowserRouter es el router principal */}
    <BrowserRouter>
      {/* 4. Routes define todas las rutas de la aplicación */}
      <Routes>
        {/* Ruta raíz: renderiza el componente principal del editor */}
        <Route path="/" element={<App />} />
        
        {/* Nueva ruta: renderiza el Explorador de Archivos */}
        <Route path="/file_explorer" element={<FileExplorerPage />} />
        <Route path="/login" element={<LoginPage onLogin={(sessionInfo) => {
          console.log('Usuario autenticado:', sessionInfo);
          // Aquí puedes manejar la información de sesión, como almacenarla en el estado global o en el almacenamiento local
        }
        } />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
);