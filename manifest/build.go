package manifest

import (
	"fmt"
	"strings"
)

func (m *Manifest) Build(dir string) error {
	builds := map[string]string{}
	pulls := map[string]string{}

	for _, service := range m.Services {
		switch {
		case service.Build != "":
			builds[fmt.Sprintf("%s|%s", service.Build, coalesce(service.Dockerfile, "Dockerfile"))] = service.Tag()
		case service.Image != "":
			pulls[service.Image] = service.Tag()
		}
	}

	for build, tag := range builds {
		parts := strings.SplitN(build, "|", 2)

		args := []string{"build"}

		args = append(args, "-f", parts[1])
		args = append(args, "-t", tag)
		args = append(args, parts[0])

		runPrefix(systemPrefix(m), Docker(args...))
	}

	for image, tag := range pulls {
		args := []string{"pull"}

		args = append(args, image)

		runPrefix(systemPrefix(m), Docker(args...))
		runPrefix(systemPrefix(m), Docker("tag", image, tag))
	}

	return nil
}
