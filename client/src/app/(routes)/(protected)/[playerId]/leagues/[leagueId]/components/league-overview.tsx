"use client";

import React from "react";

type LeagueOverviewProps = {
  name: string;
  ownerName: string;
  createdAt: string;
  membersLength: number;
};

const LeagueOverview: React.FC<LeagueOverviewProps> = ({
  name,
  ownerName,
  createdAt,
  membersLength,
}) => {
  return (
    <section className="space-y-4 w-full">
      <h1 className="text-2xl font-bold">{name}</h1>
      <div className="grid grid-cols-2 gap-4">
        <div className="p-4 border rounded-lg w-full">
          <h2 className="text-lg font-semibold mb-2">League Details</h2>
          <p>Owner: {ownerName}</p>
          <p>Created at: {new Date(createdAt).toLocaleDateString("pt-BR")}</p>
        </div>
        <div className="p-4 border rounded-lg w-full">
          <h2 className="text-lg font-semibold mb-2">Members</h2>
          {membersLength > 1 ? (
            <p className="text-gray-500">
              Esta liga conta com {membersLength} jogadores.
            </p>
          ) : (
            <p className="text-gray-500">
              Esta liga conta com {membersLength} jogador.
            </p>
          )}
        </div>
      </div>
    </section>
  );
};

export default LeagueOverview;
