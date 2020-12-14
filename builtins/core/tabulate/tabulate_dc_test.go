package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

var (
	inDC = `Builds, (re)creates, starts, and attaches to containers for a service.

Unless they are already running, this command also starts any linked services.

The ` + "`" + `docker-compose up` + "`" + ` command aggregates the output of each container. When
the command exits, all containers are stopped. Running ` + "`" + `docker-compose up -d` + "`" + `
starts the containers in the background and leaves them running.

If there are existing containers for a service, and the service's configuration
or image was changed after the container's creation, ` + "`" + `docker-compose up` + "`" + ` picks
up the changes by stopping and recreating the containers (preserving mounted
volumes). To prevent Compose from picking up changes, use the ` + "`" + `--no-recreate` + "`" + `
flag.

If you want to force Compose to stop and recreate all containers, use the
` + "`" + `--force-recreate` + "`" + ` flag.

Usage: up [options] [--scale SERVICE=NUM...] [SERVICE...]

Options:
    -d                         Detached mode: Run containers in the background,
                               print new container names.
                               Incompatible with --abort-on-container-exit.
    --no-color                 Produce monochrome output.
    --no-deps                  Don't start linked services.
    --force-recreate           Recreate containers even if their configuration
                               and image haven't changed.
                               Incompatible with --no-recreate.
    --no-recreate              If containers already exist, don't recreate them.
                               Incompatible with --force-recreate.
    --no-build                 Don't build an image, even if it's missing.
    --no-start                 Don't start the services after creating them.
    --build                    Build images before starting containers.
    --abort-on-container-exit  Stops all containers if any container was stopped.
                               Incompatible with -d.
    -t, --timeout TIMEOUT      Use this timeout in seconds for container shutdown
                               when attached or when containers are already
                               running. (default: 10)
    --remove-orphans           Remove containers for services not
                               defined in the Compose file
    --exit-code-from SERVICE   Return the exit code of the selected service container.
                               Implies --abort-on-container-exit.
    --scale SERVICE=NUM        Scale SERVICE to NUM instances. Overrides the ` + "`" + `scale` + "`" + `
                               setting in the Compose file if present.
`

	jsonDC = `{"--abort-on-container-exit":"Stops all containers if any container was stopped. Incompatible with -d.","--build":"Build images before starting containers.","--exit-code-from SERVICE":"Return the exit code of the selected service container. Implies --abort-on-container-exit.","--force-recreate":"Recreate containers even if their configuration and image haven't changed. Incompatible with --no-recreate.","--no-build":"Don't build an image, even if it's missing.","--no-color":"Produce monochrome output.","--no-deps":"Don't start linked services.","--no-recreate":"If containers already exist, don't recreate them. Incompatible with --force-recreate.","--no-start":"Don't start the services after creating them.","--remove-orphans":"Remove containers for services not defined in the Compose file","--scale SERVICE=NUM":"Scale SERVICE to NUM instances. Overrides the ` + "`" + `scale` + "`" + ` setting in the Compose file if present.","--timeout TIMEOUT":"Use this timeout in seconds for container shutdown when attached or when containers are already running. (default: 10)","-d":"Detached mode: Run containers in the background, print new container names. Incompatible with --abort-on-container-exit.","-t":"Use this timeout in seconds for container shutdown when attached or when containers are already running. (default: 10)"}`
)

func TestTabulateDockerCompose(t *testing.T) {
	/*test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inDC,
		types.Generic,
		[]string{},
		jsonDC,
		nil,
	)*/

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inDC,
		types.Generic,
		[]string{fMap, fColumnWraps, fSplitComma},
		jsonDC,
		nil,
	)
}
