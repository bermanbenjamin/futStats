import { redirect } from "next/navigation";

export default function HomePage() {
  // This will redirect to sign-in page
  redirect("/auth/sign-in");
}
