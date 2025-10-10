import React from 'react';
import ElementoSidebarFS from './ElementsSidebar'; // Usamos el nombre actualizado
import type { ElementoSidebar as IElementoSidebar } from '../../types/FileSystem';

// Definición de las props
interface Props {
  items: IElementoSidebar[];
  rutaActual: string;
  onNavegar: (ruta: string) => void;
  // Nuevas props para acciones (similares a las de tu Sidebar de VS Code)
  onRefresh?: () => void;
  onNewFolder?: () => void;
}

const SidebarFS: React.FC<Props> = ({ items, rutaActual, onNavegar, onRefresh, onNewFolder }) => {
  const ubicacionesPrincipales = items.filter(item => item.esUbicacionPrincipal);
  const otrasUbicaciones = items.filter(item => !item.esUbicacionPrincipal);

  return (
    // Contenedor principal del sidebar
    <div className="w-64 bg-[#2D2D2D] text-white h-full flex flex-col overflow-hidden shadow-xl">
      
      {/* 1. Cabecera (Similar a 'EXPLORER' en VS Code) */}
      <div className="p-3 px-4 uppercase text-sm font-semibold text-gray-400 border-b border-gray-700">
        Ubicaciones
      </div>
      
      {/* 2. Sección de Acciones (Botones) */}
      <div className="p-3 border-b border-gray-700 flex justify-around space-x-2">
        {/* Botón de Actualizar/Refresh */}
        <button 
          onClick={onRefresh} 
          className="flex-1 p-2 bg-gray-600 hover:bg-gray-500 text-white rounded-md transition-colors text-sm font-medium"
          title="Actualizar la vista"
        >
          <i className="fas fa-arrows-rotate mr-2"></i> Actualizar
        </button>
        
        {/* Botón de Nueva Carpeta */}
        <button 
          onClick={onNewFolder} 
          className="flex-1 p-2 bg-blue-600 hover:bg-blue-500 text-white rounded-md transition-colors text-sm font-medium"
          title="Crear nueva carpeta"
        >
          <i className="fas fa-folder-plus mr-2"></i> Nuevo
        </button>
      </div>

      {/* 3. Lista de Ubicaciones (Scrollable Content) */}
      <div className="p-2 flex-grow overflow-y-auto">
        
        {/* Sección principal: Ubicaciones rápidas */}
        <div className="mb-6">
          <h3 className="text-xs font-semibold uppercase text-gray-400 mb-1 ml-2 mt-2">Carpetas</h3>
          
          {ubicacionesPrincipales.map(item => (
            <ElementoSidebarFS
              key={item.ruta}
              item={item}
              // Comprueba si la ruta actual coincide o está dentro de la ruta del sidebar
              estaActivo={rutaActual.startsWith(item.ruta) && rutaActual.length >= item.ruta.length} 
              onClick={onNavegar}
            />
          ))}
        </div>

        {/* Otras Ubicaciones */}
        {otrasUbicaciones.length > 0 && (
          <div className="pt-4 border-t border-gray-700/50">
            <h3 className="text-xs font-semibold uppercase text-gray-400 mb-1 ml-2">Dispositivos</h3>
            {otrasUbicaciones.map(item => (
              <ElementoSidebarFS
                key={item.ruta}
                item={item}
                estaActivo={rutaActual.startsWith(item.ruta)}
                onClick={onNavegar}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default SidebarFS;