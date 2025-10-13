// src/components/Sidebar.tsx
import React from 'react';
import { FileText, Folder, Settings, Terminal, LayoutPanelLeft } from 'lucide-react';

// 💡 Define las herramientas que puede seleccionar el usuario.
export type ActiveTool = 'editor' | 'file_explorer' | 'terminal' | 'settings' | 'layouts';

// 💡 Interfaz para las props del componente.
interface SidebarProps {
    activeTool: ActiveTool;
    onSelect: (tool: ActiveTool) => void;
}

// 💡 Componente interno para cada botón de la barra lateral.
const NavItem: React.FC<{ 
    icon: React.ReactNode; 
    label: ActiveTool; 
    active: boolean; 
    onClick: () => void;
}> = ({ icon, label, active, onClick }) => (
    <button
        // Muestra el nombre de la herramienta al pasar el mouse.
        title={label.replace('_', ' ').split(' ').map(s => s.charAt(0).toUpperCase() + s.slice(1)).join(' ')}
        onClick={onClick}
        // Estilos para resaltar el botón activo.
        className={`p-3 text-gray-400 hover:text-gray-100 transition-colors 
                    ${active ? 'text-blue-400 border-l-2 border-blue-400 bg-[#3a3a3a]' : 'border-l-2 border-transparent'}`}
    >
        {icon}
    </button>
);

/**
 * Sidebar: Barra de navegación principal de la aplicación, inspirada en un IDE.
 * Permite al usuario cambiar entre diferentes herramientas y vistas.
 */
const Sidebar: React.FC<SidebarProps> = ({ activeTool, onSelect }) => {
    return (
        // Contenedor principal de la barra lateral vertical.
        <div className="w-16 bg-[#252526] flex flex-col items-center py-4 border-r border-[#3a3a3a] h-full">
            
            {/* Ícono para el Editor de Código */}
            <NavItem 
                icon={<FileText className="w-6 h-6" />} 
                label="editor" 
                active={activeTool === 'editor'} 
                onClick={() => onSelect('editor')} 
            />
            
            {/* Ícono para el Explorador de Archivos */}
            <NavItem 
                icon={<Folder className="w-6 h-6" />} 
                label="file_explorer" 
                active={activeTool === 'file_explorer'} 
                onClick={() => onSelect('file_explorer')} 
            />

            {/* Ícono para la Terminal (Consola) */}
            <NavItem
                icon={<Terminal className="w-6 h-6" />}
                label="terminal"
                active={activeTool === 'terminal'}
                onClick={() => onSelect('terminal')}
            />

            {/* Ícono para Layouts/Vistas */}
            <NavItem
                icon={<LayoutPanelLeft className="w-6 h-6" />}
                label="layouts"
                active={activeTool === 'layouts'}
                onClick={() => onSelect('layouts')}
            />
            
            {/* Empuja el siguiente ícono hacia abajo. */}
            <div className="flex-1"></div> 
            
            {/* Ícono para Configuración */}
            <NavItem 
                icon={<Settings className="w-6 h-6" />} 
                label="settings" 
                active={activeTool === 'settings'} 
                onClick={() => onSelect('settings')} 
            />
        </div>
    );
};

export default Sidebar;