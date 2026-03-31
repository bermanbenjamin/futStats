"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useCreateEvent, useCreateMatch } from "@/http/match/use-match-service";
import { Player } from "@/http/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryState } from "nuqs";
import { Suspense, useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

const EVENT_TYPES = [
  "Goal",
  "Assist",
  "Disarm",
  "Dribble",
  "RedCard",
  "YellowCard",
] as const;

const matchSchema = z.object({
  date: z.string().min(1, "Date is required"),
});

type MatchFormValues = z.infer<typeof matchSchema>;

interface RecordMatchModalProps {
  leagueId: string;
  leagueSlug: string;
  players: Player[];
}

function RecordMatchModalContent({
  leagueId,
  leagueSlug,
  players,
}: RecordMatchModalProps) {
  const [isOpen, setIsOpen] = useQueryState("record-match");
  const [matchId, setMatchId] = useState<string | null>(null);

  const form = useForm<MatchFormValues>({
    resolver: zodResolver(matchSchema),
    defaultValues: { date: new Date().toISOString().slice(0, 16) },
  });

  const { mutateAsync: createMatch, isPending: creatingMatch } =
    useCreateMatch({
      onSuccess: (res) => {
        setMatchId(res.data.id);
        toast.success("Match created! Now add events.");
      },
    });

  const { mutateAsync: createEvent, isPending: creatingEvent } =
    useCreateEvent({
      onSuccess: () => toast.success("Event recorded"),
    });

  async function onCreateMatch(values: MatchFormValues) {
    const date = new Date(values.date).toISOString();
    await createMatch({ league_id: leagueId, date });
  }

  async function onAddEvent(
    playerId: string,
    type: string,
    assistentId?: string
  ) {
    if (!matchId) return;
    await createEvent({
      type,
      player_id: playerId,
      match_id: matchId,
      assistent_id: assistentId,
    });
  }

  function handleClose(open: boolean) {
    setIsOpen(open ? "true" : null);
    if (!open) {
      setMatchId(null);
      form.reset();
    }
  }

  return (
    <Dialog open={isOpen === "true"} onOpenChange={handleClose}>
      <DialogContent className="sm:max-w-[540px]">
        <DialogHeader>
          <DialogTitle>
            {matchId ? "Record Events" : "Create Match"}
          </DialogTitle>
        </DialogHeader>

        {!matchId ? (
          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(onCreateMatch)}
              className="space-y-4"
            >
              <FormField
                control={form.control}
                name="date"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Match Date & Time</FormLabel>
                    <FormControl>
                      <Input type="datetime-local" {...field} />
                    </FormControl>
                  </FormItem>
                )}
              />
              <Button type="submit" className="w-full" disabled={creatingMatch}>
                {creatingMatch ? "Creating..." : "Create Match"}
              </Button>
            </form>
          </Form>
        ) : (
          <div className="space-y-4">
            <p className="text-sm text-gray-500">
              Match created. Add events below.
            </p>
            <div className="space-y-2 max-h-80 overflow-y-auto">
              {players.map((player) => (
                <div
                  key={player.id}
                  className="flex items-center gap-2 p-2 border rounded-lg"
                >
                  <span className="text-sm font-medium w-28 truncate">
                    {player.name}
                  </span>
                  <div className="flex flex-wrap gap-1">
                    {EVENT_TYPES.map((type) => (
                      <button
                        key={type}
                        onClick={() => onAddEvent(player.id, type)}
                        disabled={creatingEvent}
                        className="px-2 py-1 text-xs bg-gray-100 hover:bg-gray-200 rounded transition-colors disabled:opacity-50"
                      >
                        {type}
                      </button>
                    ))}
                  </div>
                </div>
              ))}
            </div>
            <Button
              className="w-full"
              variant="outline"
              onClick={() => handleClose(false)}
            >
              Done
            </Button>
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
}

export function RecordMatchModal(props: RecordMatchModalProps) {
  return (
    <Suspense fallback={null}>
      <RecordMatchModalContent {...props} />
    </Suspense>
  );
}
