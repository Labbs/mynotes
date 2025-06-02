import { documentApi } from '../../../../api/document';

// Garde une référence des bibliothèques en cache
let cachedLibraries: any[] | null = null;

export async function loadLibraries() {
  // Si on a déjà chargé les bibliothèques, les retourner du cache
  if (cachedLibraries) {
    return cachedLibraries;
  }

  const libraries = [];
  
  try {
    // Récupérer la liste des bibliothèques disponibles depuis l'API
    const { data: libNames } = await documentApi.listExcalidrawLibs();

    if (!Array.isArray(libNames)) {
      console.error('Invalid library names format:', libNames);
      return [];
    }
    
    // Charger chaque bibliothèque listée par le serveur
    for (const name of libNames) {
      try {
        // Charger le contenu de la bibliothèque
        const response = await documentApi.getExcalidrawLib(name+ '.excalidrawlib');
        const content = response.data;
        
        libraries.push({
          name,
          status: "published",
          libraryItems: content.libraryItems || []
        });
      } catch (error) {
        console.error(`Error loading library ${name}:`, error);
      }
    }
    
    // Mettre en cache les bibliothèques chargées
    cachedLibraries = libraries;
    return libraries;
  } catch (error) {
    console.error('Failed to fetch library list:', error);
    return [];
  }
}
