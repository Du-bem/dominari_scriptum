<script>
import SolarSystem from "./plot.js";
class Planet {
  constructor(name, diameter, orbitRadius, rotationSpeed, color) {
    this.name = name;
    this.diameter = diameter; // Actual size
    this.orbitRadius = orbitRadius;
    this.rotationSpeed = rotationSpeed;
    this.color = color;
    this.angle = 0;
  }

  update(deltaTime) {
    this.angle += this.rotationSpeed * deltaTime;
  }

  getPosition() {
    return [
      this.orbitRadius * Math.sin(this.angle),
      0,
      this.orbitRadius * Math.cos(this.angle)
    ];
  }

  getColor() {
    return this.color;
  }

  getSize() {
    // Return the actual diameter
    return this.diameter;
  }
}

// Create solar system with appropriate size parameters for our scale
const solarSystem = new SolarSystem(0, [], {
  minSize: 0.1, // Smallest visual size
  maxSize: 1.0, // Largest visual size
  sizeScaleFactor: 0.1, // Scale factor to convert real units to visual units
  sphereSegments: 24 // Detail level for spheres
});

// Create planets with realistic relative sizes
const planets = [
  // new Planet("Sun", 1.0, 0, 0, 0xffff00), // Size = 1.0 (reference)
  new Planet("Mercury", 0.038, 1.8, 0.02, 0xaaaaaa), // ~3.8% of Sun's diameter
  new Planet("Venus", 0.095, 3.0, 0.015, 0xe39e1c), // ~9.5% of Sun's diameter
  new Planet("Earth", 0.1, 5.0, 0.01, 0x2233dd), // 10% of Sun's diameter
  new Planet("Mars", 0.053, 7.0, 0.008, 0xdd3333), // ~5.3% of Sun's diameter
  new Planet("Jupiter", 0.11, 12.0, 0.002, 0xdda52c) // 11% of Sun's diameter
];

// Add planets to the solar system
for (const planet of planets) {
  solarSystem.addObject(planet.name, planet);
}

// Custom animation callback
solarSystem.setAnimationCallback((objects, time) => {
  // Update all planets
  for (const planet of planets) {
    planet.update(0.5);
  }
});

// Start animation
solarSystem.animate();
</script>

<template lang="html"> </template>

<style lang="css" scoped></style>
