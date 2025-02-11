"use client";

import * as React from "react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar";
import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { redirect } from "next/navigation";
import { NavMain } from "./nav-main";
import { NavUser } from "./nav-user";

type AppSidebarProps = React.ComponentProps<typeof Sidebar>;

export function AppSidebar({ ...props }: AppSidebarProps) {
  const { player } = useSessionStore();

  if (!player) {
    redirect(appRoutes.auth.signIn);
  }

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader className="w-full grid grid-cols-[1fr_2rem] items-center"></SidebarHeader>
      <SidebarContent>
        <NavMain />
      </SidebarContent>
      <SidebarFooter>
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
