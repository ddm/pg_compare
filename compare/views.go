package compare

import (
	"errors"
	"fmt"

	"github.com/jelmersnoeck/pg_compare/drivers"
	"github.com/jelmersnoeck/pg_compare/log"
)

func Views(source, destination drivers.Driver) error {
	sViews, err := viewDefinitions(source)
	if err != nil {
		return err
	}

	dViews, err := viewDefinitions(destination)
	if err != nil {
		return err
	}

	if len(sViews) != len(dViews) {
		return fmt.Errorf("Number of views for destination (%d) doesn't match source (%d)", len(dViews), len(sViews))
	}

	mismatches := map[string]map[string]string{}
	for view, def := range sViews {
		dst, ok := dViews[view]
		if !ok {
			return fmt.Errorf("Destination does not have source view `%s`", view)
		}

		if def != dst {
			mismatches[view] = map[string]string{
				"source":      def,
				"destination": dst,
			}
		}
	}

	if len(mismatches) != 0 {
		log.Printf("Mismatches: ", mismatches)
		return errors.New("mismatches")
	}

	return nil
}

func viewDefinitions(d drivers.Driver) (map[string]string, error) {
	views, err := d.Views()
	if err != nil {
		return nil, err
	}

	viewMap := map[string]string{}
	for _, name := range views {
		log.Printf("Getting view information for view `%s` on host `%s`", name, d.Host())

		def, err := d.ViewDefinition(name)
		if err != nil {
			return nil, err
		}

		viewMap[name] = def
	}

	return viewMap, nil
}
