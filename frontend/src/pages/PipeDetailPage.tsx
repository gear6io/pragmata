import { useNavigate, useParams } from 'react-router-dom';
import { useGetPipe, useDeletePipe } from '../api/generated/services/pipes';
import { useMemo } from 'react';
import { EditorState } from '@codemirror/state';
import { EditorView } from '@codemirror/view';
import { sql } from '@codemirror/lang-sql';
import { useEffect, useRef } from 'react';

function ReadOnlyEditor({ value }: { value: string }) {
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!containerRef.current) return;
    const view = new EditorView({
      state: EditorState.create({
        doc: value,
        extensions: [sql(), EditorView.editable.of(false), EditorView.lineWrapping],
      }),
      parent: containerRef.current,
    });
    return () => view.destroy();
  }, [value]);

  return <div ref={containerRef} className="border border-gray-200 rounded-lg overflow-hidden text-sm" />;
}

function Badge({ label }: { label: string }) {
  return (
    <span className="inline-block bg-indigo-50 text-indigo-700 px-2 py-0.5 rounded text-xs font-mono">
      {label}
    </span>
  );
}

function Field({ label, value }: { label: string; value?: string | null }) {
  if (!value) return null;
  return (
    <div>
      <dt className="text-xs font-medium text-gray-500 uppercase tracking-wide mb-0.5">{label}</dt>
      <dd className="text-sm text-gray-800">{value}</dd>
    </div>
  );
}

export default function PipeDetailPage() {
  const { name } = useParams<{ name: string }>();
  const navigate = useNavigate();
  const { data: pipe, isLoading, error } = useGetPipe({ name: name! });
  const { mutate: deletePipe, isPending: isDeleting } = useDeletePipe();

  function handleDelete() {
    if (!confirm(`Delete pipe "${name}"?`)) return;
    deletePipe(
      { pathParams: { name: name! } },
      { onSuccess: () => navigate('/pipes') },
    );
  }

  if (isLoading) return <div className="p-8 text-sm text-gray-500">Loading...</div>;
  if (error || !pipe) return <div className="p-8 text-sm text-red-600">Pipe not found.</div>;

  return (
    <div className="p-8 max-w-4xl">
      {/* Header */}
      <div className="flex items-start justify-between mb-8">
        <div>
          <h1 className="text-2xl font-semibold text-gray-900 mb-1">{pipe.name}</h1>
          {pipe.type && <Badge label={pipe.type} />}
        </div>
        <div className="flex gap-2">
          <button
            onClick={() => navigate(`/pipes/${name}/edit`)}
            className="border border-gray-300 hover:border-indigo-400 text-gray-700 hover:text-indigo-700 px-4 py-2 rounded-lg text-sm font-medium transition-colors"
          >
            Edit
          </button>
          <button
            onClick={handleDelete}
            disabled={isDeleting}
            className="bg-red-50 hover:bg-red-100 text-red-700 px-4 py-2 rounded-lg text-sm font-medium transition-colors disabled:opacity-50"
          >
            Delete
          </button>
        </div>
      </div>

      {/* Metadata */}
      <div className="bg-white border border-gray-200 rounded-xl p-6 mb-6">
        <h2 className="text-sm font-semibold text-gray-700 mb-4">Metadata</h2>
        <dl className="grid grid-cols-2 gap-4">
          <Field label="Description" value={pipe.description} />
          <Field label="Datasource" value={pipe.datasource} />
          <Field label="Target datasource" value={pipe.targetDatasource} />
          <Field label="Copy schedule" value={pipe.copySchedule} />
          <Field label="Created by" value={pipe.createdBy} />
          <Field
            label="Created at"
            value={pipe.createdAt ? new Date(pipe.createdAt).toLocaleString() : undefined}
          />
          <Field label="Updated by" value={pipe.updatedBy} />
          <Field
            label="Updated at"
            value={pipe.updatedAt ? new Date(pipe.updatedAt).toLocaleString() : undefined}
          />
        </dl>

        {pipe.tags && pipe.tags.length > 0 && (
          <div className="mt-4">
            <dt className="text-xs font-medium text-gray-500 uppercase tracking-wide mb-1">Tags</dt>
            <div className="flex flex-wrap gap-1">
              {pipe.tags.map((tag) => (
                <span
                  key={tag}
                  className="bg-gray-100 text-gray-600 text-xs px-2 py-0.5 rounded"
                >
                  {tag}
                </span>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Content */}
      {pipe.content && (
        <div className="bg-white border border-gray-200 rounded-xl p-6">
          <h2 className="text-sm font-semibold text-gray-700 mb-3">Content</h2>
          <ReadOnlyEditor value={pipe.content} />
        </div>
      )}
    </div>
  );
}
