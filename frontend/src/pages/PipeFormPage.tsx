import { useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { EditorState } from '@codemirror/state';
import { EditorView, keymap } from '@codemirror/view';
import { indentWithTab } from '@codemirror/commands';
import { autocompletion } from '@codemirror/autocomplete';
import { pipeLangCompletions } from '../editor/pipeLangCompletions';
import {
  useCreatePipe,
  useGetPipe,
  useUpdatePipe,
  invalidateListPipes,
  invalidateGetPipe,
} from '../api/generated/services/pipes';
import { useQueryClient } from '@tanstack/react-query';
import ErrorMessage from '../components/ErrorMessage';

const STARTER_TEMPLATE = `type: ENDPOINT
name: my_pipe

pipeline:
  @query:
    from table_name
`;

function CodeEditor({
  value,
  onChange,
}: {
  value: string;
  onChange: (v: string) => void;
}) {
  const containerRef = useRef<HTMLDivElement>(null);
  const viewRef = useRef<EditorView | null>(null);
  // Always-current doc text ref for the completion source callback.
  const docRef = useRef(value);

  useEffect(() => {
    const view = viewRef.current;
    if (view && value !== view.state.doc.toString()) {
      view.dispatch({
        changes: { from: 0, to: view.state.doc.length, insert: value },
      });
    }
  }, [value]);

  useEffect(() => {
    if (!containerRef.current) return;
    const view = new EditorView({
      state: EditorState.create({
        doc: value,
        extensions: [
          keymap.of([indentWithTab]),
          EditorView.lineWrapping,
          EditorView.updateListener.of((update) => {
            if (update.docChanged) {
              const text = update.state.doc.toString();
              docRef.current = text;
              onChange(text);
            }
          }),
          autocompletion({
            override: [pipeLangCompletions(() => docRef.current)],
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
      className="border border-gray-300 rounded-lg overflow-hidden text-sm flex-1 focus-within:ring-2 focus-within:ring-indigo-500 focus-within:border-indigo-500"
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

  const { data: existingResp, isLoading } = useGetPipe(
    { name: name! },
    { query: { enabled: mode === 'edit' && !!name } },
  );
  const existing = existingResp?.data;

  const [editorText, setEditorText] = useState(STARTER_TEMPLATE);

  useEffect(() => {
    if (existing) {
      setEditorText(existing.content ?? '');
    }
  }, [existing]);

  const { mutate: createPipe, isPending: isCreating, error: createError } = useCreatePipe();
  const { mutate: updatePipe, isPending: isUpdating, error: updateError } = useUpdatePipe();

  const isPending = isCreating || isUpdating;
  const mutationError = createError ?? updateError;

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    const payload = { content: editorText };
    if (mode === 'create') {
      createPipe(
        { data: payload },
        {
          onSuccess: (resp) => {
            invalidateListPipes(queryClient);
            navigate(`/pipes/${resp.data?.name}`);
          },
        },
      );
    } else {
      updatePipe(
        { pathParams: { name: name! }, data: payload },
        {
          onSuccess: (resp) => {
            invalidateListPipes(queryClient);
            invalidateGetPipe(queryClient, { name: name! });
            navigate(`/pipes/${resp.data?.name ?? name}`);
          },
        },
      );
    }
  }

  if (mode === 'edit' && isLoading) {
    return <div className="p-8 text-sm text-gray-500">Loading...</div>;
  }

  return (
    <div className="p-8 flex flex-col h-full">
      <h1 className="text-2xl font-semibold text-gray-900 mb-6">
        {mode === 'create' ? 'New Pipe' : `Edit ${name}`}
      </h1>

      <form onSubmit={handleSubmit} className="flex flex-col flex-1 gap-4">
        <CodeEditor value={editorText} onChange={setEditorText} />

        <ErrorMessage error={mutationError} />

        <div className="flex gap-3">
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
