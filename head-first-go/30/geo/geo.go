package geo

import "errors"

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("Invalid latitude")
	}
	c.latitude = latitude
	return nil
}

func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("Invalid longitude")
	}
	c.longitude = longitude
	return nil
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}
