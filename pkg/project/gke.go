package project

import (
	"github.com/chkk-io/schema/model"
	"github.com/chkk-io/schema/types"
)

var GKE = model.Project{
	ID: model.GKEKey,
	MixinTitled: model.MixinTitled{
		Title: "Google Kubernetes Engine (GKE)",
	},
	MixinTagged: model.MixinTagged{
		Tags: []string{"kube", "k8s", "gke"},
	},
	Type:    model.ProjectTypeKubeControlPlaneProvider,
	Aliases: []string{"cloud.google.com/gke", "gke"},
	Versioning: &model.Versioning{
		Scheme:                       model.VersioningSchemeCalender,
		ReleaseCycle:                 model.ReleaseCycleRegular,
		ReleaseIntervalDays:          90,
		IndividualUpgradeRecommended: true,
		ReleasePatterns: []string{
			"<YYYY>-R<MINOR>",
		},
	},
}
var GKECurationConfig = &model.ProjectCurationConfig{
	Series: &model.ReleaseCurationConfig{
		Sources: []*model.LinkTemplateCurationConfig{
			{
				Scrape: &model.SourceScrapeConfig{
					TargetCSSSelector: "#main-content",
					SectionPattern:    "\\((\\d{4}-R\\d+)\\) Version updates",
				},
				LinkTemplate: model.LinkTemplate{
					URLTemplate: "https://cloud.google.com/kubernetes-engine/docs/release-notes",
					LinkType:    types.LinkTypeProjectReleaseNotes,
				},
			},
		},
	},
}

