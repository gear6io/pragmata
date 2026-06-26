import { AxiosError } from 'axios';

function extractMessage(error: unknown): string {
  const axiosErr = error as AxiosError<{ error?: string }>;
  return (
    axiosErr.response?.data?.error ??
    (error as Error).message ??
    'An error occurred.'
  );
}

export default function ErrorMessage({ error }: { error: unknown }) {
  if (!error) return null;
  return (
    <div role="alert" className="rounded-lg bg-red-50 border border-red-200 px-4 py-3 text-sm text-red-700 flex gap-2 items-start">
      <span className="shrink-0 font-bold">Error:</span>
      <span>{extractMessage(error)}</span>
    </div>
  );
}
