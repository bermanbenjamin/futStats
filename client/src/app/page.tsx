import Image from "next/image";

export default function Home() {
  return (
    <div className="flex">
      <aside className="flex justify-center items-center h-screen w-1/3">
        <Image
          src="/logo.svg"
          alt="My Logo"
          width={150}
          height={150}
          layout="fixed"
        />
      </aside>
      <main className="flex justify-center items-center h-screen">
        <h1 className="text-6xl font-bold text-center">Welcome to Next.js!</h1>
      </main>
    </div>
  );
}
