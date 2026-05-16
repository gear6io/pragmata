import { Link, NavLink, Outlet } from 'react-router-dom';

export default function Layout() {
  return (
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
        </aside>

        {/* Main content */}
        <main className="flex-1 overflow-y-auto">
          <Outlet />
        </main>
    </div>
  );
}
