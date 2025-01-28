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
  const { player } = useSessionStore();

  const items: NavItemsProps[] = [
    {
      name: "Início",
      icon: HomeIcon,
      href: `/${player?.ID}`,
      exact: true,
    },
    {
      name: "Ligas",
      icon: Icons.shieldEllipsis,
      items: [
        ...(player?.leagues?.map(
          (league) => ({
            name: league.name,
            icon: Icons.shield,
            href: `/${player?.ID}/leagues/${league.ID}`,
            exact: true,
          }),
          {
            name: "Criar Liga",
            icon: Icons.shieldPlus,
            href: `/${player?.ID}/leagues`,
            exact: true,
          }
        ) || [
          {
            name: "Criar Liga",
            icon: Icons.shieldPlus,
            href: `/${player?.ID}/leagues`,
            exact: true,
          },
        ]),
      ],
    },
    {
      name: "Configurações",
      icon: Icons.cog,
      items: [
        {
          name: "Geral",
          icon: GearIcon,
          href: `/${player?.ID}/settings`,
          exact: true,
        },
        {
          name: "Cobrança",
          icon: ReceiptIcon,
          href: `/${player?.ID}/settings/billing`,
          exact: true,
        },
        {
          name: "Membros",
          icon: PeopleIcon,
          href: `/${player?.ID}/settings/members`,
          exact: true,
        },
      ],
    },
  ];

  return (
    <SidebarGroup>
      <SidebarGroupLabel>Menu</SidebarGroupLabel>
      <SidebarMenu>
        {items.map((item) => {
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
