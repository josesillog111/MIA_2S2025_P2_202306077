import React, { useState } from 'react';
import TerminalPage from './pages/TerminalPage';
import FileExplorerPage from './pages/FileExplorerPage';
import DiskSelectionPage from './pages/DiskSelectionPage';
import PartitionSelectionPage from './pages/PartitionSelectionPage';
import FilesystemFormatPage from './pages/FilesystemFormatPage';
import LoginPage from './pages/LoginPage';
import DashboardPage from './pages/DashboardPage';
import type { ToolPageProps } from './ToolPageProps';
// Componentes
import { LogOut, X, Play, LayoutDashboard, Terminal, Folder, Settings, Code, Monitor, CheckCircle, AlertTriangle, User } from 'lucide-react';


type ActiveTool = "terminal" | "settings" | "file_explorer" | "dashboard";
type MountPhase = 'SELECT_DISK' | 'SELECT_PARTITION' | 'LOGIN' | 'FORMAT_PARTITION';

// C4. SettingsPage, C5. LayoutsPage (Sin cambios)
const SettingsPage: React.FC<ToolPageProps> = () => (
    <div className="p-8 h-full bg-[#1e1e1e] text-gray-200"><h2 className="text-3xl font-light text-yellow-400">Configuración</h2><p className="text-gray-400">Opciones de personalización de la interfaz.</p></div>
);

interface SidebarProps {
    currentView: ActiveTool;
    onViewChange: (view: ActiveTool) => void;
}

// Datos de la navegación (Editor reemplaza Layouts)
const navItems = [
    { id: 'dashboard', icon: LayoutDashboard, label: 'Panel', tooltip: 'Vista de inicio' },
    { id: 'terminal', icon: Terminal, label: 'Terminal', tooltip: 'Consola de comandos' },
    { id: 'file_explorer', icon: Folder, label: 'Explorador', tooltip: 'Archivos del proyecto' },
] as const;


const Sidebar: React.FC<SidebarProps> = ({ currentView, onViewChange }) => {
  return (
    <div className="flex flex-col h-full w-[80px] bg-[#252526] items-center shadow-2xl z-20 transition-all duration-300">
      {/* === Logo / Icono Principal === */}
      <div className="p-3 mb-6 mt-3">
        <Code className="w-8 h-8 text-blue-500" />
      </div>

      {/* === Ítems de navegación === */}
      <nav className="flex flex-col space-y-3 w-full items-center">
        {navItems.map((item) => (
          <button
            key={item.id}
            onClick={() => onViewChange(item.id as ActiveTool)}
            title={item.tooltip}
            className={`
              flex flex-col items-center justify-center w-[60px] h-[60px] rounded-xl transition-all duration-200 
              ${
                currentView === item.id
                  ? "bg-blue-600 text-white scale-105 shadow-lg"
                  : "text-gray-400 bg-[#252526] hover:bg-[#3a3a3a] hover:text-white"
              }
            `}
          >
            <item.icon className="w-6 h-6" />
          </button>
        ))}
      </nav>

      {/* === Botón de Ajustes (parte inferior) === */}
      <div className="mt-auto pb-4 w-full flex flex-col items-center border-t border-[#3a3a3a] pt-3">
        <button
          onClick={() => onViewChange("settings")}
          title="Configuración de la aplicación"
          className={`
            flex flex-col items-center justify-center w-[60px] h-[60px] rounded-xl transition-all duration-200
            ${
              currentView === "settings"
                ? "bg-blue-600 text-white scale-105 shadow-lg"
                : "text-gray-400 bg-[#252526] hover:bg-[#3a3a3a] hover:text-white"
            }
          `}
        >
          <Settings className="w-6 h-6" />
        </button>
      </div>
    </div>
  );
};


