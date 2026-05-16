import { Link, useNavigate } from 'react-router-dom';
import { useListPipes, useDeletePipe, invalidateListPipes } from '../api/generated/services/pipes';
import { useQueryClient } from '@tanstack/react-query';

export default function PipesPage() {
  const { data, isLoading, error } = useListPipes();
  const { mutate: deletePipe, isPending: isDeleting } = useDeletePipe();
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  function handleDelete(name: string) {
    if (!confirm(`Delete pipe "${name}"?`)) return;
    deletePipe(
      { pathParams: { name } },
      { onSuccess: () => invalidateListPipes(queryClient) },
    );
  }

  return (
    <div className="p-8">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold text-gray-900">Pipes</h1>
        <Link
          to="/pipes/new"
          className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors"
        >
          New Pipe
        </Link>
      </div>

      {isLoading && <p className="text-sm text-gray-500">Loading...</p>}
      {error && <p className="text-sm text-red-600">Failed to load pipes.</p>}

      {data && (
        <div className="bg-white rounded-xl border border-gray-200 overflow-hidden">
          <table className="w-full text-sm">
            <thead className="bg-gray-50 border-b border-gray-200">
              <tr>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Name</th>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Type</th>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Description</th>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Created</th>
                <th className="px-5 py-3" />
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {data.length === 0 && (
                <tr>
                  <td colSpan={5} className="px-5 py-8 text-center text-gray-400">
                    No pipes yet.{' '}
                    <Link to="/pipes/new" className="text-indigo-600 hover:underline">
                      Create one.
                    </Link>
                  </td>
                </tr>
              )}
              {data.map((pipe) => (
                <tr key={pipe.id} className="hover:bg-gray-50 transition-colors">
                  <td className="px-5 py-3 font-medium text-indigo-700">
                    <Link to={`/pipes/${pipe.name}`} className="hover:underline">
                      {pipe.name}
                    </Link>
                  </td>
                  <td className="px-5 py-3">
                    <span className="inline-block bg-gray-100 text-gray-600 px-2 py-0.5 rounded text-xs font-mono">
                      {pipe.type}
                    </span>
                  </td>
                  <td className="px-5 py-3 text-gray-600">{pipe.description ?? '—'}</td>
                  <td className="px-5 py-3 text-gray-400">
                    {pipe.createdAt ? new Date(pipe.createdAt).toLocaleDateString() : '—'}
                  </td>
                  <td className="px-5 py-3">
                    <div className="flex gap-2 justify-end">
                      <button
                        onClick={() => navigate(`/pipes/${pipe.name}/edit`)}
                        className="text-gray-500 hover:text-indigo-600 text-xs px-2 py-1 rounded hover:bg-indigo-50 transition-colors"
                      >
                        Edit
                      </button>
                      <button
                        onClick={() => handleDelete(pipe.name!)}
                        disabled={isDeleting}
                        className="text-gray-500 hover:text-red-600 text-xs px-2 py-1 rounded hover:bg-red-50 transition-colors disabled:opacity-50"
                      >
                        Delete
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
