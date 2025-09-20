You are a Go developer working on a GKE release mapping system.

Goal: Update the `GKEProjectReleases` array in `gke.go` by adding ONLY new (latest) GKE R releases that are not yet present in the file. Do not backfill older years; stop once you hit the most recent release already recorded.

Inputs:

- Path to gke.go: pkg/project/gke.go
- Release notes base URL: https://cloud.google.com/kubernetes-engine/docs/release-notes

Definitions:

- “R release” = section titled like “(YYYY-RXX) Version updates”.
- “New” = any `(YYYY-RXX)` that is strictly greater (by year, then RXX) than the highest `(YYYY-RXX)` currently present in `GKEProjectReleases`.

Steps:

0. Read current state

   - Parse `GKEProjectReleases` from `gke.go`.
   - Extract all existing `Version` values matching regex `^\d{4}-R\d{2}$`.
   - Compute `HIGHEST_EXISTING` by comparing first on year (YYYY), then on RXX (numeric).

1. Discover new releases (reverse-chronological)

   - Crawl the GKE release notes page from newest to oldest.
   - Collect each section whose heading matches `(YYYY-RXX) Version updates`.
   - Maintain natural page order (newest first).
   - **Stop condition:** as soon as you encounter a section whose `(YYYY-RXX)` is **<= HIGHEST_EXISTING**, stop discovery. Do not collect that or any older sections.
   - The remaining collected R IDs are `NEW_R_IDS`. If empty, print “No new releases” and exit without modifying the file.

2. Extract Stable-channel versions for each new R

   - For every R in `NEW_R_IDS`:
     - Locate the **Stable** tab/panel within the same section (accept synonyms: “Stable”, “Stable channel”, “Default version for new clusters (Stable)”).
     - From Stable only, gather ALL Kubernetes versions mentioned under:
       - Default version for new clusters / control plane default / node default
       - Newly/Also available versions
       - No longer available / removed versions
     - If the Stable section is absent, log `<R> skipped: no Stable tab` and continue.
     - Extract SemVer using `\b(\d+\.\d+\.\d+)\b` from any GKE version string (ignore `-gke.*`, `-autopilot.*`, `+cos*`, etc.).
     - De-duplicate within this R; format as `kube@<semver>`; sort by full SemVer ascending (major, minor, patch).

3. Insert entries

   - For each R in `NEW_R_IDS` (process newest → oldest so the array remains newest-first):
     ```
     {
       Project: GKE.ID,
       Version: "YYYY-RXX",
       RelatedProjectReleases: []string{
         "kube@1.30.12",
         "kube@1.31.9",
         ...
       },
     },
     // source: <release-notes-URL>#<anchor-for-YYYY-RXX>
     ```
   - Only add missing R releases; do not modify existing ones.
   - Ensure overall array ordering remains **descending by (year, RXX)**.

4. Validation
   - Ensure valid Go syntax and pass `gofmt`/`goimports`.
   - Print a summary: `HIGHEST_EXISTING`, count of discovered sections, count added, and per-R version counts.

Operational notes:

- Use robust HTML parsing (don’t depend on JS). If tabs are JS-rendered, parse the static HTML; identify the Stable panel within the same section (e.g., by `aria-controls`/`role="tabpanel"` and visible “Stable” tab text).
- Treat both control plane and node defaults as in-scope (still “Stable”).
- If the site layout changes, fail gracefully and report the last successful R processed.
- Idempotency: running again with no newer releases should make no changes and print “No new releases”.

Output:

- Updated `gke.go` with the new entries only (if any), plus the console summary.
