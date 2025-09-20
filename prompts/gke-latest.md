You are a Go developer working on a GKE release mapping system.

Goal: Update the `GKEProjectReleases` array in `gke.go` by adding ONLY the GKE R releases that appear ABOVE our latest known release on the official release notes page.

Inputs:

- Path to gke.go: codecamp/pkg/project/gke.go
- Release notes base URL: https://cloud.google.com/kubernetes-engine/docs/release-notes

Definitions:

- “R release” = section titled like “(YYYY-RXX) Version updates”.
- “TOP_EXISTING” = first Version in GKEProjectReleases matching ^\d{4}-R\d{2}$.

Selectors (scoped per section):

- R header: h4[id$="\_version_updates"][data-text*="Version updates"]
- Stable tab: the tab whose visible text is “Stable” (case-insensitive); read aria-controls.
- Stable panel: section[role="tabpanel"]#<aria-controls> (fallback: section#tabpanel-stable-channel[role="tabpanel"])

Steps (high level):

- Find TOP_EXISTING in our file (first entry only).
- On the release-notes page, find the TOP_EXISTING header; every R above it is new (handle multiple).
- For each new R, extract ALL Stable-channel versions (defaults, newly available, removed).
- Normalize to kube@X.Y.Z (take SemVer head); de-duplicate; sort by full SemVer ascending.
- Return STRICT JSON only:
  { "entries": [ { "Version":"YYYY-RXX", "RelatedProjectReleases":["kube@1.30.14",...], "source":"<url#anchor>" }, ... ] }
