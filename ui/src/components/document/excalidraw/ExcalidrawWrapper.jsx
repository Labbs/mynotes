import React, { useCallback, useMemo } from 'react';
import { Excalidraw } from '@excalidraw/excalidraw';

export const ExcalidrawWrapper = ({ initialData, onChange }) => {
  // Mémoiser le handler pour éviter les re-renders inutiles
  const onChangeHandler = useCallback((elements, appState, files) => {
    onChange(elements, appState, files);
  }, [onChange]);

  // Mémoiser les options UI pour éviter les re-renders inutiles
  const uiOptions = useMemo(() => ({
    canvasActions: {
      export: false,
      changeViewBackgroundColor: false,
      clearCanvas: false,
      loadScene: false,
      saveToActiveFile: false,
      theme: false,
      saveAsImage: false,
    },
    tools: {
      image: true,
      hand: true,
      lock: true,
    }
  }), []);

  return (
    <Excalidraw
      initialData={initialData}
      onChange={onChangeHandler}
      gridModeEnabled={true}
      zenModeEnabled={false}
      viewModeEnabled={false}
      defaultSidebarDockedPreference="docked"
      renderTopRightUI={null}
      detectScroll={false}
      handleKeyboardGlobally={false}
      autoFocus={false}
      UIOptions={uiOptions}
    />
  );
};


