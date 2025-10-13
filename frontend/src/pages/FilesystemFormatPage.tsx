import React, { useState } from 'react';
import { Database, Zap, CheckCircle, ArrowRight, ArrowLeft } from 'lucide-react';
import type { WizardPageProps } from '../WizardPageProps';

// Tipos de sistemas de archivos disponibles
const FILE_SYSTEM_OPTIONS = [
    { id: 'EXT2', name: 'EXT2 (Linux)', description: 'Amplio soporte, ideal para unidades USB.' },
    { id: 'EXT3', name: 'EXT3 (Linux)', description: 'Sistema de archivos robusto y confiable, con soporte para journaling.' },
];

// B2.5. FilesystemFormatPage (Aparece si la partición NO tiene FS)
const FilesystemFormatPage: React.FC<WizardPageProps> = ({ partition, onSelect, onBack }) => {
    const [selectedFS, setSelectedFS] = useState<string | null>(null);
    const [isFormatted, setIsFormatted] = useState(false);
    const [isFormatting, setIsFormatting] = useState(false);
    
    const handleFormat = (e: React.FormEvent) => {
        e.preventDefault();
        if (!selectedFS) return;

        setIsFormatting(true);
        setTimeout(() => {
            setIsFormatting(false);
            setIsFormatted(true);
        }, 1500); // 1.5 segundos de simulación
    };

    // Manejador del botón "Siguiente"
    const handleNext = (e: React.FormEvent) => {
        e.preventDefault();
        if (isFormatted && selectedFS && onSelect) {
            // Pasamos la partición y el FS seleccionado al siguiente paso (Login)
            onSelect({ 
                id: partition?.id || `fs-${Date.now()}`,  // ID temporal
                name: partition?.name || "Partición sin nombre",
                type: partition?.type || "Desconocido",
                filesystem: selectedFS 
            });
        }
    };
    
    const partitionName = partition?.name || 'Partición Desconocida';

    return (
        <div className="flex w-full h-full bg-[#1e1e1e] text-gray-200 shadow-xl rounded-lg overflow-hidden">
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-gradient-to-br from-[#0d0d0d] via-[#1a1a1a] to-[#2c2c2c] text-center">
                <Zap className="w-16 h-16 text-yellow-400 mb-4" />
                <h1 className="text-4xl font-bold mb-4 text-white">Paso 2.5: Formato de Sistema de Archivos</h1>
                <p className="text-lg text-yellow-200 max-w-sm">La partición **{partitionName}** no tiene un sistema de archivos reconocido.</p>
                <p className="text-md text-gray-400 mt-2">Por favor, elige y aplica un formato para continuar.</p>
                <div className="mt-6 p-2 bg-[#3a3a3a] rounded-md text-sm text-gray-400">
                    <Database className="inline w-4 h-4 mr-2" />
                    Partición: {partitionName}
                </div>
            </div>
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-[#2a2a2a]">
                <div className="w-full max-w-md">
                    <h2 className="text-2xl font-semibold mb-6 text-gray-100 text-center">SELECCIONAR Y FORMATO</h2>
                    <form onSubmit={handleNext} className="space-y-4">
                        <div className="space-y-3 max-h-64 overflow-y-auto pr-2">
                            {FILE_SYSTEM_OPTIONS.map(fs => (
                                <div
                                    key={fs.id}
                                    onClick={() => { setSelectedFS(fs.id); setIsFormatted(false); }}
                                    className={`flex flex-col p-4 cursor-pointer bg-[#3a3a3a] border rounded-md transition-all ${selectedFS === fs.id ? 'border-yellow-500 ring-2 ring-yellow-500' : 'border-[#4a4a4a] hover:border-gray-500'}`}
                                >
                                    <p className="font-medium text-gray-100">{fs.name}</p>
                                    <p className="text-xs text-gray-400 mt-1">{fs.description}</p>
                                </div>
                            ))}
                        </div>
                        
                        {/* Botón de Formateo */}
                        <button
                            type="button"
                            onClick={handleFormat}
                            disabled={!selectedFS || isFormatting || isFormatted}
                            className={`w-full py-3 px-4 text-white font-semibold rounded-md transition-all flex items-center justify-center gap-2 ${
                                !selectedFS ? 'bg-gray-500 cursor-not-allowed' :
                                isFormatting ? 'bg-yellow-700 animate-pulse' :
                                isFormatted ? 'bg-green-600 hover:bg-green-700' :
                                'bg-gradient-to-r from-yellow-600 to-orange-600 hover:from-yellow-700 hover:to-orange-700'
                            }`}
                        >
                            {isFormatting && <Zap className="w-5 h-5 animate-spin" />}
                            {isFormatted ? (
                                <>
                                    <CheckCircle className="w-5 h-5" /> Formato Aplicado ({selectedFS})
                                </>
                            ) : (
                                `Aplicar Formato: ${selectedFS || 'Seleccione uno'}`
                            )}
                        </button>

                        <div className="flex gap-4 pt-2">
                            <button 
                                type="button" 
                                onClick={onBack} 
                                className="w-1/3 py-3 px-4 bg-gray-500 text-white font-semibold rounded-md hover:bg-gray-600 transition-all flex items-center justify-center"
                            >
                                <ArrowLeft className="w-4 h-4 inline mr-2" /> Atrás
                            </button>
                            <button 
                                type="submit" 
                                disabled={!isFormatted} 
                                className={`w-2/3 py-3 px-4 text-white font-semibold rounded-md transition-all ${isFormatted ? 'bg-gradient-to-r from-teal-500 to-emerald-600 hover:from-teal-600 hover:to-emerald-700' : 'bg-gray-500 cursor-not-allowed'}`}
                            >
                                Continuar al Login<ArrowRight className="w-4 h-4 inline ml-2" />
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default FilesystemFormatPage;