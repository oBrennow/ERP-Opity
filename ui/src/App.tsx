import React from 'react';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import { Box } from '@mui/material';
import Sidebar from './components/Sidebar';
import Topbar from './components/Topbar';
import MainRouter from './components/MainRouter';
import { ThemeContextProvider, useThemeContext } from './theme/ThemeContext';

function AppContent() {
  const { themeMode } = useThemeContext();
  const theme = React.useMemo(() => createTheme({
    palette: {
      mode: themeMode,
      primary: { main: '#1976d2' },
      secondary: { main: '#009688' },
      background: {
        default: themeMode === 'dark' ? '#18191A' : '#f4f6f8',
        paper: themeMode === 'dark' ? '#242526' : '#fff',
      },
    },
  }), [themeMode]);

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box sx={{ display: 'flex', height: '100vh', width: '100vw', overflow: 'hidden' }}>
        <Sidebar />
        <Box sx={{ flex: 1, display: 'flex', flexDirection: 'column', minWidth: 0 }}>
          <Topbar />
          <Box component="main" sx={{ flex: 1, p: 2, overflow: 'auto' }}>
            <MainRouter />
          </Box>
        </Box>
      </Box>
    </ThemeProvider>
  );
}

export default function App() {
  return (
    <ThemeContextProvider>
      <AppContent />
    </ThemeContextProvider>
  );
}
