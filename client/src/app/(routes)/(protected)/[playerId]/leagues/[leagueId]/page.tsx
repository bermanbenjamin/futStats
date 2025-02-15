"use client";

import { useGetLeagueService } from "@/http/league/use-league-service";
import { useParams } from "next/navigation";

export default function LeaguePage() {
  const { leagueId } = useParams();
  const { data, isLoading } = useGetLeagueService(leagueId as string);

  return (
    <div>
      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <div className="flex flex-col gap-6">
          <section className="space-y-4">
            <h1 className="text-2xl font-bold">{data?.name}</h1>
            <div className="grid grid-cols-2 gap-4">
              <div className="p-4 border rounded-lg">
                <h2 className="text-lg font-semibold mb-2">League Details</h2>
                <p>Owner: {data?.owner.name}</p>
                <p>
                  Created at:{" "}
                  {new Date(data?.CreatedAt as string).toLocaleDateString(
                    "pt-BR"
                  )}
                </p>
              </div>
              <div className="p-4 border rounded-lg">
                <h2 className="text-lg font-semibold mb-2">Members</h2>
                {/* Add member list here when available */}
                <p className="text-gray-500">No members yet</p>
              </div>
            </div>
          </section>

          <section>
            <h2 className="text-xl font-semibold mb-4">Players</h2>
            <div className="overflow-x-auto">
              <table className="min-w-full border rounded-lg">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-sm font-medium text-gray-500">
                      Name
                    </th>
                    <th className="px-6 py-3 text-left text-sm font-medium text-gray-500">
                      Position
                    </th>
                    <th className="px-6 py-3 text-left text-sm font-medium text-gray-500">
                      Team
                    </th>
                    <th className="px-6 py-3 text-left text-sm font-medium text-gray-500">
                      Points
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {/* Add player rows here when available */}
                  <tr>
                    <td className="px-6 py-4 text-gray-500" colSpan={4}>
                      No players registered yet
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>
      )}
    </div>
  );
}
