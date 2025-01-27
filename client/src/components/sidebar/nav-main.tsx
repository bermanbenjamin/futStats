"use client";

import { Icons, type IconType } from "@/components/icons";
import {
  GearIcon,
  HomeIcon,
  LayersIcon,
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
import { useParams } from "next/navigation";
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
  const { slug } = useParams() as { slug?: string };

  const items: NavItemsProps[] = [
    {
      name: "Início",
      icon: HomeIcon,
      href: `/${slug}`,
      exact: true,
    },
    {
      name: "Estoque",
      icon: Icons.bag,
      items: [
        {
          name: "Items",
          icon: Icons.list,
          href: `/${slug}/items`,
        },
        {
          name: "Categorias",
          icon: LayersIcon,
          href: `/${slug}/categories`,
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
          href: `/${slug}/settings`,
          exact: true,
        },
        {
          name: "Cobrança",
          icon: ReceiptIcon,
          href: `/${slug}/settings/billing`,
          exact: true,
        },
        {
          name: "Membros",
          icon: PeopleIcon,
          href: `/${slug}/settings/members`,
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
