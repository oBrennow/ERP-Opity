import { AppBar, Toolbar, Typography, IconButton, Box } from '@mui/material';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import Brightness7Icon from '@mui/icons-material/Brightness7';
import NotificationsIcon from '@mui/icons-material/Notifications';
import LogoutIcon from '@mui/icons-material/Logout';
import { useThemeContext } from '../theme/ThemeContext';

export default function Topbar() {
  const { themeMode, toggleTheme } = useThemeContext();
  return (
    <AppBar position="static" color="default" elevation={1} sx={{ zIndex: 1201 }}>
      <Toolbar>
        <Typography variant="h6" sx={{ flexGrow: 1 }}>
          ERP Opity
        </Typography>
        <IconButton color="inherit" onClick={toggleTheme} size="large">
          {themeMode === 'dark' ? <Brightness7Icon /> : <Brightness4Icon />}
        </IconButton>
        <IconButton color="inherit" size="large">
          <NotificationsIcon />
        </IconButton>
        <Box sx={{ ml: 2, display: 'flex', alignItems: 'center' }}>
          <Typography variant="body1" sx={{ mr: 1 }}>Usuário</Typography>
          <IconButton color="inherit" size="large">
            <LogoutIcon />
          </IconButton>
        </Box>
      </Toolbar>
    </AppBar>
  );
} 