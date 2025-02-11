"use client";

import { useGetLeagueService } from "@/http/league/use-league-service";
import { useParams } from "next/navigation";

export default function LeaguePage() {
  const { leagueId } = useParams();
  const { data: league } = useGetLeagueService(leagueId as string);

  console.log(league);

  return (
    <div>
      <h1>{league?.[0]}</h1>
    </div>
  );
}
