import { Link } from 'react-router-dom';
import { useListSources } from '../api/generated/services/sources';

export default function SourcesPage() {
  const { data: resp, isLoading, error } = useListSources();
  const data = resp?.data ?? [];

  return (
    <div className="p-8">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold text-gray-900">Sources</h1>
        <Link
          to="/sources/new"
          className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors"
        >
          New Source
        </Link>
      </div>

      {isLoading && <p className="text-sm text-gray-500">Loading...</p>}
      {error && <p className="text-sm text-red-600">Failed to load sources.</p>}

      {!isLoading && !error && (
        <div className="bg-white rounded-xl border border-gray-200 overflow-hidden">
          <table className="w-full text-sm">
            <thead className="bg-gray-50 border-b border-gray-200">
              <tr>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Name</th>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Engine</th>
                <th className="text-left px-5 py-3 font-medium text-gray-500">Fields</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {data.length === 0 && (
                <tr>
                  <td colSpan={3} className="px-5 py-8 text-center text-gray-400">
                    No sources yet.{' '}
                    <Link to="/sources/new" className="text-indigo-600 hover:underline">
                      Create one.
                    </Link>
                  </td>
                </tr>
              )}
              {data.map((source) => (
                <tr key={source.name} className="hover:bg-gray-50 transition-colors">
                  <td className="px-5 py-3 font-medium text-indigo-700">
                    <Link to={`/sources/${source.name}`} className="hover:underline">
                      {source.name}
                    </Link>
                  </td>
                  <td className="px-5 py-3">
                    <span className="inline-block bg-gray-100 text-gray-600 px-2 py-0.5 rounded text-xs font-mono">
                      {source.engine ?? '—'}
                    </span>
                  </td>
                  <td className="px-5 py-3 text-gray-600">{source.fields?.length ?? 0}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
