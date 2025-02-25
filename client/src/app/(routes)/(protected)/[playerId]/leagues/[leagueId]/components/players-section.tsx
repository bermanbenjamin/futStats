"use client";

import { Icons } from "@/components/icons";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Player } from "@/http/types";
import { useQueryState } from "nuqs";
import React from "react";

type PlayerSectionProperties = {
  players: Player[];
};

const PlayersSection: React.FC<PlayerSectionProperties> = ({ players }) => {
  const [, setIsModalOpen] = useQueryState("add-player");
  function handleNewPlayer() {
    setIsModalOpen("true");
  }

  return (
    <section className="w-full">
      <div className="w-full flex justify-between items-center mb-4">
        <h2 className="text-xl font-semibold ">Players</h2>
        <Button className="ml-auto" onClick={handleNewPlayer}>
          <Icons.circlePlus />
          Add Player
        </Button>
      </div>
      <div className="overflow-x-auto">
        <Table className="min-w-full border rounded-lg">
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Position</TableHead>
              <TableHead>Assists</TableHead>
              <TableHead>Goals</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {players.length > 0 ? (
              players.map((player) => (
                <TableRow key={player.id}>
                  <TableCell>{player.name}</TableCell>
                  <TableCell>{player.position}</TableCell>
                  <TableCell>{player.assists}</TableCell>
                  <TableCell>{player.goals}</TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell colSpan={4} className="text-gray-500">
                  No players registered yet
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </section>
  );
};

export default PlayersSection;
