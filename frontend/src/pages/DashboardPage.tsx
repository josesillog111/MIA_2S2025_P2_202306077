import type { ToolPageProps } from '../ToolPageProps';
import React from 'react';
import { LayoutDashboard, CheckCircle, AlertTriangle, Cpu, Activity, Clock, HardDrive as HDIcon, User, Lock, Play, HardDrive } from 'lucide-react';

const DashboardPage: React.FC<ToolPageProps> = ({ mountInfo, sessionInfo, isReady, onInitialize }) => {
    
    // Datos simulados para el Dashboard
    const systemMetrics = [
        { icon: Cpu, label: "Uso de CPU", value: "32%", color: "text-red-400", bgColor: "bg-red-900/30" },
        { icon: Activity, label: "Memoria RAM", value: "8.5 GB / 16 GB", color: "text-blue-400", bgColor: "bg-blue-900/30" },
        { icon: Clock, label: "Tiempo Activo", value: "12h 45m", color: "text-green-400", bgColor: "bg-green-900/30" },
        { icon: HDIcon, label: "Espacio en Disco", value: "150 GB libres", color: "text-yellow-400", bgColor: "bg-yellow-900/30" },
    ];

    const recentActivity = [
        { file: "src/App.tsx", time: "hace 5 minutos", action: "Guardado" },
        { file: "README.md", time: "hace 1 hora", action: "Editado" },
        { file: "styles/theme.css", time: "hace 3 horas", action: "Creado" },
        { file: "server.js", time: "hace 1 día", action: "Depurado" },
    ];


    return (
        <div className="p-8 h-full bg-[#1e1e1e] text-gray-200 overflow-auto">
            <h2 className="text-4xl font-extralight text-white mb-6 flex items-center">
                <LayoutDashboard className="w-8 h-8 mr-3 text-purple-400" /> Panel de Control
            </h2>

            {/* === 1. Panel de Estado de Sesión === */}
            <div className={`p-5 rounded-xl mb-8 shadow-xl ${isReady ? 'bg-green-900/20 border border-green-700/50' : 'bg-red-900/20 border border-red-700/50'}`}>
                <div className="flex items-center justify-between">
                    <h3 className="text-2xl font-semibold flex items-center">
                        {isReady ? <CheckCircle className="w-6 h-6 mr-2 text-green-400" /> : <AlertTriangle className="w-6 h-6 mr-2 text-red-400" />}
                        Estado del Sistema
                    </h3>
                    <span className={`px-3 py-1 text-sm font-bold rounded-full ${isReady ? 'bg-green-600' : 'bg-red-600'}`}>
                        {isReady ? 'ACTIVO' : 'DESCONECTADO'}
                    </span>
                </div>
                
                <div className="mt-4 grid grid-cols-1 md:grid-cols-3 gap-4 text-gray-300">
                    <div className="flex items-center space-x-3 p-3 bg-[#252526] rounded-lg border border-[#3a3a3a]">
                        <User className="w-5 h-5 text-blue-400" />
                        <div>
                            <p className="text-xs text-gray-500">Usuario Activo</p>
                            <p className="font-mono font-semibold">{sessionInfo?.user || 'Invitado'}</p>
                        </div>
                    </div>
                    <div className="flex items-center space-x-3 p-3 bg-[#252526] rounded-lg border border-[#3a3a3a]">
                        <HardDrive className="w-5 h-5 text-yellow-400" />
                        <div>
                            <p className="text-xs text-gray-500">Partición Montada</p>
                            <p className="font-mono font-semibold">{mountInfo?.name || 'N/A'}</p>
                        </div>
                    </div>
                    <div className="flex items-center space-x-3 p-3 bg-[#252526] rounded-lg border border-[#3a3a3a]">
                        <Lock className="w-5 h-5 text-purple-400" />
                        <div>
                            <p className="text-xs text-gray-500">Nivel de Acceso</p>
                            <p className="font-mono font-semibold">{isReady ? 'R/W (Desarrollo)' : 'Solo Lectura'}</p>
                        </div>
                    </div>
                </div>

                {!isReady && (
                    <button 
                        onClick={onInitialize} 
                        className="mt-6 py-2 px-4 w-full bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 transition-colors flex items-center justify-center shadow-md"
                    >
                        <Play className="w-5 h-5 mr-2" /> Iniciar Sesión y Montaje
                    </button>
                )}
            </div>

            {/* === 2. Métricas del Sistema === */}
            <h3 className="text-xl font-semibold text-gray-300 mb-4">Métricas del Servidor (Simulación)</h3>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
                {systemMetrics.map((metric, index) => (
                    <div key={index} className={`p-6 rounded-xl bg-[#252526] shadow-lg transition-transform hover:scale-[1.02] border border-[#3a3a3a] ${metric.bgColor}`}>
                        <div className={`p-3 rounded-full inline-block mb-3 ${metric.bgColor}`}>
                            <metric.icon className={`w-6 h-6 ${metric.color}`} />
                        </div>
                        <p className="text-xs text-gray-400 uppercase font-semibold">{metric.label}</p>
                        <p className="text-3xl font-bold mt-1 text-white">{metric.value}</p>
                    </div>
                ))}
            </div>

            {/* === 3. Actividad Reciente === */}
            <h3 className="text-xl font-semibold text-gray-300 mb-4">Actividad Reciente del Proyecto</h3>
            <div className="bg-[#252526] rounded-xl shadow-lg border border-[#3a3a3a] p-4">
                <ul className="divide-y divide-[#3a3a3a]">
                    {recentActivity.map((activity, index) => (
                        <li key={index} className="flex justify-between items-center py-3 hover:bg-[#2a2a2a] px-2 rounded-lg transition-colors">
                            <div className="flex items-center space-x-3">
                                <span className={`w-2 h-2 rounded-full ${activity.action === 'Guardado' ? 'bg-green-500' : activity.action === 'Editado' ? 'bg-blue-500' : 'bg-yellow-500'}`}></span>
                                <span className="font-mono text-sm text-blue-400">{activity.file}</span>
                            </div>
                            <div className="text-right">
                                <p className="text-sm text-gray-300">{activity.action}</p>
                                <p className="text-xs text-gray-500">{activity.time}</p>
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default DashboardPage;