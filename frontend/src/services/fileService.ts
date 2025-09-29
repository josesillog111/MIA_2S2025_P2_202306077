// src/services/fileService.ts
export interface EditorFile {
  name: string;
  path: string;
  content: string;
  handle?: FileSystemFileHandle;
}

function downloadFallback(file: EditorFile) {
  const blob = new Blob([file.content], { type: "text/plain" });
  const url = URL.createObjectURL(blob);

  const a = document.createElement("a");
  a.href = url;
  a.download = file.name;
  a.click();

  URL.revokeObjectURL(url);
}

function openFilesFallback(): Promise<EditorFile[]> {
  return new Promise((resolve) => {
    const input = document.createElement("input");
    input.type = "file";
    input.multiple = true;
    input.onchange = async (e: any) => {
      const files: EditorFile[] = [];
      for (const file of e.target.files) {
        const content = await file.text();
        files.push({
          name: file.name,
          path: file.name,
          content,
        });
      }
      resolve(files);
    };
    input.click();
  });
}

export async function saveFile(file: EditorFile): Promise<EditorFile> {
  if (!file) throw new Error("No hay archivo activo para guardar");

  try {
    if ("showSaveFilePicker" in window) {
      if (file.handle) {
        const writable = await file.handle.createWritable();
        await writable.write(file.content);
        await writable.close();
        return file;
      } else {
        const opts = {
          suggestedName: file.name,
          types: [
            {
              description: "Archivos de texto",
              accept: { "text/plain": [".txt", ".ts", ".tsx", ".md"] },
            },
          ],
        };

        const fileHandle = await (window as any).showSaveFilePicker(opts);
        const writable = await fileHandle.createWritable();
        await writable.write(file.content);
        await writable.close();

        return { ...file, path: fileHandle.name, handle: fileHandle };
      }
    } else {
      // 🔄 Fallback → descarga del archivo
      downloadFallback(file);
      return file;
    }
  } catch (err) {
    throw new Error(
      `Error al guardar: ${err instanceof Error ? err.message : String(err)}`
    );
  }
}

export async function openFolder(): Promise<EditorFile[]> {
  try {
    if ("showDirectoryPicker" in window) {
      const dirHandle = await (window as any).showDirectoryPicker();
      const files: EditorFile[] = [];

      for await (const [name, handle] of dirHandle.entries()) {
        if (handle.kind === "file") {
          const file = await handle.getFile();
          const content = await file.text();

          files.push({
            name,
            path: file.name,
            content,
            handle,
          });
        }
      }
      return files;
    } else {
      return await openFilesFallback();
    }
  } catch (err) {
    throw new Error(
      `Error al abrir carpeta: ${err instanceof Error ? err.message : String(err)}`
    );
  }
}