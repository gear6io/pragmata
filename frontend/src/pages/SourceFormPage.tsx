import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCreateSource } from '../api/generated/services/sources';
import ErrorMessage from '../components/ErrorMessage';

type FieldRow = { name: string; type: string };

const FIELD_TYPES = [
  'string', 'int64', 'float64', 'bool', 'datetime64', 'date',
];

export default function SourceFormPage() {
  const navigate = useNavigate();
  const { mutate: createSource, isPending, error: createError } = useCreateSource();

  const [name, setName] = useState('');
  const [engine, setEngine] = useState('');
  const [fields, setFields] = useState<FieldRow[]>([]);

  function addField() {
    setFields((f) => [...f, { name: '', type: '' }]);
  }

  function removeField(i: number) {
    setFields((f) => f.filter((_, idx) => idx !== i));
  }

  function updateField(i: number, key: keyof FieldRow, value: string) {
    setFields((f) => f.map((row, idx) => (idx === i ? { ...row, [key]: value } : row)));
  }

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    createSource(
      { data: { name, engine, fields } },
      { onSuccess: () => navigate(`/sources/${name}`) },
    );
  }

  const inputClass =
    'w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent';

  return (
    <div className="p-8 max-w-2xl">
      <h1 className="text-2xl font-semibold text-gray-900 mb-8">New Source</h1>

      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Name</label>
          <input
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
            className={inputClass}
            placeholder="my_source"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Engine</label>
          <input
            value={engine}
            onChange={(e) => setEngine(e.target.value)}
            className={inputClass}
            placeholder="MergeTree"
          />
        </div>

        <div>
          <div className="flex items-center justify-between mb-2">
            <label className="text-sm font-medium text-gray-700">Fields</label>
            <button
              type="button"
              onClick={addField}
              className="text-indigo-600 hover:text-indigo-800 text-sm font-medium"
            >
              + Add field
            </button>
          </div>

          {fields.length === 0 && (
            <p className="text-sm text-gray-400 py-2">No fields added yet.</p>
          )}

          <div className="space-y-2">
            {fields.map((field, i) => (
              <div key={i} className="flex gap-2 items-center">
                <input
                  value={field.name}
                  onChange={(e) => updateField(i, 'name', e.target.value)}
                  className={inputClass}
                  placeholder="field_name"
                />
                <select
                  value={field.type}
                  onChange={(e) => updateField(i, 'type', e.target.value)}
                  className={inputClass}
                >
                  <option value="" disabled>type</option>
                  {FIELD_TYPES.map((t) => (
                    <option key={t} value={t}>{t}</option>
                  ))}
                </select>
                <button
                  type="button"
                  onClick={() => removeField(i)}
                  className="text-gray-400 hover:text-red-600 text-sm px-2 py-2 shrink-0 transition-colors"
                >
                  ✕
                </button>
              </div>
            ))}
          </div>
        </div>

        <ErrorMessage error={createError} />

        <div className="flex gap-3 pt-2">
          <button
            type="submit"
            disabled={isPending}
            className="bg-indigo-600 hover:bg-indigo-700 text-white px-5 py-2 rounded-lg text-sm font-medium transition-colors disabled:opacity-50"
          >
            {isPending ? 'Creating…' : 'Create Source'}
          </button>
          <button
            type="button"
            onClick={() => navigate('/sources')}
            className="border border-gray-300 text-gray-700 px-5 py-2 rounded-lg text-sm font-medium hover:bg-gray-50 transition-colors"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}
