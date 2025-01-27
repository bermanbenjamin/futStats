"use client";

import { Icons } from "@/components/icons";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "@/components/ui/sidebar";
import { getInitials } from "@/lib/utils";
import { useSessionStore } from "@/stores/session-store";

export function NavUser() {
  const { isMobile } = useSidebar();
  const { player } = useSessionStore();

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <SidebarMenuButton
              size="lg"
              className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
            >
              {player ? (
                <Avatar className="size-7">
                  <AvatarFallback className="bg-black/15">
                    <span className="text-xs">
                      {getInitials(player.name ?? "AA")}
                    </span>
                  </AvatarFallback>
                </Avatar>
              ) : (
                <Avatar className="size-7">
                  <AvatarFallback className="bg-black/15">
                    <span className="text-xs">{getInitials("AA")}</span>
                  </AvatarFallback>
                </Avatar>
              )}
              {player ? (
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-semibold">
                    {player.name || player.email.split("@")[0]}
                  </span>
                  <span className="truncate text-xs">{player.email}</span>
                </div>
              ) : (
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <div className="h-3 w-12 animate-pulse rounded-full bg-neutral-200" />
                  <div className="h-3 w-20 animate-pulse rounded-full bg-neutral-200" />
                </div>
              )}
              <Icons.chevronsUpDown className="ml-auto size-4" />
            </SidebarMenuButton>
          </DropdownMenuTrigger>
          <DropdownMenuContent
            className="w-[--radix-dropdown-menu-trigger-width] min-w-56 rounded-lg"
            side={isMobile ? "bottom" : "right"}
            align="end"
            sideOffset={4}
          >
            <DropdownMenuLabel className="p-0 font-normal">
              <div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
                {player ? (
                  <Avatar className="size-7">
                    <AvatarFallback className="bg-black/15">
                      <span className="text-xs">
                        {getInitials(player.name ?? "AA")}
                      </span>
                    </AvatarFallback>
                  </Avatar>
                ) : (
                  <Avatar className="size-7">
                    <AvatarFallback className="bg-black/15">
                      <span className="text-xs">{getInitials("AA")}</span>
                    </AvatarFallback>
                  </Avatar>
                )}
                {player ? (
                  <div className="grid flex-1 text-left text-sm leading-tight">
                    <span className="truncate font-semibold">
                      {player.name || player.email.split("@")[0]}
                    </span>
                    <span className="truncate text-xs">{player.email}</span>
                  </div>
                ) : (
                  <div className="grid flex-1 text-left text-sm leading-tight">
                    <div className="h-3 w-12 animate-pulse rounded-full bg-neutral-200" />
                    <div className="h-3 w-20 animate-pulse rounded-full bg-neutral-200" />
                  </div>
                )}
              </div>
            </DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup></DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem asChild>
              <Button
                variant="ghost"
                className="w-full justify-start cursor-pointer"
                // onClick={() => signOut({ redirectUrl: appRoutes.auth.signIn })}
              >
                <Icons.signOut className="size-4 mr-2" />
                Sair
              </Button>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}
