import { FC } from 'react';

interface ExcalidrawWrapperProps {
  initialData?: {
    elements?: any[];
    appState?: any;
    files?: any;
  } | null;
  onChange?: (elements: any[], appState: any, files: any) => void;
}

export const ExcalidrawWrapper: FC<ExcalidrawWrapperProps>;
