import { useState } from 'react';
import { Link, NavLink, Outlet, useNavigate } from 'react-router-dom';
import { clearToken, getToken, setToken } from '../lib/auth';

function TokenGate({ onSave }: { onSave: () => void }) {
  const [value, setValue] = useState('');

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    const trimmed = value.trim();
    if (!trimmed) return;
    setToken(trimmed);
    onSave();
  }

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-900/60 z-50">
      <div className="bg-white rounded-xl shadow-xl p-8 w-full max-w-sm">
        <h1 className="text-xl font-semibold mb-1">Enter API Token</h1>
        <p className="text-sm text-gray-500 mb-5">
          Your token is stored locally and sent as a Bearer header on every request.
        </p>
        <form onSubmit={handleSubmit} className="flex flex-col gap-3">
          <input
            type="password"
            placeholder="Bearer token"
            value={value}
            onChange={(e) => setValue(e.target.value)}
            className="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
            autoFocus
          />
          <button
            type="submit"
            className="bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg py-2 text-sm font-medium transition-colors"
          >
            Save & Continue
          </button>
        </form>
      </div>
    </div>
  );
}

export default function Layout() {
  const [hasToken, setHasToken] = useState(() => !!getToken());
  const navigate = useNavigate();

  function handleSignOut() {
    clearToken();
    setHasToken(false);
  }

  return (
    <>
      {!hasToken && <TokenGate onSave={() => setHasToken(true)} />}
      <div className="flex h-screen bg-gray-50">
        {/* Sidebar */}
        <aside className="w-56 flex flex-col bg-white border-r border-gray-200 shrink-0">
          <Link to="/pipes" className="px-5 py-4 text-lg font-bold text-indigo-600 tracking-tight">
            Pragmata
          </Link>
          <nav className="flex-1 px-3 pb-4 space-y-1">
            <NavLink
              to="/pipes"
              className={({ isActive }) =>
                `flex items-center gap-2 px-3 py-2 rounded-lg text-sm font-medium transition-colors ${
                  isActive
                    ? 'bg-indigo-50 text-indigo-700'
                    : 'text-gray-600 hover:bg-gray-100'
                }`
              }
            >
              Pipes
            </NavLink>
          </nav>
          <div className="px-3 pb-4">
            <button
              onClick={handleSignOut}
              className="w-full text-left px-3 py-2 text-sm text-gray-500 hover:text-gray-800 rounded-lg hover:bg-gray-100 transition-colors"
            >
              Sign out
            </button>
          </div>
        </aside>

        {/* Main content */}
        <main className="flex-1 overflow-y-auto">
          <Outlet />
        </main>
      </div>
    </>
  );
}
