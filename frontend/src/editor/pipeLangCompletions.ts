import { CompletionContext, CompletionResult, CompletionSource } from '@codemirror/autocomplete';
import { getSuggestions } from '../api/generated/services/suggestions';
import type { SuggestiontypesSuggestionResponseDTO } from '../api/generated/services/pragmataAPI.schemas';

interface DetectedContext {
  type: 'source' | 'field';
  searchText: string;
  nodeRef?: string;
  /** Start position of the word being completed */
  from: number;
}

/**
 * Inspects the document text above and at the cursor to determine what kind
 * of completion is appropriate for the current position.
 *
 * PipeLang sections of interest:
 *   sources:          → SourceContext (suggest pipe/table names)
 *   pipeline:
 *     @node: |        → FieldContext (suggest column names inside SQL)
 */
function detectContext(ctx: CompletionContext): DetectedContext | null {
  const doc = ctx.state.doc.toString();
  const pos = ctx.pos;
  const textBefore = doc.slice(0, pos);

  // Determine the word (identifier fragment) the user is currently typing.
  const wordMatch = textBefore.match(/[\w.]*$/);
  const word = wordMatch ? wordMatch[0] : '';
  const from = pos - word.length;

  // Walk lines above cursor to detect which top-level section we're in.
  const lines = textBefore.split('\n');
  let inSources = false;
  let inPipeline = false;

  for (const line of lines) {
    const trimmed = line.trimStart();
    if (/^sources\s*:/.test(trimmed)) {
      inSources = true;
      inPipeline = false;
    } else if (/^pipeline\s*:/.test(trimmed)) {
      inPipeline = true;
      inSources = false;
    } else if (/^\S/.test(trimmed) && !trimmed.startsWith('#')) {
      // Any other top-level key resets context.
      inSources = false;
      inPipeline = false;
    }
  }

  if (inSources) {
    return { type: 'source', searchText: word, from };
  }

  if (inPipeline) {
    // ponytail: source branch added here; columns branch goes here next
    // lines[last] is current line text up to the cursor — anchor to it so
    // `from` anywhere else in the doc can't trigger source completions.
    const currentLine = lines[lines.length - 1];
    if (/^\s*from\s+\w*$/i.test(currentLine)) {
      return { type: 'source', searchText: word, from };
    }
    // Extract the FROM clause alias closest before the cursor.
    const nodeRef = extractFromAlias(textBefore);
    return { type: 'field', searchText: word, nodeRef, from };
  }

  return null;
}

/**
 * Scans backwards through textBefore to find the most recent
 * `FROM <alias>` or `from <alias>` at the same SQL level.
 */
function extractFromAlias(text: string): string | undefined {
  // Match the last occurrence of FROM <identifier> before the cursor.
  const match = text.match(/\bfrom\s+([\w]+)\s*$/i);
  if (match) return match[1];

  // Also handle multiline: look for the last FROM anywhere in the current node block.
  const allFroms = [...text.matchAll(/\bfrom\s+([\w]+)/gi)];
  if (allFroms.length > 0) {
    return allFroms[allFroms.length - 1][1];
  }
  return undefined;
}

/**
 * Returns a CodeMirror CompletionSource that calls the backend suggestion API.
 *
 * @param getPipeContent - Callback that returns the current full editor text.
 *   Called lazily at completion time so it always reflects the latest state.
 */
export function pipeLangCompletions(getPipeContent: () => string): CompletionSource {
  return async (ctx: CompletionContext): Promise<CompletionResult | null> => {
    // Require an explicit trigger (Ctrl+Space) or an alphanumeric character.
    if (!ctx.explicit && !ctx.matchBefore(/\w/)) return null;

    const context = detectContext(ctx);
    if (!context) return null;

    // ponytail: GeneratedAPIInstance unwraps {status,data} at runtime; cast to inner type
    let resp: SuggestiontypesSuggestionResponseDTO;
    try {
      resp = await getSuggestions({
        contextType: context.type,
        matchingType: 'fuzzy',
        searchText: context.searchText,
        nodeRef: context.nodeRef,
        pipeContent: context.type === 'field' ? getPipeContent() : undefined,
      }) as unknown as SuggestiontypesSuggestionResponseDTO;
    } catch {
      return null;
    }

    if (!resp.suggestions?.length) return null;

    return {
      from: context.from,
      options: resp.suggestions.flatMap((s) => s.label ? [{
        label: s.label,
        detail: s.detail,
        // CodeMirror completion types map to icons in the default theme.
        type: s.kind === 'field' ? 'variable' : 'keyword',
      }] : []),
      // Allow the list to update as the user keeps typing.
      validFor: /^\w*$/,
    };
  };
}
