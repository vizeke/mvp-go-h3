### H3 Sample

Heatmap view using Mapbox, Postgresql, Uber H3 and GoLang

## How to run

Add a `.env` file like `.env-sample`

run:

```go
go run ./...
```

## Application

This app was a proof of concept. Used to demostrate how we could monitor, in real time, the location of all the drivers and orders in the system.

## Endpoints

- GET: /drivers/:resolution
- GET: /orders/:resolution

**resolution** it's like a 'zoom' used by H3 sdk. Represented by an non-negative integer.

### Response

A heatmap with the necessary info to plot the map.

```json
{
  "MaxCount": 50,
  "HeatMap": {
    "1": [
      {
        "Cell": "8880104c07fffff",
        "Boundary": [
          { "Lat": -3.743333542106514, "Lng": -38.49133284031708 },
          { "Lat": -3.739956384486233, "Lng": -38.48731269019193 },
          { "Lat": -3.7347849517230314, "Lng": -38.48823129246414 },
          { "Lat": -3.7329906151638474, "Lng": -38.493170000384445 },
          { "Lat": -3.736367725764249, "Lng": -38.497190189573715 },
          { "Lat": -3.74153921993452, "Lng": -38.49627163179031 }
        ],
        "Count": 9,
        "Valid": true
      }
      ...
    ]
  }
}
```

## Screenshots

![current-drivers](/doc/sample1.png "Text to show on mouseover")
