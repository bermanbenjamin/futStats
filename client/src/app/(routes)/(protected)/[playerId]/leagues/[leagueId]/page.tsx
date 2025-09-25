import Link from "next/link";

export default function LeaguePage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            League Dashboard
          </h2>
        </div>
        <div className="text-center">
          <p className="text-gray-600 mb-4">
            This is a demo version. Please sign in to access the full league
            dashboard.
          </p>
          <Link
            href="/auth/sign-in"
            className="font-medium text-indigo-600 hover:text-indigo-500"
          >
            Sign In
          </Link>
        </div>
      </div>
    </div>
  );
}
