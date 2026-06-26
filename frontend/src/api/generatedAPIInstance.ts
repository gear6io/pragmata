import axios, { AxiosError, AxiosRequestConfig } from 'axios';

// Bundler replaces this at build time. Override in your env config.
declare const __API_BASE_URL__: string | undefined;
const API_BASE_URL: string =
	(typeof __API_BASE_URL__ !== 'undefined' && __API_BASE_URL__) || '';

const instance = axios.create({ baseURL: API_BASE_URL });

// Attach bearer token from localStorage on every request.
instance.interceptors.request.use((config) => {
	const token = localStorage.getItem('api_token');
	if (token) {
		config.headers = config.headers ?? {};
		config.headers.Authorization = `Bearer ${token}`;
	}
	return config;
});

// Required by orval: the named export used as the mutator.
export const GeneratedAPIInstance = <T>(
	config: AxiosRequestConfig,
): Promise<T> => {
	return instance(config).then(({ data: body }) => body as T);
};

// Required by orval for typed error and body shapes.
export type ErrorType<Error> = AxiosError<Error>;
export type BodyType<BodyData> = BodyData;
