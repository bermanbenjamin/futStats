"use client";

import { useGetLeagueService } from "@/http/league/use-league-service";
import { useParams } from "next/navigation";
import LeagueOverview from "./components/league-overview";
import PlayersSection from "./components/players-section";

export default function LeaguePage() {
  const { leagueId } = useParams();
  const { data, isLoading } = useGetLeagueService(leagueId as string);

  console.log(data);

  return (
    <div className="p-6 bg-white rounded-lg shadow-md w-full">
      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <div className="flex flex-col gap-6 w-full">
          <LeagueOverview
            name={data!.name}
            ownerName={data!.owner.name}
            createdAt={data!.createdAt}
            membersLength={data!.members.length}
          />
          <PlayersSection players={data!.members} />
        </div>
      )}
    </div>
  );
}
