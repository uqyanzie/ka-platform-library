package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func neo4jInit() (neo4j.DriverWithContext, error) {

	ctx := context.Background()

	uri := "neo4j+s://55b443c3.databases.neo4j.io"
	username := "neo4j"
	password := "fIHzZjQlVUAW5t4Dx_dtKNnZu18RN7sSA7sbyfm6Vrg"

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	//defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
