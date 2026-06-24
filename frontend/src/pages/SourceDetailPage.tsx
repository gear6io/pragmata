import { useParams } from 'react-router-dom';
import { useGetSource } from '../api/generated/services/sources';

export default function SourceDetailPage() {
  const { name } = useParams<{ name: string }>();
  const { data: source, isLoading, error } = useGetSource({ name: name! });

  if (isLoading) return <div className="p-8 text-sm text-gray-500">Loading...</div>;
  if (error || !source) return <div className="p-8 text-sm text-red-600">Source not found.</div>;

  return (
    <div className="p-8 max-w-4xl">
      <div className="mb-8">
        <h1 className="text-2xl font-semibold text-gray-900 mb-1">{source.name}</h1>
        {source.engine && (
          <span className="inline-block bg-indigo-50 text-indigo-700 px-2 py-0.5 rounded text-xs font-mono">
            {source.engine}
          </span>
        )}
      </div>

      <div className="bg-white border border-gray-200 rounded-xl p-6">
        <h2 className="text-sm font-semibold text-gray-700 mb-4">Fields</h2>
        {(!source.fields || source.fields.length === 0) ? (
          <p className="text-sm text-gray-400">No fields defined.</p>
        ) : (
          <table className="w-full text-sm">
            <thead className="border-b border-gray-200">
              <tr>
                <th className="text-left pb-2 font-medium text-gray-500">Name</th>
                <th className="text-left pb-2 font-medium text-gray-500">Type</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {source.fields.map((field, i) => (
                <tr key={i}>
                  <td className="py-2 font-mono text-gray-800">{field.name}</td>
                  <td className="py-2 text-gray-500">{field.type}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
}
