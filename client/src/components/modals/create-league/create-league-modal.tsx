"use client";

import CreateLeagueForm from "@/components/modals/create-league/components/create-league-form";

import { useRouter } from "next/navigation";
import { useQueryState } from "nuqs";
import { Suspense } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "../../ui/dialog";

function CreateLeagueModalContent() {
  const [isOpen, setIsOpen] = useQueryState("create-league");
  const router = useRouter();

  return (
    <Dialog
      open={isOpen === "true"}
      onOpenChange={(open) => {
        setIsOpen(open ? "true" : null);
        if (!open) {
          // Remove the query param when closing
          router.back();
        }
      }}
    >
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Criar Liga</DialogTitle>
        </DialogHeader>
        <CreateLeagueForm />
      </DialogContent>
    </Dialog>
  );
}

export function CreateLeagueModal() {
  return (
    <Suspense fallback={null}>
      <CreateLeagueModalContent />
    </Suspense>
  );
}
