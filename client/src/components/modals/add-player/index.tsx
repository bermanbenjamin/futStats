"use client";

import AddPlayerForm from "./components/add-player-form";

import { useRouter } from "next/navigation";
import { useQueryState } from "nuqs";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "../../ui/dialog";

export function AddPlayerModal() {
  const [isOpen, setIsOpen] = useQueryState("add-player");

  const router = useRouter();

  const slug = window.location.pathname.split("/").pop();

  return (
    <Dialog
      open={isOpen === "true"}
      onOpenChange={(open) => {
        setIsOpen(open ? "true" : null);
        if (!open) {
          router.push(window.location.pathname);
        }
      }}
    >
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Adicionar Jogador a Liga</DialogTitle>
        </DialogHeader>
        <AddPlayerForm slug={slug!} />
      </DialogContent>
    </Dialog>
  );
}
