// +build integration
// Copyright 2019 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testing

import (
	. "gopkg.in/check.v1"

	"github.com/kanisterio/kanister/pkg/app"
)

// Register Applications to Integration Suite

// pitr-postgresql app
type PITRPostgreSQL struct {
	IntegrationSuite
}

var _ = Suite(&PITRPostgreSQL{
	IntegrationSuite{
		name:      "pitr-postgres",
		namespace: "pitr-postgres-test",
		app:       app.NewPostgresDB("pitr-postgres"),
		bp:        app.NewPITRBlueprint("pitr-postgres"),
		profile:   newSecretProfile(),
	},
})

// postgres app
type PostgreSQL struct {
	IntegrationSuite
}

var _ = Suite(&PostgreSQL{
	IntegrationSuite{
		name:      "postgres",
		namespace: "postgres-test",
		app:       app.NewPostgresDB("postgres"),
		bp:        app.NewBlueprint("postgres"),
		profile:   newSecretProfile(),
	},
})

// mysql app
type MySQL struct {
	IntegrationSuite
}

var _ = Suite(&MySQL{
	IntegrationSuite{
		name:      "mysql",
		namespace: "mysql-test",
		app:       app.NewMysqlDB("mysql"),
		bp:        app.NewBlueprint("mysql"),
		profile:   newSecretProfile(),
	},
})

// Elasticsearch app
type Elasticsearch struct {
	IntegrationSuite
}

var _ = Suite(&Elasticsearch{
	IntegrationSuite{
		name:      "elasticsearch",
		namespace: "es-test",
		app:       app.NewElasticsearchInstance("elasticsearch"),
		bp:        app.NewBlueprint("elasticsearch"),
		profile:   newSecretProfile(),
	},
})

// MongoDB app
type MongoDB struct {
	IntegrationSuite
}

var _ = Suite(&MongoDB{
	IntegrationSuite{
		name:      "mongo",
		namespace: "mongo-test",
		app:       app.NewMongoDB("mongo"),
		bp:        app.NewBlueprint("mongo"),
		profile:   newSecretProfile(),
	},
})

// Cassandra App
type Cassandra struct {
	IntegrationSuite
}

var _ = Suite(&Cassandra{IntegrationSuite{
	name:      "cassandra",
	namespace: "cassandra-test",
	app:       app.NewCassandraInstance("cassandra"),
	bp:        app.NewBlueprint("cassandra"),
	profile:   newSecretProfile(),
},
})

// Couchbase app
type Couchbase struct {
	IntegrationSuite
}

var _ = Suite(&Couchbase{
	IntegrationSuite{
		name:      "couchbase",
		namespace: "couchbase-test",
		app:       app.NewCouchbaseDB("couchbase"),
		bp:        app.NewBlueprint("couchbase"),
		profile:   newSecretProfile(),
	},
})

// rds-postgres app
type RDSPostgreSQL struct {
	IntegrationSuite
}

var _ = Suite(&RDSPostgreSQL{
	IntegrationSuite{
		name:      "rds-postgres",
		namespace: "rds-postgres-test",
		app:       app.NewRDSPostgresDB("rds-postgres"),
		bp:        app.NewBlueprint("rds-postgres"),
		profile:   newSecretProfile(),
	},
})

type FoundationDB struct {
	IntegrationSuite
}

var _ = Suite(&FoundationDB{
	IntegrationSuite{
		name:      "foundationdb",
		namespace: "fdb-test",
		app:       app.NewFoundationDB("foundationdb"),
		bp:        app.NewBlueprint("foundationdb"),
		profile:   newSecretProfile(),
	},
})

// rds-postgres-dump app
// Create snapshot, export data and restore from dump
type RDSPostgreSQLDump struct {
	IntegrationSuite
}

var _ = Suite(&RDSPostgreSQLDump{
	IntegrationSuite{
		name:      "rds-postgres-dump",
		namespace: "rds-postgres-dump-test",
		app:       app.NewRDSPostgresDB("rds-postgres-dump"),
		bp:        app.NewBlueprint("rds-postgres-dump"),
		profile:   newSecretProfile(),
	},
})

// rds-postgres-snap app
// Create snapshot and restore from snapshot
type RDSPostgreSQLSnap struct {
	IntegrationSuite
}

var _ = Suite(&RDSPostgreSQLSnap{
	IntegrationSuite{
		name:      "rds-postgres-snap",
		namespace: "rds-postgres-snap-test",
		app:       app.NewRDSPostgresDB("rds-postgres-snap"),
		bp:        app.NewBlueprint("rds-postgres-snap"),
		profile:   newSecretProfile(),
	},
})

// Mysql Instance that is deployed through DeploymentConfig on OpenShift cluster
type MysqlDBDepConfig struct {
	IntegrationSuite
}

var _ = Suite(&MysqlDBDepConfig{
	IntegrationSuite{
		name:      "mysqldb",
		namespace: "mysqldc-test",
		app:       app.NewMysqlDepConfig("mysqldeploymentconfig"),
		bp:        app.NewBlueprint("mysql-dep-config"),
		profile:   newSecretProfile(),
	},
})
