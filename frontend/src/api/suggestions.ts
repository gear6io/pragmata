import { GeneratedAPIInstance } from './generatedAPIInstance';

export type MatchingType = 'exact' | 'fuzzy';
export type ContextType = 'source' | 'field';

export interface SuggestionRequest {
  contextType: ContextType;
  matchingType?: MatchingType;
  searchText: string;
  nodeRef?: string;
  pipeContent?: string;
}

export interface Suggestion {
  value: string;
  label: string;
  kind: string;
  detail?: string;
}

export interface SuggestionResponse {
  complete: boolean;
  suggestions: Suggestion[];
}

export function getSuggestions(req: SuggestionRequest): Promise<SuggestionResponse> {
  return GeneratedAPIInstance<SuggestionResponse>({
    url: '/v0/suggestions',
    method: 'POST',
    data: req,
    headers: { 'Content-Type': 'application/json' },
  });
}