const App: React.FC = () => {
  const [showMountWizard, setShowMountWizard] = useState(true);
  const [mountPhase, setMountPhase] = useState<MountPhase>("SELECT_DISK");

  const [activeTool, setActiveTool] = useState<ActiveTool>("terminal");
  const [mountInfo, setMountInfo] = useState<any>(null);
  const [sessionInfo, setSessionInfo] = useState<any>(null);

  const [showAlert, setShowAlert] = useState(false);

  const isReady = mountInfo !== null && sessionInfo !== null;

  // === HANDLERS DEL FLUJO DE INICIO ===
  const handleDiskSelect = (diskInfo: any) => {
    setMountInfo(diskInfo);
    setMountPhase("SELECT_PARTITION");
  };

  const checkIfPartitionNeedsFormatting = (partitionInfo: any): boolean => {
      if (!partitionInfo) return true;
      // Si el backend devuelve algo como { file_system: "EXT4" } o similar
      const fs = partitionInfo.file_system || partitionInfo.fsType || "";
      return fs === "" || fs === null || fs === "NO_FS";
  };


  const handlePartitionSelect = (partitionInfo: any) => {
      if (!partitionInfo) return;

      // 1. Actualizar la información de montaje combinando disco y partición
      const newMountInfo = { ...mountInfo, ...partitionInfo };
      setMountInfo(newMountInfo);

      // 2. Redirección según la bandera 'needsFormat'
      if (partitionInfo.needsFormat) {
          console.log("Partición sin sistema de archivos detectado, redirigiendo a FORMAT_PARTITION...");
          setMountPhase("FORMAT_PARTITION");
      } else {
          console.log("Partición con sistema de archivos detectado, redirigiendo a LOGIN...");
          setMountPhase("LOGIN");
      }
  };


  const handleLoginSuccess = (info: any) => {
    setSessionInfo(info);
    setShowMountWizard(false);
    setActiveTool("terminal");
  };

  const handleBack = () => {
    if (mountPhase === "SELECT_PARTITION") {
      setMountInfo(null);
      setMountPhase("SELECT_DISK");
    } else if (mountPhase === "FORMAT_PARTITION") {
      setMountPhase("SELECT_PARTITION");
    } else if (mountPhase === "LOGIN") {
      if (checkIfPartitionNeedsFormatting(mountInfo)) {
        setMountPhase("FORMAT_PARTITION");
      } else {
        setMountPhase("SELECT_PARTITION");
      }
    }
  };

  const handleLogout = () => {
    setSessionInfo(null);
    setMountInfo(null);
    setMountPhase("SELECT_DISK");
    setShowMountWizard(true);
    setActiveTool("terminal");
  };

  const handleInitializeSession = () => {
    setMountPhase("SELECT_DISK");
    setShowMountWizard(true);
  };

  // === HANDLER para el cambio de vista ===
  const handleViewChange = (view: ActiveTool) => {
    if (view === "file_explorer" && !isReady) {
      // Bloquear acceso si no hay sesión montada
      setShowAlert(true);
      return;
    }
    setActiveTool(view);
  };

  const handleFormatComplete = (partitionInfoWithFS: any) => {
    // Actualizar mountInfo con la nueva información del sistema de archivos
    const newMountInfo = { ...mountInfo, ...partitionInfoWithFS };
    setMountInfo(newMountInfo);
    // Ya formateado, ir a Login
    setMountPhase("LOGIN");
  };


  // === RENDER DEL WIZARD ===
  const renderMountWizard = () => {
    let CurrentComponent;
    let props: any = { onBack: handleBack };

    switch (mountPhase) {
      case "SELECT_DISK":
        CurrentComponent = DiskSelectionPage;
        props.mountInfo = mountInfo;
        props.onSelect = handleDiskSelect;
        break;
      case "SELECT_PARTITION":
        CurrentComponent = PartitionSelectionPage;
        props.disk = mountInfo;
        props.onSelect = handlePartitionSelect; // Redirige a FORMAT_PARTITION o LOGIN
        break;
      case "FORMAT_PARTITION":
        CurrentComponent = FilesystemFormatPage;
        props.partition = mountInfo;
        props.onSelect = handleFormatComplete; 
        break;
      case "LOGIN":
        CurrentComponent = LoginPage;
        props.mountInfo = mountInfo; // Pasar info de disco/partición
        props.onLogin = handleLoginSuccess;
        break;
      default:
        return (
          <div className="p-10 text-white bg-red-800">
            Error: Estado de fase de montaje desconocido.
          </div>
        );
    }

    return <CurrentComponent {...props} />;
  };

  // === RENDER PRINCIPAL ===
  const renderToolContent = () => {
    const commonProps = { mountInfo, sessionInfo, isReady, onInitialize: handleInitializeSession };

    switch (activeTool) {
      case "file_explorer":
        return <FileExplorerPage />;
      case "terminal":
        return <TerminalPage {...commonProps} />;
      case "settings":
        return <SettingsPage {...commonProps} />;
      case "dashboard":
        return <DashboardPage {...commonProps} />;
      default:
        return <DashboardPage {...commonProps} />;
    }
  };


  const projectName = mountInfo?.name || 'Sistema Desconectado';
  const headerBg = 'bg-[#3a3a3a] border-[#4a4a4a]';
  const headerText = 'text-gray-400';
  const headerAppText = 'text-blue-400';

  return (
    <div className="flex flex-col h-screen">
      {/* === Barra superior === */}
      <div className={`h-12 border-b flex items-center justify-between px-4 text-sm font-medium ${headerBg} ${headerText} shadow-md`}>
                
                {/* 1. Nombre del Proyecto/IDE */}
                <div className="flex items-center">
                    <span className={`text-lg font-bold mr-4 flex items-center ${headerAppText}`}>
                        <Monitor className="w-5 h-5 mr-2" /> 
                        {isReady ? (
                            <span className="text-white font-semibold">{projectName}</span>
                        ) : (
                            <span className="text-gray-400">IDE App</span>
                        )}
                    </span>
                </div>
                
                {/* 2. Indicador Central de Estado */}
                <div className="flex items-center space-x-2 p-1.5 rounded-full px-4" style={{ background: "#2e2e2e" }}>
                    {isReady ? (
                        <>
                            <CheckCircle className="w-4 h-4 text-green-500" />
                            <span className="text-green-500 font-semibold text-xs">SESIÓN ACTIVA</span>
                        </>
                    ) : (
                        <>
                            <AlertTriangle className="w-4 h-4 text-red-500" />
                            <span className="text-red-500 font-semibold text-xs">SIN MONTAJE</span>
                        </>
                    )}
                </div>
                
                {/* 3. Botón de Sesión y Usuario */}
                <div className='flex items-center space-x-4'>
                    <span className={`text-sm flex items-center text-gray-300`}>
                        <User className='w-4 h-4 mr-1.5 text-blue-400'/>
                        {sessionInfo?.user || 'Invitado'}
                    </span>
                    <button 
                        onClick={isReady ? handleLogout : handleInitializeSession} 
                        title={isReady ? "Cerrar Sesión y Desmontar" : "Iniciar Flujo de Montaje"}
                        className={`flex items-center space-x-1 px-3 py-1 rounded-md transition-colors font-semibold text-sm shadow-md ${isReady ? 'text-red-400 hover:text-red-300 bg-red-800/20 border border-red-800' : 'text-white bg-blue-600 hover:bg-blue-700'}`}
                    >
                        {isReady ? <LogOut className="w-4 h-4" /> : <Play className="w-4 h-4" />}
                        <span>{isReady ? 'Logout' : 'Iniciar'}</span>
                    </button>
                </div>
            </div>
      {/* === Contenido principal === */}
      <div className="flex flex-1 overflow-hidden">
        <Sidebar currentView={activeTool} onViewChange={handleViewChange} />
        <div className="flex-1 overflow-auto bg-[#1e1e1e]">{renderToolContent()}</div>
      </div>

      {/* === WIZARD === */}
      {showMountWizard && (
        <div className="absolute inset-0 bg-black bg-opacity-90 flex justify-center items-center z-50">
          <div className="w-4/5 md:w-3/5 h-3/4 shadow-2xl rounded-xl overflow-hidden relative">
            <button
              onClick={() => setShowMountWizard(false)}
              className="absolute top-4 right-4 text-white text-lg bg-red-600 p-2 rounded-full hover:bg-red-700 z-10"
              title="Cerrar Asistente"
            >
              <X className="w-5 h-5" />
            </button>
            {renderMountWizard()}
          </div>
        </div>
      )}

      {showAlert && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-[60]">
          <div className="bg-[#2c2c2c] text-white p-6 rounded-lg shadow-xl text-center w-80">
            <h2 className="text-lg font-bold mb-2">Acceso Denegado</h2>
            <p className="text-gray-300 mb-4">
              Debes iniciar sesión y montar un sistema antes de abrir el explorador de archivos.
            </p>
            <button
              onClick={() => setShowAlert(false)}
              className="bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded-md font-semibold transition"
            >
              Entendido
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default App;

