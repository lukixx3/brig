package config

// Defaults is the default validation for brig
var Defaults = DefaultMapping{
	"daemon": DefaultMapping{
		"port": DefaultEntry{
			Default:      6666,
			NeedsRestart: true,
			Docs:         "Port of the daemon",
		},
	},
	"sync": DefaultMapping{
		"ignored_removed": DefaultEntry{
			Default:      false,
			NeedsRestart: false,
			Docs:         "Do not remove wha the remote removed",
		},
		"conflict_strategy": DefaultEntry{
			Default:      "marker",
			NeedsRestart: false,
		},
	},
	"data": DefaultMapping{
		"ipfs": DefaultMapping{
			"path": DefaultEntry{
				Default:      "",
				NeedsRestart: true,
				Docs:         "Root directory of the ipfs repository",
			},
		},
		"compress": DefaultMapping{
			"default_algo": DefaultEntry{
				Default:      "snappy",
				NeedsRestart: false,
				Docs:         "What compression algorithm to use by default",
			},
		},
	},
}
