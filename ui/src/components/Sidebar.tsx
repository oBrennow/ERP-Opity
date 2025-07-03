import { Drawer, List, ListItem, ListItemIcon, ListItemText, Toolbar } from '@mui/material';
import PeopleIcon from '@mui/icons-material/People';
import InventoryIcon from '@mui/icons-material/Inventory';
import PointOfSaleIcon from '@mui/icons-material/PointOfSale';
import AssessmentIcon from '@mui/icons-material/Assessment';
import SettingsIcon from '@mui/icons-material/Settings';
import { useNavigate } from 'react-router-dom';

const drawerWidth = 220;

const menu = [
  { text: 'Clientes', icon: <PeopleIcon />, path: '/clientes' },
  { text: 'Produtos', icon: <InventoryIcon />, path: '/produtos' },
  { text: 'Vendas', icon: <PointOfSaleIcon />, path: '/vendas' },
  { text: 'Relatórios', icon: <AssessmentIcon />, path: '/relatorios' },
  { text: 'Configurações', icon: <SettingsIcon />, path: '/configuracoes' },
];

export default function Sidebar() {
  const navigate = useNavigate();
  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box', background: 'inherit' },
      }}
    >
      <Toolbar />
      <List>
        {menu.map((item) => (
          <ListItem component="button" key={item.text} onClick={() => navigate(item.path)}>
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.text} />
          </ListItem>
        ))}
      </List>
    </Drawer>
  );
} 