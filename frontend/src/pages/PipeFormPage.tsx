import { useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { EditorState } from '@codemirror/state';
import { EditorView } from '@codemirror/view';
import { sql } from '@codemirror/lang-sql';
import {
  useCreatePipe,
  useGetPipe,
  useUpdatePipe,
  invalidateListPipes,
  invalidateGetPipe,
} from '../api/generated/services/pipes';
import type { PipetypesPipeDTO } from '../api/generated/services/pragmataAPI.schemas';
import { useQueryClient } from '@tanstack/react-query';

const PIPE_TYPES = [
  'ENDPOINT',
  'MATERIALIZED',
  'COPY',
  'TABLE',
  'VIEW',
  'INCREMENTAL',
  'SNAPSHOT',
] as const;

function CodeEditor({
  value,
  onChange,
}: {
  value: string;
  onChange: (v: string) => void;
}) {
  const containerRef = useRef<HTMLDivElement>(null);
  const viewRef = useRef<EditorView | null>(null);

  useEffect(() => {
    if (!containerRef.current) return;
    const view = new EditorView({
      state: EditorState.create({
        doc: value,
        extensions: [
          sql(),
          EditorView.lineWrapping,
          EditorView.updateListener.of((update) => {
            if (update.docChanged) onChange(update.state.doc.toString());
          }),
        ],
      }),
      parent: containerRef.current,
    });
    viewRef.current = view;
    return () => {
      view.destroy();
      viewRef.current = null;
    };
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <div
      ref={containerRef}
      className="border border-gray-300 rounded-lg overflow-hidden text-sm min-h-48 focus-within:ring-2 focus-within:ring-indigo-500 focus-within:border-indigo-500"
    />
  );
}

interface Props {
  mode: 'create' | 'edit';
}

export default function PipeFormPage({ mode }: Props) {
  const { name } = useParams<{ name: string }>();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const { data: existing, isLoading } = useGetPipe(
    { name: name! },
    { query: { enabled: mode === 'edit' && !!name } },
  );

  const [formName, setFormName] = useState('');
  const [formType, setFormType] = useState<string>(PIPE_TYPES[0]);
  const [formDescription, setFormDescription] = useState('');
  const [formTags, setFormTags] = useState('');
  const [formContent, setFormContent] = useState('');

  useEffect(() => {
    if (existing) {
      setFormName(existing.name ?? '');
      setFormType(existing.type ?? PIPE_TYPES[0]);
      setFormDescription(existing.description ?? '');
      setFormTags((existing.tags ?? []).join(', '));
      setFormContent(existing.content ?? '');
    }
  }, [existing]);

  const { mutate: createPipe, isPending: isCreating, error: createError } = useCreatePipe();
  const { mutate: updatePipe, isPending: isUpdating, error: updateError } = useUpdatePipe();

  const isPending = isCreating || isUpdating;
  const mutationError = createError ?? updateError;

  function buildPayload(): PipetypesPipeDTO {
    const tags = formTags
      .split(',')
      .map((t) => t.trim())
      .filter(Boolean);
    return {
      id: existing?.id ?? '',
      name: formName.trim(),
      type: formType,
      description: formDescription.trim() || undefined,
      tags: tags.length ? tags : undefined,
      content: formContent,
    };
  }

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    const payload = buildPayload();
    if (mode === 'create') {
      createPipe(
        { data: payload },
        {
          onSuccess: (created) => {
            invalidateListPipes(queryClient);
            navigate(`/pipes/${created.name}`);
          },
        },
      );
    } else {
      updatePipe(
        { pathParams: { name: name! }, data: payload },
        {
          onSuccess: (updated) => {
            invalidateListPipes(queryClient);
            invalidateGetPipe(queryClient, { name: name! });
            navigate(`/pipes/${updated.name}`);
          },
        },
      );
    }
  }

  if (mode === 'edit' && isLoading) {
    return <div className="p-8 text-sm text-gray-500">Loading...</div>;
  }

  return (
    <div className="p-8 max-w-3xl">
      <h1 className="text-2xl font-semibold text-gray-900 mb-6">
        {mode === 'create' ? 'New Pipe' : `Edit ${name}`}
      </h1>

      <form onSubmit={handleSubmit} className="space-y-5">
        {/* Name */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Name</label>
          <input
            type="text"
            required
            disabled={mode === 'edit'}
            value={formName}
            onChange={(e) => setFormName(e.target.value)}
            placeholder="my_pipe"
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 disabled:bg-gray-100 disabled:text-gray-500"
          />
        </div>

        {/* Type */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Type</label>
          <select
            value={formType}
            onChange={(e) => setFormType(e.target.value)}
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 bg-white"
          >
            {PIPE_TYPES.map((t) => (
              <option key={t} value={t}>
                {t}
              </option>
            ))}
          </select>
        </div>

        {/* Description */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Description</label>
          <input
            type="text"
            value={formDescription}
            onChange={(e) => setFormDescription(e.target.value)}
            placeholder="What this pipe does"
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>

        {/* Tags */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Tags</label>
          <input
            type="text"
            value={formTags}
            onChange={(e) => setFormTags(e.target.value)}
            placeholder="tag1, tag2, tag3"
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
          <p className="mt-1 text-xs text-gray-400">Comma-separated</p>
        </div>

        {/* Content */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Content</label>
          <CodeEditor value={formContent} onChange={setFormContent} />
        </div>

        {mutationError && (
          <p className="text-sm text-red-600">
            {(mutationError as Error).message ?? 'An error occurred.'}
          </p>
        )}

        <div className="flex gap-3 pt-2">
          <button
            type="submit"
            disabled={isPending}
            className="bg-indigo-600 hover:bg-indigo-700 text-white px-5 py-2 rounded-lg text-sm font-medium transition-colors disabled:opacity-50"
          >
            {isPending ? 'Saving...' : mode === 'create' ? 'Create Pipe' : 'Save Changes'}
          </button>
          <button
            type="button"
            onClick={() => navigate(mode === 'edit' ? `/pipes/${name}` : '/pipes')}
            className="border border-gray-300 hover:bg-gray-50 text-gray-700 px-5 py-2 rounded-lg text-sm font-medium transition-colors"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}
