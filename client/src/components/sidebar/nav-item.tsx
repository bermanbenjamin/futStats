import { type IconType } from "@/components/icons";
import { cn } from "@/lib/utils";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Suspense, useMemo, useState } from "react";
import type { NavItemsProps } from "./nav-main";

export type NavItemType = {
  name: string;
  icon: IconType;
  href: string;
  exact?: boolean;
};

function NavItemContent({ item }: { item: NavItemsProps }) {
  const { name, icon: Icon, href, exact } = item;

  const [hovered, setHovered] = useState(false);

  const pathname = usePathname();

  const isActive = useMemo(() => {
    return exact ? pathname === href : pathname.startsWith(href!);
  }, [pathname, href, exact]);

  return (
    <Link
      href={href!}
      data-active={isActive}
      onPointerEnter={() => setHovered(true)}
      onPointerLeave={() => setHovered(false)}
      className={cn(
        "group flex items-center gap-2.5 rounded-md p-2 text-sm leading-none text-neutral-600 dark:text-neutral-100 transition-[background-color,color,font-weight] duration-75 hover:bg-neutral-200/50 dark:hover:bg-neutral-800 active:bg-neutral-200/80 dark:active:bg-neutral-700/80",
        "outline-none focus-visible:ring-2 focus-visible:ring-black/50",
        isActive &&
          "bg-indigo-100/50 font-medium text-indigo-900 dark:text-neutral-100 hover:bg-indigo-100/80 active:bg-indigo-100 dark:bg-neutral-700 dark:hover:bg-neutral-700/60 transition-all"
      )}
    >
      <Icon
        className="size-4 text-neutral-500 dark:text-neutral-100 transition-colors duration-75 group-data-[active=true]:text-indigo-900 dark:group-data-[active=true]:text-neutral-100"
        data-hovered={hovered}
      />
      {name}
    </Link>
  );
}

export function NavItem({ item }: { item: NavItemsProps }) {
  return (
    <Suspense
      fallback={
        <div className="group flex items-center gap-2.5 rounded-md p-2 text-sm leading-none text-neutral-600 dark:text-neutral-100">
          <div className="size-4 bg-neutral-300 dark:bg-neutral-600 rounded animate-pulse" />
          <div className="h-4 w-20 bg-neutral-300 dark:bg-neutral-600 rounded animate-pulse" />
        </div>
      }
    >
      <NavItemContent item={item} />
    </Suspense>
  );
}
