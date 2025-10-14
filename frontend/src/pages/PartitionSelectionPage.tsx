import React, { useState, useEffect } from 'react';
import { HardDrive, ArrowRight, Server, Database, Loader2, ArrowLeft, CheckCircle, XCircle } from 'lucide-react';
import type { WizardPageProps } from '../WizardPageProps';

// Interfaces para tipado estricto
interface PartitionRecord {
    id: string;      // Usaremos path o una combinación de nombre/size como ID
    name: string;
    size: number;    // El backend devuelve int64, usemos number
    type: 'Primaria' | 'Lógica'; // Para diferenciar MBR y EBR
    used: string;    // Simulado para la barra de progreso (e.g., "50%")
}

// Interfaz para los datos del backend
interface BackendPartition {
    status: number;
    type: string; // P, E
    fit: string;
    start: number;
    size: number;
    name: string;
    correlative?: number; // Solo en MBR
    id?: string;         // Solo en MBR
    next?: number;       // Solo en EBR (Next)
    usage: number; // Solo en EBR
}

// B2. PartitionSelectionPage
const PartitionSelectionPage: React.FC<WizardPageProps> = ({ disk, onSelect, onBack }) => {
    // Estado para las particiones cargadas
    const [partitions, setPartitions] = useState<PartitionRecord[]>([]);
    const [selectedPartition, setSelectedPartition] = useState<string | null>(null);
    
    // Nuevo estado para la verificación del sistema de archivos
    const [hasFs, setHasFs] = useState<boolean | null>(null); // null = no verificado, true/false = resultado
    const [isCheckingFs, setIsCheckingFs] = useState(false); // Estado de carga de la verificación
    
    const [isLoading, setIsLoading] = useState(true); // Estado de carga inicial de particiones
    const [error, setError] = useState<string | null>(null);

    // 1. Obtener la ruta del disco del prop
    const diskPath = disk?.id; 

    // Función para llamar al backend y verificar el sistema de archivos
    const checkPartitionFilesystem = async (diskPath: string, partitionName: string) => { setIsCheckingFs(true); setHasFs(null); // Reiniciar el estado FS al iniciar la verificación try { // Implementar exponencial backoff para el reintento de la llamada a la API const MAX_RETRIES = 2; let response; for (let i = 0; i < MAX_RETRIES; i++) { const API_URL = /disk/checkfs?diskPath=${encodeURIComponent(diskPath)}&partitionName=${encodeURIComponent(partitionName)}; response = await fetch(API_URL); if (response.ok) { break; } // Esperar con backoff exponencial antes de reintentar if (i < MAX_RETRIES - 1) { await new Promise(resolve => setTimeout(resolve, Math.pow(2, i) * 1000)); } } if (!response || !response.ok) { console.error(Error HTTP al verificar FS: ${response?.status}); setHasFs(false); return false; } const data: { hasFS: boolean } = await response.json(); setHasFs(data.hasFS); return data.hasFS; } catch (err) { console.error("Error al verificar el sistema de archivos:", err); // Si hay un error de red o de parsing, asumimos que no hay FS para permitir el formato. setHasFs(false); return false; } finally { setIsCheckingFs(false); } };
        setIsCheckingFs(true);
        setHasFs(null); // Reiniciar el estado FS al iniciar la verificación

        try {
            // Implementar exponencial backoff para el reintento de la llamada a la API
            const MAX_RETRIES = 2;
            let response;

            for (let i = 0; i < MAX_RETRIES; i++) {
                const API_URL = `/disk/checkfs?diskPath=${encodeURIComponent(diskPath)}&partitionName=${encodeURIComponent(partitionName)}`;
                response = await fetch(API_URL);

                if (response.ok) {
                    break;
                }

                // Esperar con backoff exponencial antes de reintentar
                if (i < MAX_RETRIES - 1) {
                    await new Promise(resolve => setTimeout(resolve, Math.pow(2, i) * 1000));
                }
            }

            if (!response || !response.ok) {
                console.error(`Error HTTP al verificar FS: ${response?.status}`);
                setHasFs(false);
                return false;
            }

            const data: { hasFS: boolean } = await response.json();
            setHasFs(data.hasFS);
            return data.hasFS;

        } catch (err) {
            console.error("Error al verificar el sistema de archivos:", err);
            // Si hay un error de red o de parsing, asumimos que no hay FS para permitir el formato.
            setHasFs(false);
            return false;
        } finally {
            setIsCheckingFs(false);
        }
    };


    // Función auxiliar para obtener un ícono basado en el tipo de partición
    const getPartitionIcon = (type: string) => {
        if (type.startsWith('Primaria')) return <Server className="w-5 h-5 text-purple-400" />;
        if (type.startsWith('Lógica')) return <Database className="w-5 h-5 text-green-400" />;
        return <HardDrive className="w-5 h-5 text-gray-400" />;
    };
    
    // Función auxiliar para formatear tamaño
    const formatPartitionSize = (bytes: number) => {
        if (bytes >= 1e12) return `${(bytes / 1e12).toFixed(2)} TB`;
        if (bytes >= 1e9) return `${(bytes / 1e9).toFixed(2)} GB`;
        if (bytes >= 1e6) return `${(bytes / 1e6).toFixed(2)} MB`;
        if (bytes >= 1e3) return `${(bytes / 1e3).toFixed(2)} KB`;
        if (bytes === 0) return '0 B';
        return `${bytes} B`;
    };
    
    // ⬇️ EFECTO 1: Carga inicial de particiones 
    useEffect(() => {
        if (!diskPath) {
            setError("No se proporcionó la ruta del disco.");
            setIsLoading(false);
            return;
        }

        const fetchPartitions = async () => {
            setIsLoading(true);
            setError(null);

            const API_URL = `/disk/partitions?diskPath=${encodeURIComponent(diskPath)}`;

            try {
                const response = await fetch(API_URL);
                if (!response.ok) {
                    throw new Error(`Error HTTP: ${response.status}`);
                }

                const data: { 
                    mbr_partitions: BackendPartition[] | null, 
                    ebr_partitions: BackendPartition[] | null 
                } = await response.json();

                const combinedPartitions: PartitionRecord[] = [];

                // 🔹 Asignador automático de IDs
                let counter = 1;

                const mbrPartitions = data.mbr_partitions ?? [];
                mbrPartitions.forEach(p => {
                    if (p.type === 'E') return;

                    combinedPartitions.push({
                        id: `${diskPath}-auto-${counter++}`,
                        name: p.name || `Partición_Primaria_${counter}`,
                        size: p.size,
                        type: 'Primaria',
                        used: `${p.usage}%`,
                    });
                });

                const ebrPartitions = data.ebr_partitions ?? [];
                ebrPartitions.forEach(p => {
                    combinedPartitions.push({
                        id: `${diskPath}-auto-${counter++}`,
                        name: p.name || `Partición_Lógica_${counter}`,
                        size: p.size,
                        type: 'Lógica',
                        used: `${p.usage}%`,
                    });
                });

                setPartitions(combinedPartitions);
                setIsLoading(false);

                if (combinedPartitions.length > 0) {
                    setSelectedPartition(combinedPartitions[0].id);
                } else {
                    setSelectedPartition(null);
                }

            } catch (e: any) {
                console.error("Error al cargar las particiones:", e);
                setError(`No se pudieron cargar las particiones. Detalles: ${e.message}`);
                setIsLoading(false);
            }
        };


        fetchPartitions();
    }, [diskPath]);

    // ⬇️ EFECTO 2: Verificar el sistema de archivos cuando cambia la partición seleccionada
    useEffect(() => {
        if (!diskPath || !selectedPartition) {
            setHasFs(null);
            return;
        }

        const partition = partitions.find(p => p.id === selectedPartition);
        if (partition) {
            checkPartitionFilesystem(diskPath, partition.name);
        }
    }, [diskPath, selectedPartition, partitions]);


    // Manejador del botón "Siguiente/Formatear"
    const handleNext = async (e: React.FormEvent) => {
        e.preventDefault();
        const partition = partitions.find(p => p.id === selectedPartition);
        
        // Bloquear si no hay selección, o si la verificación está en curso o no ha terminado
        if (!partition || !onSelect || isCheckingFs || hasFs === null) return;

        const partitionData: {
            id: string;
            name: string;
            type: 'Primaria' | 'Lógica';
            needsFormat: boolean;
            message?: string;
        } = {
            id: partition.id,
            name: partition.name,
            type: partition.type,
            // Agregamos una bandera para que el componente principal sepa a dónde ir
            needsFormat: !hasFs
        };

        const URL_API = `/disk/mounted?diskPath=${encodeURIComponent(diskPath || '')}&partitionName=${encodeURIComponent(partition.name)}`;
        try {
            const response = await fetch(URL_API);
            if (!response.ok) {
                throw new Error(`Error HTTP: ${response.status}`);
            }
            const data: {
                message: string;
                partitionId: string;
            } = await response.json();

            console.log(`Partición montada: ${data.partitionId}`);
            partitionData.id = data.partitionId;
        } catch (e: any) {
            console.error("Error al montar la partición:", e);
            // Continuamos aunque falle el montaje, ya que no es crítico para el flujo
            partitionData.id = `error-${Date.now()}`;
            partitionData.message = e.message;

            // Continuamos aunque falle el montaje, ya que no es crítico para el flujo
            partitionData.id = `error-${Date.now()}`;
            partitionData.message = e.message;

            console.error("Error al montar la partición:", e);
            // Continuamos aunque falle el montaje, ya que no es crítico para el flujo
            partitionData.id = `error-${Date.now()}`;
            partitionData.message = e.message;
        }
        
        // La función onSelect ahora decide ir a LOGIN o FORMAT_PARTITION
        onSelect(partitionData);
    };

    const handlePartitionClick = (partitionId: string) => {
        if (selectedPartition !== partitionId) {
            // 1. Actualiza la selección
            setSelectedPartition(partitionId);
            // 2. Prepara la UI para el nuevo chequeo
            setHasFs(null);
            setIsCheckingFs(true); 
            // El useEffect se encargará de hacer la llamada API
        }
    };
    
    // Renderizado condicional del contenido de la lista
    let content;

    if (isLoading) {
        content = (
            <div className="flex flex-col items-center justify-center h-48 text-gray-400">
                <Loader2 className="w-8 h-8 animate-spin text-purple-400" />
                <p className="mt-4">Cargando particiones...</p>
            </div>
        );
    } else if (error) {
        content = (
            <div className="text-red-400 text-center p-4">
                <p>{error}</p>
            </div>
        );
    } else if (partitions.length === 0) {
        content = (
            <div className="text-gray-400 text-center p-4">
                <p>El disco no contiene particiones utilizables.</p>
            </div>
        );
    } else {
        // Lista de particiones
        content = (
            <form onSubmit={handleNext} className="space-y-4">
                <div className="space-y-3 max-h-64 overflow-y-auto pr-2">
                    {partitions.map(partition => (
                        <div 
                            key={partition.id} 
                            onClick={() => { handlePartitionClick(partition.id)}} 
                            className={`flex flex-col p-4 cursor-pointer bg-[#3a3a3a] border rounded-md transition-all ${selectedPartition === partition.id ? 'border-purple-500 ring-2 ring-purple-500' : 'border-[#4a4a4a] hover:border-gray-500'}`}
                        >
                            <div className="flex justify-between items-center">
                                <div className="flex items-center gap-3">
                                    {getPartitionIcon(partition.type)}
                                    <div>
                                        <p className="font-medium text-gray-100">{partition.name}</p>
                                        <p className="text-xs text-gray-400">Tipo: {partition.type} | {formatPartitionSize(partition.size)}</p>
                                    </div>
                                </div>
                                {/* Indicador de Estado FS */}
                                {selectedPartition === partition.id && (
                                    <span className="text-sm flex items-center gap-2">
                                        {isCheckingFs ? (
                                            <>
                                                <Loader2 className="w-4 h-4 animate-spin text-purple-400" /> Verificando FS...
                                            </>
                                        ) : hasFs === true ? (
                                            <>
                                                <CheckCircle className="w-4 h-4 text-green-500" /> Listo
                                            </>
                                        ) : hasFs === false ? (
                                            <>
                                                <XCircle className="w-4 h-4 text-yellow-500" /> Requiere Formato
                                            </>
                                        ) : (
                                            <span className="text-gray-400">Seleccionada</span>
                                        )}
                                    </span>
                                )}
                            </div>
                            <div className="w-full bg-[#1e1e1e] rounded-full h-2.5 mt-2">
                                <div className="bg-purple-600 h-2.5 rounded-full" style={{ width: partition.used }} title={`Uso: ${partition.used}`}></div>
                            </div>
                        </div>
                    ))}
                </div>
                
                <div className="flex gap-4 pt-4">
                    <button type="button" onClick={onBack} className="w-1/3 py-3 px-4 bg-gray-500 text-white font-semibold rounded-md hover:bg-gray-600 transition-all flex items-center justify-center">
                        <ArrowLeft className="w-4 h-4 inline mr-2" /> Atrás
                    </button>
                    <button 
                        type="submit" 
                        disabled={!selectedPartition || isCheckingFs || hasFs === null} 
                        className={`w-2/3 py-3 px-4 text-white font-semibold rounded-md transition-all flex items-center justify-center gap-2 
                            ${!selectedPartition || isCheckingFs || hasFs === null ? 'bg-gray-500 cursor-not-allowed' :
                            hasFs ? 'bg-gradient-to-r from-purple-600 to-indigo-600 hover:from-purple-700 hover:to-indigo-700' :
                            'bg-gradient-to-r from-yellow-600 to-orange-600 hover:from-yellow-700 hover:to-orange-700'
                        }`}
                    >
                        {isCheckingFs ? (
                             <>
                                <Loader2 className="w-4 h-4 animate-spin" /> Verificando...
                             </>
                        ) : hasFs ? (
                            <>
                                Continuar al Login <ArrowRight className="w-4 h-4 inline ml-2" />
                            </>
                        ) : (
                            <>
                                Formatear Partición <ArrowRight className="w-4 h-4 inline ml-2" />
                            </>
                        )}
                    </button>
                </div>
            </form>
        );
    }

    // Estructura principal de la página
    return (
        <div className="flex w-full h-full bg-[#1e1e1e] text-gray-200 shadow-xl rounded-lg overflow-hidden">
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-gradient-to-br from-[#2c2c2c] via-[#1a1a1a] to-[#0d0d0d] text-center">
                <Database className="w-16 h-16 text-purple-400 mb-4" />
                <h1 className="text-4xl font-bold mb-4 text-white">Paso 2: Selección de Partición</h1>
                <p className="text-lg text-gray-300 max-w-sm">Seleccionaste el disco **{disk?.name || 'Desconocido'}**.</p>
                <div className="mt-4 p-2 bg-[#3a3a3a] rounded-md text-sm text-gray-400"><HardDrive className="inline w-4 h-4 mr-2" />Ruta: {disk?.id || 'N/A'}</div>
            </div>
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-[#2a2a2a]">
                <div className="w-full max-w-md">
                    <h2 className="text-2xl font-semibold mb-6 text-gray-100 text-center">ELIGE UNA PARTICIÓN</h2>
                    {content}
                </div>
            </div>
        </div>
    );
};

export default PartitionSelectionPage;
