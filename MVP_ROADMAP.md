# futStats (PitchStats) — MVP Roadmap

> Last updated: 2026-03-30

## Current State Analysis

### What's working
- Auth: login / signup / logout with JWT ✅
- Player CRUD ✅
- League CRUD (create, get, update, delete, add member) ✅
- Event model + CRUD endpoints ✅
- Player stats computation service ✅
- Clean architecture (Handler → Service → Repository) ✅
- Frontend auth pages (sign-in, sign-up) ✅
- Frontend: player dashboard shell + stats components ✅
- Frontend: league detail shell (overview, players, transfer-history) ✅
- Create league modal + add player modal (UI exists) ✅
- Railway + Vercel deployment config ✅

### Critical gaps (can't launch without)
- **No Match API** — Match model exists but zero endpoints. Events require a `matchId`, so right now there is no way to record any statistics.
- **No event recording UI** — Even if the API existed, there's no UI flow to log a match and add events.
- **Protected routes unverified** — The `(protected)` route group needs to be confirmed to redirect unauthenticated users.

---

## Roadmap

### Phase 0 — Core Gameplay Loop (P0 Blockers)
These must ship before anything else has value.

| # | Task | Owner | Status |
|---|------|-------|--------|
| 1 | **Match API** — POST/GET endpoints for matches under a league | backend | ⬜ |
| 2 | **Event recording flow** — validate matchId on event creation, trigger stats recompute | backend | ⬜ |
| 3 | **Frontend protected routes** — verify auth guard, handle 401 redirect | frontend | ⬜ |
| 4 | **Match recording UI** — create match modal + event entry per player | frontend | ⬜ |

### Phase 1 — Data & UX (P1 Core MVP)
The platform works but data must be visible.

| # | Task | Owner | Status |
|---|------|-------|--------|
| 5 | **Player stats dashboard** — wire StatsCard/StatsChart to real API data | frontend | ⬜ |
| 6 | **League detail page** — wire members list, match history, add-player flow | frontend | ⬜ |
| 7 | **Season API + UI** — CRUD seasons, season selector on league page | full-stack | ⬜ |
| 8 | **Server-side validation** — audit all mutation endpoints, return 400 with field errors | backend | ⬜ |

### Phase 2 — Quality & Launch (P2)

| # | Task | Owner | Status |
|---|------|-------|--------|
| 9 | **Frontend error handling** — error boundaries, sonner toasts, Zod form errors | frontend | ⬜ |
| 10 | **Deployment verification** — Railway env vars, Vercel env vars, CORS, DB migrations | devops | ⬜ |
| 11 | **Pagination** — page/limit on list endpoints + frontend handling | full-stack | ⬜ |

---

## Architecture Notes

```
client (Next.js 15 / Vercel)
  └── src/
      ├── app/(routes)/(protected)/[playerId]/        ← player dashboard
      ├── app/(routes)/(protected)/[playerId]/leagues/[leagueId]/  ← league view
      ├── app/(routes)/auth/                          ← sign-in / sign-up
      ├── http/{auth,league,player}/                  ← API service + hooks
      └── stores/session-store.ts                     ← Zustand JWT session

server (Go / Gin / Railway)
  └── cmd/
      ├── main.go                                     ← bootstrap
      └── api/
          ├── router.go                               ← routes
          ├── handlers/                               ← HTTP layer
          ├── requests/                               ← request types
          └── constants/
  └── internal/
      ├── models/                                     ← GORM models
      ├── repository/                                 ← DB queries
      ├── services/                                   ← business logic
      ├── middlewares/                                ← auth, logging, owner
      └── config/ db/ errors/ logger/ validation/
```

## Event Types (enums/eventType.go)
Goal, Assist, Disarm, Dribble, RedCard, YellowCard

## Player Computed Stats
goals, assists, disarms, dribbles, matches, red_cards, yellow_cards
(computed in `player_stats_service.go` from Events table)

## Deployment
- **Backend**: Railway (`railway.json` + `Dockerfile`) — needs DATABASE_URL, SECRET_KEY, ALLOWED_ORIGINS
- **Frontend**: Vercel (`.vercel/`) — needs NEXT_PUBLIC_API_BASE_URL

---

## Definition of MVP Done

- [ ] A user can sign up, create a league, add players, record a match with events, and see updated stats on the dashboard
- [ ] Deployment is live (Railway API + Vercel frontend)
- [ ] No critical unhandled errors in the happy path
