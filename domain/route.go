package domain

import (
	"fmt"
	"sort"

	"github.com/boreq/errors"
	"github.com/boreq/velo/internal/eventsourcing"
)

type Route struct {
	uuid   RouteUUID
	points []Point

	es eventsourcing.EventSourcing
}

func NewRoute(uuid RouteUUID, points []Point) (*Route, error) {
	if uuid.IsZero() {
		return nil, errors.New("zero value of uuid")
	}

	if len(points) == 0 {
		return nil, errors.New("missing points")
	}

	points = NormaliseRoutePoints(points)

	if len(points) < 2 {
		return nil, errors.New("a route has to have at least 2 points")
	}

	route := &Route{}

	if err := route.update(
		RouteCreated{
			UUID:   uuid,
			Points: points,
		},
	); err != nil {
		return nil, errors.Wrap(err, "could not consume the initial event")

	}

	return route, nil
}

func NewRouteFromHistory(events eventsourcing.EventSourcingEvents) (*Route, error) {
	route := &Route{}

	for _, event := range events {
		if err := route.update(event.Event); err != nil {
			return nil, errors.Wrap(err, "could not consume an event")
		}
		route.es.LoadVersion(event)
	}

	route.PopChanges()

	return route, nil
}

func (r Route) UUID() RouteUUID {
	return r.uuid
}

func (r Route) Points() []Point {
	points := make([]Point, len(r.points))
	copy(points, r.points)
	return points
}

func (r Route) IsZero() bool {
	return r.uuid.IsZero() // if uuid is set then everything else must be as well
}

func (r *Route) PopChanges() eventsourcing.EventSourcingEvents {
	return r.es.PopChanges()
}

func (r *Route) update(event eventsourcing.Event) error {
	switch e := event.(type) {
	case RouteCreated:
		r.handleRouteCreated(e)
	default:
		return fmt.Errorf("unknown event type '%T'", event)
	}

	return r.es.Record(event)
}

func (r *Route) handleRouteCreated(event RouteCreated) {
	r.uuid = event.UUID
	r.points = event.Points
}

type RouteCreated struct {
	UUID   RouteUUID
	Points []Point
}

func (r RouteCreated) EventType() eventsourcing.EventType {
	return "RouteCreated_v1"
}

func NormaliseRoutePoints(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		return points[i].Time().Before(points[j].Time())
	})

	var normalised []Point

	for i, point := range points {
		if len(normalised) == 0 || i == len(points)-1 {
			normalised = append(normalised, point)
		} else {
			previous := normalised[len(normalised)-1]
			if shouldAdd(previous, point) {
				normalised = append(normalised, point)
			}
		}
	}

	return normalised
}

// Points that occur more often than that will be dropped.
const distanceBetweenPointsInMeters = 10

func shouldAdd(previous Point, next Point) bool {
	return next.Position().Distance(previous.Position()).Float64() >= distanceBetweenPointsInMeters
}
