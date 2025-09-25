import Link from "next/link";

export default function SignInPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Sign in to your account
          </h2>
        </div>
        <div className="text-center">
          <p className="text-gray-600 mb-4">
            This is a demo version. Please use the full application.
          </p>
          <Link
            href="/auth/sign-up"
            className="font-medium text-indigo-600 hover:text-indigo-500"
          >
            Create an account
          </Link>
        </div>
      </div>
    </div>
  );
}
