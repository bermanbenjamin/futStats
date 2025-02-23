"use client";

import { Icons, type IconType } from "@/components/icons";
import {
  GearIcon,
  HomeIcon,
  PeopleIcon,
  ReceiptIcon,
} from "@/components/icons/dub";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import {
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar";
import { useSessionStore } from "@/stores/session-store";
import { useEffect, useState } from "react";
import { NavItem } from "./nav-item";

export type NavItemsProps = {
  name: string;
  icon: IconType;
  exact?: boolean;
} & (
  | { href: string; items?: never }
  | { href?: never; items: NavItemsProps[] }
);

export function NavMain() {
  const player = useSessionStore((state) => state.player);
  const [navItems, setNavItems] = useState<NavItemsProps[]>([]);

  useEffect(() => {
    if (!player) return;
    const items: NavItemsProps[] = [
      {
        name: "Início",
        icon: HomeIcon,
        href: `/${player.id}`,
        exact: true,
      },
      {
        name: "Ligas",
        icon: Icons.shieldEllipsis,
        items: [
          ...(player.owned_leagues?.map((league) => ({
            name: league.name,
            icon: Icons.shield,
            href: `/${player.id}/leagues/${league.slug}`,
            exact: true,
          })) || []),
          ...(player.member_leagues?.map((league) => ({
            name: league.name,
            icon: Icons.shield,
            href: `/${player.id}/leagues/${league.slug}`,
            exact: true,
          })) || []),
          {
            name: "Criar Liga",
            icon: Icons.shieldPlus,
            href: `/${player.id}?create-league=true`,
            exact: true,
          },
        ],
      },
      {
        name: "Configurações",
        icon: Icons.cog,
        items: [
          {
            name: "Geral",
            icon: GearIcon,
            href: `/${player.id}/settings`,
            exact: true,
          },
          {
            name: "Cobrança",
            icon: ReceiptIcon,
            href: `/${player.id}/settings/billing`,
            exact: true,
          },
          {
            name: "Membros",
            icon: PeopleIcon,
            href: `/${player.id}/settings/members`,
            exact: true,
          },
        ],
      },
    ];
    setNavItems(items);
  }, [player]);

  return (
    <SidebarGroup>
      <SidebarGroupLabel>Menu</SidebarGroupLabel>
      <SidebarMenu>
        {navItems.map((item) => {
          const hasItems = item.items && item.items.length > 0;

          if (!hasItems) {
            return (
              <SidebarMenuItem key={item.name}>
                <SidebarMenuButton asChild>
                  <NavItem item={item} />
                </SidebarMenuButton>
              </SidebarMenuItem>
            );
          }

          return (
            <Collapsible key={item.name} asChild className="group/collapsible">
              <SidebarMenuItem>
                <CollapsibleTrigger asChild>
                  <SidebarMenuButton tooltip={item.name}>
                    <item.icon />
                    <span>{item.name}</span>
                    <Icons.chevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
                  </SidebarMenuButton>
                </CollapsibleTrigger>
                <CollapsibleContent>
                  <SidebarMenuSub>
                    {item.items?.map((subItem) => (
                      <SidebarMenuSubItem key={subItem.name}>
                        <SidebarMenuSubButton asChild>
                          <NavItem item={subItem} />
                        </SidebarMenuSubButton>
                      </SidebarMenuSubItem>
                    ))}
                  </SidebarMenuSub>
                </CollapsibleContent>
              </SidebarMenuItem>
            </Collapsible>
          );
        })}
      </SidebarMenu>
    </SidebarGroup>
  );
}
