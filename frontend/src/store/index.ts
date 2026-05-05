import { create } from 'zustand';

interface AppState {
  versions: any[];
  evalReports: any[];
  alerts: any[];
  costSummary: any;
  setVersions: (versions: any[]) => void;
  setEvalReports: (reports: any[]) => void;
  setAlerts: (alerts: any[]) => void;
  setCostSummary: (summary: any) => void;
}

export const useAppStore = create<AppState>((set) => ({
  versions: [],
  evalReports: [],
  alerts: [],
  costSummary: null,
  setVersions: (versions) => set({ versions }),
  setEvalReports: (evalReports) => set({ evalReports }),
  setAlerts: (alerts) => set({ alerts }),
  setCostSummary: (costSummary) => set({ costSummary }),
}));