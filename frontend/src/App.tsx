import { Navigate, Route, Routes } from 'react-router-dom';
import Layout from './components/Layout';
import PipesPage from './pages/PipesPage';
import PipeDetailPage from './pages/PipeDetailPage';
import PipeFormPage from './pages/PipeFormPage';
import SourcesPage from './pages/SourcesPage';
import SourceDetailPage from './pages/SourceDetailPage';
import SourceFormPage from './pages/SourceFormPage';

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/pipes" replace />} />
      <Route element={<Layout />}>
        <Route path="/pipes" element={<PipesPage />} />
        <Route path="/pipes/new" element={<PipeFormPage mode="create" />} />
        <Route path="/pipes/:name" element={<PipeDetailPage />} />
        <Route path="/pipes/:name/edit" element={<PipeFormPage mode="edit" />} />
        <Route path="/sources" element={<SourcesPage />} />
        <Route path="/sources/new" element={<SourceFormPage />} />
        <Route path="/sources/:name" element={<SourceDetailPage />} />
      </Route>
    </Routes>
  );
}
