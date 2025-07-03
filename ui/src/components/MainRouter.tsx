import { Routes, Route, Navigate } from 'react-router-dom';
import ClientesPage from '../pages/ClientesPage';
import ProdutosPage from '../pages/ProdutosPage';
import VendasPage from '../pages/VendasPage';
import RelatoriosPage from '../pages/RelatoriosPage';
import ConfiguracoesPage from '../pages/ConfiguracoesPage';

export default function MainRouter() {
  return (
    <Routes>
      <Route path="/clientes" element={<ClientesPage />} />
      <Route path="/produtos" element={<ProdutosPage />} />
      <Route path="/vendas" element={<VendasPage />} />
      <Route path="/relatorios" element={<RelatoriosPage />} />
      <Route path="/configuracoes" element={<ConfiguracoesPage />} />
      <Route path="*" element={<Navigate to="/clientes" />} />
    </Routes>
  );
} 