var GKEProjectReleases = []model.ProjectRelease{
	{
		Project: GKE.ID,
		Version: "2025-R37",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.30.14",
			"kube@1.31.11",
			"kube@1.32.6",
			"kube@1.32.7",
			"kube@1.33.2",
			"kube@1.33.3",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R36",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.10",
			"kube@1.31.11",
			"kube@1.32.6",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R35",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.10",
			"kube@1.31.11",
			"kube@1.32.6",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R34",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.10",
			"kube@1.32.6",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R33",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.9",
			"kube@1.31.10",
			"kube@1.32.4",
			"kube@1.32.6",
			"kube@1.33.2",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R32",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.9",
			"kube@1.31.10",
			"kube@1.32.4",
			"kube@1.32.6",
			"kube@1.33.2",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R31",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.8",
			"kube@1.31.9",
			"kube@1.32.2",
			"kube@1.32.4",
			"kube@1.33.2",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R30",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.9",
			"kube@1.32.4",
		},
	},
	{
		Project: GKE.ID,
		Version: "2025-R29",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.9",
			"kube@1.32.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r29_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R27",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.9",
			"kube@1.32.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r27_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R26",
		RelatedProjectReleases: []string{
			"kube@1.30.12",
			"kube@1.31.8",
			"kube@1.31.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r26_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R25",
		RelatedProjectReleases: []string{
			"kube@1.30.11",
			"kube@1.30.12",
			"kube@1.31.7",
			"kube@1.31.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r25_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R24",
		RelatedProjectReleases: []string{
			"kube@1.30.11",
			"kube@1.30.12",
			"kube@1.31.7",
			"kube@1.31.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r24_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R22",
		RelatedProjectReleases: []string{
			"kube@1.30.11",
			"kube@1.31.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r22_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R20",
		RelatedProjectReleases: []string{
			"kube@1.30.10",
			"kube@1.30.11",
			"kube@1.31.6",
			"kube@1.31.7",
			"kube@1.32.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r20_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R19",
		RelatedProjectReleases: []string{
			"kube@1.30.11",
			"kube@1.31.7",
			"kube@1.32.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r19_version_updates
	{
		Project:                GKE.ID,
		Version:                "2025-R18",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r18_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R17",
		RelatedProjectReleases: []string{
			"kube@1.31.6",
			"kube@1.32.1",
			"kube@1.32.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r17_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R16",
		RelatedProjectReleases: []string{
			"kube@1.29.13",
			"kube@1.32.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r16_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R15",
		RelatedProjectReleases: []string{
			"kube@1.29.13",
			"kube@1.30.9",
			"kube@1.30.10",
			"kube@1.31.5",
			"kube@1.31.6",
			"kube@1.32.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r15_version_updates
	{
		Project:                GKE.ID,
		Version:                "2025-R14",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r14_version_updates
	{
		Project:                GKE.ID,
		Version:                "2025-R13",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r13_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R12",
		RelatedProjectReleases: []string{
			"kube@1.30.10",
			"kube@1.31.6",
			"kube@1.32.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r12_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R11",
		RelatedProjectReleases: []string{
			"kube@1.30.10",
			"kube@1.31.6",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r11_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R10",
		RelatedProjectReleases: []string{
			"kube@1.29.13",
			"kube@1.30.9",
			"kube@1.31.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r10_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R09",
		RelatedProjectReleases: []string{
			"kube@1.29.13",
			"kube@1.30.9",
			"kube@1.31.5",
			"kube@1.32.1",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r09_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R08",
		RelatedProjectReleases: []string{
			"kube@1.29.12",
			"kube@1.29.13",
			"kube@1.30.8",
			"kube@1.30.9",
			"kube@1.31.4",
			"kube@1.31.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r08_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R07",
		RelatedProjectReleases: []string{
			"kube@1.29.12",
			"kube@1.29.13",
			"kube@1.30.8",
			"kube@1.30.9",
			"kube@1.31.4",
			"kube@1.31.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r07_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R06",
		RelatedProjectReleases: []string{
			"kube@1.28.15",
			"kube@1.29.12",
			"kube@1.30.8",
			"kube@1.31.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r06_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R05",
		RelatedProjectReleases: []string{
			"kube@1.28.15",
			"kube@1.29.12",
			"kube@1.30.8",
			"kube@1.31.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r05_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R04",
		RelatedProjectReleases: []string{
			"kube@1.28.15",
			"kube@1.29.10",
			"kube@1.29.12",
			"kube@1.30.5",
			"kube@1.30.8",
			"kube@1.31.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r04_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R03",
		RelatedProjectReleases: []string{
			"kube@1.28.15",
			"kube@1.29.12",
			"kube@1.30.6",
			"kube@1.30.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r03_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R02",
		RelatedProjectReleases: []string{
			"kube@1.28.15",
			"kube@1.29.10",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r02_version_updates
	{
		Project: GKE.ID,
		Version: "2025-R01",
		RelatedProjectReleases: []string{
			"kube@1.28.15",
			"kube@1.29.10",
			"kube@1.30.5",
			"kube@1.30.6",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2025-r01_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R50",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r50_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R49",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r49_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R48",
		RelatedProjectReleases: []string{
			"kube@1.28.14",
			"kube@1.28.15",
			"kube@1.29.9",
			"kube@1.29.10",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r48_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R47",
		RelatedProjectReleases: []string{
			"kube@1.28.14",
			"kube@1.28.15",
			"kube@1.29.9",
			"kube@1.29.10",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r47_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R46",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r46_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R45",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r45_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R44",
		RelatedProjectReleases: []string{
			"kube@1.28.14",
			"kube@1.29.9",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r44_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R43",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r43_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R42",
		RelatedProjectReleases: []string{
			"kube@1.28.14",
			"kube@1.29.8",
			"kube@1.29.9",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r42_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R41",
		RelatedProjectReleases: []string{
			"kube@1.28.13",
			"kube@1.28.14",
			"kube@1.29.8",
			"kube@1.29.9",
			"kube@1.30.4",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r41_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R40",
		RelatedProjectReleases: []string{
			"kube@1.28.13",
			"kube@1.28.14",
			"kube@1.29.8",
			"kube@1.30.3",
			"kube@1.30.4",
			"kube@1.30.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r40_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R39",
		RelatedProjectReleases: []string{
			"kube@1.27.16",
			"kube@1.28.13",
			"kube@1.29.8",
			"kube@1.30.3",
			"kube@1.30.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r39_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R38",
		RelatedProjectReleases: []string{
			"kube@1.27.16",
			"kube@1.28.13",
			"kube@1.29.8",
			"kube@1.30.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r38_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R36",
		RelatedProjectReleases: []string{
			"kube@1.28.13",
			"kube@1.29.8",
			"kube@1.30.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r36_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R35",
		RelatedProjectReleases: []string{
			"kube@1.27.16",
			"kube@1.28.12",
			"kube@1.29.7",
			"kube@1.30.2",
			"kube@1.30.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r35_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R32",
		RelatedProjectReleases: []string{
			"kube@1.27.16",
			"kube@1.28.12",
			"kube@1.29.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r32_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R27",
		RelatedProjectReleases: []string{
			"kube@1.27.14",
			"kube@1.28.10",
			"kube@1.28.11",
			"kube@1.29.5",
			"kube@1.29.6",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r27_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R25",
		RelatedProjectReleases: []string{
			"kube@1.27.13",
			"kube@1.27.14",
			"kube@1.28.9",
			"kube@1.28.10",
			"kube@1.29.4",
			"kube@1.29.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r25_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R23",
		RelatedProjectReleases: []string{
			"kube@1.26.15",
			"kube@1.27.13",
			"kube@1.28.9",
			"kube@1.29.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r23_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R21",
		RelatedProjectReleases: []string{
			"kube@1.27.11",
			"kube@1.27.13",
			"kube@1.28.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r21_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R20",
		RelatedProjectReleases: []string{
			"kube@1.26.15",
			"kube@1.27.13",
			"kube@1.28.7",
			"kube@1.28.8",
			"kube@1.28.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r20_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R19",
		RelatedProjectReleases: []string{
			"kube@1.27.13",
			"kube@1.28.8",
			"kube@1.28.9",
			"kube@1.29.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r19_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R18",
		RelatedProjectReleases: []string{
			"kube@1.27.11",
			"kube@1.27.12",
			"kube@1.27.13",
			"kube@1.28.8",
			"kube@1.28.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r18_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R16",
		RelatedProjectReleases: []string{
			"kube@1.27.11",
			"kube@1.28.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r16_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R14",
		RelatedProjectReleases: []string{
			"kube@1.26.8",
			"kube@1.26.14",
			"kube@1.26.15",
			"kube@1.27.11",
			"kube@1.27.12",
			"kube@1.28.7",
			"kube@1.28.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r14_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R13",
		RelatedProjectReleases: []string{
			"kube@1.26.13",
			"kube@1.26.14",
			"kube@1.27.8",
			"kube@1.27.11",
			"kube@1.28.3",
			"kube@1.28.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r13_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R12",
		RelatedProjectReleases: []string{
			"kube@1.26.14",
			"kube@1.27.11",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r12_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R11",
		RelatedProjectReleases: []string{
			"kube@1.25.16",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r11_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R10",
		RelatedProjectReleases: []string{
			"kube@1.25.15",
			"kube@1.25.16",
			"kube@1.26.11",
			"kube@1.26.13",
			"kube@1.26.14",
			"kube@1.27.7",
			"kube@1.27.11",
			"kube@1.28.3",
			"kube@1.28.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r10_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R09",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r09_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R08",
		RelatedProjectReleases: []string{
			"kube@1.25.16",
			"kube@1.26.10",
			"kube@1.26.13",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#March_20_2024
	{
		Project: GKE.ID,
		Version: "2024-R07",
		RelatedProjectReleases: []string{
			"kube@1.26.11",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#March_07_2024
	{
		Project: GKE.ID,
		Version: "2024-R06",
		RelatedProjectReleases: []string{
			"kube@1.27.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#March_04_2024
	{
		Project:                GKE.ID,
		Version:                "2024-R05",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r05_version_updates
	{
		Project:                GKE.ID,
		Version:                "2024-R04",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r04_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R03",
		RelatedProjectReleases: []string{
			"kube@1.27.3",
			"kube@1.27.7",
			"kube@1.28.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r03_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R02",
		RelatedProjectReleases: []string{
			"kube@1.24.17",
			"kube@1.25.10",
			"kube@1.25.13",
			"kube@1.25.15",
			"kube@1.25.16",
			"kube@1.26.10",
			"kube@1.26.11",
			"kube@1.27.4",
			"kube@1.27.5",
			"kube@1.27.7",
			"kube@1.27.8",
			"kube@1.28.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r02_version_updates
	{
		Project: GKE.ID,
		Version: "2024-R01",
		RelatedProjectReleases: []string{
			"kube@1.24.16",
			"kube@1.24.17",
			"kube@1.25.12",
			"kube@1.25.15",
			"kube@1.26.7",
			"kube@1.26.8",
			"kube@1.26.10",
			"kube@1.27.5",
			"kube@1.27.7",
			"kube@1.28.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2024-r01_version_updates
	{
		Project:                GKE.ID,
		Version:                "2023-R26",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r26_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R25",
		RelatedProjectReleases: []string{
			"kube@1.24.15",
			"kube@1.24.16",
			"kube@1.24.17",
			"kube@1.25.13",
			"kube@1.26.5",
			"kube@1.26.7",
			"kube@1.26.8",
			"kube@1.27.4",
			"kube@1.27.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r25_version_updates
	{
		Project:                GKE.ID,
		Version:                "2023-R24",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r24_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R23",
		RelatedProjectReleases: []string{
			"kube@1.24.14",
			"kube@1.24.15",
			"kube@1.26.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r23_version_updates
	{
		Project:                GKE.ID,
		Version:                "2023-R22",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r22_version_updates
	{
		Project:                GKE.ID,
		Version:                "2023-R20",
		RelatedProjectReleases: []string{},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r20_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R19",
		RelatedProjectReleases: []string{
			"kube@1.26.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r19_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R18",
		RelatedProjectReleases: []string{
			"kube@1.23.17",
			"kube@1.24.14",
			"kube@1.24.16",
			"kube@1.25.10",
			"kube@1.25.12",
			"kube@1.26.7",
			"kube@1.27.3",
			"kube@1.27.4",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r18_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R17",
		RelatedProjectReleases: []string{
			"kube@1.22.17",
			"kube@1.23.17",
			"kube@1.24.14",
			"kube@1.24.15",
			"kube@1.25.10",
			"kube@1.26.5",
			"kube@1.27.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r17_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R16",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.23.17",
			"kube@1.24.13",
			"kube@1.24.14",
			"kube@1.25.9",
			"kube@1.25.10",
			"kube@1.26.5",
			"kube@1.27.2",
			"kube@1.27.3",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r16_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R15",
		RelatedProjectReleases: []string{
			"kube@1.23.17",
			"kube@1.24.12",
			"kube@1.24.13",
			"kube@1.24.14",
			"kube@1.25.8",
			"kube@1.25.9",
			"kube@1.25.10",
			"kube@1.26.5",
			"kube@1.27.2",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r15_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R14",
		RelatedProjectReleases: []string{
			"kube@1.22.17",
			"kube@1.23.17",
			"kube@1.24.12",
			"kube@1.24.14",
			"kube@1.25.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r14_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R13",
		RelatedProjectReleases: []string{
			"kube@1.24.11",
			"kube@1.24.13",
			"kube@1.25.8",
			"kube@1.26.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r13_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R12",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.25.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r12_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R11",
		RelatedProjectReleases: []string{
			"kube@1.22.17",
			"kube@1.23.17",
			"kube@1.24.10",
			"kube@1.24.12",
			"kube@1.25.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r11_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R10",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.22.17",
			"kube@1.23.16",
			"kube@1.23.17",
			"kube@1.24.9",
			"kube@1.24.11",
			"kube@1.24.12",
			"kube@1.25.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r10_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R09",
		RelatedProjectReleases: []string{
			"kube@1.24.10",
			"kube@1.24.11",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r09_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R08",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.22.17",
			"kube@1.23.16",
			"kube@1.24.10",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r08_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R07",
		RelatedProjectReleases: []string{
			"kube@1.22.17",
			"kube@1.23.16",
			"kube@1.24.10",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r07_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R06",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.23.14",
			"kube@1.23.16",
			"kube@1.24.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r06_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R05",
		RelatedProjectReleases: []string{
			"kube@1.22.16",
			"kube@1.22.17",
			"kube@1.24.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r05_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R04",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.22.15",
			"kube@1.22.16",
			"kube@1.23.13",
			"kube@1.23.14",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r04_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R03",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.24.8",
			"kube@1.24.9",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r03_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R02",
		RelatedProjectReleases: []string{
			"kube@1.23.14",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r02_version_updates
	{
		Project: GKE.ID,
		Version: "2023-R01",
		RelatedProjectReleases: []string{
			"kube@1.22.15",
			"kube@1.22.16",
			"kube@1.23.11",
			"kube@1.23.14",
			"kube@1.24.7",
			"kube@1.24.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2023-r01_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R28",
		RelatedProjectReleases: []string{
			"kube@1.23.13",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r28_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R27",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.22.12",
			"kube@1.22.15",
			"kube@1.23.11",
			"kube@1.23.13",
			"kube@1.24.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r27_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R26",
		RelatedProjectReleases: []string{
			"kube@1.22.15",
			"kube@1.24.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r26_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R25",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.22.12",
			"kube@1.22.15",
			"kube@1.23.8",
			"kube@1.23.11",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r25_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R24",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r24_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R23",
		RelatedProjectReleases: []string{
			"kube@1.21.14",
			"kube@1.22.12",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r23_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R22",
		RelatedProjectReleases: []string{
			"kube@1.20.15",
			"kube@1.21.13",
			"kube@1.21.14",
			"kube@1.22.12",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r22_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R21",
		RelatedProjectReleases: []string{
			"kube@1.21.12",
			"kube@1.21.14",
			"kube@1.22.10",
			"kube@1.22.12",
			"kube@1.23.7",
			"kube@1.23.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r21_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R20",
		RelatedProjectReleases: []string{
			"kube@1.20.15",
			"kube@1.21.13",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r20_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R19",
		RelatedProjectReleases: []string{
			"kube@1.21.12",
			"kube@1.21.14",
			"kube@1.22.8",
			"kube@1.22.10",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r19_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R18",
		RelatedProjectReleases: []string{
			"kube@1.20.15",
			"kube@1.21.12",
			"kube@1.21.13",
			"kube@1.22.8",
			"kube@1.22.10",
			"kube@1.23.6",
			"kube@1.23.7",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r18_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R17",
		RelatedProjectReleases: []string{
			"kube@1.20.15",
			"kube@1.21.12",
			"kube@1.23.6",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r17_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R16",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.11",
			"kube@1.21.12",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r16_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R15",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.11",
			"kube@1.21.12",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r15_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R14",
		RelatedProjectReleases: []string{
			"kube@1.21.11",
			"kube@1.22.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r14_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R13",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.11",
			"kube@1.22.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r13_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R12",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.10",
			"kube@1.21.11",
			"kube@1.22.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r12_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R11",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.11",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r11_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R10",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.10",
			"kube@1.22.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r10_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R9",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.10",
			"kube@1.22.8",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r9_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R8",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.15",
			"kube@1.21.5",
			"kube@1.21.10",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r8_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R7",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.12",
			"kube@1.20.15",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r7_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R6",
		RelatedProjectReleases: []string{
			"kube@1.20.15",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r6_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R5",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.11",
			"kube@1.20.15",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r5_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R4",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
			"kube@1.20.11",
			"kube@1.20.15",
			"kube@1.21.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r4_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R3",
		RelatedProjectReleases: []string{
			"kube@1.19.15",
			"kube@1.19.16",
			"kube@1.21.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r3_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R02",
		RelatedProjectReleases: []string{
			"kube@1.19.16",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r02_version_updates
	{
		Project: GKE.ID,
		Version: "2022-R01",
		RelatedProjectReleases: []string{
			"kube@1.20.12",
			"kube@1.21.5",
		},
	},
	// source: https://cloud.google.com/kubernetes-engine/docs/release-notes#2022-r01_version_updates
}

func init() {
	RegisterProject(&GKE)
	RegisterProjectReleases(GKE.ID, GKEProjectReleases)
	RegisterCurationConfig(GKE.ID, GKECurationConfig)
}
